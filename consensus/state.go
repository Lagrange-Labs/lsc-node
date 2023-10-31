package consensus

import (
	"context"
	"fmt"
	"sort"
	"sync"
	"time"

	"github.com/Lagrange-Labs/lagrange-node/consensus/types"
	govtypes "github.com/Lagrange-Labs/lagrange-node/governance/types"
	"github.com/Lagrange-Labs/lagrange-node/logger"
	networktypes "github.com/Lagrange-Labs/lagrange-node/network/types"
	sequencertypes "github.com/Lagrange-Labs/lagrange-node/sequencer/types"
	storetypes "github.com/Lagrange-Labs/lagrange-node/store/types"
	"github.com/Lagrange-Labs/lagrange-node/utils"
	"github.com/umbracle/go-eth-consensus/bls"
)

const CheckInterval = 1 * time.Second

// State handles the consensus process.
type State struct {
	validators *types.ValidatorSet

	rounds          map[uint64]*types.RoundState
	lastBlockNumber uint64
	rwMutex         *sync.RWMutex

	proposer      *bls.SecretKey
	storage       storageInterface
	roundLimit    time.Duration
	roundInterval time.Duration
	batchSize     uint32
	chainID       uint32
	lastCommittee *govtypes.CommitteeRoot

	chStop chan struct{}
}

// NewState returns a new State.
func NewState(cfg *Config, storage storageInterface, chainID uint32) *State {
	privKey, err := utils.HexToBlsPrivKey(cfg.ProposerPrivateKey)
	if err != nil {
		logger.Fatalf("failed to parse the proposer private key: %v", err)
	}

	if err := storage.AddNode(context.Background(),
		&networktypes.ClientNode{
			StakeAddress: cfg.OperatorAddress,
			PublicKey:    utils.BlsPubKeyToHex(privKey.GetPublicKey()),
			ChainID:      chainID,
		},
	); err != nil {
		logger.Fatalf("failed to add the proposer node: %v", err)
	}

	chStop := make(chan struct{})

	return &State{
		proposer:      privKey,
		storage:       storage,
		roundLimit:    time.Duration(cfg.RoundLimit),
		roundInterval: time.Duration(cfg.RoundInterval),
		chainID:       chainID,
		batchSize:     cfg.BatchSize,
		chStop:        chStop,
		rwMutex:       &sync.RWMutex{},
	}
}

// OnStart loads the first unverified block and starts the round.
func (s *State) OnStart() {
	logger.Info("Consensus process is started with the batch size: ", s.batchSize)

	for {
		// check if chStop is triggered
		select {
		case <-s.chStop:
			return
		default:
		}

		lastBlock, err := s.storage.GetLastFinalizedBlock(context.Background(), s.chainID)
		if err != nil {
			logger.Errorf("failed to get the last finalized block: %v", err)
			return
		}
		if lastBlock != nil {
			s.lastBlockNumber = lastBlock.BlockNumber()
			if len(lastBlock.AggSignature) > 0 {
				// the last block is finalized
				s.lastBlockNumber += 1
			}
			break
		}
		logger.Info("waiting for the first block")
		time.Sleep(CheckInterval)
	}

	for {
		// check if chStop is triggered
		select {
		case <-s.chStop:
			return
		default:
		}

		logger.Infof("start the batch rounds from the block number %v", s.lastBlockNumber)
		if err := s.startRound(s.lastBlockNumber); err != nil {
			logger.Errorf("failed to start the round: %v", err)
			time.Sleep(s.roundInterval)
			continue
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(s.roundLimit))
		defer cancel()
		isVoted := s.processRound(ctx)
		if !isVoted {
			logger.Error("the current batch is not finalized within the round limit")
		}

		// store the evidences
		for _, round := range s.rounds {
			evidences, err := round.GetEvidences()
			if err != nil {
				logger.Errorf("failed to get the evidences: %v", err)
				continue
			}
			if len(evidences) > 0 {
				if err := s.storage.AddEvidences(ctx, evidences); err != nil {
					logger.Errorf("failed to add the evidences: %v", err)
					continue
				}
			}
		}

		// store the finalized block
		failedRounds := make(map[uint64]*types.RoundState)
		lastBlockNumber := uint64(0)
		for blockNumber, round := range s.rounds {
			if !round.IsFinalized() {
				logger.Errorf("the block %d is not finalized", round.GetCurrentBlockNumber())
				failedRounds[blockNumber] = round
				continue
			}
			if err := s.storage.UpdateBlock(context.Background(), round.GetCurrentBlock()); err != nil {
				logger.Errorf("failed to update the block %d: %v", round.GetCurrentBlockNumber(), err)
				failedRounds[blockNumber] = round
				continue
			}
			if lastBlockNumber < blockNumber {
				lastBlockNumber = blockNumber
			}
			logger.Infof("the block %d is finalized", blockNumber)
		}

		// update the last block number
		s.rwMutex.Lock()
		s.rounds = failedRounds
		s.rwMutex.Unlock()

		if !isVoted {
			// TODO: handle the case when the batch is not finalized, now it will be run forever
			logger.Error("the infinite loop is started!")
			_ = s.processRound(context.Background())
			for blockNumber, round := range s.rounds {
				if err := s.storage.UpdateBlock(context.Background(), round.GetCurrentBlock()); err != nil {
					logger.Errorf("failed to update the block %d: %v", round.GetCurrentBlockNumber(), err)
					continue
				}
				if lastBlockNumber < blockNumber {
					lastBlockNumber = blockNumber
				}
				logger.Infof("the block %d is finalized", round.GetCurrentBlockNumber())
			}
		}

		s.rwMutex.Lock()
		s.lastBlockNumber = lastBlockNumber + 1
		s.rwMutex.Unlock()
	}
}

// OnStop stops the consensus process.
func (s *State) OnStop() {
	logger.Infof("OnStop() called")
	s.chStop <- struct{}{}
	close(s.chStop)
}

// AddCommit adds the commit to the round.
func (s *State) AddCommit(commit *sequencertypes.BlsSignature, pubKey, stakeAddr string) error {
	s.rwMutex.Lock()
	defer s.rwMutex.Unlock()

	if s.validators.GetVotingPower(stakeAddr) == 0 {
		return fmt.Errorf("the stake address %s is not registered", stakeAddr)
	}

	round, ok := s.rounds[commit.BlockNumber()]
	if !ok {
		return fmt.Errorf("the round for the block %d is not found", commit.BlockNumber())
	}
	round.AddCommit(commit, pubKey, stakeAddr)
	return nil
}

// GetOpenRoundBlocks returns the blocks that are not finalized yet.
func (s *State) GetOpenRoundBlocks(blockNumber uint64) []*sequencertypes.Block {
	s.rwMutex.RLock()
	defer s.rwMutex.RUnlock()

	if blockNumber > s.lastBlockNumber {
		return nil
	}

	blocks := make([]*sequencertypes.Block, 0)
	for _, round := range s.rounds {
		if !round.IsFinalized() {
			blocks = append(blocks, round.GetCurrentBlock())
		}
	}

	// sort the blocks by the block number
	sort.Slice(blocks, func(i, j int) bool {
		return blocks[i].BlockNumber() < blocks[j].BlockNumber()
	})

	return blocks
}

// IsFinalized returns true if all batch blocks are finalized.
func (s *State) IsFinalized(blockNumber uint64) bool {
	s.rwMutex.RLock()
	defer s.rwMutex.RUnlock()

	return blockNumber < s.lastBlockNumber
}

// startRound loads the next block batch and initializes the round state.
func (s *State) startRound(blockNumber uint64) error {
	s.rwMutex.Lock()
	defer s.rwMutex.Unlock()

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(s.roundLimit))
	defer cancel()
	blocks, err := s.getNextBlocks(ctx, blockNumber)
	if err != nil {
		return fmt.Errorf("getting the next block batch from %d is failed: %v", blockNumber, err)
	}

	epochNumber, err := s.storage.GetLastCommitteeEpochNumber(context.Background(), s.chainID)
	if err != nil {
		return fmt.Errorf("failed to get the last committee epoch number: %v", err)
	}
	if epochNumber == 0 {
		return fmt.Errorf("the last committee epoch number is 0")
	}

	committee := s.lastCommittee
	if committee == nil || epochNumber > committee.EpochNumber {
		// update the last committee
		lastCommittee, err := s.storage.GetLastCommitteeRoot(context.Background(), s.chainID, false)
		if err != nil {
			return fmt.Errorf("failed to get the last committee root: %v", err)
		}
		if lastCommittee == nil {
			return fmt.Errorf("the last committee root is nil")
		}
		if lastCommittee.TotalVotingPower == 0 {
			return fmt.Errorf("the last committee root has 0 voting power")
		}
		s.lastCommittee = lastCommittee
		if committee == nil {
			lastCommittee.IsFinalized = true
			committee = lastCommittee
			if err := s.storage.UpdateCommitteeRoot(context.Background(), lastCommittee); err != nil {
				return fmt.Errorf("failed to update the last committee root: %v", err)
			}
		}
	}

	logger.Infof("the registered nodes are loaded: %v", committee.Operators)

	s.validators = types.NewValidatorSet(committee.Operators, committee.TotalVotingPower)
	if s.validators.GetTotalVotingPower()*3 < committee.TotalVotingPower*2 {
		return fmt.Errorf("the voting power of the registered nodes voting power %d is less than 2/3 of the total voting power %d", s.validators.GetTotalVotingPower(), committee.TotalVotingPower)
	}

	pubKey := utils.BlsPubKeyToHex(s.proposer.GetPublicKey())
	s.rounds = make(map[uint64]*types.RoundState)

	blockNumbers := make([]uint64, 0)
	for _, block := range blocks {
		// epochNumber, err := s.
		block.BlockHeader = &sequencertypes.BlockHeader{}
		block.BlockHeader.CurrentCommittee = committee.CurrentCommitteeRoot
		block.BlockHeader.NextCommittee = committee.CurrentCommitteeRoot
		block.BlockHeader.EpochBlockNumber = committee.EpochBlockNumber
		block.BlockHeader.TotalVotingPower = committee.TotalVotingPower

		// generate a proposer signature
		blsSigHash := block.BlsSignature().Hash()
		signature, err := s.proposer.Sign(blsSigHash)
		if err != nil {
			return err
		}
		block.BlockHeader.ProposerSignature = utils.BlsSignatureToHex(signature)
		block.BlockHeader.ProposerPubKey = pubKey

		round := types.NewEmptyRoundState()
		round.UpdateRoundState(block)
		s.rounds[block.BlockNumber()] = round
		blockNumbers = append(blockNumbers, block.BlockNumber())
	}

	if committee.EpochNumber < s.lastCommittee.EpochNumber {
		// update the next committee of the last block
		blocks[len(blocks)-1].BlockHeader.NextCommittee = s.lastCommittee.CurrentCommitteeRoot
		s.lastCommittee.IsFinalized = true
		if err := s.storage.UpdateCommitteeRoot(context.Background(), s.lastCommittee); err != nil {
			return fmt.Errorf("failed to update the last committee root: %v", err)
		}
	}
	logger.Infof("the next block batch is loaded: %v", blockNumbers)

	return nil
}

// getNextBlocks returns the next block batch from the storage.
// NOTE: it will return blocks more than 1 to parallelize.
func (s *State) getNextBlocks(ctx context.Context, blockNumber uint64) ([]*sequencertypes.Block, error) {
	blocks, err := s.storage.GetBlocks(ctx, uint32(s.chainID), blockNumber, s.batchSize)
	if err != nil && err != storetypes.ErrBlockNotFound {
		return nil, err
	}
	if len(blocks) > 0 {
		return blocks, nil
	}
	// in case the number of blocks is less than 2, wait for it to be added from the sequencer
	ticker := time.NewTicker(CheckInterval)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-ticker.C:
			blocks, err := s.storage.GetBlocks(context.Background(), s.chainID, blockNumber, s.batchSize)
			if err != nil {
				if err == storetypes.ErrBlockNotFound {
					continue
				}
				return nil, err
			}
			if len(blocks) < 2 {
				continue
			}
			return blocks, nil
		}
	}
}

// processRound processes the round.
func (s *State) processRound(ctx context.Context) bool {
	checkCommit := func(round *types.RoundState) (bool, error) {
		if round.CheckEnoughVotingPower(s.validators) {
			round.BlockCommit()
			err := round.CheckAggregatedSignature()
			if err != nil {
				round.UnblockCommit()
				if err == types.ErrInvalidAggregativeSignature {
					logger.Warnf("the aggregated signature is invalid for the block %d", round.GetCurrentBlockNumber())
					return false, nil
				}
				return false, err
			}
			return true, nil
		}
		return false, nil
	}

	wg := sync.WaitGroup{}
	wg.Add(len(s.rounds))

	for _, round := range s.rounds {
		go func(round *types.RoundState) {
			ticker := time.NewTicker(s.roundInterval)
			defer ticker.Stop()
			defer wg.Done()

			for {
				select {
				case <-ctx.Done():
					return
				case <-ticker.C:
					isFinalized, err := checkCommit(round)
					if err != nil {
						logger.Errorf("failed to check the commit for the block %d: %v", round.GetCurrentBlockNumber(), err)
						return
					}
					if isFinalized {
						return
					}
				}
			}
		}(round)
	}
	wg.Wait()

	isAllFinalized := true
	for _, round := range s.rounds {
		if !round.IsFinalized() {
			isAllFinalized = false
		}
	}

	return isAllFinalized
}

package consensus

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/Lagrange-Labs/lagrange-node/consensus/types"
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
	var err error
	logger.Info("Consensus process is started with the batch size: ", s.batchSize)

	for {
		// check if chStop is triggered
		select {
		case <-s.chStop:
			return
		default:
		}

		isFinalized := false
		s.lastBlockNumber, isFinalized, err = s.storage.GetLastFinalizedBlockNumber(context.Background(), s.chainID)
		if err != nil {
			logger.Errorf("failed to get the last finalized block number: %v", err)
			return
		}
		if s.lastBlockNumber > 0 {
			if isFinalized {
				// the last block is not finalized yet
				s.lastBlockNumber = s.lastBlockNumber + 1
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
		for blockNumber, round := range s.rounds {
			if !round.IsFinalized() {
				logger.Errorf("the block %d is not finalized", round.GetCurrentBlockNumber())
				failedRounds[blockNumber] = round
				continue
			}
			if err := s.storage.UpdateBlock(context.Background(), round.GetCurrentBlock()); err != nil {
				logger.Errorf("failed to update the block %d: %v", round.GetCurrentBlockNumber(), err)
				continue
			}

			logger.Infof("the block %d is finalized", round.GetCurrentBlockNumber())
		}

		// update the last block number
		s.rwMutex.Lock()
		s.lastBlockNumber += uint64(len(s.rounds))
		s.rounds = failedRounds
		s.rwMutex.Unlock()

		if !isVoted {
			// TODO: handle the case when the batch is not finalized, now it will be run forever
			logger.Error("the infinite loop is started!")
			_ = s.processRound(context.Background())
			for _, round := range s.rounds {
				if err := s.storage.UpdateBlock(context.Background(), round.GetCurrentBlock()); err != nil {
					logger.Errorf("failed to update the block %d: %v", round.GetCurrentBlockNumber(), err)
					continue
				}

				logger.Infof("the block %d is finalized", round.GetCurrentBlockNumber())
			}
		}
	}
}

// OnStop stops the consensus process.
func (s *State) OnStop() {
	logger.Infof("OnStop() called")
	s.chStop <- struct{}{}
	close(s.chStop)
}

// AddCommit adds the commit to the round.
func (s *State) AddCommit(commit *sequencertypes.BlsSignature, pubKey string) error {
	s.rwMutex.Lock()
	defer s.rwMutex.Unlock()

	round, ok := s.rounds[commit.BlockNumber()]
	if !ok {
		return fmt.Errorf("the round for the block %d is not found", commit.BlockNumber())
	}
	round.AddCommit(commit, pubKey)
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
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(s.roundLimit))
	defer cancel()
	blocks, err := s.getNextBlocks(ctx, blockNumber)
	if err != nil {
		return fmt.Errorf("getting the next block batch from %d is failed: %v", blockNumber, err)
	}

	committee, err := s.storage.GetLastCommitteeRoot(context.Background(), s.chainID)
	if err != nil {
		return fmt.Errorf("failed to get the last committee root: %v", err)
	}
	if committee == nil {
		return fmt.Errorf("the last committee root is nil")
	}
	if committee.TotalVotingPower == 0 {
		return fmt.Errorf("the total voting power of the last committee is 0")
	}

	nodes, err := s.storage.GetNodesByStatuses(context.Background(), []networktypes.NodeStatus{networktypes.NodeRegistered}, s.chainID)
	if err != nil {
		return err
	}

	s.validators = types.NewValidatorSet(nodes)
	if s.validators.GetTotalVotingPower()*3 < committee.TotalVotingPower*2 {
		return fmt.Errorf("the voting power of the registered nodes is less than 2/3 of the total voting power")
	}

	pubKey := utils.BlsPubKeyToHex(s.proposer.GetPublicKey())
	s.rounds = make(map[uint64]*types.RoundState)

	blockNumbers := make([]uint64, 0)
	for _, block := range blocks {
		block.BlockHeader = &sequencertypes.BlockHeader{}
		block.BlockHeader.CurrentCommittee = committee.CurrentCommitteeRoot
		block.BlockHeader.NextCommittee = committee.NextCommitteeRoot
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
	if len(blocks) > 1 {
		return blocks, err
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

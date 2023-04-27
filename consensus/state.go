package consensus

import (
	"context"
	"fmt"
	"time"

	"github.com/Lagrange-Labs/lagrange-node/consensus/types"
	"github.com/Lagrange-Labs/lagrange-node/logger"
	networktypes "github.com/Lagrange-Labs/lagrange-node/network/types"
	sequencertypes "github.com/Lagrange-Labs/lagrange-node/sequencer/types"
	storetypes "github.com/Lagrange-Labs/lagrange-node/store/types"
	"github.com/Lagrange-Labs/lagrange-node/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/umbracle/go-eth-consensus/bls"
)

const CheckInterval = 500 * time.Millisecond

// State handles the consensus process.
type State struct {
	*types.RoundState

	proposer      *bls.SecretKey
	storage       storageInterface
	roundLimit    time.Duration
	roundInterval time.Duration

	chCommit <-chan *networktypes.CommitBlockRequest
	chStop   chan struct{}
}

// NewState returns a new State.
func NewState(cfg *Config, storage storageInterface) *State {
	privKey, err := utils.HexToBlsPrivKey(cfg.ProposerPrivateKey)
	if err != nil {
		logger.Fatalf("failed to parse the proposer private key: %v", err)
	}

	chStop := make(chan struct{}, 1)

	return &State{
		proposer:      privKey,
		storage:       storage,
		roundLimit:    time.Duration(cfg.RoundLimit),
		roundInterval: time.Duration(cfg.RoundInterval),
		chCommit:      make(<-chan *networktypes.CommitBlockRequest, 1000),
		chStop:        chStop,
		RoundState:    types.NewEmptyRoundState(),
	}
}

// OnStart loads the first unverified block and starts the round.
func (s *State) OnStart() {
	lastBlockNumber, err := s.storage.GetLastFinalizedBlockNumber(context.Background())
	if err != nil {
		logger.Errorf("failed to get the last finalized block number: %v", err)
		return
	}

	for {
		// check if chStop is triggered
		select {
		case <-s.chStop:
			return
		default:
		}

		logger.Infof("start the round for the block number %v", lastBlockNumber+1)
		if err := s.startRound(lastBlockNumber); err != nil {
			logger.Errorf("failed to start the round: %v", err)
			time.Sleep(s.roundInterval)
			continue
		}

		logger.Infof("the proposal block %v is ready", s.ProposalBlock)
		logger.Infof("the validator set %v is set", s.Validators)

		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(s.roundLimit))
		defer cancel()
		isVoted, err := s.processRound(ctx)
		if err != nil {
			// TODO handle timeout error, restart the round
			logger.Errorf("failed to process the round: %v", err)
			continue
		}
		if !isVoted {
			logger.Errorf("the current block %d is not finalized", s.ProposalBlock.BlockNumber())
		}

		// store the finalized block
		if err := s.storage.UpdateBlock(context.Background(), s.ProposalBlock); err != nil {
			logger.Errorf("failed to update the block %v: %v", s.ProposalBlock, err)
			continue
		}

		logger.Infof("the block %d is finalized", s.ProposalBlock.BlockNumber())

		lastBlockNumber++
	}
}

// OnStop stops the consensus process.
func (s *State) OnStop() {
	logger.Infof("OnStop() called")
	s.chStop <- struct{}{}
}

// startRound loads the next block and initializes the round state.
func (s *State) startRound(blockNumber uint64) error {
	nodes, err := s.storage.GetNodesByStatuses(context.Background(), []sequencertypes.NodeStatus{sequencertypes.NodeRegistered})
	if err != nil {
		return err
	}

	// TODO how to determince nodes status?
	if len(nodes) == 0 {
		return fmt.Errorf("no nodes are registered")
	}

	pubkey := utils.BlsPubKeyToHex(s.proposer.GetPublicKey())
	validatorSet := types.NewValidatorSet(&types.Validator{PublicKey: pubkey}, nodes)

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(s.roundLimit))
	defer cancel()
	block, err := s.getNextBlock(ctx, blockNumber)
	if err != nil {
		// TODO handle timeout error
		return fmt.Errorf("getting the block %d is failed: %v", blockNumber, err)
	}
	// generate a proposer signature
	signature, err := s.proposer.Sign(common.FromHex(block.Header.BlockHash))
	if err != nil {
		return err
	}
	block.Header.ProposerSignature = utils.BlsSignatureToHex(signature)
	block.Header.ProposerPubKey = pubkey

	s.UpdateRoundState(validatorSet, block)

	return nil
}

// getNextBlock returns the next block from the storage.
func (s *State) getNextBlock(ctx context.Context, blockNumber uint64) (*sequencertypes.Block, error) {
	block, err := s.storage.GetBlock(ctx, blockNumber+1)
	if err == nil || err != storetypes.ErrBlockNotFound {
		return block, err
	}
	// in case the block is not found, wait for it to be added from the sequencer
	ticker := time.NewTicker(CheckInterval)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-ticker.C:
			block, err := s.storage.GetBlock(context.Background(), blockNumber+1)
			if err != nil {
				if err == storetypes.ErrBlockNotFound {
					continue
				}
				return nil, err
			}

			return block, nil
		}
	}
}

// processRound processes the round.
func (s *State) processRound(ctx context.Context) (bool, error) {
	isAfterRoundInterval := false
	isBlocked := false

	checkCommit := func() (bool, error) {
		if s.CheckEnoughVotingPower() {
			pubkeys, aggSignature, err := s.CheckAggregatedSignature()
			if err != nil {
				// TODO handle error
				return true, err
			}
			isBlocked = true
			s.BlockCommit()

			pubKeys := make([]string, 0)
			for _, pubkey := range pubkeys {
				pubKeys = append(pubKeys, utils.BlsPubKeyToHex(pubkey))
			}
			s.UpdateAggregatedSignature(pubKeys, utils.BlsSignatureToHex(aggSignature))
			return true, nil
		} else if isBlocked {
			isBlocked = false
			s.UnblockCommit()
		}

		logger.Warnf("the current block %d doesn't get enough power", s.ProposalBlock.BlockNumber())
		return false, nil
	}

	timer := time.NewTimer(s.roundInterval)
	defer timer.Stop()
	ticker := time.NewTicker(CheckInterval)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return false, ctx.Err()
		case <-timer.C:
			logger.Infof("check the commit after the round interval")
			isAfterRoundInterval = true
			isFinalized, err := checkCommit()
			if err != nil || isFinalized {
				return isFinalized, err
			}
			logger.Warnf("the current block %d is not finalized for the round interval", s.ProposalBlock.BlockNumber())
		case <-ticker.C:
			if isAfterRoundInterval {
				isFinalized, err := checkCommit()
				if err != nil || isFinalized {
					return isFinalized, err
				}
			}
		}
	}
}

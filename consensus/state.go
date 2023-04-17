package consensus

import (
	"context"
	"time"

	"github.com/Lagrange-Labs/lagrange-node/consensus/types"
	"github.com/Lagrange-Labs/lagrange-node/logger"
	networktypes "github.com/Lagrange-Labs/lagrange-node/network/types"
	sequencertypes "github.com/Lagrange-Labs/lagrange-node/sequencer/types"
	storetypes "github.com/Lagrange-Labs/lagrange-node/store/types"
	"github.com/Lagrange-Labs/lagrange-node/utils"
)

const CheckInterval = 500 * time.Millisecond

// State handles the consensus process.
type State struct {
	types.RoundState

	proposer      *types.Validator
	storage       storageInterface
	roundLimit    time.Duration
	roundInterval time.Duration

	chCommit <-chan *networktypes.CommitBlockRequest
	chStop   chan struct{}
}

// NewState returns a new State.
func NewState(cfg *Config, storage storageInterface, chCommit chan *networktypes.CommitBlockRequest) *State {
	chStop := make(chan struct{}, 1)

	return &State{
		proposer: &types.Validator{
			PublicKey: cfg.ProposerPubKey,
		},
		storage:       storage,
		roundLimit:    time.Duration(cfg.RoundLimit),
		roundInterval: time.Duration(cfg.RoundInterval),
		chCommit:      chCommit,
		chStop:        chStop,
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
		if err := s.startRound(lastBlockNumber); err != nil {
			logger.Errorf("failed to start the round: %v", err)
			return
		}

		chBlocked := make(chan bool)
		chDone := make(chan struct{})

		// It starts the receiveRoutine to receive the commit from the gRPC server.
		go s.receiveRoutine(chBlocked, chDone)

		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(s.roundLimit))
		defer cancel()
		isVoted, err := s.processRound(ctx, chBlocked)
		if err != nil {
			// TODO handle timeout error, restart the round
			logger.Errorf("failed to process the round: %v", err)
			return
		}
		if !isVoted {
			logger.Errorf("the current block %v is not finalized", s.ProposalBlock)
		}

		// close the receiveRoutine
		chDone <- struct{}{}

		lastBlockNumber++

		// check if chStop is triggered
		select {
		case <-s.chStop:
			return
		default:
		}
	}
}

// OnStop stops the consensus process.
func (s *State) OnStop() {
	logger.Infof("OnStop() called")
	s.chStop <- struct{}{}
}

// receiveRoutine receives the commit from the gRPC server.
func (s *State) receiveRoutine(chBlocked chan bool, chDone chan struct{}) {
	isBlocked := false

	for {
		select {
		case commit := <-s.chCommit:
			if isBlocked {
				continue
			}
			s.AddCommit(commit)
		case <-chDone:
			return
		case blocked := <-chBlocked:
			isBlocked = blocked
		}
	}
}

// startRound loads the next block and initializes the round state.
func (s *State) startRound(blockNumber uint64) error {
	nodes, err := s.storage.GetNodesByStatuses(context.Background(), []sequencertypes.NodeStatus{sequencertypes.NodeRegistered})
	if err != nil {
		return err
	}
	validatorSet := types.NewValidatorSet(s.proposer, nodes)

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(s.roundLimit))
	defer cancel()
	block, err := s.getNextBlock(ctx, blockNumber)
	if err != nil {
		// TODO handle timeout error
		return err
	}

	s.RoundState = *types.NewRoundState(validatorSet, block)

	return nil
}

// getNextBlock returns the next block from the storage.
func (s *State) getNextBlock(ctx context.Context, blockNumber uint64) (*sequencertypes.Block, error) {
	block, err := s.storage.GetBlock(ctx, blockNumber+1)
	if err == nil || err != storetypes.ErrBlockNotFound {
		return block, err
	}
	// in case the block is not found, wait for it to be added from the sequencer
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-time.After(CheckInterval):
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
func (s *State) processRound(ctx context.Context, chBlocked chan bool) (bool, error) {
	isAfterRoundInterval := false
	isBlocked := false

	checkCommit := func() (bool, error) {
		if s.CheckEnoughVotingPower() {
			pubkeys, aggSignature, err := s.CheckAggregatedSignature()
			if err != nil {
				// TODO handle error
				return true, err
			}
			chBlocked <- true
			isBlocked = true
			s.ProposalBlock.AggSignature = utils.BlsSignatureToHex(aggSignature)
			for _, pubkey := range pubkeys {
				s.ProposalBlock.PubKeys = append(s.ProposalBlock.PubKeys, utils.BlsPubKeyToHex(pubkey))
			}
			return true, nil
		} else if isBlocked {
			chBlocked <- false
			isBlocked = false
		}

		return false, nil
	}

	t := time.NewTimer(s.roundInterval)

	for {
		select {
		case <-ctx.Done():
			return false, ctx.Err()
		case <-t.C:
			t.Stop()
			isAfterRoundInterval = true
			isFinalized, err := checkCommit()
			if err != nil || isFinalized {
				return isFinalized, err
			}
		case <-time.After(CheckInterval):
			if isAfterRoundInterval {
				isFinalized, err := checkCommit()
				if err != nil || isFinalized {
					return isFinalized, err
				}
			}
		}
	}
}

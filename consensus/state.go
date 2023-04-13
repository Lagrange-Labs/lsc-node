package consensus

import (
	"context"
	"time"

	"github.com/Lagrange-Labs/lagrange-node/consensus/types"
	sequencertypes "github.com/Lagrange-Labs/lagrange-node/sequencer/types"
	storetypes "github.com/Lagrange-Labs/lagrange-node/store/types"
)

const CheckInterval = 500 * time.Millisecond

// State handles the consensus process.
type State struct {
	types.RoundState

	proposer      *types.Validator
	storage       storageInterface
	roundLimit    time.Duration
	roundInterval time.Duration
}

// NewState returns a new State.
func NewState(cfg *Config, storage storageInterface) *State {
	return &State{
		proposer: &types.Validator{
			PublicKey: cfg.ProposerPubKey,
		},
		storage:    storage,
		roundLimit: time.Duration(cfg.RoundLimit),
	}
}

// OnStart loads the first unverified block and starts the round.
func (s *State) OnStart() error {
	lastBlockNumber, err := s.storage.GetLastFinalizedBlockNumber(context.Background())
	if err != nil {
		return err
	}

	for {
		if err := s.startRound(lastBlockNumber); err != nil {
			return err
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(s.roundLimit))
		defer cancel()
		if err := s.processRound(ctx); err != nil {
			// TODO handle timeout error
			return err
		}

		lastBlockNumber++
	}
}

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

func (s *State) getNextBlock(ctx context.Context, blockNumber uint64) (*sequencertypes.Block, error) {
	block, err := s.storage.GetBlock(context.Background(), blockNumber+1)
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

func (s *State) processRound(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(s.roundInterval):

		}
	}
}

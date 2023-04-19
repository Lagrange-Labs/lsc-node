package sequencer

import (
	"context"
)

// Sequencer is the main component of the lagrange node.
// - It is responsible for generating new proofs.
// - It is responsible for stacking and slashing.
type Sequencer struct {
	storage storageInterface

	blockNumber uint64
}

// NewSequencer creates a new sequencer instance.
func NewSequencer(cfg *Config, storage storageInterface) (*Sequencer, error) {
	blockNumber, err := storage.GetLastBlockNumber(context.Background())
	if err != nil {
		return nil, err
	}

	return &Sequencer{
		storage:     storage,
		blockNumber: blockNumber,
	}, nil
}

// Start starts the sequencer.
func (s *Sequencer) Start() error {
	for {
		// TODO generate new proof
		break
	}
	return nil
}

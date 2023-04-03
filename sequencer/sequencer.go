package sequencer

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/Lagrange-Labs/Lagrange-Node/sequencer/nodestaking"
)

// Sequencer is the main component of the lagrange node.
// - It is responsible for generating new proofs.
// - It is responsible for stacking and slashing.
type Sequencer struct {
	stackingSC      *nodestaking.Nodestaking
	storage         storageInterface
	stakingInterval uint32
	blockNumber     uint64
}

// NewSequencer creates a new sequencer instance.
func NewSequencer(cfg *Config, storage storageInterface) (*Sequencer, error) {
	client, err := ethclient.Dial(cfg.EthereumURL)
	if err != nil {
		return nil, err
	}
	stakingSC, err := nodestaking.NewNodestaking(common.HexToAddress(cfg.StakingSCAddress), client)
	if err != nil {
		return nil, err
	}
	blockNumber, err := storage.GetLastBlockNumber(context.Background())
	if err != nil {
		return nil, err
	}

	return &Sequencer{
		stackingSC:      stakingSC,
		storage:         storage,
		stakingInterval: cfg.StackingCheckInterval,
		blockNumber:     blockNumber,
	}, nil
}

// Start starts the sequencer.
func (s *Sequencer) Start() error {
	for {
		// Begin Block
		// stacking status check
		if err := s.checkStacking(); err != nil {
			return err
		}
		// TODO generate new proof

		// End Block
		// TODO slashing
	}
}

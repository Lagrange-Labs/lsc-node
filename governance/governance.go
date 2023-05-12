package governance

import (
	"context"
	"time"

	"github.com/Lagrange-Labs/lagrange-node/scinterface/lagrange"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	SyncInterval = 10 * time.Second
)

// Governance is the module which is responsible for the staking and slashing.
type Governance struct {
	lagrangeSC      *lagrange.Lagrange
	stakingInterval uint32

	ctx    context.Context
	cancel context.CancelFunc
}

// NewGovernance creates a new Governance instance.
func NewGovernance(cfg *Config) (*Governance, error) {
	client, err := ethclient.Dial(cfg.EthereumURL)
	if err != nil {
		return nil, err
	}
	lagrangeSC, err := lagrange.NewLagrange(common.HexToAddress(cfg.StakingSCAddress), client)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithCancel(context.Background())
	return &Governance{
		lagrangeSC:      lagrangeSC,
		stakingInterval: cfg.StakingCheckInterval,
		ctx:             ctx,
		cancel:          cancel,
	}, nil
}

// Start starts the governance process.
func (g *Governance) Start() error {
	for {
		select {
		case <-g.ctx.Done():
			return nil
		case <-time.After(SyncInterval):

		}
	}
}

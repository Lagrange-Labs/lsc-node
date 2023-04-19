package governance

import (
	"github.com/Lagrange-Labs/lagrange-node/governance/nodestaking"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Governance is the module which is responsible for the staking and slashing.
type Governance struct {
	stackingSC      *nodestaking.Nodestaking
	stakingInterval uint32
}

// NewGovernance creates a new Governance instance.
func NewGovernance(cfg *Config) (*Governance, error) {
	client, err := ethclient.Dial(cfg.EthereumURL)
	if err != nil {
		return nil, err
	}
	stakingSC, err := nodestaking.NewNodestaking(common.HexToAddress(cfg.StakingSCAddress), client)
	if err != nil {
		return nil, err
	}
	return &Governance{
		stackingSC:      stakingSC,
		stakingInterval: cfg.StakingCheckInterval,
	}, nil
}

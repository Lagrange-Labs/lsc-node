package sequencer

import "github.com/Lagrange-Labs/lsc-node/core"

type Config struct {
	// L2 Chain name
	Chain string `mapstructure:"Chain"`
	// FromL1BlockNumber is the starting L1 block number.
	FromL1BlockNumber uint64 `mapstructure:"FromL1BlockNumber"`
	// FromL2BlockNumber is the L2 block number to start fetching from.
	FromL2BlockNumber uint64 `mapstructure:"FromL2BlockNumber"`
	// EthereumURL is the endpoint of the ethereum node.
	EthereumURL string `mapstructure:"EthereumURL"`
	// CommitteeSCAddress is the address of the committee smart contract.
	CommitteeSCAddress string `mapstructure:"CommitteeSCAddress"`
	// EigenDMSCAddress is the address of the Eigen DelegationManager smart contract.
	EigenDMSCAddress string `mapstructure:"EigenDMSCAddress"`
	// StakingCheckInterval is the interval to check the staking status.
	StakingCheckInterval core.TimeDuration `mapstructure:"StakingCheckInterval"`
}

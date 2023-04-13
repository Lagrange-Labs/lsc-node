package governance

// Config is the configuration for the Governance module.
type Config struct {
	// EthereumURL is the endpoint of the ethereum node.
	EthereumURL string `mapstructure:"EthereumURL"`
	// StakingSCAddress is the address of the staking smart contract.
	StakingSCAddress string `mapstructure:"StakingSCAddress"`
	// StakingCheckInterval is the interval to check the staking status.
	StakingCheckInterval uint32 `mapstructure:"StakingCheckInterval"`
}

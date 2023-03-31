package sequencer

type Config struct {
	// EthereumURL is the endpoint of the ethereum node.
	EthereumURL string `mapstructure:"EthereumURL"`
	// StakingSCAddress is the address of the staking smart contract.
	StakingSCAddress string `mapstructure:"StakingSCAddress"`
	// StackingCheckInterval is the interval to check the stacking status.
	StackingCheckInterval uint32 `mapstructure:"StackingCheckInterval"`
}

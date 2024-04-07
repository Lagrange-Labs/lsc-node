package optimism

// Config is the configuration for the Optimism client.
type Config struct {
	// RPCURL is the URL of the Optimism RPC node
	RPCURL string `mapstructure:"RPCURL"`
	// L1RPCURL is the URL of the L1 Ethereum RPC node
	L1RPCURL string `mapstructure:"L1RPCURL"`
	// BeaconURL is the URL of the Beacon RPC node
	BeaconURL string `mapstructure:"BeaconURL"`
	// BatchInbox is the address of the BatchInbox EOA
	BatchInbox string `mapstructure:"BatchInbox"`
	// BatchSender is the address of the Batcher
	BatchSender string `mapstructure:"BatchSender"`
}

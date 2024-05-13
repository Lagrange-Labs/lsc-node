package arbitrum

// Config is the configuration for the Arbitrum client.
type Config struct {
	// RPCURL is the URL of the Arbitrum RPC node
	RPCURL string `mapstructure:"RPCURL"`
	// L1RPCURL is the URL of the L1 Ethereum RPC node
	L1RPCURL string `mapstructure:"L1RPCURL"`
	// BeaconURL is the URL of the Beacon RPC node
	BeaconURL string `mapstructure:"BeaconURL"`
	// BatchInbox is the address of the BatchInbox EOA
	BatchInbox string `mapstructure:"BatchInbox"`
	// ConcurrentFetchers is the number of concurrent fetchers
	ConcurrentFetchers int `mapstructure:"ConcurrentFetchers"`
}

package optimism

// Config is the configuration for the Optimism client.
type Config struct {
	// RPCURLs is the URL list of the Optimism RPC node
	RPCURLs []string `mapstructure:"RPCURL"`
	// L1RPCURLs is the URL list of the L1 Ethereum RPC node
	L1RPCURLs []string `mapstructure:"L1RPCURL"`
	// BeaconURL is the URL of the Beacon RPC node
	BeaconURL string `mapstructure:"BeaconURL"`
	// BatchInbox is the address of the BatchInbox EOA
	BatchInbox string `mapstructure:"BatchInbox"`
	// BatchSender is the address of the Batcher
	BatchSender string `mapstructure:"BatchSender"`
	// ConcurrentFetchers is the number of concurrent fetchers
	ConcurrentFetchers int `mapstructure:"ConcurrentFetchers"`
	// L1ParallelBlocks is the number of blocks to fetch in parallel from L1
	L1ParallelBlocks int `mapstructure:"L1ParallelBlocks"`
	// L2ParallelBlocks is the number of blocks to fetch in parallel from L2
	L2ParallelBlocks int `mapstructure:"L2ParallelBlocks"`
}

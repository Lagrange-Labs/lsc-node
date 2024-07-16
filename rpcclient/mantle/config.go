package mantle

// Config is the configuration for the Mantle client.
type Config struct {
	// RPCURLs is the URL list of the Mantle RPC node
	RPCURLs []string `mapstructure:"RPCURL"`
	// L1RPCURLs is the URL list of the L1 Ethereum RPC node
	L1RPCURLs []string `mapstructure:"L1RPCURL"`
	// BatchStorageAddr is the address of the L1BatchStorage contract
	BatchStorageAddr string `mapstructure:"BatchStorageAddr"`
}

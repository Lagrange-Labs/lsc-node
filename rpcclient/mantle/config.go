package mantle

// Config is the configuration for the Mantle client.
type Config struct {
	// RPCURL is the URL of the Mantle RPC node
	RPCURL string `mapstructure:"RPCURL"`
	// L1RPCURL is the URL of the L1 Ethereum RPC node
	L1RPCURL string `mapstructure:"L1RPCURL"`
	// BatchStorageAddr is the address of the L1BatchStorage contract
	BatchStorageAddr string `mapstructure:"BatchStorageAddr"`
}

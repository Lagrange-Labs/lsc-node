package arbitrum

// Config is the configuration for the Arbitrum client.
type Config struct {
	// RPCURL is the URL of the Arbitrum RPC node
	RPCURL string `mapstructure:"RPCURL"`
}

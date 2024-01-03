package mock

// Config is the configuration for the Mock client.
type Config struct {
	// RPCURL is the URL of the Mock RPC node
	RPCURL string `mapstructure:"RPCURL"`
}

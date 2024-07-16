package mock

// Config is the configuration for the Mock client.
type Config struct {
	// RPCURLs is the URL list of the Mock RPC node
	RPCURLs []string `mapstructure:"RPCURL"`
}

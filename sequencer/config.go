package sequencer

type Config struct {
	Chain           string `mapstructure:"Chain"`
	RPCURL          string `mapstructure:"RPCURL"`
	EthURL          string `mapstructure:"EthURL"`
	NewRPCURL       string `mapstructure:"NewRPCURL"`
	FromBlockNumber uint64 `mapstructure:"FromBlockNumber"`
}

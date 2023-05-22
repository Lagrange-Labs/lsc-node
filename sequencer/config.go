package sequencer

type Config struct {
	Chain           string `mapstructure:"Chain"`
	RPCURL          string `mapstructure:"RPCURL"`
	FromBlockNumber uint64 `mapstructure:"FromBlockNumber"`
}

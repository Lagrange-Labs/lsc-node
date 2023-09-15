package sequencer

type Config struct {
	Chain           string `mapstructure:"Chain"`
	RPCURL          string `mapstructure:"RPCURL"`
	EthURL          string `mapstructure:"EthURL"`
	BatchStorage    string `mapstructure:"BatchStorage"`
	FromBlockNumber uint64 `mapstructure:"FromBlockNumber"`
}

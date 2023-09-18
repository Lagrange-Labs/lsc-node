package sequencer

type Config struct {
	Chain            string `mapstructure:"Chain"`
	RPCURL           string `mapstructure:"RPCURL"`
	EthURL           string `mapstructure:"EthURL"`
	BatchStorageAddr string `mapstructure:"BatchStorageAddr"`
	FromBlockNumber  uint64 `mapstructure:"FromBlockNumber"`
}

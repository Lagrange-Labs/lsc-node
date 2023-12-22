package sequencer

type Config struct {
	Chain           string `mapstructure:"Chain"`
	FromBlockNumber uint64 `mapstructure:"FromBlockNumber"`
}

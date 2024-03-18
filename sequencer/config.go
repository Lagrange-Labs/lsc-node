package sequencer

type Config struct {
	Chain             string `mapstructure:"Chain"`
	FromL1BlockNumber uint64 `mapstructure:"FromL1BlockNumber"`
	FromL2BlockNumber uint64 `mapstructure:"FromL2BlockNumber"`
}

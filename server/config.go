package server

// Config is the configuration for the sequencer server.
type Config struct {
	// GRPCPort is TCP port to listen by gRPC server
	GRPCPort string `mapstructure:"GRPCPort"`
}

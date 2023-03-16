package network

import (
	"github.com/Lagrange-Labs/Lagrange-Node/utils"
)

// ServerConfig is the configuration for the sequencer server.
type ServerConfig struct {
	// GRPCPort is TCP port to listen by gRPC server
	GRPCPort string `mapstructure:"GRPCPort"`
	// PublicKey is the public key of the sequencer
	PublicKey string `mapstructure:"PublicKey"`
}

// ClientConfig is the configuration for the client node.
type ClientConfig struct {
	// GrpcURL is the URL of the gRPC server
	GrpcURL string `mapstructure:"GrpcURL"`
	// PrivateKey is the private key of the client node
	PrivateKey string `mapstructure:"PrivateKey"`
	// PullInterval is the interval to pull the latest proof
	PullInterval utils.TimeDuration `mapstructure:"PullInterval"`
}

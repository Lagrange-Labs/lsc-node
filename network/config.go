package network

import (
	"github.com/Lagrange-Labs/lagrange-node/utils"
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
	// Chain is the chain name of the blockchain
	Chain string `mapstructure:"Chain"`
	// RPCEndpoint is the endpoint of the blockchain node
	RPCEndpoint string `mapstructure:"RPCEndpoint"`
	// BLSPrivateKey is the private key of the client node
	BLSPrivateKey string `mapstructure:"BLSPrivateKey"`
	// ECDSAPrivateKey is the private key of the client node
	ECDSAPrivateKey string `mapstructure:"ECDSAPrivateKey"`
	// StakeAddress is the ethereum address of the staking
	StakeAddress string `mapstructure:"StakeAddress"`
	// PullInterval is the interval to pull the latest proof
	PullInterval utils.TimeDuration `mapstructure:"PullInterval"`
}

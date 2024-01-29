package network

import (
	"github.com/Lagrange-Labs/lagrange-node/utils"
)

// ServerConfig is the configuration for the sequencer server.
type ServerConfig struct {
	// GRPCPort is TCP port to listen by gRPC server
	GRPCPort string `mapstructure:"GRPCPort"`
}

// ClientConfig is the configuration for the client node.
type ClientConfig struct {
	// GrpcURL is the URL of the gRPC server
	GrpcURL string `mapstructure:"GrpcURL"`
	// Chain is the chain name of the blockchain
	Chain string `mapstructure:"Chain"`
	// EthereumURL is the endpoint of the ethereum node
	EthereumURL string `mapstructure:"EthereumURL"`
	// CommitteeSCAddress is the address of the committee smart contract
	CommitteeSCAddress string `mapstructure:"CommitteeSCAddress"`
	// BLSPrivateKey is the private key of the client node
	BLSPrivateKey string `mapstructure:"BLSPrivateKey"`
	// ECDSAPrivateKey is the private key of the client node
	ECDSAPrivateKey string `mapstructure:"ECDSAPrivateKey"`
	// PullInterval is the interval to pull the latest proof
	PullInterval utils.TimeDuration `mapstructure:"PullInterval"`
	// BLSCurve is the curve used for BLS signature
	BLSCurve string `mapstructure:"BLSCurve"`
}

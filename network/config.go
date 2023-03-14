package network

import "time"

// ServerConfig is the configuration for the sequencer server.
type ServerConfig struct {
	// GRPCPort is TCP port to listen by gRPC server
	GRPCPort string
	// PublicKey is the public key of the sequencer
	PublicKey string
}

// ClientConfig is the configuration for the client node.
type ClientConfig struct {
	// GrpcURL is the URL of the gRPC server
	GrpcURL string
	// PrivateKey is the private key of the client node
	PrivateKey string
	// PullInterval is the interval to pull the latest proof
	PullInterval time.Duration
}

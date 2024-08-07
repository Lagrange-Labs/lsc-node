package client

import "github.com/Lagrange-Labs/lagrange-node/utils"

// Config is the configuration for the client node.
type Config struct {
	// GrpcURLs is the list of gRPC URLs
	GrpcURLs []string `mapstructure:"GrpcURLs"`
	// Chain is the chain name of the blockchain
	Chain string `mapstructure:"Chain"`
	// OperatorAddress is the address of the operator
	OperatorAddress string `mapstructure:"OperatorAddress"`
	// EthereumURL is the endpoint of the ethereum node
	EthereumURL string `mapstructure:"EthereumURL"`
	// CommitteeSCAddress is the address of the committee smart contract
	CommitteeSCAddress string `mapstructure:"CommitteeSCAddress"`
	// PullInterval is the interval to pull the latest proof
	PullInterval utils.TimeDuration `mapstructure:"PullInterval"`
	// BLSCurve is the curve used for BLS signature
	BLSCurve string `mapstructure:"BLSCurve"`
	// BLSKeyAccountID is the account ID of the BLS key
	BLSKeyAccountID string `mapstructure:"BLSKeyAccountID"`
	// SignerKeyAccountID is the account ID of the signer ECDSA key
	SignerKeyAccountID string `mapstructure:"SignerKeyAccountID"`
	// SignerServerURL is the URL of the signer server
	SignerServerURL string `mapstructure:"SignerServerURL"`
}

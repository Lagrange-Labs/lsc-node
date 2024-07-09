package consensus

import "github.com/Lagrange-Labs/lagrange-node/utils"

// Config is the configuration for the consensus module.
type Config struct {
	// ProposerBLSKeystorePath is the path of the BLS keystore file
	ProposerBLSKeystorePath string `mapstructure:"ProposerBLSKeystorePath"`
	// ProposerBLSKeystorePassword is the password of the BLS keystore file
	ProposerBLSKeystorePassword string `mapstructure:"ProposerBLSKeystorePassword"`
	// ProposerBLSKeystorePasswordPath is the path of the password file of the BLS keystore file
	ProposerBLSKeystorePasswordPath string `mapstructure:"ProposerBLSKeystorePasswordPath"`
	// RoundLimit is the maximum time to wait for the block finalization.
	RoundLimit utils.TimeDuration `mapstructure:"RoundLimit"`
	// RoundInterval is the interval to wait for the next round.
	RoundInterval utils.TimeDuration `mapstructure:"RoundInterval"`
	// BLSCurve is the curve used for BLS signature
	BLSCurve string `mapstructure:"BLSCurve"`
}

// ChainInfo is the information of the chain.
type ChainInfo struct {
	// ChainID is the ID of the chain.
	ChainID uint32 `mapstructure:"ChainID"`
	// EthereumURL is the endpoint of the ethereum node.
	EthereumURL string `mapstructure:"EthereumURL"`
	// CommitteeSCAddress is the address of the committee smart contract.
	CommitteeSCAddress string `mapstructure:"CommitteeSCAddress"`
}

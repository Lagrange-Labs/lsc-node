package consensus

import "github.com/Lagrange-Labs/lagrange-node/utils"

// Config is the configuration for the consensus module.
type Config struct {
	// OperatorAddress is the address of the operator.
	OperatorAddress string `mapstructure:"OperatorAddress"`
	// BatchSize is the number of blocks in a batch.
	BatchSize uint32 `mapstructure:"BatchSize"`
	// ProposerPrivateKey is the BLS private key of the proposer node.
	ProposerPrivateKey string `mapstructure:"ProposerPrivateKey"`
	// RoundLimit is the maximum time to wait for the block finalization.
	RoundLimit utils.TimeDuration `mapstructure:"RoundLimit"`
	// RoundInterval is the interval to wait for the next round.
	RoundInterval utils.TimeDuration `mapstructure:"RoundInterval"`
	// BLSCurve is the curve used for BLS signature
	BLSCurve string `mapstructure:"BLSCurve"`
}

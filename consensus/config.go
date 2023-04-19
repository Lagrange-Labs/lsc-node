package consensus

import "github.com/Lagrange-Labs/lagrange-node/utils"

// Config is the configuration for the consensus module.
type Config struct {
	// ProposerPubKey is the public key of the proposer node.
	ProposerPubKey string `mapstructure:"ProposerPubKey"`
	// RoundLimit is the maximum time to wait for the block finalization.
	RoundLimit utils.TimeDuration `mapstructure:"RoundLimit"`
	// RoundInterval is the interval to wait for the next round.
	RoundInterval utils.TimeDuration `mapstructure:"RoundInterval"`
}

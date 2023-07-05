package types

import (
	"github.com/Lagrange-Labs/lagrange-node/utils"
)

// Config is the configuration for the Governance module.
type Config struct {
	// EthereumURL is the endpoint of the ethereum node.
	EthereumURL string `mapstructure:"EthereumURL"`
	// PrivateKey is the private key of the sequencer.
	PrivateKey string `mapstructure:"PrivateKey"`
	// StakingSCAddress is the address of the staking smart contract.
	StakingSCAddress string `mapstructure:"StakingSCAddress"`
	// CommitteeSCAddress is the address of the committee smart contract.
	CommitteeSCAddress string `mapstructure:"CommitteeSCAddress"`
	// StakingCheckInterval is the interval to check the staking status.
	StakingCheckInterval utils.TimeDuration `mapstructure:"StakingCheckInterval"`
	// EvidenceUploadInterval is the interval to upload the evidence.
	EvidenceUploadInterval utils.TimeDuration `mapstructure:"EvidenceUploadInterval"`
}

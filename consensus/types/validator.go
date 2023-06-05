package types

import (
	"github.com/Lagrange-Labs/lagrange-node/logger"
	networktypes "github.com/Lagrange-Labs/lagrange-node/network/types"
)

// Validator defines a validator state.
type Validator struct {
	PublicKey   string
	VotingPower uint64
}

// ValidatorSet defines a set of validators.
type ValidatorSet struct {
	Validators []*Validator
	Proposer   *Validator

	votingPowerMap map[string]uint64
}

// NewValidatorSet creates a new validator set.
func NewValidatorSet(proposer *Validator, nodes []networktypes.ClientNode) *ValidatorSet {
	validators := make([]*Validator, len(nodes))
	votingPowerMap := make(map[string]uint64)

	for i, node := range nodes {
		validators[i] = &Validator{
			PublicKey:   node.PublicKey,
			VotingPower: node.VotingPower,
		}
		logger.Infof("validator: %s, voting power: %d", node.PublicKey, node.VotingPower)
		votingPowerMap[node.PublicKey] = node.VotingPower
	}

	return &ValidatorSet{
		Validators:     validators,
		Proposer:       proposer,
		votingPowerMap: votingPowerMap,
	}
}

// GetVotingPower returns the voting power of a validator.
func (vs *ValidatorSet) GetVotingPower(pubKey string) uint64 {
	return vs.votingPowerMap[pubKey]
}

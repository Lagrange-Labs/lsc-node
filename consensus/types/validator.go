package types

import (
	networktypes "github.com/Lagrange-Labs/lagrange-node/network/types"
)

// Validator defines a validator state.
type Validator struct {
	PublicKey   string
	VotingPower uint64
}

// ValidatorSet defines a set of validators.
type ValidatorSet struct {
	validators []*Validator

	votingPowerMap   map[string]uint64
	totalVotingPower uint64
}

// NewValidatorSet creates a new validator set.
func NewValidatorSet(nodes []networktypes.ClientNode) *ValidatorSet {
	validators := make([]*Validator, len(nodes))
	votingPowerMap := make(map[string]uint64)

	for i, node := range nodes {
		validators[i] = &Validator{
			PublicKey:   node.PublicKey,
			VotingPower: node.VotingPower,
		}
		votingPowerMap[node.PublicKey] = node.VotingPower
	}

	return &ValidatorSet{
		validators:     validators,
		votingPowerMap: votingPowerMap,
	}
}

// GetVotingPower returns the voting power of a validator.
func (vs *ValidatorSet) GetVotingPower(pubKey string) uint64 {
	return vs.votingPowerMap[pubKey]
}

// GetTotalVotingPower returns the total voting power.
func (vs *ValidatorSet) GetTotalVotingPower() uint64 {
	return vs.totalVotingPower
}

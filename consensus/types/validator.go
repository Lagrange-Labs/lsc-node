package types

import (
	networktypes "github.com/Lagrange-Labs/lagrange-node/network/types"
)

// Validator defines a validator state.
type Validator struct {
	BlsPubKey    []byte
	StakeAddress string
	VotingPower  uint64
}

// ValidatorSet defines a set of validators.
type ValidatorSet struct {
	validators           map[string]*Validator
	totalVotingPower     uint64
	committeeVotingPower uint64
}

// NewValidatorSet creates a new validator set.
func NewValidatorSet(nodes []networktypes.ClientNode, committeeVotingPower uint64) *ValidatorSet {
	validators := make(map[string]*Validator)
	totalVotingPower := uint64(0)

	for _, node := range nodes {
		validators[node.StakeAddress] = &Validator{
			BlsPubKey:    node.PublicKey,
			StakeAddress: node.StakeAddress,
			VotingPower:  node.VotingPower,
		}
		totalVotingPower += node.VotingPower
	}

	return &ValidatorSet{
		validators:           validators,
		totalVotingPower:     totalVotingPower,
		committeeVotingPower: committeeVotingPower,
	}
}

// GetValidatorCount returns the number of validators.
func (vs *ValidatorSet) GetValidatorCount() int {
	return len(vs.validators)
}

// GetVotingPower returns the voting power of a validator.
func (vs *ValidatorSet) GetVotingPower(stakeAddr string) uint64 {
	return vs.validators[stakeAddr].VotingPower
}

// GetPublicKey returns the public key of a validator.
func (vs *ValidatorSet) GetPublicKey(stakeAddr string) []byte {
	return vs.validators[stakeAddr].BlsPubKey
}

// GetTotalVotingPower returns the total committee voting power.
func (vs *ValidatorSet) GetTotalVotingPower() uint64 {
	return vs.totalVotingPower
}

// GetCommitteeVotingPower returns the total committee voting power.
func (vs *ValidatorSet) GetCommitteeVotingPower() uint64 {
	return vs.committeeVotingPower
}

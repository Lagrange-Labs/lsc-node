package types

import (
	networktypes "github.com/Lagrange-Labs/lagrange-node/network/types"
)

// Validator defines a validator state.
type Validator struct {
	StakeAddress string
	VotingPower  uint64
	SignAddress  string
}

// ValidatorSet defines a set of validators.
type ValidatorSet struct {
	validators           map[string]map[string]*Validator
	totalVotingPower     uint64
	committeeVotingPower uint64
}

// NewValidatorSet creates a new validator set.
func NewValidatorSet(nodes []networktypes.ClientNode, committeeVotingPower uint64) *ValidatorSet {
	validators := make(map[string]map[string]*Validator)
	totalVotingPower := uint64(0)

	for _, node := range nodes {
		if _, ok := validators[node.StakeAddress]; !ok {
			validators[node.StakeAddress] = make(map[string]*Validator)
		}
		validators[node.StakeAddress][node.PublicKey] = &Validator{
			StakeAddress: node.StakeAddress,
			VotingPower:  node.VotingPower,
			SignAddress:  node.SignAddress,
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
func (vs *ValidatorSet) GetVotingPower(stakeAddr, pubKey string) uint64 {
	if ops, ok := vs.validators[stakeAddr]; !ok {
		return 0
	} else if vp, ok := ops[pubKey]; !ok {
		return 0
	} else {
		return vp.VotingPower
	}
}

// GetSignAddress returns the sign address of a validator.
func (vs *ValidatorSet) GetSignAddress(stakeAddr string) string {
	for _, node := range vs.validators[stakeAddr] {
		return node.SignAddress
	}
	return ""
}

// GetTotalVotingPower returns the total committee voting power.
func (vs *ValidatorSet) GetTotalVotingPower() uint64 {
	return vs.totalVotingPower
}

// GetCommitteeVotingPower returns the total committee voting power.
func (vs *ValidatorSet) GetCommitteeVotingPower() uint64 {
	return vs.committeeVotingPower
}

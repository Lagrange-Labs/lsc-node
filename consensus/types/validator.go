package types

// Validator defines a validator state
type Validator struct {
	PublicKey    string
	StakeAddress string
	VotingPower  int64
}

// ValidatorSet defines a set of validators
type ValidatorSet struct {
	Validators       []*Validator
	Proposer         *Validator
	TotalVotingPower int64
}

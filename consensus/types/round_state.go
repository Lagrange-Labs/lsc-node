package types

import (
	networktypes "github.com/Lagrange-Labs/lagrange-node/network/types"
	sequencertypes "github.com/Lagrange-Labs/lagrange-node/sequencer/types"
)

// RoundState defines the internal consensus state.
type RoundState struct {
	Height        uint64
	Validators    *ValidatorSet
	ProposalBlock *sequencertypes.Block

	CommitSignatures []*networktypes.CommitBlockRequest // to determine slashing
}

// NewRoundState creates a new round state.
func NewRoundState(validators *ValidatorSet, proposalBlock *sequencertypes.Block) *RoundState {
	return &RoundState{
		Height:        proposalBlock.Header.BlockNumber,
		Validators:    validators,
		ProposalBlock: proposalBlock,
	}
}

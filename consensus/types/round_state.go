package types

import (
	networktypes "github.com/Lagrange-Labs/lagrange-node/network/types"
	sequencertypes "github.com/Lagrange-Labs/lagrange-node/sequencer/types"
)

// RoundState defines the internal consensus state.
type RoundState struct {
	Height        int64
	Validators    *ValidatorSet
	ProposalBlock *sequencertypes.Block

	CommitSignatures []*networktypes.CommitBlockRequest // to determine slashing
}

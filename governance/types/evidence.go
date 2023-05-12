package types

import (
	"math/big"

	"github.com/Lagrange-Labs/lagrange-node/scinterface/lagrange"
	"github.com/ethereum/go-ethereum/common"
)

type Evidence struct {
	Operator                    common.Address
	BlockHash                   [32]byte
	CorrectBlockHash            [32]byte
	StateRoot                   [32]byte
	CorrectStateRoot            [32]byte
	Prooves                     [][32]byte
	CurrentCommitteeRoot        [32]byte
	CorrectCurrentCommitteeRoot [32]byte
	NextCommitteeRoot           [32]byte
	CorrectNextCommitteeRoot    [32]byte
	BlockNumber                 uint64
	EpochNumber                 uint64
	BlockSignature              [96]byte
	CommitSignature             [96]byte
	ChainID                     uint32
	Status                      bool
}

// GetSCEvidence converts the evidence to a smart contract input evidence.
func (e *Evidence) GetSCEvidence() *lagrange.LagrangeServiceEvidence {
	return &lagrange.LagrangeServiceEvidence{
		Operator:                    e.Operator,
		BlockHash:                   e.BlockHash,
		CorrectBlockHash:            e.CorrectBlockHash,
		StateRoot:                   e.StateRoot,
		CorrectStateRoot:            e.CorrectStateRoot,
		Prooves:                     e.Prooves,
		CurrentCommitteeRoot:        e.CurrentCommitteeRoot,
		CorrectCurrentCommitteeRoot: e.CorrectCurrentCommitteeRoot,
		NextCommitteeRoot:           e.NextCommitteeRoot,
		CorrectNextCommitteeRoot:    e.CorrectNextCommitteeRoot,
		BlockNumber:                 new(big.Int).SetInt64(int64(e.BlockNumber)),
		EpochNumber:                 new(big.Int).SetInt64(int64(e.EpochNumber)),
		BlockSignature:              e.BlockSignature[:],
		CommitSignature:             e.CommitSignature[:],
		ChainID:                     e.ChainID,
	}
}

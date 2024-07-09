package v2

import (
	servertypes "github.com/Lagrange-Labs/lagrange-node/server/types"
)

// CommitteeRoot is the root of the committee.
type CommitteeRoot struct {
	ChainID               uint32                    `json:"chain_id" bson:"chain_id"`
	CurrentCommitteeRoot  string                    `json:"current_committee_root" bson:"current_committee_root"`
	EpochNumber           uint64                    `json:"epoch_number" bson:"epoch_number"`
	EpochStartBlockNumber uint64                    `json:"epoch_start_block_number" bson:"epoch_start_block_number"`
	TotalVotingPower      uint64                    `json:"total_voting_power" bson:"total_voting_power"`
	Operators             []networktypes.ClientNode `json:"operators" bson:"operators"`
}

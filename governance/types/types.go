package types

// CommitteeRoot is the root of the committee.
type CommitteeRoot struct {
	ChainID              uint32 `json:"chain_id" bson:"chain_id"`
	CurrentCommitteeRoot string `json:"current_committee_root" bson:"current_committee_root"`
	NextCommitteeRoot    string `json:"next_committee_root" bson:"next_committee_root"`
	EpochBlockNumber     uint64 `json:"epoch_block_number" bson:"epoch_block_number"`
	TotalVotingPower     uint64 `json:"total_voting_power" bson:"total_voting_power"`
}

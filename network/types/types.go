package types

// NodeStatus is the status of a node.
type NodeStatus string

const (
	NodeStacked    = NodeStatus("stacked")
	NodeUnstacking = NodeStatus("unstacking")
	NodeUnstacked  = NodeStatus("unstacked")
	NodeSlashed    = NodeStatus("slashed")
	NodeJoined     = NodeStatus("joined")
	NodeRegistered = NodeStatus("registered")
)

// ClientNode is a struct to store the information of a node.
type ClientNode struct {
	// PublicKey is the bls public key of the node.
	PublicKey string `json:"public_key" bson:"public_key"`
	// IPAddress is the IP address of the client node.
	IPAddress string `json:"ip_address" bson:"ip_address"`
	// StakeAddress is the ethereum address of the staking.
	StakeAddress string `json:"stake_address" bson:"stake_address"`
	// VotingPower is the voting power of the node.
	VotingPower uint64 `json:"voting_power" bson:"voting_power"`
	// Status is the status of the node.
	Status NodeStatus `json:"status" bson:"status"`
}

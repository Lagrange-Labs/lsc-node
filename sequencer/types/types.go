package types

// NodeStatus is the status of a node.
type NodeStatus string

const (
	NodeStacked    = NodeStatus("stacked")
	NodeUnstacking = NodeStatus("unstacking")
	NodeUnstacked  = NodeStatus("unstacked")
	NodeSlashed    = NodeStatus("slashed")
	NodeRegistered = NodeStatus("registered")
)

// ClientNode is a struct to store the information of a node.
type ClientNode struct {
	// PublicKey is the bls public key of the node.
	PublicKey string
	// IPAddress is the IP address of the client node.
	IPAddress string
	// StakeAddress is the ethereum address of the staking.
	StakeAddress string
	// VotingPower is the voting power of the node.
	VotingPower uint64
	// Status is the status of the node.
	Status NodeStatus
}

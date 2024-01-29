package types

// NodeStatus is the status of a node.
type NodeStatus string

const (
	NodeStaked     = NodeStatus("staked")
	NodeUnstaking  = NodeStatus("unstaking")
	NodeUnstaked   = NodeStatus("unstaked")
	NodeSlashed    = NodeStatus("slashed")
	NodeJoined     = NodeStatus("joined")
	NodeRegistered = NodeStatus("registered")
)

// ClientNode is a struct to store the information of a node.
type ClientNode struct {
	// PublicKey is the bls public key of the node.
	PublicKey []byte `json:"public_key" bson:"public_key"`
	// IPAddress is the IP address of the client node.
	IPAddress string `json:"ip_address" bson:"ip_address"`
	// StakeAddress is the ethereum address of the staking.
	StakeAddress string `json:"stake_address" bson:"stake_address"`
	// VotingPower is the voting power of the node.
	VotingPower uint64 `json:"voting_power" bson:"voting_power"`
	// ChainID is the chain id of the node.
	ChainID uint32 `json:"chain_id" bson:"chain_id"`
	// Status is the status of the node.
	Status NodeStatus `json:"status" bson:"status"`
}

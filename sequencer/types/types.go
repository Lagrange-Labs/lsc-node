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

// BlockHash returns the block hash of the chain header.
func (b *Block) BlockHash() string {
	return b.ChainHeader.BlockHash
}

// BlockNumber returns the block number of the chain header.
func (b *Block) BlockNumber() uint64 {
	return b.ChainHeader.BlockNumber
}

// CurrentCommittee returns the current committee of the block.
func (b *Block) CurrentCommittee() string {
	return b.BlockHeader.CurrentCommittee
}

// NextCommittee returns the next committee of the block.
func (b *Block) NextCommittee() string {
	return b.BlockHeader.NextCommittee
}

// ProposerPubKey returns the proposer public key of the block.
func (b *Block) ProposerPubKey() string {
	return b.BlockHeader.ProposerPubKey
}

// ProposerSignature returns the proposer signature of the block.
func (b *Block) ProposerSignature() string {
	return b.BlockHeader.ProposerSignature
}

// BlsSignature returns the bls signature of the block.
func (b *Block) BlsSignature() *BlsSignature {
	return &BlsSignature{
		ChainHeader:      b.ChainHeader,
		CurrentCommittee: b.CurrentCommittee(),
		NextCommittee:    b.NextCommittee(),
	}
}

// BlockNumber returns the block number of the bls signature.
func (b *BlsSignature) BlockNumber() uint64 {
	return b.ChainHeader.BlockNumber
}

// Clone returns a clone of the bls signature.
func (b *BlsSignature) Clone() *BlsSignature {
	return &BlsSignature{
		ChainHeader:      b.ChainHeader,
		CurrentCommittee: b.CurrentCommittee,
		NextCommittee:    b.NextCommittee,
	}
}

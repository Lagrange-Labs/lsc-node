package types

import (
	"github.com/Lagrange-Labs/lagrange-node/utils"
	"github.com/ethereum/go-ethereum/common"
	"google.golang.org/protobuf/proto"
)

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

// Hash returns the hash of the node.
func (b *Block) Hash() string {
	return b.Header.BlockHash
}

// BlockNumber returns the block number.
func (b *Block) BlockNumber() uint64 {
	return b.Header.BlockNumber
}

// VerifyBlockHash verifies the block hash.
func (b *Block) VerifyBlockHash() bool {
	tempBlock := &Block{
		ChainHeader: b.ChainHeader,
		Header: &BlockHeader{
			BlockNumber: b.Header.BlockNumber,
			ParentHash:  b.Header.ParentHash,
		},
		Proof: b.Proof,
	}
	blockMsg, err := proto.Marshal(tempBlock)
	if err != nil {
		return false
	}
	return b.Hash() == common.Bytes2Hex(utils.Hash(blockMsg))
}

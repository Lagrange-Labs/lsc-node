package store

import (
	"context"

	sequencertypes "github.com/Lagrange-Labs/Lagrange-Node/sequencer/types"
)

type Storage interface {
	// AddNode adds a new node to the database.
	AddNode(ctx context.Context, stakeAdr, pubKey, ipAdr string) error
	// GetNode returns the node for the given IP address.
	GetNode(ctx context.Context, ip string) (*sequencertypes.ClientNode, error)
	// GetLastBlock returns the last block that was submitted to the network.
	GetLastBlock(ctx context.Context) (*sequencertypes.Block, error)
	// GetBlock returns the block for the given block number.
	GetBlock(ctx context.Context, blockNumber uint64) (*sequencertypes.Block, error)
	// GetNodeCount returns the number of nodes in the network.
	GetNodeCount(ctx context.Context) (uint16, error)
	// AddBlock adds a new block to the database.
	AddBlock(ctx context.Context, block *sequencertypes.Block) error
	// UpdateNode updates the node status in the database.
	UpdateNode(ctx context.Context, node *sequencertypes.ClientNode) error
	// GetLastBlockNumber returns the last block number that was submitted to the network.
	GetLastBlockNumber(ctx context.Context) (uint64, error)
}

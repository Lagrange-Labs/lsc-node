package store

import (
	"context"

	"github.com/Lagrange-Labs/Lagrange-Node/network"
	"github.com/Lagrange-Labs/Lagrange-Node/network/types"
)

type DB interface {
	// AddNode adds a new node to the database.
	AddNode(ctx context.Context, stakeAdr, pubKey, ipAdr string) error
	// GetNode returns the node for the given IP address.
	GetNode(ctx context.Context, ip string) (*network.ClientNode, error)
	// GetLastBlock returns the last block that was submitted to the network.
	GetLastBlock(ctx context.Context) (*types.Block, error)
	// GetBlock returns the block for the given block number.
	GetBlock(ctx context.Context, blockNumber uint64) (*types.Block, error)
	// GetNodeCount returns the number of nodes in the network.
	GetNodeCount(ctx context.Context) (uint16, error)
}

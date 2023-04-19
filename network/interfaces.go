package network

import (
	"context"

	sequencertypes "github.com/Lagrange-Labs/lagrange-node/sequencer/types"
)

type storageInterface interface {
	AddNode(ctx context.Context, node *sequencertypes.ClientNode) error
	GetNode(ctx context.Context, ip string) (*sequencertypes.ClientNode, error)
	GetBlock(ctx context.Context, blockNumber uint64) (*sequencertypes.Block, error)
	GetLastBlock(ctx context.Context) (*sequencertypes.Block, error)
	GetNodeCount(ctx context.Context) (uint16, error)
}

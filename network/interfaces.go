package network

import (
	"context"

	"github.com/Lagrange-Labs/lagrange-node/network/types"
	sequencertypes "github.com/Lagrange-Labs/lagrange-node/sequencer/types"
)

type storageInterface interface {
	AddNode(ctx context.Context, node *types.ClientNode) error
	GetNode(ctx context.Context, ip string) (*types.ClientNode, error)
	GetBlock(ctx context.Context, blockNumber uint64) (*sequencertypes.Block, error)
}

type consensusInterface interface {
	GetCurrentBlock() *sequencertypes.Block
	GetCurrentBlockNumber() uint64
	AddCommit(commit *types.CommitBlockRequest)
}

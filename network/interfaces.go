package network

import (
	"context"

	"github.com/Lagrange-Labs/lagrange-node/network/types"
	sequencertypes "github.com/Lagrange-Labs/lagrange-node/sequencer/types"
	storetypes "github.com/Lagrange-Labs/lagrange-node/store/types"
)

var _ storageInterface = (storetypes.Storage)(nil)

type storageInterface interface {
	AddNode(ctx context.Context, node *types.ClientNode) error
	GetNodeByStakeAddr(ctx context.Context, stakeAddress string) (*types.ClientNode, error)
	GetBlock(ctx context.Context, chainID uint32, blockNumber uint64) (*sequencertypes.Block, error)
}

type consensusInterface interface {
	GetCurrentBlock() *sequencertypes.Block
	GetCurrentBlockNumber() uint64
	GetCurrentEpochBlockNumber() uint64
	AddCommit(commit *types.CommitBlockRequest)
}

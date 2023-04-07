package network

import (
	"context"

	"github.com/Lagrange-Labs/lagrange-node/network/types"
	sequencertypes "github.com/Lagrange-Labs/lagrange-node/sequencer/types"
)

type storageInterface interface {
	AddNode(ctx context.Context, stakeAdr, pubKey, ipAdr string) error
	GetNode(ctx context.Context, ip string) (*sequencertypes.ClientNode, error)
	GetBlock(ctx context.Context, blockNumber uint64) (*types.Block, error)
	GetLastBlock(ctx context.Context) (*types.Block, error)
	GetNodeCount(ctx context.Context) (uint16, error)
}

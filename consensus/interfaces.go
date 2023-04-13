package consensus

import (
	"context"

	sequencertypes "github.com/Lagrange-Labs/lagrange-node/sequencer/types"
)

type storageInterface interface {
	GetLastFinalizedBlockNumber(ctx context.Context) (uint64, error)
	GetBlock(ctx context.Context, blockNumber uint64) (*sequencertypes.Block, error)
	UpdateBlock(ctx context.Context, block *sequencertypes.Block) error
	GetNodesByStatuses(ctx context.Context, statuses []sequencertypes.NodeStatus) ([]*sequencertypes.ClientNode, error)
}

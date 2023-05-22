package consensus

import (
	"context"

	"github.com/Lagrange-Labs/lagrange-node/consensus/types"
	networktypes "github.com/Lagrange-Labs/lagrange-node/network/types"
	sequencertypes "github.com/Lagrange-Labs/lagrange-node/sequencer/types"
	storetypes "github.com/Lagrange-Labs/lagrange-node/store/types"
)

var _ storageInterface = (storetypes.Storage)(nil)

type storageInterface interface {
	GetLastFinalizedBlockNumber(ctx context.Context, chainID int32) (uint64, error)
	GetBlock(ctx context.Context, blockNumber uint64) (*sequencertypes.Block, error)
	UpdateBlock(ctx context.Context, block *sequencertypes.Block) error
	GetNodesByStatuses(ctx context.Context, statuses []networktypes.NodeStatus) ([]networktypes.ClientNode, error)
	AddEvidences(ctx context.Context, evidences []*types.Evidence) error
}

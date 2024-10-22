package consensus

import (
	"context"

	"github.com/Lagrange-Labs/lsc-node/consensus/types"
	sequencertypes "github.com/Lagrange-Labs/lsc-node/sequencer/types"
	sequencerv2types "github.com/Lagrange-Labs/lsc-node/sequencer/types/v2"
	servertypes "github.com/Lagrange-Labs/lsc-node/server/types"
	storetypes "github.com/Lagrange-Labs/lsc-node/store/types"
)

var _ storageInterface = (storetypes.Storage)(nil)

type storageInterface interface {
	GetLastFinalizedBatchNumber(ctx context.Context, chainID uint32) (uint64, error)
	GetBatch(ctx context.Context, chainID uint32, batchNumber uint64) (*sequencerv2types.Batch, error)
	UpdateBatch(ctx context.Context, batch *sequencerv2types.Batch) error

	GetLastFinalizedBlock(ctx context.Context, chainID uint32) (*sequencertypes.Block, error)
	GetBlock(ctx context.Context, chainID uint32, blockNumber uint64) (*sequencertypes.Block, error)
	GetBlocks(ctx context.Context, chainID uint32, fromBlockNumber uint64, count uint32) ([]*sequencertypes.Block, error)
	UpdateBlock(ctx context.Context, block *sequencertypes.Block) error
	GetNodesByStatuses(ctx context.Context, statuses []servertypes.NodeStatus, chainID uint32) ([]servertypes.ClientNode, error)
	AddEvidences(ctx context.Context, evidences []*types.Evidence) error
	GetCommitteeRoot(ctx context.Context, chainID uint32, l1BlockNumber uint64) (*sequencerv2types.CommitteeRoot, error)
	GetCommitteeRootByEpochNumber(ctx context.Context, chainID uint32, epochNumber uint64) (*sequencerv2types.CommitteeRoot, error)
	GetLastCommitteeEpochNumber(ctx context.Context, chainID uint32) (uint64, error)
}

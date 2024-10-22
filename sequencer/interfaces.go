package sequencer

import (
	context "context"

	v2types "github.com/Lagrange-Labs/lsc-node/sequencer/types/v2"
	storetypes "github.com/Lagrange-Labs/lsc-node/store/types"
)

var _ storageInterface = (storetypes.Storage)(nil)

type storageInterface interface {
	GetLastBatchNumber(ctx context.Context, chainID uint32) (uint64, error)
	AddBatch(ctx context.Context, batch *v2types.Batch) error
	GetBatch(ctx context.Context, chainID uint32, batchNumber uint64) (*v2types.Batch, error)
	UpdateCommitteeRoot(ctx context.Context, committeeRoot *v2types.CommitteeRoot) error
	GetLastCommitteeEpochNumber(ctx context.Context, chainID uint32) (uint64, error)
	GetCommitteeRootByEpochNumber(ctx context.Context, chainID uint32, epochNumber uint64) (*v2types.CommitteeRoot, error)
}

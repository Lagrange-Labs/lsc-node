package sequencer

import (
	context "context"

	v2types "github.com/Lagrange-Labs/lagrange-node/sequencer/types/v2"
	storetypes "github.com/Lagrange-Labs/lagrange-node/store/types"
)

var _ storageInterface = (storetypes.Storage)(nil)

type storageInterface interface {
	GetLastBatchNumber(ctx context.Context, chainID uint32) (uint64, error)
	AddBatch(ctx context.Context, batch *v2types.Batch) error
	GetBatch(ctx context.Context, chainID uint32, batchNumber uint64) (*v2types.Batch, error)
}

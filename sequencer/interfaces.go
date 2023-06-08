package sequencer

import (
	context "context"

	"github.com/Lagrange-Labs/lagrange-node/sequencer/types"
	storetypes "github.com/Lagrange-Labs/lagrange-node/store/types"
)

var _ storageInterface = (storetypes.Storage)(nil)

type storageInterface interface {
	GetLastBlockNumber(ctx context.Context, chainID uint32) (uint64, error)
	AddBlock(ctx context.Context, block *types.Block) error
}

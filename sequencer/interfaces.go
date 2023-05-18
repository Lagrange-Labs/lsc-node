package sequencer

import (
	context "context"

	"github.com/Lagrange-Labs/lagrange-node/sequencer/types"
)

type storageInterface interface {
	GetLastBlockNumber(ctx context.Context, chainID int32) (uint64, error)
	AddBlock(ctx context.Context, block *types.Block) error
}

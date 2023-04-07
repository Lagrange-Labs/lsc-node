package sequencer

import (
	context "context"

	"github.com/Lagrange-Labs/lagrange-node/sequencer/types"
)

type storageInterface interface {
	GetLastBlockNumber(ctx context.Context) (uint64, error)
	AddBlock(ctx context.Context, block *types.Block) error
	UpdateNode(ctx context.Context, node *types.ClientNode) error
}

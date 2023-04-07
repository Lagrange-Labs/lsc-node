package sequencer

import (
	context "context"

	networktypes "github.com/Lagrange-Labs/lagrange-node/network/types"
	"github.com/Lagrange-Labs/lagrange-node/sequencer/types"
)

type storageInterface interface {
	GetLastBlockNumber(ctx context.Context) (uint64, error)
	AddBlock(ctx context.Context, block *networktypes.Block) error
	UpdateNode(ctx context.Context, node *types.ClientNode) error
}

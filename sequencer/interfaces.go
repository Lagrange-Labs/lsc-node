package sequencer

import (
	context "context"

	networktypes "github.com/Lagrange-Labs/Lagrange-Node/network/types"
	"github.com/Lagrange-Labs/Lagrange-Node/sequencer/types"
)

type storageInterface interface {
	GetLastBlockNumber(ctx context.Context) (uint64, error)
	AddBlock(ctx context.Context, block *networktypes.Block) error
	UpdateNode(ctx context.Context, node *types.ClientNode) error
}

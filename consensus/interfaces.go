package consensus

import (
	"context"

	"github.com/Lagrange-Labs/lagrange-node/consensus/types"
	govtypes "github.com/Lagrange-Labs/lagrange-node/governance/types"
	networktypes "github.com/Lagrange-Labs/lagrange-node/network/types"
	sequencertypes "github.com/Lagrange-Labs/lagrange-node/sequencer/types"
	storetypes "github.com/Lagrange-Labs/lagrange-node/store/types"
)

var _ storageInterface = (storetypes.Storage)(nil)

type storageInterface interface {
	GetLastFinalizedBlock(ctx context.Context, chainID uint32) (*sequencertypes.Block, error)
	GetBlock(ctx context.Context, chainID uint32, blockNumber uint64) (*sequencertypes.Block, error)
	GetBlocks(ctx context.Context, chainID uint32, fromBlockNumber uint64, count uint32) ([]*sequencertypes.Block, error)
	UpdateBlock(ctx context.Context, block *sequencertypes.Block) error
	GetNodesByStatuses(ctx context.Context, statuses []networktypes.NodeStatus, chainID uint32) ([]networktypes.ClientNode, error)
	AddEvidences(ctx context.Context, evidences []*types.Evidence) error
	GetCommitteeRoot(ctx context.Context, chainID uint32, l1BlockNumber uint64) (*govtypes.CommitteeRoot, error)
	GetLastCommitteeEpochNumber(ctx context.Context, chainID uint32) (uint64, error)
}

package governance

import (
	"context"

	contypes "github.com/Lagrange-Labs/lagrange-node/consensus/types"
	"github.com/Lagrange-Labs/lagrange-node/governance/types"
	networktypes "github.com/Lagrange-Labs/lagrange-node/network/types"
)

type storageInterface interface {
	GetNodesByStatuses(ctx context.Context, statuses []networktypes.NodeStatus, chainID uint32) ([]networktypes.ClientNode, error)
	UpdateNode(ctx context.Context, node *networktypes.ClientNode) error
	GetEvidences(ctx context.Context, fromBlockNumber, toBlockNumber uint64) ([]*contypes.Evidence, error)
	UpdateEvidence(ctx context.Context, evidence *contypes.Evidence) error
	UpdateCommitteeRoot(ctx context.Context, committeeRoot *types.CommitteeRoot) error
	GetLastCommitteeEpochNumber(ctx context.Context, chainID uint32) (uint64, error)
}

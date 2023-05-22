package governance

import (
	"context"

	contypes "github.com/Lagrange-Labs/lagrange-node/consensus/types"
	networktypes "github.com/Lagrange-Labs/lagrange-node/network/types"
)

type storageInterface interface {
	GetNodesByStatuses(ctx context.Context, statuses []networktypes.NodeStatus) ([]networktypes.ClientNode, error)
	UpdateNode(ctx context.Context, node *networktypes.ClientNode) error
	GetEvidences(ctx context.Context) ([]*contypes.Evidence, error)
	UpdateEvidence(ctx context.Context, evidence *contypes.Evidence) error
}

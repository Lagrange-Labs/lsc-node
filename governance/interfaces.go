package governance

import (
	"context"

	"github.com/Lagrange-Labs/lagrange-node/governance/types"
)

type storageInterface interface {
	UpdateCommitteeRoot(ctx context.Context, committeeRoot *types.CommitteeRoot) error
	GetLastCommitteeEpochNumber(ctx context.Context, chainID uint32) (uint64, error)
}

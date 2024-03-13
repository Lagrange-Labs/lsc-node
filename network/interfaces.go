package network

import (
	"context"

	"github.com/Lagrange-Labs/lagrange-node/crypto"
	"github.com/Lagrange-Labs/lagrange-node/network/types"
	sequencertypes "github.com/Lagrange-Labs/lagrange-node/sequencer/types"
	sequencerv2types "github.com/Lagrange-Labs/lagrange-node/sequencer/types/v2"
	storetypes "github.com/Lagrange-Labs/lagrange-node/store/types"
)

var _ storageInterface = (storetypes.Storage)(nil)

type storageInterface interface {
	AddNode(ctx context.Context, node *types.ClientNode) error
	GetNodeByStakeAddr(ctx context.Context, stakeAddress string, chainID uint32) (*types.ClientNode, error)
	GetBlock(ctx context.Context, chainID uint32, blockNumber uint64) (*sequencertypes.Block, error)
}

type consensusInterface interface {
	GetOpenBatch(batchNumber uint64) *sequencerv2types.Batch
	AddBatchCommit(commit *sequencerv2types.BlsSignature, stakeAddr string) error
	CheckCommitteeMember(stakeAddr string, pubKey []byte) bool
	IsFinalized(batchNumber uint64) bool
	GetBLSScheme() crypto.BLSScheme
}

package server

import (
	"context"
	"time"

	"github.com/Lagrange-Labs/lagrange-node/crypto"
	sequencertypes "github.com/Lagrange-Labs/lagrange-node/sequencer/types"
	sequencerv2types "github.com/Lagrange-Labs/lagrange-node/sequencer/types/v2"
	"github.com/Lagrange-Labs/lagrange-node/server/types"
	storetypes "github.com/Lagrange-Labs/lagrange-node/store/types"
)

var _ storageInterface = (storetypes.Storage)(nil)

type storageInterface interface {
	AddNode(ctx context.Context, node *types.ClientNode) error
	GetNodeByStakeAddr(ctx context.Context, stakeAddress string, chainID uint32) (*types.ClientNode, error)
	GetBlock(ctx context.Context, chainID uint32, blockNumber uint64) (*sequencertypes.Block, error)
}

type consensusInterface interface {
	Start()
	GetOpenBatch() *sequencerv2types.Batch
	GetPrevBatch() *sequencerv2types.Batch
	GetRoundInterval() time.Duration
	AddBatchCommit(commit *sequencerv2types.BlsSignature, stakeAddr, pubKey string) error
	CheckCommitteeMember(stakeAddr, pubKey string) (bool, error)
	CheckSignAddress(stakeAddr, signAddr string) bool
	IsFinalized(batchNumber uint64) bool
	GetBLSScheme() crypto.BLSScheme
}

package types

import (
	"context"

	contypes "github.com/Lagrange-Labs/lagrange-node/consensus/types"
	govtypes "github.com/Lagrange-Labs/lagrange-node/governance/types"
	networktypes "github.com/Lagrange-Labs/lagrange-node/network/types"
	sequencertypes "github.com/Lagrange-Labs/lagrange-node/sequencer/types"
)

type Storage interface {
	// AddNode adds a new node to the database.
	AddNode(ctx context.Context, node *networktypes.ClientNode) error
	// GetNodeByStakeAddr returns the node for the given stake address.
	GetNodeByStakeAddr(ctx context.Context, stakeAddress string, chainID uint32) (*networktypes.ClientNode, error)
	// GetLastBlock returns the last block that was submitted to the network.
	GetLastFinalizedBlock(ctx context.Context, chainID uint32) (*sequencertypes.Block, error)
	// GetBlock returns the block for the given block number.
	GetBlock(ctx context.Context, chainID uint32, blockNumber uint64) (*sequencertypes.Block, error)
	// GetBlocks returns the `count` blocks starting from `fromBlockNumber`.
	GetBlocks(ctx context.Context, chainID uint32, fromBlockNumber uint64, count uint32) ([]*sequencertypes.Block, error)
	// AddBlock adds a new block to the database.
	AddBlock(ctx context.Context, block *sequencertypes.Block) error
	// UpdateNode updates the node status in the database.
	UpdateNode(ctx context.Context, node *networktypes.ClientNode) error
	// GetLastBlockNumber returns the last block number that was submitted to the network.
	GetLastBlockNumber(ctx context.Context, chainID uint32) (uint64, error)
	// GetLastFinalizedBlockNumber returns the last block number that was finalized.
	GetLastFinalizedBlockNumber(ctx context.Context, chainID uint32) (uint64, bool, error)
	// GetNodesByStatuses returns the nodes with the given statuses.
	GetNodesByStatuses(ctx context.Context, statuses []networktypes.NodeStatus, chainID uint32) ([]networktypes.ClientNode, error)
	// UpdateBlock updates the block in the database.
	UpdateBlock(ctx context.Context, block *sequencertypes.Block) error
	// AddEvidences adds new evidences to the database.
	AddEvidences(ctx context.Context, evidences []*contypes.Evidence) error
	// GetEvidences returns the pending evidences for the given block number range.
	GetEvidences(ctx context.Context, chainID uint32, fromBlockNumber, toBlockNumber uint64) ([]*contypes.Evidence, error)
	// UpdateEvidence updates the evidence in the database.
	UpdateEvidence(ctx context.Context, evidence *contypes.Evidence) error
	// UpdateCommitteeRoot updates the committee root in the database.
	UpdateCommitteeRoot(ctx context.Context, committeeRoot *govtypes.CommitteeRoot) error
	// GetLastCommitteeRoot returns the last committee root for the given chainID.
	GetLastCommitteeRoot(ctx context.Context, chainID uint32, isFinalized bool) (*govtypes.CommitteeRoot, error)
	// GetCommitteeRoot returns the committee root for the given epoch block number.
	GetCommitteeRoot(ctx context.Context, chainID uint32, epochBlockNumber uint64) (*govtypes.CommitteeRoot, error)
	// GetLastCommitteeEpochNumber returns the last committee epoch number for the given chainID.
	GetLastCommitteeEpochNumber(ctx context.Context, chainID uint32) (uint64, error)
	// GetLastEvidenceBlockNumber returns the last submitted evidence block number for the given chainID.
	GetLastEvidenceBlockNumber(ctx context.Context, chainID uint32) (uint64, error)
}

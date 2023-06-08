package types

import (
	"context"

	contypes "github.com/Lagrange-Labs/lagrange-node/consensus/types"
	networktypes "github.com/Lagrange-Labs/lagrange-node/network/types"
	sequencertypes "github.com/Lagrange-Labs/lagrange-node/sequencer/types"
)

type Storage interface {
	// AddNode adds a new node to the database.
	AddNode(ctx context.Context, node *networktypes.ClientNode) error
	// GetNodeByStakeAddr returns the node for the given stake address.
	GetNodeByStakeAddr(ctx context.Context, stakeAddress string) (*networktypes.ClientNode, error)
	// GetLastBlock returns the last block that was submitted to the network.
	GetLastFinalizedBlock(ctx context.Context, chainID uint32) (*sequencertypes.Block, error)
	// GetBlock returns the block for the given block number.
	GetBlock(ctx context.Context, chainID uint32, blockNumber uint64) (*sequencertypes.Block, error)
	// AddBlock adds a new block to the database.
	AddBlock(ctx context.Context, block *sequencertypes.Block) error
	// UpdateNode updates the node status in the database.
	UpdateNode(ctx context.Context, node *networktypes.ClientNode) error
	// GetLastBlockNumber returns the last block number that was submitted to the network.
	GetLastBlockNumber(ctx context.Context, chainID uint32) (uint64, error)
	// GetLastFinalizedBlockNumber returns the last block number that was finalized.
	GetLastFinalizedBlockNumber(ctx context.Context, chainID uint32) (uint64, error)
	// GetNodesByStatuses returns the nodes with the given statuses.
	GetNodesByStatuses(ctx context.Context, statuses []networktypes.NodeStatus) ([]networktypes.ClientNode, error)
	// UpdateBlock updates the block in the database.
	UpdateBlock(ctx context.Context, block *sequencertypes.Block) error
	// AddEvidences adds new evidences to the database.
	AddEvidences(ctx context.Context, evidences []*contypes.Evidence) error
	// GetEvidences returns the pending evidences for the given block number.
	GetEvidences(ctx context.Context) ([]*contypes.Evidence, error)
	// UpdateEvidence updates the evidence in the database.
	UpdateEvidence(ctx context.Context, evidence *contypes.Evidence) error
}

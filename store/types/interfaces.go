package types

import (
	"context"

	contypes "github.com/Lagrange-Labs/lagrange-node/consensus/types"
	sequencertypes "github.com/Lagrange-Labs/lagrange-node/sequencer/types"
	sequencerv2types "github.com/Lagrange-Labs/lagrange-node/sequencer/types/v2"
	servertypes "github.com/Lagrange-Labs/lagrange-node/server/types"
)

// Storage is the interface for the database storage.
type Storage interface {
	// AddNode adds a new node to the database.
	AddNode(ctx context.Context, node *servertypes.ClientNode) error
	// GetNodesByStatuses returns the nodes with the given statuses.
	GetNodesByStatuses(ctx context.Context, statuses []servertypes.NodeStatus, chainID uint32) ([]servertypes.ClientNode, error)
	// GetNodeByStakeAddr returns the node for the given stake address.
	GetNodeByStakeAddr(ctx context.Context, stakeAddress string, chainID uint32) (*servertypes.ClientNode, error)
	// GetLastFinalizedBlock returns the last finalized block for the given chainID.
	GetLastFinalizedBlock(ctx context.Context, chainID uint32) (*sequencertypes.Block, error)
	// GetLastFinalizedBlockNumber returns the last finalized block number for the given chainID.
	GetLastFinalizedBlockNumber(ctx context.Context, chainID uint32) (uint64, error)
	// GetBlock returns the block for the given block number.
	GetBlock(ctx context.Context, chainID uint32, blockNumber uint64) (*sequencertypes.Block, error)
	// GetBlocks returns the `count` blocks starting from `fromBlockNumber`.
	GetBlocks(ctx context.Context, chainID uint32, fromBlockNumber uint64, count uint32) ([]*sequencertypes.Block, error)
	// AddBlock adds a new block to the database.
	AddBlock(ctx context.Context, block *sequencertypes.Block) error
	// UpdateBlock updates the block in the database.
	UpdateBlock(ctx context.Context, block *sequencertypes.Block) error
	// GetLastBlockNumber returns the last block number that was stored to the db.
	GetLastBlockNumber(ctx context.Context, chainID uint32) (uint64, error)
	// GetLastFinalizedBatchNumber returns the last finalized batch number for the given chainID.
	GetLastFinalizedBatchNumber(ctx context.Context, chainID uint32) (uint64, error)
	// GetLastBatchNumber returns the last batch number that was stored to the db.
	GetLastBatchNumber(ctx context.Context, chainID uint32) (uint64, error)
	// GetBatch returns the batch for the given batch number.
	GetBatch(ctx context.Context, chainID uint32, batchNumber uint64) (*sequencerv2types.Batch, error)
	// AddBatch adds a new batch to the database.
	AddBatch(ctx context.Context, batch *sequencerv2types.Batch) error
	// UpdateBatch updates the batch in the database.
	UpdateBatch(ctx context.Context, batch *sequencerv2types.Batch) error
	// AddEvidences adds new evidences to the database.
	AddEvidences(ctx context.Context, evidences []*contypes.Evidence) error
	// GetEvidences returns the pending evidences for the given block number range.
	GetEvidences(ctx context.Context, chainID uint32, fromBlockNumber, toBlockNumber uint64, limit, offset int64) ([]*contypes.Evidence, error)
	// UpdateEvidence updates the evidence in the database.
	UpdateEvidence(ctx context.Context, evidence *contypes.Evidence) error
	// UpdateCommitteeRoot updates the committee root in the database.
	UpdateCommitteeRoot(ctx context.Context, committeeRoot *sequencerv2types.CommitteeRoot) error
	// GetCommitteeRoot returns the first committee root which EpochBlockNumber is greater than or equal to the given l1BlockNumber.
	GetCommitteeRoot(ctx context.Context, chainID uint32, l1BlockNumber uint64) (*sequencerv2types.CommitteeRoot, error)
	// GetLastCommitteeEpochNumber returns the last committee epoch number for the given chainID.
	GetLastCommitteeEpochNumber(ctx context.Context, chainID uint32) (uint64, error)
	// GetLastEvidenceBlockNumber returns the last submitted evidence block number for the given chainID.
	GetLastEvidenceBlockNumber(ctx context.Context, chainID uint32) (uint64, error)
	// CleanUp cleans up the database.
	CleanUp(ctx context.Context) error
}

// KVStorage is the interface for the key-value database storage.
type KVStorage interface {
	// Get returns the value for the given key.
	Get(key []byte) ([]byte, error)
	// Put puts the value for the given key.
	Put(key, value []byte) error
	// Next returns the next key-value pair.
	Next(key []byte) ([]byte, []byte, error)
	// Prev returns the previous key-value pair.
	Prev(key []byte) ([]byte, []byte, error)
	// Iterate iterates over the key-value pairs.
	Iterate(prefix []byte, f func(key, value []byte) error) error
	// Prune prunes the key-value pairs.
	Prune(prefix []byte) error
	// Close closes the database.
	Close() error
}

package types

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"

	sequencerv2types "github.com/Lagrange-Labs/lagrange-node/sequencer/types/v2"
)

var (
	// ErrBlockNotFound is returned when the block is not found.
	ErrBlockNotFound = fmt.Errorf("block not found")
	// ErrBatchNotFound is returned when the batch is not found.
	ErrBatchNotFound = fmt.Errorf("batch not found")
	// ErrNoResult is returned when there is no result.
	ErrNoResult = fmt.Errorf("no result")
)

type RpcClient interface {
	// GetCurrentBlockNumber returns the current L2 block number.
	GetCurrentBlockNumber() (uint64, error)
	// GetFinalizedBlockNumber returns the L2 finalized block number.
	GetFinalizedBlockNumber() (uint64, error)
	// GetChainID returns the chain ID.
	GetChainID() (uint32, error)
	// SetBeginBlockNumber sets the begin L1 & L2 block number.
	SetBeginBlockNumber(l1BlockNumber, l2BlockNumber uint64)
	// NextBatch returns the next batch after SetBeginBlockNumber.
	NextBatch() (*sequencerv2types.BatchHeader, error)
}

type EvmClient interface {
	// GetChainID returns the chain ID.
	GetChainID() (uint32, error)
	// GetBlockHashByNumber returns the block hash by the given block number.
	GetBlockHashByNumber(blockNumber uint64) (common.Hash, error)
	// GetBlockNumberByHash returns the block number by the given block hash.
	GetBlockNumberByHash(blockHash common.Hash) (uint64, error)
	// GetBlockNumberByTxHash returns the block number by the given transaction hash.
	GetBlockNumberByTxHash(txHash common.Hash) (uint64, error)
	// GetFinalizedBlockNumber returns the finalized block number.
	GetFinalizedBlockNumber() (uint64, error)
	// GetBlockHashesByRange returns the block hashes by the given range.
	GetBlockHashesByRange(startBlockNumber, endBlockNumber uint64) ([]common.Hash, error)
}

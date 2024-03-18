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
	// GetBlockHeaderByNumber returns the L2 block header by the given block number.
	GetBlockHeaderByNumber(blockNumber uint64, l1TxHash common.Hash) (*L2BlockHeader, error)
	// GetFinalizedBlockNumber returns the L2 finalized block number.
	GetFinalizedBlockNumber() (uint64, error)
	// GetChainID returns the chain ID.
	GetChainID() (uint32, error)
	// SetBeginBlockNumber sets the begin L1 block number.
	SetBeginBlockNumber(l1BlockNumber uint64)
	// GetBatchHeaderByNumber returns the batch header for the given L2 block number.
	GetBatchHeaderByNumber(l2BlockNumber uint64) (*sequencerv2types.BatchHeader, error)
}

// L2 Block Header
type L2BlockHeader struct {
	L1BlockNumber uint64      `json:"l1BlockNumber"`
	L2BlockHash   common.Hash `json:"l2BlockHash"`
}

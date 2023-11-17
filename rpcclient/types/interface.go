package types

import "fmt"

var (
	// ErrBlockNotFound is returned when the block is not found.
	ErrBlockNotFound = fmt.Errorf("block not found")
)

type RpcClient interface {
	// GetCurrentBlockNumber returns the current L2 block number.
	GetCurrentBlockNumber() (uint64, error)
	// GetBlockHashByNumber returns the L2 block hash by the given block number.
	GetBlockHashByNumber(blockNumber uint64) (string, error)
	// GetFinalizedBlockNumber returns the L2 finalized block number.
	GetFinalizedBlockNumber() (uint64, error)
	// GetChainID returns the chain ID.
	GetChainID() (uint32, error)
	// GetL1BlockNumber returns the L1 block number for the given L2 block number.
	GetL1BlockNumber(l2BlockNumber uint64) (uint64, error)
}

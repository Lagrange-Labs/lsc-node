package types

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
)

var (
	// ErrBlockNotFound is returned when the block is not found.
	ErrBlockNotFound = fmt.Errorf("block not found")
)

type RpcClient interface {
	// GetCurrentBlockNumber returns the current L2 block number.
	GetCurrentBlockNumber() (uint64, error)
	// GetBlockHeaderByNumber returns the L2 block header by the given block number.
	GetBlockHeaderByNumber(blockNumber uint64) (*L2BlockHeader, error)
	// GetFinalizedBlockNumber returns the L2 finalized block number.
	GetFinalizedBlockNumber() (uint64, error)
	// GetChainID returns the chain ID.
	GetChainID() (uint32, error)
}

// L2 Block Header
type L2BlockHeader struct {
	L1BlockNumber uint64      `json:"l1BlockNumber"`
	L2BlockHash   common.Hash `json:"l2BlockHash"`
}

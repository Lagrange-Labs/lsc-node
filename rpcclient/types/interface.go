package types

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	coretypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

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
	// SetBeginBlockNumber sets the begin L1 and L2 block number.
	SetBeginBlockNumber(l1BlockNumber, l2BlockNumber uint64) bool
	// NextBatch returns the next batch after SetBeginBlockNumber.
	NextBatch() (*sequencerv2types.BatchHeader, error)
	// GetBlockHashFromRLPHeader returns the block hash and the parent hash from the rlp encoded block header.
	GetBlockHashFromRLPHeader(rlpHeader []byte) (common.Hash, common.Hash, error)
}

type EvmClient interface {
	// GetEthClient returns the eth client.
	GetEthClient() *ethclient.Client
	// GetChainID returns the chain ID.
	GetChainID() (uint32, error)
	// GetBlockByNumber returns the block by the given block number.
	GetBlockByNumber(blockNumber uint64) (*coretypes.Block, error)
	// GetBlockHashByNumber returns the block hash by the given block number.
	GetBlockHashByNumber(blockNumber uint64) (common.Hash, error)
	// GetBlockNumberByHash returns the block number by the given block hash.
	GetBlockNumberByHash(blockHash common.Hash) (uint64, error)
	// GetBlockNumberByTxHash returns the block number by the given transaction hash.
	GetBlockNumberByTxHash(txHash common.Hash) (uint64, error)
	// GetFinalizedBlockNumber returns the finalized block number.
	GetFinalizedBlockNumber() (uint64, error)
	// GetBlockHeadersByRange returns the block headers by the given range.
	GetBlockHeadersByRange(startBlockNumber, endBlockNumber uint64) ([]sequencerv2types.BlockHeader, error)
}

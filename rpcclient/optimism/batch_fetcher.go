package optimism

import (
	"context"
	"math/big"
	"sync"
	"sync/atomic"
	"time"

	"github.com/Lagrange-Labs/lagrange-node/logger"
	"github.com/Lagrange-Labs/lagrange-node/rpcclient/types"
	"github.com/ethereum-optimism/optimism/op-node/rollup/derive"
	"github.com/ethereum/go-ethereum/common"
	coretypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/sync/errgroup"
)

const (
	ConcurrentFetchers    = 4
	EthereumFinalityDepth = 64
	BlockHeaderCacheSize  = 1024
	CacheFullWaitInterval = 10 * time.Microsecond
	FetchInterval         = 5 * time.Second
)

// Fetcher is a synchronizer for the BatchInbox EOA.
type Fetcher struct {
	l1Client          *ethclient.Client
	l2Client          *ethclient.Client
	beginBlockNumber  uint64
	batchInboxAddress common.Address
	batchSender       common.Address
	signer            coretypes.Signer
	cache             sync.Map // block number -> block header
	cacheCount        atomic.Int32
}

// NewFetcher creates a new Fetcher instance.
func NewFetcher(cfg *Config) (*Fetcher, error) {
	l1Client, err := ethclient.Dial(cfg.L1RPCURL)
	if err != nil {
		return nil, err
	}
	l2Client, err := ethclient.Dial(cfg.RPCURL)
	if err != nil {
		return nil, err
	}
	chainID, err := l1Client.ChainID(context.Background())
	if err != nil {
		return nil, err
	}

	return &Fetcher{
		l1Client:          l1Client,
		l2Client:          l2Client,
		beginBlockNumber:  cfg.BeginBlockNumber,
		batchInboxAddress: common.HexToAddress(cfg.BatchInbox),
		batchSender:       common.HexToAddress(cfg.BatchSender),
		signer:            coretypes.LatestSignerForChainID(chainID),
	}, nil
}

// Fetch fetches the block data from the Ethereum and analyzes the
// transactions which are sent to the BatchInbox EOA.
func (f *Fetcher) Fetch() error {
	lastSyncedBlockNumber := f.beginBlockNumber

	for {
		g, ctx := errgroup.WithContext(context.Background())
		g.SetLimit(ConcurrentFetchers)
		// Fetch the latest finalized block number.
		blockNumber, err := f.l1Client.BlockNumber(ctx)
		if err != nil {
			return err
		}
		if lastSyncedBlockNumber > blockNumber-EthereumFinalityDepth {
			time.Sleep(FetchInterval)
			continue
		}
		for i := lastSyncedBlockNumber; i <= blockNumber-EthereumFinalityDepth; i++ {
			if err := ctx.Err(); err != nil {
				return err
			}

			number := i
			g.Go(func() error {
				return f.fetchBlock(ctx, number)
			})
		}
		if err := g.Wait(); err != nil {
			return err
		}
		lastSyncedBlockNumber = blockNumber - EthereumFinalityDepth + 1
	}
}

// fetchBlock fetches the given block and analyzes the transactions
// which are sent to the BatchInbox EOA.
func (f *Fetcher) fetchBlock(ctx context.Context, blockNumber uint64) error {
	block, err := f.l1Client.BlockByNumber(ctx, big.NewInt(int64(blockNumber)))
	if err != nil {
		return err
	}

	for _, tx := range block.Transactions() {
		if !f.validTransaction(tx) {
			continue
		}
		if err := f.decodeBatchTx(blockNumber, tx); err != nil {
			return err
		}
	}

	return nil
}

// decodeBatchTx decodes the given transaction and stores the batch data into the cache.
func (f *Fetcher) decodeBatchTx(blockNumber uint64, tx *coretypes.Transaction) error {
	frames, err := derive.ParseFrames(tx.Data())
	if err != nil {
		return err
	}

	batches, err := handleFrames(blockNumber, frames)
	if err != nil {
		return err
	}

	for _, batch := range batches {
		parentL2Block, err := f.l2Client.BlockByHash(context.Background(), batch.ParentHash)
		if err != nil {
			return err
		}
		// wait for the cache to be consumed
		for f.cacheCount.Load() > BlockHeaderCacheSize {
			time.Sleep(CacheFullWaitInterval)
		}
		f.cache.Store(blockNumber, types.L2BlockHeader{
			L2BlockHash:   parentL2Block.Header().Hash(),
			L1BlockNumber: blockNumber,
		})
		f.cacheCount.Add(1)
	}

	logger.Infof("block number: %v, batch count: %v\n", blockNumber, len(batches))

	return nil
}

// getL2BlockHeader returns the L2 block header for the give block number
// and removes the header from the cache.
func (f *Fetcher) getL2BlockHeader(blockNumber uint64) (*types.L2BlockHeader, error) {
	raw, ok := f.cache.Load(blockNumber)
	if !ok {
		return nil, types.ErrBlockNotFound
	}
	f.cache.Delete(blockNumber)
	f.cacheCount.Add(-1)

	return raw.(*types.L2BlockHeader), nil
}

// getL2BlockHeaderByTxHash returns the L2 block header for the given L1 transaction hash.
func (f *Fetcher) getL2BlockHeaderByTxHash(blockNumber uint64, l1TxHash common.Hash) (*types.L2BlockHeader, error) {
	tx, _, err := f.l1Client.TransactionByHash(context.Background(), l1TxHash)
	if err != nil {
		return nil, err
	}
	if !f.validTransaction(tx) {
		return nil, types.ErrBlockNotFound
	}
	receipt, err := f.l1Client.TransactionReceipt(context.Background(), l1TxHash)
	if err != nil {
		return nil, err
	}
	if err := f.decodeBatchTx(receipt.BlockNumber.Uint64(), tx); err != nil {
		return nil, err
	}

	return f.getL2BlockHeader(blockNumber)
}

// validTransaction returns true if the given transaction is valid.
func (f *Fetcher) validTransaction(tx *coretypes.Transaction) bool {
	if tx == nil || tx.To() == nil {
		return false
	}
	if *tx.To() != f.batchInboxAddress {
		return false
	}
	from, err := f.signer.Sender(tx)
	if err != nil {
		return false
	}
	if from != f.batchSender {
		return false
	}
	return true
}

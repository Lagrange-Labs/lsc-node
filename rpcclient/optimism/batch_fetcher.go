package optimism

import (
	"context"
	"math/big"
	"time"

	"github.com/Lagrange-Labs/lagrange-node/logger"
	"github.com/ethereum-optimism/optimism/op-node/rollup/derive"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/sync/errgroup"
)

const (
	ConcurrentFetchers    = 4
	EthereumFinalityDepth = 64
	FetchInterval         = 5 * time.Second
)

// Fetcher is a synchronizer for the BatchInbox EOA.
type Fetcher struct {
	client            ethclient.Client
	beginBlockNumber  uint64
	batchInboxAddress common.Address
	batchSender       common.Address
	chainID           *big.Int
}

// NewFetcher creates a new Fetcher instance.
func NewFetcher(cfg *Config) (*Fetcher, error) {
	client, err := ethclient.Dial(cfg.L1RPCURL)
	if err != nil {
		return nil, err
	}
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return nil, err
	}

	return &Fetcher{
		client:            *client,
		beginBlockNumber:  cfg.BeginBlockNumber,
		batchInboxAddress: common.HexToAddress(cfg.BatchInbox),
		batchSender:       common.HexToAddress(cfg.BatchSender),
		chainID:           chainID,
	}, nil
}

// Fetch fetches the block data from the Ethereum and analyzes the
// transactions which are sent to the BatchInbox EOA.
func (f *Fetcher) Fetch() error {
	g, ctx := errgroup.WithContext(context.Background())
	g.SetLimit(ConcurrentFetchers)

	lastSyncedBlockNumber := f.beginBlockNumber

	for {
		if err := ctx.Err(); err != nil {
			return err
		}
		// Fetch the latest finalized block number.
		blockNumber, err := f.client.BlockNumber(ctx)
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
	block, err := f.client.BlockByNumber(ctx, big.NewInt(int64(blockNumber)))
	if err != nil {
		return err
	}

	for _, tx := range block.Transactions() {
		if !f.validTransaction(tx) {
			continue
		}
		frames, err := derive.ParseFrames(tx.Data())
		if err != nil {
			return err
		}

		batches, err := handleFrames(blockNumber, frames)
		if err != nil {
			return err
		}

		for _, batch := range batches {
			logger.Infof("batch %v", batch)
		}
	}

	return nil
}

// validTransaction returns true if the given transaction is valid.
func (f *Fetcher) validTransaction(tx *types.Transaction) bool {
	if tx.To().Hex() != f.batchInboxAddress.Hex() {
		return false
	}
	from, err := types.Sender(types.NewEIP155Signer(f.chainID), tx)
	if err != nil {
		return false
	}
	if from.Hex() != f.batchSender.Hex() {
		return false
	}
	return true
}

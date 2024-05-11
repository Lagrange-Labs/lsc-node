package arbitrum

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"sort"
	"sync"
	"sync/atomic"
	"time"

	"github.com/ethereum-optimism/optimism/op-service/client"
	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum-optimism/optimism/op-service/sources"
	"github.com/ethereum/go-ethereum/common"
	coretypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/log"
	"golang.org/x/sync/errgroup"

	"github.com/Lagrange-Labs/lagrange-node/logger"
	"github.com/Lagrange-Labs/lagrange-node/rpcclient/evmclient"
	"github.com/Lagrange-Labs/lagrange-node/rpcclient/types"
	sequencerv2types "github.com/Lagrange-Labs/lagrange-node/sequencer/types/v2"
	"github.com/Lagrange-Labs/lagrange-node/telemetry"
	"github.com/Lagrange-Labs/lagrange-node/utils"
)

const (
	ParallelBlocks = 32
	cacheLimit     = 1024
	maxTxBlobCount = 10000
	FetchInterval  = 5 * time.Second
)

// BatchesRef is a struct to represent the batch reference.
type BatchesRef struct {
	L1BlockNumber     uint64
	L1TxHash          common.Hash
	L1TxIndex         uint
	FromL2BlockNumber uint64
	ToL2BlockNumber   uint64
}

// Fetcher is a synchronizer for the BatchInbox EOA.
type Fetcher struct {
	l1Client          *ethclient.Client
	l1EvmClient       types.EvmClient
	l2Client          types.EvmClient
	l1BlobFetcher     *sources.L1BeaconClient
	sequencerInbox    *SequencerInbox
	concurrentFetcher int
	signer            coretypes.Signer
	l2BlockCache      *utils.Cache
	batchHeaders      chan *BatchesRef

	chainID                 *big.Int
	lastSyncedL1BlockNumber atomic.Uint64
	lastSyncedL2BlockNumber uint64

	mtx    sync.Mutex
	ctx    context.Context
	cancel context.CancelFunc
	done   chan struct{}
}

// NewFetcher creates a new Fetcher instance.
func NewFetcher(cfg *Config) (*Fetcher, error) {
	l1Client, err := ethclient.Dial(cfg.L1RPCURL)
	if err != nil {
		return nil, err
	}
	l1EvmClient, err := evmclient.NewClient(cfg.L1RPCURL)
	if err != nil {
		return nil, err
	}

	l2Client, err := evmclient.NewClient(cfg.RPCURL)
	if err != nil {
		return nil, err
	}
	chainID, err := l1Client.ChainID(context.Background())
	if err != nil {
		return nil, err
	}
	l2ChainID, err := l2Client.GetChainID()
	if err != nil {
		return nil, err
	}

	l1Beacon := sources.NewBeaconHTTPClient(client.NewBasicHTTPClient(cfg.BeaconURL, log.New()))
	l1BlobFetcher := sources.NewL1BeaconClient(l1Beacon, sources.L1BeaconClientConfig{FetchAllSidecars: false})

	sequencerInbox, err := NewSequencerInbox(common.HexToAddress(cfg.BatchInbox), l1Client)
	if err != nil {
		return nil, err
	}

	return &Fetcher{
		l1Client:          l1Client,
		l1EvmClient:       l1EvmClient,
		l2Client:          l2Client,
		l1BlobFetcher:     l1BlobFetcher,
		sequencerInbox:    sequencerInbox,
		chainID:           big.NewInt(int64(l2ChainID)),
		concurrentFetcher: cfg.ConcurrentFetchers,
		signer:            coretypes.LatestSignerForChainID(chainID),
		l2BlockCache:      utils.NewCache(cacheLimit),
		batchHeaders:      make(chan *BatchesRef, 64),

		done: make(chan struct{}, 2),
	}, nil
}

// GetFetchedBlockNumber returns the last fetched L1 block number.
func (f *Fetcher) GetFetchedBlockNumber() uint64 {
	return f.lastSyncedL1BlockNumber.Load()
}

// InitFetch inits the fetcher context.
func (f *Fetcher) InitFetch(l2BlockNumber uint64) {
	f.ctx, f.cancel = context.WithCancel(context.Background())
	f.lastSyncedL2BlockNumber = l2BlockNumber
}

// Fetch fetches the block data from the Ethereum and analyzes the
// transactions which are sent to the BatchSequencer Contractor.
func (f *Fetcher) Fetch(l1BeginBlockNumber uint64) error {
	defer func() {
		logger.Infof("l1 fetcher is stopped")
		f.done <- struct{}{}
	}()

	f.lastSyncedL1BlockNumber.Store(l1BeginBlockNumber)

	for {
		select {
		case <-f.ctx.Done():
			return nil
		default:
			// Fetch the latest finalized block number.
			blockNumber, err := f.l1EvmClient.GetFinalizedBlockNumber()
			if err != nil {
				return err
			}
			lastSyncedL1BlockNumber := f.lastSyncedL1BlockNumber.Load()
			nextBlockNumber := lastSyncedL1BlockNumber + ParallelBlocks
			if blockNumber < nextBlockNumber {
				nextBlockNumber = blockNumber
			}
			if lastSyncedL1BlockNumber > nextBlockNumber {
				time.Sleep(FetchInterval)
				continue
			}
			ti := time.Now()
			batches, err := f.sequencerInbox.fetchBatchTransactions(f.ctx, big.NewInt(int64(lastSyncedL1BlockNumber)), big.NewInt(int64(nextBlockNumber)))
			if err != nil {
				return err
			}
			telemetry.MeasureSince(ti, "rpc_arbitrum", "l1_filter_logs")

			// sort the batches by L1 block number and L1 tx index
			sort.Slice(batches, func(i, j int) bool {
				if batches[i].BlockNumber == batches[j].BlockNumber {
					return batches[i].TxIndex < batches[j].TxIndex
				}
				return batches[i].BlockNumber < batches[j].BlockNumber
			})
			for _, batch := range batches {
				var rawMsg []byte
				if batch.serialized[0] == BlobHashesHeaderFlag {
					rawMsg, err = f.fetchBlock(f.ctx, batch.BlockNumber, batch.TxHash)
					if err != nil {
						return err
					}
				} else {
					rawMsg = batch.serialized
				}
				batch.segments, err = decompress(rawMsg)
				if err != nil {
					return err
				}
				_, err := f.sequencerInbox.parseL2Transactions(batch)
				if err != nil {
					return err
				}
				batchesRef, err := f.getBatchRef(batch)
				if err != nil {
					return err
				}
				logger.Infof("batch reference is fetched: %+v", batchesRef)
				f.batchHeaders <- batchesRef
			}

			f.lastSyncedL1BlockNumber.Store(nextBlockNumber + 1)
		}
	}
}

// FetchL2Blocks fetches the L2 blocks from the given L2 block number.
func (f *Fetcher) FetchL2Blocks() error {
	defer func() {
		logger.Info("l2 fetcher is stopped")
		f.done <- struct{}{}
	}()

	l2BeginBlockNumber := f.lastSyncedL2BlockNumber
	for {
		select {
		case <-f.ctx.Done():
			return nil
		default:
			// Fetch the latest finalized block number.
			blockNumber, err := f.l2Client.GetFinalizedBlockNumber()
			if err != nil {
				return err
			}
			if l2BeginBlockNumber >= blockNumber {
				time.Sleep(FetchInterval)
				continue
			}
			g, ctx := errgroup.WithContext(context.Background())
			g.SetLimit(f.concurrentFetcher)
			nextBlockNumber := l2BeginBlockNumber + ParallelBlocks
			if blockNumber < nextBlockNumber {
				nextBlockNumber = blockNumber
			}
			ti := time.Now()
			for i := l2BeginBlockNumber; i < nextBlockNumber; i++ {
				if err := ctx.Err(); err != nil {
					logger.Errorf("fetch l2 block context error: %v", err)
					return err
				}
				number := i
				g.Go(func() error {
					blockHash, err := f.l2Client.GetBlockHashByNumber(number)
					if err != nil {
						return err
					}
					f.l2BlockCache.Set(number, blockHash)
					return nil
				})
			}
			if err := g.Wait(); err != nil {
				return err
			}
			telemetry.MeasureSince(ti, "rpc_optimism", "fetch_l2_blocks")
			logger.Infof("L2 blocks are fetched from %d to %d in %v milliseconds", l2BeginBlockNumber, nextBlockNumber, time.Since(ti).Milliseconds())
			l2BeginBlockNumber = nextBlockNumber
		}
	}
}

// Stop stops the Fetcher.
func (f *Fetcher) Stop() {
	f.mtx.Lock()
	defer f.mtx.Unlock()

	if f.cancel == nil {
		return
	}

	f.lastSyncedL1BlockNumber.Store(0)
	// close L1 & L2 fetcher
	f.cancel()
	<-f.done
	<-f.done
	// release retains and close batch headers channel to notify the outside
	close(f.batchHeaders)
	for range f.batchHeaders {
	}

	f.cancel = nil
	f.ctx = nil
}

// fetchBlock fetches the given block and analyzes the transactions
// which are sent to the BatchInbox EOA.
func (f *Fetcher) fetchBlock(ctx context.Context, blockNumber uint64, txHash common.Hash) ([]byte, error) {
	block, err := f.l1Client.BlockByNumber(ctx, big.NewInt(int64(blockNumber)))
	if err != nil {
		return nil, err
	}

	res := make([]byte, 0)
	var hashes []eth.IndexedBlobHash
	blobIndex := 0
	for _, tx := range block.Transactions() {
		if tx.Hash() != txHash {
			blobIndex += len(tx.BlobHashes())
			continue
		}
		if tx.Type() != coretypes.BlobTxType {
			return nil, fmt.Errorf("unexpected tx type: %v", tx.Type())
		}
		for _, hash := range tx.BlobHashes() {
			hashes = append(hashes, eth.IndexedBlobHash{
				Index: uint64(blobIndex),
				Hash:  hash,
			})
			blobIndex++
		}
		break
	}

	if len(hashes) > 0 {
		blockRef := eth.L1BlockRef{
			Number:     blockNumber,
			Hash:       block.Hash(),
			ParentHash: block.ParentHash(),
			Time:       block.Time(),
		}
		ti := time.Now()
		blobs, err := f.l1BlobFetcher.GetBlobs(ctx, blockRef, hashes)
		if err != nil {
			logger.Errorf("failed to get blobs: %v", err)
			return nil, err
		}
		telemetry.MeasureSince(ti, "rpc_arbitrum", "fetch_beacon_blobs")
		if len(blobs) != len(hashes) {
			logger.Errorf("blobs length is not matched: %d, %d", len(blobs), len(hashes))
			return nil, fmt.Errorf("blobs length is not matched: %d, %d", len(blobs), len(hashes))
		}
		res, err = decodeBlobs(blobs)
		if err != nil {
			logger.Errorf("failed to decode blobs: %v", err)
			return nil, err
		}
	}

	return res, nil
}

// getL2BlockHash returns the L2 block hash for the given block number.
func (f *Fetcher) getL2BlockHash(blockNumber uint64) (common.Hash, error) {
	hash, ok := f.l2BlockCache.Get(blockNumber)
	if ok {
		return hash.(common.Hash), nil
	}
	blockHash, err := f.l2Client.GetBlockHashByNumber(blockNumber)
	if err != nil {
		return common.Hash{}, err
	}
	f.l2BlockCache.Set(blockNumber, blockHash)
	return blockHash, nil
}

// getBatchRef returns the batch reference from the SequencerBatch.
func (f *Fetcher) getBatchRef(batch *SequencerBatch) (*BatchesRef, error) {
	if len(batch.txes) == 0 {
		return nil, errors.New("no transactions in the batch")
	}

	startBlockNumber, err := f.l2Client.GetBlockNumberByTxHash(batch.txes[0].Hash())
	if err != nil {
		return nil, err
	}
	endBlockNumber, err := f.l2Client.GetBlockNumberByTxHash(batch.txes[len(batch.txes)-1].Hash())
	if err != nil {
		return nil, err
	}

	return &BatchesRef{
		L1BlockNumber:     batch.BlockNumber,
		L1TxHash:          batch.TxHash,
		L1TxIndex:         batch.TxIndex,
		FromL2BlockNumber: startBlockNumber,
		ToL2BlockNumber:   endBlockNumber,
	}, nil
}

// nextBatchHeader returns the L2 batch header.
func (f *Fetcher) nextBatchHeader() (*sequencerv2types.BatchHeader, error) {
	f.mtx.Lock()
	defer f.mtx.Unlock()

	batchesRef, ok := <-f.batchHeaders
	if !ok {
		return nil, errors.New("batch headers channel is closed")
	}

	header := sequencerv2types.BatchHeader{
		L1BlockNumber: batchesRef.L1BlockNumber,
		L1TxHash:      batchesRef.L1TxHash.Hex(),
		L1TxIndex:     uint32(batchesRef.L1TxIndex),
		ChainId:       uint32(f.chainID.Uint64()),
	}

	l2Blocks := make([]*sequencerv2types.BlockHeader, 0)
	for i := batchesRef.FromL2BlockNumber; i <= batchesRef.ToL2BlockNumber; i++ {
		blockHash, err := f.getL2BlockHash(i)
		if err != nil {
			return nil, err
		}
		l2Blocks = append(l2Blocks, &sequencerv2types.BlockHeader{
			BlockNumber: i,
			BlockHash:   blockHash.Hex(),
		})

	}
	header.L2Blocks = l2Blocks

	return &header, nil
}

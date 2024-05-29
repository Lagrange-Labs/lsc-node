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
	fetchInterval  = 5 * time.Second
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
	l1Client          types.EvmClient
	l2Client          types.EvmClient
	l1BlobFetcher     *sources.L1BeaconClient
	sequencerInbox    *SequencerInbox
	concurrentFetcher int
	batchHeaders      chan *BatchesRef

	chainID                 *big.Int
	lastSyncedL1BlockNumber atomic.Uint64

	ctx    context.Context
	cancel context.CancelFunc
	done   chan struct{}
}

// NewFetcher creates a new Fetcher instance.
func NewFetcher(cfg *Config) (*Fetcher, error) {
	l1Client, err := evmclient.NewClient(cfg.L1RPCURL)
	if err != nil {
		return nil, err
	}

	l2Client, err := evmclient.NewClient(cfg.RPCURL)
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
		l2Client:          l2Client,
		l1BlobFetcher:     l1BlobFetcher,
		sequencerInbox:    sequencerInbox,
		chainID:           big.NewInt(int64(l2ChainID)),
		concurrentFetcher: cfg.ConcurrentFetchers,
		batchHeaders:      make(chan *BatchesRef, ParallelBlocks),

		done: make(chan struct{}, 1),
	}, nil
}

// GetFetchedBlockNumber returns the last fetched L1 block number.
func (f *Fetcher) GetFetchedBlockNumber() uint64 {
	return f.lastSyncedL1BlockNumber.Load()
}

// InitFetch inits the fetcher context.
func (f *Fetcher) InitFetch() {
	f.ctx, f.cancel = context.WithCancel(context.Background())
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
			blockNumber, err := f.l1Client.GetFinalizedBlockNumber()
			if err != nil {
				return err
			}
			lastSyncedL1BlockNumber := f.lastSyncedL1BlockNumber.Load()
			nextBlockNumber := lastSyncedL1BlockNumber + ParallelBlocks
			if blockNumber < nextBlockNumber {
				nextBlockNumber = blockNumber
			}
			if lastSyncedL1BlockNumber > nextBlockNumber {
				time.Sleep(fetchInterval)
				continue
			}
			ti := time.Now()
			batches, err := f.sequencerInbox.fetchBatchTransactions(big.NewInt(int64(lastSyncedL1BlockNumber)), big.NewInt(int64(nextBlockNumber)))
			if err != nil {
				return err
			}
			telemetry.MeasureSince(ti, "rpc", "l1_filter_logs")

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
					rawMsg, err = f.fetchBlock(batch.BlockNumber, batch.TxHash)
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
				if _, err := f.sequencerInbox.parseL2Transactions(batch); err != nil {
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

// getL2BlockHashes gets the L2 block hashes for the given range.
// The range is [start, end].
func (f *Fetcher) getL2BlockHashes(start, end uint64) ([]*sequencerv2types.BlockHeader, error) {
	ti := time.Now()
	defer telemetry.MeasureSince(ti, "rpc", "fetch_l2_block_hashes")

	g, ctx := errgroup.WithContext(context.Background())
	g.SetLimit(f.concurrentFetcher)
	hashesMap := sync.Map{}
	for i := start; i <= end; i += ParallelBlocks {
		if err := ctx.Err(); err != nil {
			logger.Errorf("fetch l2 block context error: %v", err)
			return nil, err
		}
		startBlockNumber := i
		endBlockNumber := i + ParallelBlocks
		if endBlockNumber > end {
			endBlockNumber = (end + 1)
		}
		g.Go(func() error {
			blockHashes, err := f.l2Client.GetBlockHashesByRange(startBlockNumber, endBlockNumber)
			if err != nil {
				return err
			}
			for i, hash := range blockHashes {
				hashesMap.Store(startBlockNumber+uint64(i), hash)
			}
			return nil
		})
	}
	if err := g.Wait(); err != nil {
		return nil, err
	}
	blockHashes := make([]*sequencerv2types.BlockHeader, 0, end-start+1)
	for i := start; i <= end; i++ {
		hash, ok := hashesMap.Load(i)
		if !ok {
			return nil, errors.New("block hash is not found")
		}
		blockHashes = append(blockHashes, &sequencerv2types.BlockHeader{
			BlockNumber: i,
			BlockHash:   hash.(common.Hash).Hex(),
		})
	}
	return blockHashes, nil
}

// Stop stops the Fetcher.
func (f *Fetcher) Stop() {
	f.StopFetch()
	// close the batch headers channel to notify the outside
	close(f.batchHeaders)
	close(f.done)
}

// StopFetch stops the Fetching logic.
func (f *Fetcher) StopFetch() {
	if f.cancel == nil {
		return
	}

	f.lastSyncedL1BlockNumber.Store(0)
	// close L1 fetcher
	f.cancel()

	func() {
		// wait for the fetcher to finish
		ctx, cancel := context.WithTimeout(context.Background(), fetchInterval*5)
		defer cancel()
		for {
			select {
			case <-f.done: // wait for the fetcher to finish
				return
			case <-ctx.Done():
				panic("failed to stop the fetcher")
			default:
				time.Sleep(10 * time.Millisecond)
				// drain channel, if the `batchHeaders` channel is full, it will block the fetcher
				// and the fetcher will not stop.
				for len(f.batchHeaders) > 0 {
					<-f.batchHeaders
				}
			}
		}
	}()
	// drain channel to clean up the batches while stopping the fetcher
	for len(f.batchHeaders) > 0 {
		<-f.batchHeaders
	}

	f.cancel = nil
	f.ctx = nil
}

// fetchBlock fetches the given block and analyzes the transactions
// which are sent to the BatchInbox EOA.
func (f *Fetcher) fetchBlock(blockNumber uint64, txHash common.Hash) ([]byte, error) {
	block, err := f.l1Client.GetBlockByNumber(blockNumber)
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
		blobs, err := f.l1BlobFetcher.GetBlobs(utils.GetContext(), blockRef, hashes)
		if err != nil {
			logger.Errorf("failed to get blobs: %v", err)
			return nil, err
		}
		telemetry.MeasureSince(ti, "rpc", "fetch_beacon_blobs")
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
	l2Blocks, err := f.getL2BlockHashes(batchesRef.FromL2BlockNumber, batchesRef.ToL2BlockNumber)
	if err != nil {
		return nil, err
	}
	header.L2Blocks = l2Blocks

	return &header, nil
}

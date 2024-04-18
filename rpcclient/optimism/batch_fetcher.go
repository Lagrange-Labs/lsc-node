package optimism

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"sort"
	"sync"
	"sync/atomic"
	"time"

	"github.com/ethereum-optimism/optimism/op-node/rollup/derive"
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
	"github.com/Lagrange-Labs/lagrange-node/utils"
)

const (
	EthereumFinalityDepth = 64
	ParallelBlocks        = 32
	cacheLimit            = 1024
	maxTxBlobCount        = 10000
	FetchInterval         = 5 * time.Second
)

// TxDataRef is a the list of transaction data with tx metadata.
type TxDataRef struct {
	Data    []byte
	TxType  uint8
	TxHash  common.Hash
	TxIndex int
}

// BatchesRef is a the list of batches with the L1 metadata.
type BatchesRef struct {
	Batches       []L2BlockBatch
	L1BlockNumber uint64
	L2BlockNumber uint64
	L1TxHash      common.Hash
	L1TxIndex     int
	L2BlockCount  int
}

// FramesRef is a the list of frames with the L1 metadata.
type FramesRef struct {
	Frames        []derive.Frame
	L1BlockNumber uint64
	L1TxHash      common.Hash
	TxIndex       int
}

// Fetcher is a synchronizer for the BatchInbox EOA.
type Fetcher struct {
	l1Client          *ethclient.Client
	l2Client          types.EvmClient
	l1BlobFetcher     *sources.L1BeaconClient
	batchInboxAddress common.Address
	batchSender       common.Address
	concurrentFetcher int
	signer            coretypes.Signer
	l2BlockCache      *utils.Cache
	batchHeaders      chan *BatchesRef

	lastSyncedL1BlockNumber atomic.Uint64
	lastSyncedL2BlockNumber uint64

	// decoder
	chFramesRef chan *FramesRef
	chainID     *big.Int

	mtx    sync.Mutex
	ctx    context.Context
	cancel context.CancelFunc
	done   chan struct{}
	chErr  chan error
}

// NewFetcher creates a new Fetcher instance.
func NewFetcher(cfg *Config) (*Fetcher, error) {
	l1Client, err := ethclient.Dial(cfg.L1RPCURL)
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

	return &Fetcher{
		l1Client:          l1Client,
		l2Client:          l2Client,
		l1BlobFetcher:     l1BlobFetcher,
		chainID:           big.NewInt(int64(l2ChainID)),
		batchInboxAddress: common.HexToAddress(cfg.BatchInbox),
		batchSender:       common.HexToAddress(cfg.BatchSender),
		concurrentFetcher: cfg.ConcurrentFetchers,
		signer:            coretypes.LatestSignerForChainID(chainID),
		l2BlockCache:      utils.NewCache(cacheLimit),
		batchHeaders:      make(chan *BatchesRef, 64),

		chErr: make(chan error, 1),
		done:  make(chan struct{}, 3),
	}, nil
}

// GetFetchedBlockNumber returns the last fetched L1 block number.
func (f *Fetcher) GetFetchedBlockNumber() uint64 {
	return f.lastSyncedL1BlockNumber.Load()
}

// InitFetch inits the fetcher context.
func (f *Fetcher) InitFetch(l2BlockNumber uint64) {
	f.chFramesRef = make(chan *FramesRef, 64)
	f.ctx, f.cancel = context.WithCancel(context.Background())
	f.lastSyncedL2BlockNumber = l2BlockNumber
}

// Fetch fetches the block data from the Ethereum and analyzes the
// transactions which are sent to the BatchInbox EOA.
func (f *Fetcher) Fetch(l1BeginBlockNumber uint64) error {
	go func() {
		if err := f.handleFrames(); err != nil {
			logger.Errorf("failed to handle frames: %v", err)
			f.chErr <- err
		}
		logger.Infof("decoder is stopped")
	}()

	defer func() {
		f.done <- struct{}{}
	}()

	f.lastSyncedL1BlockNumber.Store(l1BeginBlockNumber)

	for {
		select {
		case <-f.ctx.Done():
			logger.Infof("fetcher is stopped")
			return nil
		case err := <-f.chErr:
			return err
		default:
			g, ctx := errgroup.WithContext(context.Background())
			g.SetLimit(f.concurrentFetcher)
			// Fetch the latest finalized block number.
			blockNumber, err := f.l1Client.BlockNumber(ctx)
			if err != nil {
				return err
			}
			lastSyncedL1BlockNumber := f.lastSyncedL1BlockNumber.Load()
			nextBlockNumber := lastSyncedL1BlockNumber + ParallelBlocks
			if blockNumber-EthereumFinalityDepth < nextBlockNumber {
				nextBlockNumber = blockNumber - EthereumFinalityDepth
			}
			if lastSyncedL1BlockNumber >= nextBlockNumber {
				time.Sleep(FetchInterval)
				continue
			}
			m := sync.Map{}
			for i := lastSyncedL1BlockNumber; i < nextBlockNumber; i++ {
				if err := ctx.Err(); err != nil {
					logger.Errorf("fetch context error: %v", err)
					return err
				}

				number := i
				g.Go(func() error {
					res, err := f.fetchBlock(ctx, number)
					if err != nil {
						return err
					}
					for _, ref := range res {
						m.Store(fmt.Sprintf("%x_%d", ref.L1TxHash, ref.TxIndex), ref)
					}
					return nil
				})
			}
			if err := g.Wait(); err != nil {
				return err
			}
			framesRefs := make([]*FramesRef, 0)
			m.Range(func(_, ref interface{}) bool {
				framesRefs = append(framesRefs, ref.(*FramesRef))
				return true
			})
			sort.Slice(framesRefs, func(i, j int) bool {
				if framesRefs[i].L1BlockNumber == framesRefs[j].L1BlockNumber {
					return framesRefs[i].TxIndex < framesRefs[j].TxIndex
				}
				return framesRefs[i].L1BlockNumber < framesRefs[j].L1BlockNumber
			})
			for _, framesRef := range framesRefs {
				f.chFramesRef <- framesRef
			}
			f.lastSyncedL1BlockNumber.Store(nextBlockNumber)
		}
	}
}

// FetchL2Blocks fetches the L2 blocks from the given L2 block number.
func (f *Fetcher) FetchL2Blocks() error {
	defer func() {
		f.done <- struct{}{}
	}()

	l2BeginBlockNumber := f.lastSyncedL2BlockNumber
	for {
		select {
		case <-f.ctx.Done():
			logger.Infof("l2 fetcher is stopped")
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
	// close batch decoder
	close(f.chFramesRef)
	<-f.done
	// close batch headers channel to notify the outside
	close(f.batchHeaders)
	f.cancel = nil
	f.ctx = nil
}

// fetchBlock fetches the given block and analyzes the transactions
// which are sent to the BatchInbox EOA.
func (f *Fetcher) fetchBlock(ctx context.Context, blockNumber uint64) ([]*FramesRef, error) {
	block, err := f.l1Client.BlockByNumber(ctx, big.NewInt(int64(blockNumber)))
	if err != nil {
		return nil, err
	}

	res := make([]*FramesRef, 0)
	txDatas := make([]TxDataRef, 0)
	var hashes []eth.IndexedBlobHash
	blobIndex := 0

	for i, tx := range block.Transactions() {
		if !f.validTransaction(tx) {
			blobIndex += len(tx.BlobHashes())
			continue
		}

		if tx.Type() != coretypes.BlobTxType {
			txDatas = append(txDatas, TxDataRef{
				Data:    tx.Data(),
				TxType:  tx.Type(),
				TxHash:  tx.Hash(),
				TxIndex: i * maxTxBlobCount,
			})
		} else {
			if len(tx.Data()) > 0 {
				logger.Warnf("blob tx has calldata which will be ignored %v", tx.Hash().Hex())
			}
			for bi, hash := range tx.BlobHashes() {
				hashes = append(hashes, eth.IndexedBlobHash{
					Index: uint64(blobIndex),
					Hash:  hash,
				})
				blobIndex++
				txDatas = append(txDatas, TxDataRef{
					Data:    nil,
					TxType:  tx.Type(),
					TxHash:  tx.Hash(),
					TxIndex: i*maxTxBlobCount + bi,
				})
			}
		}
	}

	if len(hashes) > 0 {
		blockRef := eth.L1BlockRef{
			Number:     blockNumber,
			Hash:       block.Hash(),
			ParentHash: block.ParentHash(),
			Time:       block.Time(),
		}
		blobs, err := f.l1BlobFetcher.GetBlobs(ctx, blockRef, hashes)
		if err != nil {
			logger.Errorf("failed to get blobs: %v", err)
			return nil, err
		}
		if len(blobs) != len(hashes) {
			logger.Errorf("blobs length is not matched: %d, %d", len(blobs), len(hashes))
			return nil, fmt.Errorf("blobs length is not matched: %d, %d", len(blobs), len(hashes))
		}
		blobIndex := 0
		for i := range txDatas {
			if txDatas[i].TxType == coretypes.BlobTxType {
				logger.Infof("L1 blob tx is loaded from %d TxIndex: %d TxHash: %v", blockNumber, txDatas[i].TxIndex, txDatas[i].TxHash.Hex())
				data, err := blobs[blobIndex].ToData()
				if err != nil {
					logger.Errorf("failed to convert blob data: %v", err)
					return nil, err
				}
				txDatas[i].Data = data
				blobIndex++
			}
		}
	}
	for _, data := range txDatas {
		frames, err := derive.ParseFrames(data.Data)
		if err != nil {
			logger.Errorf("failed to parse frames: %v", err)
			return nil, err
		}
		framesRef := &FramesRef{
			Frames:        frames,
			L1BlockNumber: blockNumber,
			L1TxHash:      data.TxHash,
			TxIndex:       data.TxIndex,
		}
		res = append(res, framesRef)
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

// getL2BlockNumberByHash returns the L2 block number for the given block hash.
func (f *Fetcher) getL2BlockNumberByHash(blockHash common.Hash) (uint64, error) {
	return f.l2Client.GetBlockNumberByHash(blockHash)
}

// getL2BlockNumberByTxHash returns the L2 block number for the given transaction hash.
func (f *Fetcher) getL2BlockNumberByTxHash(txHash common.Hash) (uint64, error) {
	return f.l2Client.GetBlockNumberByTxHash(txHash)
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

	l2Blocks := make([]*sequencerv2types.BlockHeader, 0)
	blockNumberIndex := batchesRef.L2BlockNumber
	for _, batch := range batchesRef.Batches {
		for i := uint64(0); i < uint64(batch.BlockCount); i++ {
			blockHash, err := f.getL2BlockHash(blockNumberIndex)
			if err != nil {
				return nil, err
			}
			l2Blocks = append(l2Blocks, &sequencerv2types.BlockHeader{
				BlockNumber: blockNumberIndex,
				BlockHash:   blockHash.Hex(),
			})
			blockNumberIndex++
		}
	}
	header.L2Blocks = l2Blocks

	return &header, nil
}

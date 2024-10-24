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
	"github.com/ethereum/go-ethereum/log"
	"golang.org/x/sync/errgroup"

	"github.com/Lagrange-Labs/lsc-node/core"
	"github.com/Lagrange-Labs/lsc-node/core/logger"
	"github.com/Lagrange-Labs/lsc-node/core/telemetry"
	"github.com/Lagrange-Labs/lsc-node/rpcclient/evmclient"
	"github.com/Lagrange-Labs/lsc-node/rpcclient/types"
	sequencerv2types "github.com/Lagrange-Labs/lsc-node/sequencer/types/v2"
)

const (
	searchLimit             = 1024
	maxTxBlobCount          = 10000
	fetchInterval           = 5 * time.Second
	getL2BatchHeaderTimeout = 30 * time.Second
	iterateWarningLevel     = 10
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
	l1Client          types.EvmClient
	l2Client          types.EvmClient
	l1BlobFetcher     *sources.L1BeaconClient
	batchInboxAddress common.Address
	batchSenders      []string
	concurrentFetcher int
	l1ParallelBlocks  int
	l2ParallelBlocks  int
	signer            coretypes.Signer
	batchHeaders      chan *BatchesRef

	isLight                 bool
	lastSyncedL1BlockNumber atomic.Uint64
	lastPulledL1BlockNumber atomic.Uint64
	lastSyncedL2BlockNumber atomic.Uint64

	// decoder
	chFramesRef chan *FramesRef
	chainID     *big.Int

	ctx         context.Context
	cancel      context.CancelFunc
	fetcherDone chan struct{}
	decoderDone chan struct{}
	chErr       chan error
}

// NewFetcher creates a new Fetcher instance.
func NewFetcher(cfg *Config, isLight bool) (*Fetcher, error) {
	l1Client, err := evmclient.NewClient(cfg.L1RPCURLs)
	if err != nil {
		return nil, err
	}
	l2Client, err := evmclient.NewClient(cfg.RPCURLs)
	if err != nil {
		return nil, err
	}
	chainID, err := l1Client.GetChainID()
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
		isLight:           isLight,
		chainID:           big.NewInt(int64(l2ChainID)),
		batchInboxAddress: common.HexToAddress(cfg.BatchInbox),
		batchSenders:      cfg.BatchSender,
		concurrentFetcher: cfg.ConcurrentFetchers,
		l1ParallelBlocks:  cfg.L1ParallelBlocks,
		l2ParallelBlocks:  cfg.L2ParallelBlocks,
		signer:            coretypes.LatestSignerForChainID(big.NewInt(int64(chainID))),
		batchHeaders:      make(chan *BatchesRef, cfg.L1ParallelBlocks),

		chErr:       make(chan error, 1),
		fetcherDone: make(chan struct{}, 1),
		decoderDone: make(chan struct{}, 1),
	}, nil
}

// GetFetchedBlockNumber returns the last fetched L1 block number.
func (f *Fetcher) GetFetchedBlockNumber() uint64 {
	return f.lastSyncedL1BlockNumber.Load()
}

// GetPulledBlockNumber returns the last pulled batch L1 block number.
func (f *Fetcher) GetPulledBlockNumber() uint64 {
	return f.lastPulledL1BlockNumber.Load()
}

// InitFetch inits the fetcher context.
//
// NOTE: This should be called before calling Fetch after the Stop.
func (f *Fetcher) InitFetch() {
	f.chFramesRef = make(chan *FramesRef, f.l1ParallelBlocks)
	f.ctx, f.cancel = context.WithCancel(context.Background())

	go func() {
		if err := f.handleFrames(); err != nil {
			logger.Errorf("failed to handle frames: %v", err)
			f.chErr <- err
		}
	}()
}

// GetL2BatchHeader returns the next L2 batch header for the given L1 block number
// and the Tx Hash
func (f *Fetcher) GetL2BatchHeader(l1BlockNumber, l2BlockNumber uint64, txHash string) (*sequencerv2types.BatchHeader, error) {
	chBatchHeader := make(chan *sequencerv2types.BatchHeader, 1) // buffered to avoid blocking

	go func() {
		defer close(chBatchHeader)
		for {
			batchHeader, err := f.nextBatchHeader()
			if err != nil {
				f.chErr <- err
				return
			}
			if batchHeader.L1BlockNumber == l1BlockNumber {
				if len(txHash) > 0 && batchHeader.L1TxHash == txHash {
					chBatchHeader <- batchHeader
					return
				}
				if l2BlockNumber > 0 && batchHeader.L2FromBlockNumber == l2BlockNumber {
					chBatchHeader <- batchHeader
					return
				}
			}
		}
	}()

	checkTxHash := func(ctx context.Context) (*sequencerv2types.BatchHeader, error) {
		for {
			select {
			case <-ctx.Done():
				return nil, nil
			case err := <-f.chErr:
				return nil, fmt.Errorf("batch decoder err: %v", err)
			case batchHeader := <-chBatchHeader:
				return batchHeader, nil
			}
		}
	}

	iterL1 := l1BlockNumber
	for {
		if l2BlockNumber > 0 {
			f.lastSyncedL2BlockNumber.Store(l2BlockNumber - 1)
		}

		if f.lastSyncedL1BlockNumber.Load() == l1BlockNumber {
			bh, err := checkTxHash(core.GetContextWithTimeout(getL2BatchHeaderTimeout))
			if err != nil {
				return nil, err
			}
			if bh != nil {
				return bh, nil
			}
			iterL1 += 1
			continue
		}

		frames, err := f.fetchBlock(iterL1)
		if err != nil {
			return nil, err
		}
		for _, framesRef := range frames {
			f.chFramesRef <- framesRef
		}

		if len(frames) == 0 {
			iterL1 += 1
			continue
		}

		bh, err := checkTxHash(core.GetContextWithTimeout(getL2BatchHeaderTimeout))
		if err != nil {
			return nil, err
		}
		if bh != nil {
			f.lastSyncedL1BlockNumber.Store(iterL1)
			return bh, nil
		}

		iterL1 += 1
		if (iterL1-l1BlockNumber)%iterateWarningLevel == iterateWarningLevel-1 {
			logger.Warnf("no batch header found for L1 block number: %d Iteration: %d", l1BlockNumber, iterL1)
		}
		if iterL1-l1BlockNumber > searchLimit {
			return nil, errors.New("batch header not found after searching the limit")
		}
	}
}

// Fetch fetches the block data from the Ethereum and analyzes the
// transactions which are sent to the BatchInbox EOA.
func (f *Fetcher) Fetch(l1BeginBlockNumber, l2BeginBlockNumber uint64) error {
	f.lastSyncedL2BlockNumber.Store(l2BeginBlockNumber)

	defer func() {
		f.fetcherDone <- struct{}{}
		logger.Infof("fetcher is stopped")
	}()

	f.lastSyncedL1BlockNumber.Store(l1BeginBlockNumber)

	for {
		select {
		case <-f.ctx.Done():
			return nil
		case err := <-f.chErr:
			return err
		default:
			g, ctx := errgroup.WithContext(context.Background())
			g.SetLimit(f.concurrentFetcher)
			// Fetch the latest finalized block number.
			blockNumber, err := f.l1Client.GetFinalizedBlockNumber()
			if err != nil {
				return err
			}
			lastSyncedL1BlockNumber := f.lastSyncedL1BlockNumber.Load()
			nextBlockNumber := lastSyncedL1BlockNumber + uint64(f.l1ParallelBlocks)
			if blockNumber < nextBlockNumber {
				nextBlockNumber = blockNumber
			}
			if lastSyncedL1BlockNumber >= nextBlockNumber {
				time.Sleep(fetchInterval)
				continue
			}
			ti := time.Now()
			m := sync.Map{}
			for i := lastSyncedL1BlockNumber; i <= nextBlockNumber; i++ {
				if err := ctx.Err(); err != nil {
					logger.Errorf("fetch context error: %v", err)
					return err
				}

				number := i
				g.Go(func() error {
					res, err := f.fetchBlock(number)
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
			telemetry.MeasureSince(ti, "rpc", "fetch_l1_blocks")
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
			f.lastSyncedL1BlockNumber.Store(nextBlockNumber + 1)
		}
	}
}

// getL2BlockHeaders gets the L2 block headers for the given range.
// The range is [start, end].
func (f *Fetcher) getL2BlockHeaders(start, end uint64) ([]*sequencerv2types.BlockHeader, error) {
	ti := time.Now()
	defer telemetry.MeasureSince(ti, "rpc", "fetch_l2_block_headers")

	g, ctx := errgroup.WithContext(context.Background())
	g.SetLimit(f.concurrentFetcher)
	headersMap := sync.Map{}
	for i := start; i <= end; i += uint64(f.l2ParallelBlocks) {
		if err := ctx.Err(); err != nil {
			logger.Errorf("fetch l2 block context error: %v", err)
			return nil, err
		}
		startBlockNumber := i
		endBlockNumber := i + uint64(f.l2ParallelBlocks)
		if endBlockNumber > end {
			endBlockNumber = (end + 1)
		}
		g.Go(func() error {
			blockHeaders, err := f.l2Client.GetBlockHeadersByRange(startBlockNumber, endBlockNumber)
			if err != nil {
				return err
			}
			for i := range blockHeaders {
				headersMap.Store(startBlockNumber+uint64(i), &blockHeaders[i])
			}
			return nil
		})
	}
	if err := g.Wait(); err != nil {
		return nil, err
	}
	blockHeaders := make([]*sequencerv2types.BlockHeader, end-start+1)
	for i := start; i <= end; i++ {
		blockHeader, ok := headersMap.Load(i)
		if !ok {
			return nil, errors.New("block hash is not found")
		}
		blockHeaders[i-start] = blockHeader.(*sequencerv2types.BlockHeader)
	}
	return blockHeaders, nil
}

// Stop stops the Fetcher.
func (f *Fetcher) Stop() {
	f.StopFetch()

	// close the batch headers channel to notify the outside
	close(f.batchHeaders)
	close(f.fetcherDone)
	close(f.decoderDone)
	close(f.chErr)
}

// StopFetch stops the fetching logic.
func (f *Fetcher) StopFetch() {
	if f.cancel == nil {
		return
	}

	f.lastSyncedL1BlockNumber.Store(0)

	// close L1 fetcher
	f.cancel()

	func() {
		doneCount := 0
		// wait for the fetcher to finish
		ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
		defer cancel()
		for {
			select {
			case <-f.fetcherDone: // wait for the fetcher to finish
				// close batch decoder
				close(f.chFramesRef)
				doneCount++
			case <-f.decoderDone: // wait for the batch decoder to finish
				doneCount++
			case <-ctx.Done():
				panic("failed to stop the fetcher")
			default:
				if doneCount == 2 {
					return
				}
				time.Sleep(10 * time.Millisecond)
				// drain channel, if the `batchHeaders` channel is full, it will block the fetcher
				// and the fetcher will not stop.
				for len(f.batchHeaders) > 0 {
					<-f.batchHeaders
				}
				for len(f.chFramesRef) > 0 { // drain channel
					<-f.chFramesRef
				}
			}
		}
	}()
	// drain channel to clean up the batches while stopping the fetcher
	for len(f.batchHeaders) > 0 {
		<-f.batchHeaders
	}
	for len(f.chFramesRef) > 0 { // drain channel
		<-f.chFramesRef
	}

	f.cancel = nil
	f.ctx = nil
}

// StopDecoder stops the decoder.
func (f *Fetcher) StopDecoder() {
	if f.cancel == nil {
		return
	}

	f.lastSyncedL1BlockNumber.Store(0)
	close(f.chFramesRef)

	func() {
		ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
		defer cancel()
		for {
			select {
			case <-f.decoderDone: // wait for the batch decoder to finish
				return
			case <-ctx.Done():
				panic("failed to stop the decoder")
			default:
				time.Sleep(10 * time.Millisecond)
			}
		}
	}()

	// drain channel to clean up the batches while stopping the fetcher
	for len(f.batchHeaders) > 0 {
		<-f.batchHeaders
	}
	for len(f.chFramesRef) > 0 { // drain channel
		<-f.chFramesRef
	}

	f.cancel = nil
	f.ctx = nil
}

// fetchBlock fetches the given block and analyzes the transactions
// which are sent to the BatchInbox EOA.
func (f *Fetcher) fetchBlock(blockNumber uint64) ([]*FramesRef, error) {
	block, err := f.l1Client.GetBlockByNumber(blockNumber)
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
		ti := time.Now()
		blobs, err := f.l1BlobFetcher.GetBlobs(core.GetContext(), blockRef, hashes)
		if err != nil {
			logger.Errorf("failed to get blobs: %v", err)
			return nil, err
		}
		telemetry.MeasureSince(ti, "rpc", "fetch_beacon_blobs")
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
			continue
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
	for _, sender := range f.batchSenders {
		if from == common.HexToAddress(sender) {
			return true
		}
	}

	return false
}

// nextBatchHeader returns the L2 batch header.
func (f *Fetcher) nextBatchHeader() (*sequencerv2types.BatchHeader, error) {
	batchesRef, ok := <-f.batchHeaders
	if !ok {
		return nil, errors.New("batch headers channel is closed")
	}
	defer f.lastPulledL1BlockNumber.Store(batchesRef.L1BlockNumber)

	header := sequencerv2types.BatchHeader{
		L1BlockNumber:     batchesRef.L1BlockNumber,
		L1TxHash:          batchesRef.L1TxHash.Hex(),
		L1TxIndex:         uint32(batchesRef.L1TxIndex),
		ChainId:           uint32(f.chainID.Uint64()),
		L2FromBlockNumber: batchesRef.L2BlockNumber,
		L2ToBlockNumber:   batchesRef.L2BlockNumber + uint64(batchesRef.L2BlockCount) - 1,
	}

	if f.isLight {
		lastHash, err := f.l2Client.GetBlockHashByNumber(header.L2ToBlockNumber)
		if err != nil {
			return nil, err
		}
		header.L2Blocks = []*sequencerv2types.BlockHeader{
			{
				BlockNumber: header.L2ToBlockNumber,
				BlockHash:   lastHash.Hex(),
			},
		}
		return &header, nil
	}
	l2Blocks := make([]*sequencerv2types.BlockHeader, 0)
	fromBlockNumber := batchesRef.L2BlockNumber
	for _, batch := range batchesRef.Batches {
		toBlockNumber := fromBlockNumber + uint64(batch.BlockCount) - 1
		l2BlockHeaders, err := f.getL2BlockHeaders(fromBlockNumber, toBlockNumber)
		if err != nil {
			return nil, err
		}
		l2Blocks = append(l2Blocks, l2BlockHeaders...)
		fromBlockNumber = toBlockNumber + 1
	}
	header.L2Blocks = l2Blocks

	return &header, nil
}

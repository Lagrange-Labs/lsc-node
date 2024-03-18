package optimism

import (
	"context"
	"fmt"
	"math/big"
	"sort"
	"sync"
	"time"

	"github.com/ethereum-optimism/optimism/op-node/rollup/derive"
	"github.com/ethereum/go-ethereum/common"
	coretypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/sync/errgroup"

	"github.com/Lagrange-Labs/lagrange-node/logger"
	sequencerv2types "github.com/Lagrange-Labs/lagrange-node/sequencer/types/v2"
	"github.com/Lagrange-Labs/lagrange-node/utils"
)

const (
	ConcurrentFetchers    = 4
	EthereumFinalityDepth = 64
	cacheLimit            = 2048
	CacheFullWaitInterval = 10 * time.Microsecond
	FetchInterval         = 5 * time.Second
)

// BatchesRef is a the list of batches with the L1 metadata.
type BatchesRef struct {
	Batches       []L2BlockBatch
	L1BlockNumber uint64
	L1TxHash      common.Hash
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
	l2Client          *ethclient.Client
	batchInboxAddress common.Address
	batchSender       common.Address
	signer            coretypes.Signer
	batchCache        *utils.Cache
	blockCache        *utils.Cache
	// decoder
	chFramesRef             chan *FramesRef
	lastSyncedL2BlockNumber uint64
	pendingBatchesRefs      []*BatchesRef
	chainID                 *big.Int

	ctx    context.Context
	cancel context.CancelFunc
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
	l2ChainID, err := l2Client.ChainID(context.Background())
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithCancel(context.Background())
	return &Fetcher{
		l1Client:          l1Client,
		l2Client:          l2Client,
		chainID:           l2ChainID,
		batchInboxAddress: common.HexToAddress(cfg.BatchInbox),
		batchSender:       common.HexToAddress(cfg.BatchSender),
		signer:            coretypes.LatestSignerForChainID(chainID),
		batchCache:        utils.NewCache(cacheLimit),
		blockCache:        utils.NewCache(cacheLimit),

		chFramesRef: make(chan *FramesRef, 64),

		ctx:    ctx,
		cancel: cancel,
	}, nil
}

// GetFetchedBlockNumber returns the last fetched block number.
func (f *Fetcher) GetFetchedBlockNumber() uint64 {
	return f.batchCache.GetHeadKey()
}

// Fetch fetches the block data from the Ethereum and analyzes the
// transactions which are sent to the BatchInbox EOA.
func (f *Fetcher) Fetch(l1BeginBlockNumber uint64) error {
	go func() {
		if err := f.handleFrames(); err != nil {
			logger.Errorf("failed to handle frames: %v", err)
		}
		logger.Infof("decoder is stopped")
	}()

	lastSyncedL1BlockNumber := l1BeginBlockNumber

	for {
		select {
		case <-f.ctx.Done():
			return nil
		default:
			g, ctx := errgroup.WithContext(context.Background())
			g.SetLimit(ConcurrentFetchers)
			// Fetch the latest finalized block number.
			blockNumber, err := f.l1Client.BlockNumber(ctx)
			if err != nil {
				return err
			}
			nextBlockNumber := lastSyncedL1BlockNumber + EthereumFinalityDepth
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
					logger.Errorf("context error: %v", err)
					return err
				}

				number := i
				g.Go(func() error {
					res, err := f.fetchBlock(ctx, number)
					if err != nil {
						return err
					}
					for _, ref := range res {
						m.Store(ref.L1TxHash, ref)
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
			lastSyncedL1BlockNumber = nextBlockNumber
		}
	}
}

// Stop stops the Fetcher.
func (f *Fetcher) Stop() {
	f.cancel()
	close(f.chFramesRef)
}

// fetchBlock fetches the given block and analyzes the transactions
// which are sent to the BatchInbox EOA.
func (f *Fetcher) fetchBlock(ctx context.Context, blockNumber uint64) ([]*FramesRef, error) {
	block, err := f.l1Client.BlockByNumber(ctx, big.NewInt(int64(blockNumber)))
	if err != nil {
		return nil, err
	}

	res := make([]*FramesRef, 0)
	for i, tx := range block.Transactions() {
		if !f.validTransaction(tx) {
			continue
		}
		frames, err := derive.ParseFrames(tx.Data())
		if err != nil {
			logger.Errorf("failed to parse frames: %v", err)
			return nil, err
		}
		framesRef := &FramesRef{
			Frames:        frames,
			L1BlockNumber: blockNumber,
			L1TxHash:      tx.Hash(),
			TxIndex:       i,
		}
		res = append(res, framesRef)
	}

	return res, nil
}

// getL2BlockHash returns the L2 block hash for the given block number.
func (f *Fetcher) getL2BlockHash(blockNumber uint64) (common.Hash, error) {
	raw, ok := f.blockCache.Get(blockNumber)
	if !ok {
		block, err := f.l2Client.BlockByNumber(context.Background(), big.NewInt(int64(blockNumber)))
		if err != nil {
			return common.Hash{}, err
		}
		hash := block.Hash()
		f.blockCache.Set(blockNumber, hash)
		return hash, nil
	}
	return raw.(common.Hash), nil
}

// getL2BlockNumberByHash returns the L2 block number for the given block hash.
func (f *Fetcher) getL2BlockNumberByHash(blockHash common.Hash) (uint64, error) {
	block, err := f.l2Client.BlockByHash(context.Background(), blockHash)
	if err != nil {
		return 0, err
	}
	number := block.NumberU64()
	f.blockCache.Set(number, blockHash)
	return number, nil
}

// getL2BlockNumberByTxHash returns the L2 block number for the given transaction hash.
func (f *Fetcher) getL2BlockNumberByTxHash(txHash common.Hash) (uint64, error) {
	receipt, err := f.l2Client.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		return 0, err
	}

	blockNumber := receipt.BlockNumber.Uint64()
	f.blockCache.Set(blockNumber, receipt.BlockHash)
	return blockNumber, nil
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
	if len(tx.Data()) == 0 {
		return false
	}
	return true
}

// getL2BatchData returns the L2 batch data for the given L2 block number.
func (f *Fetcher) getL2BatchData(blockNumber uint64) (*sequencerv2types.BatchHeader, error) {
	raw, ok := f.batchCache.Get(blockNumber)
	if !ok {
		return nil, fmt.Errorf("no batch data")
	}
	batchesRef := raw.(*BatchesRef)
	header := sequencerv2types.BatchHeader{
		L1BlockNumber: batchesRef.L1BlockNumber,
		L1TxHash:      batchesRef.L1TxHash.Hex(),
		BatchNumber:   blockNumber,
		ChainId:       uint32(f.chainID.Uint64()),
	}

	l2Blocks := make([]*sequencerv2types.BlockHeader, 0)
	blockNumberIndex := blockNumber
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

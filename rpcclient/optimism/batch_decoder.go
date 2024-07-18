package optimism

import (
	"bytes"
	"errors"
	"fmt"
	"io"

	"github.com/ethereum-optimism/optimism/op-node/rollup/derive"
	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/Lagrange-Labs/lagrange-node/logger"
)

var (
	// ErrNoTransaction is the error for no transaction.
	ErrNoTransaction = errors.New("no transaction")
)

// L2BlockBatch represents an L2 block batch.
type L2BlockBatch struct {
	// ParentHash is the parent hash of the first block.
	ParentHash common.Hash
	// ParentHashCheck is the first 20 bytes of parent hash of the first block for the check.
	ParentHashCheck []byte
	// TxHash is the transaction hash of the first transaction.
	TxHash common.Hash
	// BlockCount is the number of L2 blocks in the batch.
	BlockCount int
}

// handleFrames returns BatchData items from the given frames.
func (f *Fetcher) handleFrames() error {
	var (
		channels = make(map[derive.ChannelID]struct {
			Channel   *derive.Channel
			L1TxHash  common.Hash
			L1TxIndex int
		})
	)

	defer func() {
		f.decoderDone <- struct{}{}
		logger.Infof("decoder is stopped")
	}()

	for framesRef := range f.chFramesRef {
		blockRef := eth.L1BlockRef{Number: framesRef.L1BlockNumber}
		for _, frame := range framesRef.Frames {
			if _, ok := channels[frame.ID]; !ok {
				channels[frame.ID] = struct {
					Channel   *derive.Channel
					L1TxHash  common.Hash
					L1TxIndex int
				}{
					derive.NewChannel(frame.ID, blockRef),
					framesRef.L1TxHash,
					framesRef.TxIndex,
				}
			}
			ch := channels[frame.ID]

			if ch.Channel.IsReady() {
				logger.Errorf("Invaild Frame: channel %v is ready", frame.ID)
				break
			}

			if err := ch.Channel.AddFrame(frame, blockRef); err != nil {
				logger.Errorf("Failed to add frame to channel %v : %v", frame.ID, err)
				continue
			}

			if ch.Channel.IsReady() {
				// Optimism Fjord upgrade
				br, err := derive.BatchReader(ch.Channel.Reader(), 100_000_000, true)
				if err != nil {
					logger.Errorf("Failed to create zlib reader: %v", err)
					return err
				}
				batchesRef := &BatchesRef{
					L1BlockNumber: ch.Channel.OpenBlockNumber(),
					L1TxHash:      ch.L1TxHash,
					L1TxIndex:     ch.L1TxIndex,
					Batches:       make([]L2BlockBatch, 0),
					L2BlockCount:  0,
				}
				for batchData, err := br(); err != io.EOF; batchData, err = br() {
					if err != nil {
						logger.Errorf("Failed to read batch data: %v", err)
						return err
					}
					batch, err := f.parseBatch(batchData)
					if err != nil {
						return err
					}
					batchesRef.Batches = append(batchesRef.Batches, *batch)
					batchesRef.L2BlockCount += batch.BlockCount
				}
				delete(channels, frame.ID)
				if err := f.pushBatchesRef(batchesRef); err != nil {
					return err
				}
			}
		}
	}

	return nil
}

// parseBatch parses the batch data and returns the L2 block batch.
func (f *Fetcher) parseBatch(batchData *derive.BatchData) (*L2BlockBatch, error) {
	batchType := batchData.GetBatchType()
	switch batchType {
	case derive.SingularBatchType:
		batch, err := derive.GetSingularBatch(batchData)
		if err != nil {
			logger.Errorf("Failed to get singular batch: %v", err)
			return nil, err
		}
		return &L2BlockBatch{
			ParentHash:      batch.ParentHash,
			ParentHashCheck: batch.ParentHash[:20],
			BlockCount:      1,
		}, nil
	case derive.SpanBatchType:
		batch, err := derive.DeriveSpanBatch(batchData, 1, 1, f.chainID)
		if err != nil {
			logger.Errorf("Failed to derive span batch: %v", err)
			return nil, err
		}
		var txHash common.Hash
		for _, b := range batch.Batches {
			isTx := false
			for _, txData := range b.Transactions {
				var tx types.Transaction
				if err := tx.UnmarshalBinary(txData); err != nil {
					logger.Errorf("Failed to unmarshal transaction: %v", err)
					return nil, err
				}
				txHash = tx.Hash()
				isTx = true
				break
			}
			if isTx {
				break
			}
		}
		return &L2BlockBatch{
			ParentHashCheck: batch.ParentCheck[:],
			TxHash:          txHash,
			BlockCount:      len(batch.Batches),
		}, nil
	default:
		logger.Errorf("Unsupported batch type: %+v", batchData)
		return nil, fmt.Errorf("unsupported batch type: %+v", batchData)
	}
}

// pushBatch pushes the L2 block batch to the cache.
func (f *Fetcher) pushBatchesRef(batchesRef *BatchesRef) error {
	pushCache := func(l2BlockNumber uint64, batchesRef *BatchesRef) error {
		// check the parent hash of the first block is correct
		parentHash, err := f.l2Client.GetBlockHashByNumber(l2BlockNumber - 1)
		if err != nil {
			logger.Errorf("failed to get L2 block hash: %v", err)
			return err
		}
		if bytes.Equal(batchesRef.Batches[0].ParentHashCheck, parentHash[:20]) {
			logger.Infof("L2 batch is loaded from %d BlockCount: %d L1 BlockNumber:%d TxHash: %v", l2BlockNumber, batchesRef.L2BlockCount, batchesRef.L1BlockNumber, batchesRef.L1TxHash.Hex())
			batchesRef.L2BlockNumber = l2BlockNumber
			f.batchHeaders <- batchesRef
			f.lastSyncedL2BlockNumber = l2BlockNumber + uint64(batchesRef.L2BlockCount) - 1
			return nil
		} else {
			logger.Errorf("parent hash mismatch L2 BlockNumber: %d, Parent Hash: %v, Ref: %+v", l2BlockNumber, parentHash, batchesRef)
			return fmt.Errorf("parent hash mismatch")
		}
	}

	// determine the last synced L2 block number
	bn, err := f.getBeginL2BlockNumber(batchesRef)
	if err == ErrNoTransaction {
		// no transaction in the batch
		logger.Warnf("no transaction in the batch: %+v", batchesRef)
		for i := 0; i < searchLimit; i++ {
			bn := f.lastSyncedL2BlockNumber + uint64(i)
			hash, err := f.l2Client.GetBlockHashByNumber(bn)
			if err != nil {
				logger.Errorf("failed to get L2 block hash: %v", err)
				return err
			}
			if bytes.Equal(batchesRef.Batches[0].ParentHashCheck, hash[:20]) {
				return pushCache(bn+1, batchesRef)
			}
			if i == 0 {
				continue
			}
			bn = f.lastSyncedL2BlockNumber - uint64(i)
			hash, err = f.l2Client.GetBlockHashByNumber(bn)
			if err != nil {
				logger.Errorf("failed to get L2 block hash: %v", err)
				return err
			}
			if bytes.Equal(batchesRef.Batches[0].ParentHashCheck, hash[:20]) {
				return pushCache(bn+1, batchesRef)
			}
		}
		logger.Errorf("no L2 block number found for the batch L1BlockNumber: %d , L1TxHash: %s", batchesRef.L1BlockNumber, batchesRef.L1TxHash.Hex())
		return fmt.Errorf("no L2 block number found")
	} else if err != nil {
		return err
	}

	return pushCache(bn+1, batchesRef)
}

// getBeginL2BlockNumber returns the begin L2 block number for the given BatchesRef.
func (f *Fetcher) getBeginL2BlockNumber(batchesRef *BatchesRef) (uint64, error) {
	l2BlockNumber := uint64(0)
	forwardCount := uint64(0)
	var err error
	for _, batch := range batchesRef.Batches {
		forwardCount += uint64(batch.BlockCount)
		if batch.ParentHash.Cmp((common.Hash{})) != 0 { // singular batch
			l2BlockNumber, err = f.getL2BlockNumberByHash(batch.ParentHash)
			if err != nil {
				logger.Errorf("failed to get L2 block number by block hash: %v", err)
				return 0, err
			}
			break
		}
		if batch.TxHash.Cmp((common.Hash{})) != 0 {
			l2BlockNumber, err = f.getL2BlockNumberByTxHash(batch.TxHash)
			if err != nil {
				logger.Errorf("failed to get L2 block number by tx hash: %v", err)
				return 0, err
			}
			break
		}
	}
	if l2BlockNumber == 0 {
		logger.Warnf("no L2 block number found: %+v", batchesRef)
		return 0, ErrNoTransaction
	}
	for bn := l2BlockNumber; bn >= l2BlockNumber-forwardCount; bn-- {
		blockHash, err := f.l2Client.GetBlockHashByNumber(bn)
		if err != nil {
			logger.Errorf("failed to get L2 block hash: %v", err)
			return 0, err
		}
		if bytes.Equal(batchesRef.Batches[0].ParentHashCheck, blockHash[:20]) {
			return bn, nil
		}
	}

	logger.Errorf("no L2 block number found parentHashCheck %x from l2BlockNumber %d forwardCount %d", batchesRef.Batches[0].ParentHashCheck, l2BlockNumber, forwardCount)
	return 0, fmt.Errorf("no L2 block number found")
}

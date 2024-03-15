package optimism

import (
	"fmt"
	"io"
	"math/big"

	"github.com/ethereum-optimism/optimism/op-node/rollup/derive"
	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/Lagrange-Labs/lagrange-node/logger"
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
func handleFrames(blockNumber uint64, chainID *big.Int, frames []derive.Frame) ([]L2BlockBatch, error) {
	var (
		batches         []L2BlockBatch
		framesByChannel = make(map[derive.ChannelID][]derive.Frame)
	)

	for _, frame := range frames {
		framesByChannel[frame.ID] = append(framesByChannel[frame.ID], frame)
	}

	blockRef := eth.L1BlockRef{Number: blockNumber}
	for channelID, frames := range framesByChannel {
		ch := derive.NewChannel(channelID, blockRef)
		if ch.IsReady() {
			logger.Errorf("Invaild Frame: channel %v is ready", channelID)
			return nil, fmt.Errorf("Invaild Frame: channel %v is ready", channelID)
		}
		for _, frame := range frames {
			if ch.IsReady() {
				logger.Errorf("Invaild Frame: channel %v is ready", channelID)
				return nil, fmt.Errorf("Invaild Frame: channel %v is ready", channelID)
			}
			if err := ch.AddFrame(frame, blockRef); err != nil {
				logger.Errorf("Failed to add frame: %v", err)
				return nil, err
			}
		}
		if ch.IsReady() {
			br, err := derive.BatchReader(ch.Reader())
			if err != nil {
				logger.Errorf("Failed to create zlib reader: %v", err)
				return nil, err
			}
			for batchData, err := br(); err != io.EOF; batchData, err = br() {
				if err != nil {
					logger.Errorf("Failed to read batch data: %v", err)
					return nil, err
				}
				batchType := batchData.GetBatchType()
				switch batchType {
				case derive.SingularBatchType:
					batch, err := derive.GetSingularBatch(batchData)
					if err != nil {
						logger.Errorf("Failed to get singular batch: %v", err)
						return nil, err
					}
					batches = append(batches, L2BlockBatch{
						ParentHash:      batch.ParentHash,
						ParentHashCheck: batch.ParentHash[:20],
						BlockCount:      1,
					})
				case derive.SpanBatchType:
					batch, err := derive.DeriveSpanBatch(batchData, 1, 1, chainID)
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
					batches = append(batches, L2BlockBatch{
						ParentHashCheck: batch.ParentCheck[:],
						TxHash:          txHash,
						BlockCount:      len(batch.Batches),
					})
				default:
					logger.Errorf("Unsupported batch type: %+v", batchData)
					return nil, fmt.Errorf("Unsupported batch type: %+v", batchData)
				}
			}
		}
	}

	return batches, nil
}

package optimism

import (
	"compress/zlib"
	"fmt"
	"io"

	"github.com/Lagrange-Labs/lagrange-node/logger"
	"github.com/ethereum-optimism/optimism/op-node/rollup/derive"
	"github.com/ethereum-optimism/optimism/op-service/eth"
	"github.com/ethereum/go-ethereum/rlp"
)

// handleFrames returns BatchData items from the given frames.
func handleFrames(blockNumber uint64, frames []derive.Frame) ([]derive.SingularBatch, error) {
	var (
		batches         []derive.SingularBatch
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
			if err := ch.AddFrame(frame, blockRef); err != nil {
				logger.Errorf("Failed to add frame: %v", err)
				return nil, err
			}
		}
		if ch.IsReady() {
			zr, err := zlib.NewReader(ch.Reader())
			if err != nil {
				logger.Errorf("Failed to create zlib reader: %v", err)
				return nil, err
			}
			rlpReader := rlp.NewStream(zr, derive.MaxRLPBytesPerChannel)
			for {
				v, err := rlpReader.Bytes()
				if err != nil {
					if err == io.EOF {
						break
					}
				}
				var batch derive.SingularBatch
				if uint8(v[0]) != derive.SingularBatchType {
					logger.Errorf("Invalid batch type: %v", uint8(v[0]))
					return nil, fmt.Errorf("Invalid batch type: %v", uint8(v[0]))
				}
				if err := rlp.DecodeBytes(v[1:], &batch); err != nil {
					logger.Errorf("Failed to decode batch: %v", err)
					return nil, err
				}
				batches = append(batches, batch)
			}
		}
	}

	return batches, nil
}

package optimism

import (
	"fmt"
	"io"

	"github.com/ethereum-optimism/optimism/op-node/rollup/derive"
	"github.com/ethereum-optimism/optimism/op-service/eth"
)

// handleFrames returns BatchData items from the given frames.
func handleFrames(blockNumber uint64, frames []derive.Frame) ([]*derive.BatchData, error) {
	var (
		batches         []*derive.BatchData
		framesByChannel = make(map[derive.ChannelID][]derive.Frame)
	)

	for _, frame := range frames {
		framesByChannel[frame.ID] = append(framesByChannel[frame.ID], frame)
	}

	blockRef := eth.L1BlockRef{Number: blockNumber}
	for channelID, frames := range framesByChannel {
		ch := derive.NewChannel(channelID, blockRef)
		if ch.IsReady() {
			return nil, fmt.Errorf("Invaild Frame: channel %v is ready", channelID)
		}
		for _, frame := range frames {
			if err := ch.AddFrame(frame, blockRef); err != nil {
				return nil, err
			}
		}
		if ch.IsReady() {
			br, err := derive.BatchReader(ch.Reader())
			if err != nil {
				return nil, err
			}
			for batch, err := br(); err != io.EOF; batch, err = br() {
				if err != nil {
					return nil, err
				}
				batches = append(batches, batch)
			}
		}
	}

	return batches, nil
}

package v2

// BatchNumber returns the batch number of the batch.
func (b *Batch) BatchNumber() uint64 {
	return b.BatchHeader.BatchNumber
}

// ChainID returns the chain ID of the batch.
func (b *Batch) ChainID() uint32 {
	return b.BatchHeader.ChainId
}

// BatchHash returns the hash of the batch.
func (b *Batch) BatchHash() string {
	return b.BatchHeader.BatchHash
}

// L1BlockNumber returns the L1 block number of the batch.
func (b *Batch) L1BlockNumber() uint64 {
	return b.BatchHeader.L1BlockNumber
}

// L1TxHash returns the L1 transaction hash of the batch.
func (b *Batch) L1TxHash() string {
	return b.BatchHeader.L1TxHash
}

// L2BlockHash returns the hash of the block with the given block number.
func (b *Batch) BlockHash(blockNumber uint64) string {
	for _, block := range b.BatchHeader.L2Blocks {
		if block.BlockNumber == blockNumber {
			return block.BlockHash
		}
	}

	return ""
}

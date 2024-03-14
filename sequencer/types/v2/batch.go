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

// BlsSignature returns the BLS signature of the batch.
func (b *Batch) BlsSignature() *BlsSignature {
	return &BlsSignature{
		BatchHeader:     b.BatchHeader,
		CommitteeHeader: b.CommitteeHeader,
	}
}

// FromBlockNumber returns the block number of the first block in the batch header.
func (bh *BatchHeader) FromBlockNumber() uint64 {
	if len(bh.L2Blocks) == 0 {
		return 0
	}

	return bh.L2Blocks[0].BlockNumber
}

// ToBlockNumber returns the block number of the last block in the batch header.
func (bh *BatchHeader) ToBlockNumber() uint64 {
	if len(bh.L2Blocks) == 0 {
		return 0
	}

	return bh.L2Blocks[len(bh.L2Blocks)-1].BlockNumber
}

// Hash returns the hash of the batch header.
func (bh *BatchHeader) Hash() []byte {
	return nil // TODO: implement
}

// GetBatchNumber returns the batch number of the bls signature.
func (b *BlsSignature) GetBatchNumber() uint64 {
	return b.BatchHeader.BatchNumber
}

// Hash returns the hash of the bls signature.
func (b *BlsSignature) Hash() []byte {
	return nil // TODO: implement
}

// CommitHash returns the hash of the commit bls signature.
func (b *BlsSignature) CommitHash() []byte {
	return nil // TODO: implement
}

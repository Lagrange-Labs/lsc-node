package v2

import (
	"github.com/Lagrange-Labs/lagrange-node/core"
	"github.com/Lagrange-Labs/lagrange-node/core/crypto"
)

// BatchNumber returns the batch number of the batch.
func (b *Batch) BatchNumber() uint64 {
	if b.BatchHeader == nil {
		return 0
	}
	return b.BatchHeader.BatchNumber
}

// ChainID returns the chain ID of the batch.
func (b *Batch) ChainID() uint32 {
	if b.BatchHeader == nil {
		return 0
	}
	return b.BatchHeader.ChainId
}

// L1BlockNumber returns the L1 block number of the batch.
func (b *Batch) L1BlockNumber() uint64 {
	if b.BatchHeader == nil {
		return 0
	}
	return b.BatchHeader.L1BlockNumber
}

// L1TxHash returns the L1 transaction hash of the batch.
func (b *Batch) L1TxHash() string {
	if b.BatchHeader == nil {
		return ""
	}
	return b.BatchHeader.L1TxHash
}

// L2BlockHash returns the hash of the block with the given block number.
func (b *Batch) BlockHash(blockNumber uint64) string {
	if b.BatchHeader == nil {
		return ""
	}

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

// CurrentCommittee returns the current committee root of the batch.
func (b *Batch) CurrentCommittee() string {
	if b.CommitteeHeader == nil {
		return ""
	}
	return b.CommitteeHeader.CurrentCommittee
}

// NextCommittee returns the next committee root of the batch.
func (b *Batch) NextCommittee() string {
	if b.CommitteeHeader == nil {
		return ""
	}
	return b.CommitteeHeader.NextCommittee
}

// TotalVotingPower returns the total voting power of the batch.
func (b *Batch) TotalVotingPower() uint64 {
	if b.CommitteeHeader == nil {
		return 0
	}
	return b.CommitteeHeader.TotalVotingPower
}

// FromBlockNumber returns the block number of the first block in the batch header.
func (bh *BatchHeader) FromBlockNumber() uint64 {
	if bh.L2FromBlockNumber > 0 {
		return bh.L2FromBlockNumber
	}

	if len(bh.L2Blocks) == 0 {
		return 0
	}

	return bh.L2Blocks[0].BlockNumber
}

// ToBlockNumber returns the block number of the last block in the batch header.
func (bh *BatchHeader) ToBlockNumber() uint64 {
	if bh.L2ToBlockNumber > 0 {
		return bh.L2ToBlockNumber
	}

	if len(bh.L2Blocks) == 0 {
		return 0
	}

	return bh.L2Blocks[len(bh.L2Blocks)-1].BlockNumber
}

// Hash returns the hash of the batch header.
func (bh *BatchHeader) Hash() []byte {
	h := append([]byte{}, core.Hex2Bytes(bh.L1TxHash)...)
	h = append(h, core.Uint64ToBytes(bh.L1BlockNumber)...)
	for _, block := range bh.L2Blocks {
		h = append(h, core.Uint64ToBytes(block.BlockNumber)...)
		h = append(h, core.Hex2Bytes(block.BlockHash)...)
	}
	return crypto.Hash(h)
}

// MerkleHash returns the hash of the batch header.
func (bh *BatchHeader) MerkleHash() []byte {
	h := append([]byte{}, core.Uint64ToBytes(uint64(bh.ChainId))...)
	h = append(h, core.Hex2Bytes(bh.L1TxHash)...)
	h = append(h, core.Uint64ToBytes(bh.L1BlockNumber)...)
	h = append(h, core.Uint64ToBytes(bh.FromBlockNumber())...)
	h = append(h, core.Uint64ToBytes(bh.ToBlockNumber())...)
	hashes := make([][]byte, 0, len(bh.L2Blocks))
	for _, block := range bh.L2Blocks {
		hashes = append(hashes, core.Hex2Bytes(block.BlockHash))
	}
	h = append(h, crypto.MerkleRoot(hashes)...)
	return crypto.Hash(h)
}

// BatchNumber returns the batch number of the bls signature.
func (b *BlsSignature) BatchNumber() uint64 {
	if b.BatchHeader == nil {
		return 0
	}
	return b.BatchHeader.BatchNumber
}

// Hash returns the hash of the bls signature.
func (b *BlsSignature) Hash() []byte {
	h := append([]byte{}, b.BatchHeader.MerkleHash()...)
	h = append(h, core.Hex2Bytes(b.CommitteeHeader.CurrentCommittee)...)
	h = append(h, core.Hex2Bytes(b.CommitteeHeader.NextCommittee)...)
	h = append(h, core.Uint64ToBytes(b.CommitteeHeader.TotalVotingPower)...)
	return crypto.Hash(h)
}

// CommitHash returns the hash of the commit bls signature.
func (b *BlsSignature) CommitHash() []byte {
	h := append([]byte{}, b.Hash()...)
	h = append(h, core.Hex2Bytes(b.BlsSignature)...)
	return crypto.Hash(h)
}

// CurrentCommittee returns the current committee root of the bls signature.
func (b *BlsSignature) CurrentCommittee() string {
	if b.CommitteeHeader == nil {
		return ""
	}
	return b.CommitteeHeader.CurrentCommittee
}

// NextCommittee returns the next committee root of the bls signature.
func (b *BlsSignature) NextCommittee() string {
	if b.CommitteeHeader == nil {
		return ""
	}
	return b.CommitteeHeader.NextCommittee
}

// TotalVotingPower returns the total voting power of the bls signature.
func (b *BlsSignature) TotalVotingPower() uint64 {
	if b.CommitteeHeader == nil {
		return 0
	}
	return b.CommitteeHeader.TotalVotingPower
}

// Clone returns a clone of the bls signature.
// NOTE: Only used for testing.
func (b *BlsSignature) Clone() *BlsSignature {
	return &BlsSignature{
		BatchHeader: &BatchHeader{
			BatchNumber:   b.BatchHeader.BatchNumber,
			L1BlockNumber: b.BatchHeader.L1BlockNumber,
			L1TxHash:      b.BatchHeader.L1TxHash,
			ChainId:       b.BatchHeader.ChainId,
			L2Blocks:      b.BatchHeader.L2Blocks,
		},
		CommitteeHeader: &CommitteeHeader{
			CurrentCommittee: b.CommitteeHeader.CurrentCommittee,
			NextCommittee:    b.CommitteeHeader.NextCommittee,
			TotalVotingPower: b.CommitteeHeader.TotalVotingPower,
		},
	}
}

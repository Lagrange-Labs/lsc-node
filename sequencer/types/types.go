package types

import (
	"encoding/binary"
	"math/big"

	"github.com/Lagrange-Labs/lagrange-node/utils"
	"github.com/ethereum/go-ethereum/common"
)

// BlockHash returns the block hash of the chain header.
func (b *Block) BlockHash() string {
	return b.ChainHeader.BlockHash
}

// BlockNumber returns the L2 block number of the chain header.
func (b *Block) BlockNumber() uint64 {
	return b.ChainHeader.BlockNumber
}

// L1TxHash returns the L1 tx hash of the chain header.
func (b *Block) L1TxHash() common.Hash {
	return common.HexToHash(b.ChainHeader.L1TxHash)
}

// L1BlockNumber returns the L1 block number of the chain header.
func (b *Block) L1BlockNumber() uint64 {
	return b.ChainHeader.L1BlockNumber
}

// TotalVotingPower returns the total voting power of the block.
func (b *Block) TotalVotingPower() uint64 {
	return b.BlockHeader.TotalVotingPower
}

// CurrentCommittee returns the current committee of the block.
func (b *Block) CurrentCommittee() string {
	return b.BlockHeader.CurrentCommittee
}

// NextCommittee returns the next committee of the block.
func (b *Block) NextCommittee() string {
	return b.BlockHeader.NextCommittee
}

// ProposerPubKey returns the proposer public key of the block.
func (b *Block) ProposerPubKey() string {
	if b.BlockHeader == nil {
		return ""
	}
	return b.BlockHeader.ProposerPubKey
}

// ProposerSignature returns the proposer signature of the block.
func (b *Block) ProposerSignature() string {
	return b.BlockHeader.ProposerSignature
}

// BlsSignature returns the bls signature of the block.
func (b *Block) BlsSignature() *BlsSignature {
	return &BlsSignature{
		ChainHeader:      b.ChainHeader,
		CurrentCommittee: b.CurrentCommittee(),
		NextCommittee:    b.NextCommittee(),
	}
}

// Hash returns the hash of the chain header.
func (c *ChainHeader) Hash() []byte {
	var blockNumberBuf common.Hash
	blockHash := common.FromHex(c.BlockHash)[:]
	blockNumber := big.NewInt(int64(c.BlockNumber)).FillBytes(blockNumberBuf[:])
	chainID := make([]byte, 4)
	binary.BigEndian.PutUint32(chainID, c.ChainId)
	return utils.Hash(blockHash, blockNumber, chainID)
}

// Hash returns the hash of the bls signature.
func (b *BlsSignature) Hash() []byte {
	currentCommitteeRoot := common.FromHex(b.CurrentCommittee)[:]
	nextCommitteeRoot := common.FromHex(b.NextCommittee)[:]
	chainHash := b.ChainHeader.Hash()

	return utils.PoseidonHash(chainHash, currentCommitteeRoot, nextCommitteeRoot)
}

// BlockNumber returns the block number of the bls signature.
func (b *BlsSignature) BlockNumber() uint64 {
	return b.ChainHeader.BlockNumber
}

// L1BlockNumber returns the L1 block number of the bls signature.
func (b *BlsSignature) L1BlockNumber() uint64 {
	return b.ChainHeader.L1BlockNumber
}

// Clone returns a clone of the bls signature.
func (b *BlsSignature) Clone() *BlsSignature {
	return &BlsSignature{
		ChainHeader:      b.ChainHeader,
		CurrentCommittee: b.CurrentCommittee,
		NextCommittee:    b.NextCommittee,
	}
}

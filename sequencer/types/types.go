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

// BlockNumber returns the block number of the chain header.
func (b *Block) BlockNumber() uint64 {
	return b.ChainHeader.BlockNumber
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
	return b.BlockHeader.ProposerPubKey
}

// ProposerSignature returns the proposer signature of the block.
func (b *Block) ProposerSignature() string {
	return b.BlockHeader.ProposerSignature
}

// EpochNumber returns the epoch number of the block.
func (b *Block) EpochNumber() uint64 {
	return b.BlockHeader.EpochNumber
}

// BlsSignature returns the bls signature of the block.
func (b *Block) BlsSignature() *BlsSignature {
	return &BlsSignature{
		ChainHeader:      b.ChainHeader,
		CurrentCommittee: b.CurrentCommittee(),
		NextCommittee:    b.NextCommittee(),
	}
}

// Hash returns the hash of the bls signature.
func (b *BlsSignature) Hash() []byte {
	var blockNumberBuf, tvpBuf common.Hash
	blockHash := common.FromHex(b.ChainHeader.BlockHash)[:]
	currentCommitteeRoot := common.FromHex(b.CurrentCommittee)[:]
	nextCommitteeRoot := common.FromHex(b.NextCommittee)[:]
	blockNumber := big.NewInt(int64(b.ChainHeader.BlockNumber)).FillBytes(blockNumberBuf[:])
	tvp := big.NewInt(int64(b.TotalVotingPower)).FillBytes(tvpBuf[:])
	chainID := make([]byte, 4)
	binary.BigEndian.PutUint32(chainID, b.ChainHeader.ChainId)

	return utils.Hash(
		blockHash,
		currentCommitteeRoot,
		nextCommitteeRoot,
		blockNumber,
		tvp,
		chainID,
	)
}

// BlockNumber returns the block number of the bls signature.
func (b *BlsSignature) BlockNumber() uint64 {
	return b.ChainHeader.BlockNumber
}

// Clone returns a clone of the bls signature.
func (b *BlsSignature) Clone() *BlsSignature {
	return &BlsSignature{
		ChainHeader:      b.ChainHeader,
		CurrentCommittee: b.CurrentCommittee,
		NextCommittee:    b.NextCommittee,
	}
}

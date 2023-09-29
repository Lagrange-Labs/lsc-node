package types

import (
	"encoding/binary"
	"math/big"

	sequencertypes "github.com/Lagrange-Labs/lagrange-node/sequencer/types"
	"github.com/Lagrange-Labs/lagrange-node/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

// Evidence defines an evidence.
type Evidence struct {
	Operator             string   `json:"operator" bson:"operator"`
	BlockHash            [32]byte `json:"block_hash" bson:"block_hash"`
	CurrentCommitteeRoot [32]byte `json:"current_committee_root" bson:"current_committee_root"`
	NextCommitteeRoot    [32]byte `json:"next_committee_root" bson:"next_committee_root"`
	BlockNumber          uint64   `json:"block_number" bson:"block_number"`
	EpochBlockNumber     uint64   `json:"epoch_block_number" bson:"epoch_block_number"`
	BlockSignature       []byte   `json:"block_signature" bson:"block_signature"`
	CommitSignature      []byte   `json:"commit_signature" bson:"commit_signature"`
	ChainID              uint32   `json:"chain_id" bson:"chain_id"`
	Status               bool     `json:"status" bson:"status"`
}

// GetCommitRequestHash returns the hash of the commit block request.
func GetCommitRequestHash(sig *sequencertypes.BlsSignature) []byte {
	var blockNumberBuf, epochNumberBuf common.Hash
	blockHash := common.FromHex(sig.ChainHeader.BlockHash)[:]
	currentCommitteeRoot := common.FromHex(sig.CurrentCommittee)[:]
	nextCommitteeRoot := common.FromHex(sig.NextCommittee)[:]
	blockNumber := big.NewInt(int64(sig.BlockNumber())).FillBytes(blockNumberBuf[:])
	epochNumber := big.NewInt(int64(sig.EpochBlockNumber)).FillBytes(epochNumberBuf[:])
	blockSignature := common.FromHex(sig.BlsSignature)[:]
	chainID := make([]byte, 4)
	binary.BigEndian.PutUint32(chainID, sig.ChainHeader.ChainId)

	return utils.Hash(
		blockHash,
		currentCommitteeRoot,
		nextCommitteeRoot,
		blockNumber,
		epochNumber,
		blockSignature,
		chainID,
	)
}

// GetEvidence returns the evidence from the commit block request.
func GetEvidence(sig *sequencertypes.BlsSignature) (*Evidence, error) {
	hash := GetCommitRequestHash(sig)
	signature := common.FromHex(sig.EcdsaSignature)
	pubKey, err := crypto.SigToPub(hash, signature)
	if err != nil {
		return nil, err
	}
	// convert the signature to the legacy format which be able to be verified in Solidity
	if signature[64] == 0 || signature[64] == 1 {
		signature[64] += 27
	}
	addr := crypto.PubkeyToAddress(*pubKey).Hex()
	return &Evidence{
		Operator:             addr,
		BlockHash:            common.HexToHash(sig.ChainHeader.BlockHash),
		CurrentCommitteeRoot: common.HexToHash(sig.CurrentCommittee),
		NextCommitteeRoot:    common.HexToHash(sig.NextCommittee),
		BlockNumber:          sig.BlockNumber(),
		EpochBlockNumber:     sig.EpochBlockNumber,
		BlockSignature:       common.FromHex(sig.BlsSignature),
		CommitSignature:      signature,
		ChainID:              sig.ChainHeader.ChainId,
	}, nil
}

// GetBlsSignature returns the bls signature from the evidence.
func GetBlsSignature(evidence *Evidence) *sequencertypes.BlsSignature {
	return &sequencertypes.BlsSignature{
		ChainHeader: &sequencertypes.ChainHeader{
			ChainId:     evidence.ChainID,
			BlockHash:   common.Bytes2Hex(evidence.BlockHash[:]),
			BlockNumber: evidence.BlockNumber,
		},
		CurrentCommittee: common.Bytes2Hex(evidence.CurrentCommitteeRoot[:]),
		NextCommittee:    common.Bytes2Hex(evidence.NextCommitteeRoot[:]),
		BlsSignature:     common.Bytes2Hex(evidence.BlockSignature),
		EpochBlockNumber: evidence.EpochBlockNumber,
		EcdsaSignature:   common.Bytes2Hex(evidence.CommitSignature),
	}
}

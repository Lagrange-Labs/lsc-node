package types

import (
	"encoding/binary"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/Lagrange-Labs/lagrange-node/core"
	corecrypto "github.com/Lagrange-Labs/lagrange-node/core/crypto"
	"github.com/Lagrange-Labs/lagrange-node/core/logger"
	sequencertypes "github.com/Lagrange-Labs/lagrange-node/sequencer/types"
	sequencerv2types "github.com/Lagrange-Labs/lagrange-node/sequencer/types/v2"
)

// TODO: refactor the evidence to use the new sequencer types
// Evidence defines an evidence.
type Evidence struct {
	Operator             string   `json:"operator" bson:"operator"`
	Signer               string   `json:"signer" bson:"signer"`
	BlsPublicKey         string   `json:"bls_public_key" bson:"bls_public_key"`
	BlockHash            [32]byte `json:"block_hash" bson:"block_hash"`
	CurrentCommitteeRoot [32]byte `json:"current_committee_root" bson:"current_committee_root"`
	NextCommitteeRoot    [32]byte `json:"next_committee_root" bson:"next_committee_root"`
	BlockNumber          uint64   `json:"block_number" bson:"block_number"`
	L1BlockNumber        uint64   `json:"l1_block_number" bson:"l1_block_number"`
	BlockSignature       []byte   `json:"block_signature" bson:"block_signature"`
	CommitSignature      []byte   `json:"commit_signature" bson:"commit_signature"`
	ChainID              uint32   `json:"chain_id" bson:"chain_id"`
	Status               bool     `json:"status" bson:"status"`
}

// GetCommitRequestHash returns the hash of the commit block request.
func GetCommitRequestHash(sig *sequencertypes.BlsSignature) []byte {
	var blockNumberBuf, l1BlockNumberBuf common.Hash
	blockHash := common.FromHex(sig.ChainHeader.BlockHash)[:]
	currentCommitteeRoot := common.FromHex(sig.CurrentCommittee)[:]
	nextCommitteeRoot := common.FromHex(sig.NextCommittee)[:]
	blockNumber := big.NewInt(int64(sig.BlockNumber())).FillBytes(blockNumberBuf[:])
	l1BlockNumber := big.NewInt(int64(sig.L1BlockNumber())).FillBytes(l1BlockNumberBuf[:])
	blockSignature := core.Hex2Bytes(sig.BlsSignature)
	chainID := make([]byte, 4)
	binary.BigEndian.PutUint32(chainID, sig.ChainHeader.ChainId)

	return corecrypto.Hash(
		blockHash,
		currentCommitteeRoot,
		nextCommitteeRoot,
		blockNumber,
		l1BlockNumber,
		blockSignature,
		chainID,
	)
}

// GetEvidence returns the evidence from the commit block request.
func GetEvidence(operator string, blsPubKey string, sig *sequencerv2types.BlsSignature) (*Evidence, error) {
	hash := sig.CommitHash()
	signature := common.FromHex(sig.EcdsaSignature)
	pubKey, err := crypto.SigToPub(hash, signature)
	if err != nil {
		logger.Errorf("failed to recover public key from signature: %v", err)
		return nil, err
	}
	// convert the signature to the legacy format which be able to be verified in Solidity
	if signature[64] == 0 || signature[64] == 1 {
		signature[64] += 27
	}
	addr := crypto.PubkeyToAddress(*pubKey).Hex()
	return &Evidence{
		Operator:             operator,
		Signer:               addr,
		BlsPublicKey:         blsPubKey,
		CurrentCommitteeRoot: common.HexToHash(sig.CurrentCommittee()),
		NextCommitteeRoot:    common.HexToHash(sig.NextCommittee()),
		BlockNumber:          sig.BatchNumber(),
		BlockSignature:       common.FromHex(sig.BlsSignature),
		CommitSignature:      signature,
		ChainID:              sig.BatchHeader.ChainId,
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

		EcdsaSignature: common.Bytes2Hex(evidence.CommitSignature),
	}
}

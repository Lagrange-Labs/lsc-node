package types

import (
	"encoding/binary"
	"math/big"

	evm "github.com/Lagrange-Labs/lagrange-node/rpcclient/evmclient"
	"github.com/Lagrange-Labs/lagrange-node/scinterface/lagrange"
	sequencertypes "github.com/Lagrange-Labs/lagrange-node/sequencer/types"
	"github.com/Lagrange-Labs/lagrange-node/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// Evidence defines an evidence.
type Evidence struct {
	Operator                    string   `json:"operator" bson:"operator"`
	BlockHash                   [32]byte `json:"block_hash" bson:"block_hash"`
	CorrectBlockHash            [32]byte `json:"correct_block_hash" bson:"correct_block_hash"`
	CurrentCommitteeRoot        [32]byte `json:"current_committee_root" bson:"current_committee_root"`
	CorrectCurrentCommitteeRoot [32]byte `json:"correct_current_committee_root" bson:"correct_current_committee_root"`
	NextCommitteeRoot           [32]byte `json:"next_committee_root" bson:"next_committee_root"`
	CorrectNextCommitteeRoot    [32]byte `json:"correct_next_committee_root" bson:"correct_next_committee_root"`
	BlockNumber                 uint64   `json:"block_number" bson:"block_number"`
	EpochBlockNumber            uint64   `json:"epoch_block_number" bson:"epoch_block_number"`
	BlockSignature              []byte   `json:"block_signature" bson:"block_signature"`
	CommitSignature             []byte   `json:"commit_signature" bson:"commit_signature"`
	ChainID                     uint32   `json:"chain_id" bson:"chain_id"`
	Status                      bool     `json:"status" bson:"status"`
	CorrectRawHeader	    []byte   
	CheckpointBlockHash	    [32]byte
	HeaderProof		    []byte
	ExtraData                   []byte     // network-specific data
}

// GetLagrangeServiceEvidence returns the lagrange service evidence.
func GetLagrangeServiceEvidence(e *Evidence) lagrange.EvidenceVerifierEvidence {
	return lagrange.EvidenceVerifierEvidence{
		Operator:                    common.HexToAddress(e.Operator),
		BlockHash:                   e.BlockHash,
		CorrectBlockHash:            e.CorrectBlockHash,
		CurrentCommitteeRoot:        e.CurrentCommitteeRoot,
		CorrectCurrentCommitteeRoot: e.CorrectCurrentCommitteeRoot,
		NextCommitteeRoot:           e.NextCommitteeRoot,
		CorrectNextCommitteeRoot:    e.CorrectNextCommitteeRoot,
		BlockNumber:                 big.NewInt(int64(e.BlockNumber)),
		EpochBlockNumber:            big.NewInt(int64(e.EpochBlockNumber)),
		BlockSignature:              e.BlockSignature,
		CommitSignature:             e.CommitSignature,
		ChainID:                     e.ChainID,
	}
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
func GetEvidence(sig *sequencertypes.BlsSignature, correctBlockHash, correctCurrentCommitteeRoot, correctNextCommitteeRoot string) (*Evidence, error) {
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

	rawHeader,err := evm.GetRawAttestBlockHeader(int(sig.BlockNumber()))
	if err != nil {
	    return nil,err
	}

	hex, l2hash, err := evm.GetExtraDataByNetwork(int(sig.BlockNumber()))
	if err != nil {
	    return nil,err
	}
	hexbytes, err := hexutil.Decode(hex)
	if err != nil {
	    return nil,err
	}
	headerbytes, err := hexutil.Decode(rawHeader)
	if err != nil {
	    return nil,err
	}

	return &Evidence{
		Operator:                    addr,
		BlockHash:                   common.HexToHash(sig.ChainHeader.BlockHash),
		CorrectBlockHash:            common.HexToHash(correctBlockHash),
		CurrentCommitteeRoot:        common.HexToHash(sig.CurrentCommittee),
		CorrectCurrentCommitteeRoot: common.HexToHash(correctCurrentCommitteeRoot),
		NextCommitteeRoot:           common.HexToHash(sig.NextCommittee),
		CorrectNextCommitteeRoot:    common.HexToHash(correctNextCommitteeRoot),
		BlockNumber:                 sig.BlockNumber(),
		EpochBlockNumber:            sig.EpochBlockNumber,
		BlockSignature:              common.FromHex(sig.BlsSignature),
		CommitSignature:             signature,
		ChainID:                     sig.ChainHeader.ChainId,
		Status:			     false, //TODO
		CorrectRawHeader:	     headerbytes,
		CheckpointBlockHash:	     l2hash,
		HeaderProof:    	     []byte(""),
		ExtraData:      	     hexbytes,
	}, nil
}
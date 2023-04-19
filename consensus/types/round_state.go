package types

import (
	"bytes"
	"fmt"

	"github.com/Lagrange-Labs/lagrange-node/logger"
	networktypes "github.com/Lagrange-Labs/lagrange-node/network/types"
	sequencertypes "github.com/Lagrange-Labs/lagrange-node/sequencer/types"
	"github.com/ethereum/go-ethereum/common"
	"github.com/umbracle/go-eth-consensus/bls"
	"google.golang.org/protobuf/proto"
)

// RoundState defines the internal consensus state.
type RoundState struct {
	Height        uint64
	Validators    *ValidatorSet
	ProposalBlock *sequencertypes.Block

	commitSignatures map[string]*networktypes.CommitBlockRequest
	evidences        []*networktypes.CommitBlockRequest // to determine slashing
}

// NewRoundState creates a new round state.
func NewRoundState(validators *ValidatorSet, proposalBlock *sequencertypes.Block) *RoundState {
	return &RoundState{
		Height:           proposalBlock.Header.BlockNumber,
		Validators:       validators,
		ProposalBlock:    proposalBlock,
		commitSignatures: make(map[string]*networktypes.CommitBlockRequest),
	}
}

// AddCommit adds a commit to the round state.
func (rs *RoundState) AddCommit(commit *networktypes.CommitBlockRequest) {
	rs.commitSignatures[commit.PubKey] = commit
}

// CheckEnoughVotingPower checks if there is enough voting power to finalize the block.
func (rs *RoundState) CheckEnoughVotingPower() bool {
	votingPower := uint64(0)
	for _, signature := range rs.commitSignatures {
		votingPower += rs.Validators.GetVotingPower(signature.PubKey)
	}

	return votingPower*3 > rs.Validators.TotalVotingPower*2
}

// CheckAggregatedSignature checks if the aggregated signature is valid.
func (rs *RoundState) CheckAggregatedSignature() ([]*bls.PublicKey, *bls.Signature, error) {
	sigMessage, err := proto.Marshal(&sequencertypes.Signature{
		ChainHeader:      rs.ProposalBlock.ChainHeader,
		CurrentCommittee: rs.ProposalBlock.Header.CurrentCommittee,
		NextCommittee:    rs.ProposalBlock.Header.NextCommittee,
	})
	if err != nil {
		logger.Fatalf("failed to marshal signature message: %v", err)
		return nil, nil, err
	}
	signatures := make([]*bls.Signature, 0)
	pubkeys := make([]*bls.PublicKey, 0)
	invalid_keys := make([]string, 0)
	for _, commit := range rs.commitSignatures {
		pubkey := new(bls.PublicKey)
		if err := pubkey.Deserialize(common.FromHex(commit.PubKey)); err != nil {
			return nil, nil, err
		}
		sig := new(bls.Signature)
		if err := sig.Deserialize(common.FromHex(commit.BlsSignature.Signature)); err != nil {
			logger.Errorf("failed to deserialize signature %v: %v", commit, err)
			invalid_keys = append(invalid_keys, commit.PubKey)
			continue
		}
		signatures = append(signatures, sig)
		pubkeys = append(pubkeys, pubkey)
	}
	aggSig := bls.AggregateSignatures(signatures)
	verified, err := aggSig.FastAggregateVerify(pubkeys, sigMessage)
	if err != nil || !verified {
		// find the invalid signature
		for i, pubkey := range pubkeys {
			pubkey_raw := pubkey.Serialize()
			commit := rs.commitSignatures[common.Bytes2Hex(pubkey_raw[:])]
			commitMessage, err := proto.Marshal(&sequencertypes.Signature{
				ChainHeader:      commit.BlsSignature.ChainHeader,
				CurrentCommittee: commit.BlsSignature.CurrentCommittee,
				NextCommittee:    commit.BlsSignature.NextCommittee,
			})
			if err != nil {
				logger.Fatalf("failed to marshal commit singature message %v: %v", commit, err)
				return nil, nil, err
			}
			if !bytes.Equal(commitMessage, sigMessage) {
				logger.Errorf("wrong commit message: %v, original: %v", common.Bytes2Hex(commitMessage), common.Bytes2Hex(sigMessage))
				invalid_keys = append(invalid_keys, common.Bytes2Hex(pubkey_raw[:]))
				continue
			}
			verified, err := signatures[i].VerifyByte(pubkey, commitMessage)
			if err != nil {
				return nil, nil, err
			}
			if !verified {
				logger.Errorf("invalid signature: %v", commit)
				invalid_keys = append(invalid_keys, common.Bytes2Hex(pubkey_raw[:]))
			}
		}
		err = ErrInvalidAggregativeSignature
	}
	// add invalid signatures to evidences
	for _, key := range invalid_keys {
		rs.evidences = append(rs.evidences, rs.commitSignatures[key])
		delete(rs.commitSignatures, key)
	}
	return pubkeys, aggSig, err
}

// GetEvidences returns the evidences.
func (rs *RoundState) GetEvidences() []*networktypes.CommitBlockRequest {
	return rs.evidences
}

var (
	ErrInvalidAggregativeSignature = fmt.Errorf("invalid aggregative signature")
)

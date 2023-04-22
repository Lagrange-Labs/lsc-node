package types

import (
	"bytes"
	"fmt"
	"sync"

	"github.com/Lagrange-Labs/lagrange-node/logger"
	networktypes "github.com/Lagrange-Labs/lagrange-node/network/types"
	sequencertypes "github.com/Lagrange-Labs/lagrange-node/sequencer/types"
	"github.com/Lagrange-Labs/lagrange-node/utils"
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

	rwMutex   *sync.RWMutex // to protect the round state updates
	isBlocked bool          // to prevent the block commit
}

// NewEmptyRoundState creates a new empty round state for rwMutex.
func NewEmptyRoundState() *RoundState {
	return &RoundState{
		rwMutex:   &sync.RWMutex{},
		isBlocked: true,
	}
}

// UpdateRoundState updates a new round state.
func (rs *RoundState) UpdateRoundState(validators *ValidatorSet, proposalBlock *sequencertypes.Block) {
	rs.rwMutex.Lock()
	defer rs.rwMutex.Unlock()

	rs.Height = proposalBlock.BlockNumber()
	rs.Validators = validators
	rs.ProposalBlock = proposalBlock
	rs.commitSignatures = make(map[string]*networktypes.CommitBlockRequest)
	rs.isBlocked = false
}

// AddCommit adds a commit to the round state.
func (rs *RoundState) AddCommit(commit *networktypes.CommitBlockRequest) {
	rs.rwMutex.Lock()
	defer rs.rwMutex.Unlock()

	if rs.isBlocked {
		logger.Warnf("the add commit is blocked")
		return
	}

	rs.commitSignatures[commit.PubKey] = commit
}

// BlockCommit blocks adds a commit to the round state.
func (rs *RoundState) BlockCommit() {
	rs.rwMutex.Lock()
	defer rs.rwMutex.Unlock()

	rs.isBlocked = true
}

// UnblockCommit unblocks adds a commit to the round state.
func (rs *RoundState) UnblockCommit() {
	rs.rwMutex.Lock()
	defer rs.rwMutex.Unlock()

	rs.isBlocked = false
}

// GetCurrentBlock returns the current proposal block.
func (rs *RoundState) GetCurrentBlock() *sequencertypes.Block {
	rs.rwMutex.RLock()
	defer rs.rwMutex.RUnlock()

	return rs.ProposalBlock
}

// GetCurrentBlockNumber returns the current block number.
func (rs *RoundState) GetCurrentBlockNumber() uint64 {
	rs.rwMutex.RLock()
	defer rs.rwMutex.RUnlock()

	return rs.ProposalBlock.BlockNumber()
}

// CheckEnoughVotingPower checks if there is enough voting power to finalize the block.
func (rs *RoundState) CheckEnoughVotingPower() bool {
	rs.rwMutex.RLock()
	defer rs.rwMutex.RUnlock()

	votingPower := uint64(0)
	for _, signature := range rs.commitSignatures {
		votingPower += rs.Validators.GetVotingPower(signature.PubKey)
	}

	logger.Infof("voting power: %v, total voting power: %v", votingPower, rs.Validators.TotalVotingPower)
	return votingPower*3 > rs.Validators.TotalVotingPower*2
}

// CheckAggregatedSignature checks if the aggregated signature is valid.
func (rs *RoundState) CheckAggregatedSignature() ([]*bls.PublicKey, *bls.Signature, error) {
	rs.rwMutex.Lock()
	defer rs.rwMutex.Unlock()

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
		pubkey, err := utils.HexToBlsPubKey(commit.PubKey)
		if err != nil {
			return nil, nil, err
		}
		sig, err := utils.HexToBlsSignature(commit.BlsSignature.Signature)
		if err != nil {
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

// UpdateAggregatedSignature updates the aggregated signature.
func (rs *RoundState) UpdateAggregatedSignature(pubkeys []string, aggSig string) {
	rs.rwMutex.Lock()
	defer rs.rwMutex.Unlock()

	rs.ProposalBlock.PubKeys = pubkeys
	rs.ProposalBlock.AggSignature = aggSig
}

// GetEvidences returns the evidences.
func (rs *RoundState) GetEvidences() []*networktypes.CommitBlockRequest {
	rs.rwMutex.RLock()
	defer rs.rwMutex.RUnlock()

	return rs.evidences
}

var (
	ErrInvalidAggregativeSignature = fmt.Errorf("invalid aggregative signature")
)

package types

import (
	"bytes"
	"fmt"
	"sync"

	"github.com/Lagrange-Labs/lagrange-node/crypto"
	"github.com/Lagrange-Labs/lagrange-node/logger"
	sequencerv2types "github.com/Lagrange-Labs/lagrange-node/sequencer/types/v2"
	"github.com/Lagrange-Labs/lagrange-node/utils"
)

// CommitSignature defines the commit signature.
type CommitSignature struct {
	Signature *sequencerv2types.BlsSignature
	PubKey    []byte
	StakeAddr string
}

// RoundState defines the internal consensus state.
type RoundState struct {
	blsScheme     crypto.BLSScheme
	proposedBatch *sequencerv2types.Batch

	commitSignatures map[string]CommitSignature
	evidences        []*sequencerv2types.BlsSignature // to determine slashing

	rwMutex   sync.RWMutex // to protect the round state updates
	isBlocked bool         // to prevent the block commit
}

// NewEmptyRoundState creates a new empty round state for rwMutex.
func NewEmptyRoundState(blsScheme crypto.BLSScheme) *RoundState {
	return &RoundState{
		isBlocked: true,
		blsScheme: blsScheme,
	}
}

// UpdateRoundState updates a new round state.
func (rs *RoundState) UpdateRoundState(prposedBatch *sequencerv2types.Batch) {
	rs.rwMutex.Lock()
	defer rs.rwMutex.Unlock()

	rs.proposedBatch = prposedBatch
	rs.commitSignatures = make(map[string]CommitSignature)
	rs.isBlocked = false
}

// AddCommit adds a commit to the round state.
func (rs *RoundState) AddCommit(commit *sequencerv2types.BlsSignature, pubKey []byte, stakeAddr string) error {
	rs.rwMutex.Lock()
	defer rs.rwMutex.Unlock()

	if rs.isBlocked {
		return fmt.Errorf("the current round is blocked")
	}

	rs.commitSignatures[stakeAddr] = CommitSignature{
		Signature: commit,
		PubKey:    pubKey,
		StakeAddr: stakeAddr,
	}

	return nil
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

// IsFinalized checks if the block is finalized.
func (rs *RoundState) IsFinalized() bool {
	rs.rwMutex.RLock()
	defer rs.rwMutex.RUnlock()

	return len(rs.proposedBatch.PubKeys) > 0
}

// GetCurrentBatchNumber returns the current batch number.
func (rs *RoundState) GetCurrentBatchNumber() uint64 {
	rs.rwMutex.RLock()
	defer rs.rwMutex.RUnlock()

	return rs.proposedBatch.BatchNumber()
}

// GetCurrentBatch returns the current batch.
func (rs *RoundState) GetCurrentBatch() *sequencerv2types.Batch {
	rs.rwMutex.RLock()
	defer rs.rwMutex.RUnlock()

	return rs.proposedBatch
}

// CheckEnoughVotingPower checks if there is enough voting power to finalize the block.
func (rs *RoundState) CheckEnoughVotingPower(vs *ValidatorSet) bool {
	rs.rwMutex.RLock()
	defer rs.rwMutex.RUnlock()

	votingPower := uint64(0)
	for stakeAddr := range rs.commitSignatures {
		votingPower += vs.GetVotingPower(stakeAddr)
	}

	logger.Infof("committed count: %d, committed voting power: %v, total voting power: %v", len(rs.commitSignatures), votingPower, vs.GetCommitteeVotingPower())
	return len(rs.commitSignatures)*3 > len(vs.validators)*2 && votingPower*3 > vs.GetCommitteeVotingPower()*2
}

// CheckAggregatedSignature checks if the aggregated signature is valid.
func (rs *RoundState) CheckAggregatedSignature() error {
	rs.rwMutex.Lock()
	defer rs.rwMutex.Unlock()

	blsSignature := rs.proposedBatch.BlsSignature()
	sigHash := blsSignature.Hash()
	signatures := make([][]byte, 0)
	pubKeys := make([][]byte, 0)

	// aggregate the signatures of client nodes
	for _, commit := range rs.commitSignatures {
		signatures = append(signatures, utils.Hex2Bytes(commit.Signature.BlsSignature))
		pubKeys = append(pubKeys, commit.PubKey)
	}

	aggSig, err := rs.blsScheme.AggregateSignatures(signatures)
	if err != nil {
		logger.Errorf("failed to aggregate the signatures: %v", err)
	} else {
		verified, err := rs.blsScheme.VerifyAggregatedSignature(pubKeys, sigHash, aggSig)
		if err == nil && verified {
			rs.proposedBatch.AggSignature = utils.Bytes2Hex(aggSig)
			for _, pubKey := range pubKeys {
				rs.proposedBatch.PubKeys = append(rs.proposedBatch.PubKeys, utils.Bytes2Hex(pubKey))
			}
			return nil
		}
		if err != nil {
			logger.Errorf("failed to verify the aggregated signature: %v", err)
		}
	}

	// find the invalid signature
	for operator, commit := range rs.commitSignatures {
		commitHash := commit.Signature.Hash()
		if !bytes.Equal(commitHash, sigHash) {
			logger.Errorf("wrong commit message: %v, original: %v", utils.Bytes2Hex(commitHash), utils.Bytes2Hex(sigHash))
			rs.addEvidence(operator, commit.Signature)
			continue
		}
		verified, err := rs.blsScheme.VerifySignature(commit.PubKey, commitHash, utils.Hex2Bytes(commit.Signature.BlsSignature))
		if err != nil {
			logger.Errorf("failed to verify the signature: %v", err)
			rs.addEvidence(operator, commit.Signature)
			continue
		}
		if !verified {
			logger.Errorf("invalid signature: %v", commit)
			rs.addEvidence(operator, commit.Signature)
		}
	}

	logger.Errorf("invalid aggregated signature: %v", rs.proposedBatch)

	return ErrInvalidAggregativeSignature
}

func (rs *RoundState) addEvidence(operator string, signature *sequencerv2types.BlsSignature) {
	rs.evidences = append(rs.evidences, signature)
	delete(rs.commitSignatures, operator)
}

// GetEvidences returns the evidences.
func (rs *RoundState) GetEvidences() ([]*Evidence, error) {
	rs.rwMutex.RLock()
	defer rs.rwMutex.RUnlock()
	var evidences []*Evidence

	for _, req := range rs.evidences {
		evidence, err := GetEvidence(req)
		if err != nil {
			return nil, err
		}
		evidences = append(evidences, evidence)
	}
	return evidences, nil
}

var (
	ErrInvalidAggregativeSignature = fmt.Errorf("invalid aggregative signature")
)

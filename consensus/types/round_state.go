package types

import (
	"bytes"
	"fmt"
	"sync"

	"github.com/Lagrange-Labs/lagrange-node/crypto"
	"github.com/Lagrange-Labs/lagrange-node/logger"
	sequencertypes "github.com/Lagrange-Labs/lagrange-node/sequencer/types"
	"github.com/Lagrange-Labs/lagrange-node/utils"
)

// CommitSignature defines the commit signature.
type CommitSignature struct {
	Signature *sequencertypes.BlsSignature
	PubKey    []byte
	StakeAddr string
}

// RoundState defines the internal consensus state.
type RoundState struct {
	blsScheme     crypto.BLSScheme
	proposalBlock *sequencertypes.Block

	commitSignatures []CommitSignature
	evidences        []*sequencertypes.BlsSignature // to determine slashing

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
func (rs *RoundState) UpdateRoundState(proposalBlock *sequencertypes.Block) {
	rs.rwMutex.Lock()
	defer rs.rwMutex.Unlock()

	rs.proposalBlock = proposalBlock
	rs.commitSignatures = make([]CommitSignature, 0)
	rs.isBlocked = false
}

// AddCommit adds a commit to the round state.
func (rs *RoundState) AddCommit(commit *sequencertypes.BlsSignature, pubKey []byte, stakeAddr string) {
	rs.rwMutex.Lock()
	defer rs.rwMutex.Unlock()

	if rs.isBlocked {
		logger.Info("the add commit is blocked")
		return
	}

	rs.commitSignatures = append(rs.commitSignatures, CommitSignature{
		Signature: commit,
		PubKey:    pubKey,
		StakeAddr: stakeAddr,
	})
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

	return len(rs.proposalBlock.PubKeys) > 0
}

// GetCurrentBlockNumber returns the current block number.
func (rs *RoundState) GetCurrentBlockNumber() uint64 {
	rs.rwMutex.RLock()
	defer rs.rwMutex.RUnlock()

	return rs.proposalBlock.BlockNumber()
}

// GetCurrentBlock returns the current block.
func (rs *RoundState) GetCurrentBlock() *sequencertypes.Block {
	rs.rwMutex.RLock()
	defer rs.rwMutex.RUnlock()

	return rs.proposalBlock
}

// GetCurrentEpochBlockNumber returns the current epoch block number.
func (rs *RoundState) GetCurrentEpochBlockNumber() uint64 {
	rs.rwMutex.RLock()
	defer rs.rwMutex.RUnlock()

	return rs.proposalBlock.EpochBlockNumber()
}

// CheckEnoughVotingPower checks if there is enough voting power to finalize the block.
func (rs *RoundState) CheckEnoughVotingPower(vs *ValidatorSet) bool {
	rs.rwMutex.RLock()
	defer rs.rwMutex.RUnlock()

	votingPower := uint64(0)
	for _, commit := range rs.commitSignatures {
		votingPower += vs.GetVotingPower(commit.StakeAddr)
	}

	logger.Infof("committed voting power: %v, validator set voting power: %v", votingPower, vs.GetTotalVotingPower())
	return votingPower*3 > vs.GetCommitteeVotingPower()*2
}

// CheckAggregatedSignature checks if the aggregated signature is valid.
func (rs *RoundState) CheckAggregatedSignature() error {
	rs.rwMutex.Lock()
	defer rs.rwMutex.Unlock()

	blsSignature := rs.proposalBlock.BlsSignature()
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
			rs.proposalBlock.AggSignature = utils.Bytes2Hex(aggSig)
			for _, pubKey := range pubKeys {
				//TODO: refactor the block structure, this iteration is too expensive
				rs.proposalBlock.PubKeys = append(rs.proposalBlock.PubKeys, utils.Bytes2Hex(pubKey))
			}
			return nil
		}
		if err != nil {
			logger.Errorf("failed to verify the aggregated signature: %v", err)
		}
	}

	// find the invalid signature
	for i, pubKeyRaw := range pubKeys {
		commit := rs.commitSignatures[i]
		commitHash := commit.Signature.Hash()
		if !bytes.Equal(commitHash, sigHash) {
			logger.Errorf("wrong commit message: %v, original: %v", utils.Bytes2Hex(commitHash), utils.Bytes2Hex(sigHash))
			rs.evidences = append(rs.evidences, rs.commitSignatures[i].Signature)
			continue
		}
		verified, err := rs.blsScheme.VerifySignature(pubKeyRaw, commitHash, signatures[i])
		if err != nil {
			logger.Errorf("failed to verify the signature: %v", err)
			rs.evidences = append(rs.evidences, rs.commitSignatures[i].Signature)
			continue
		}
		if !verified {
			logger.Errorf("invalid signature: %v", commit)
			rs.evidences = append(rs.evidences, rs.commitSignatures[i].Signature)
		}
	}

	logger.Errorf("invalid aggregated signature: %v", rs.proposalBlock)

	return ErrInvalidAggregativeSignature
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

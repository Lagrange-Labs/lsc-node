package types

import (
	"bytes"
	"fmt"
	"sync"

	"github.com/Lagrange-Labs/lagrange-node/logger"
	sequencertypes "github.com/Lagrange-Labs/lagrange-node/sequencer/types"
	"github.com/Lagrange-Labs/lagrange-node/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/umbracle/go-eth-consensus/bls"
)

// RoundState defines the internal consensus state.
type RoundState struct {
	proposalBlock *sequencertypes.Block

	commitSignatures map[string]*sequencertypes.BlsSignature
	evidences        []*sequencertypes.BlsSignature // to determine slashing

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
func (rs *RoundState) UpdateRoundState(proposalBlock *sequencertypes.Block) {
	rs.rwMutex.Lock()
	defer rs.rwMutex.Unlock()

	rs.proposalBlock = proposalBlock
	rs.commitSignatures = make(map[string]*sequencertypes.BlsSignature)
	rs.isBlocked = false
}

// AddCommit adds a commit to the round state.
func (rs *RoundState) AddCommit(commit *sequencertypes.BlsSignature, pubKey string) {
	rs.rwMutex.Lock()
	defer rs.rwMutex.Unlock()

	if rs.isBlocked {
		logger.Info("the add commit is blocked")
		return
	}

	rs.commitSignatures[pubKey] = commit
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
	for pubKey := range rs.commitSignatures {
		votingPower += vs.GetVotingPower(pubKey)
	}

	totalVotingPower := vs.GetTotalVotingPower()
	logger.Infof("committed voting power: %v, committee voting power: %v", votingPower, totalVotingPower)
	return votingPower*3 > totalVotingPower*2
}

// CheckAggregatedSignature checks if the aggregated signature is valid.
func (rs *RoundState) CheckAggregatedSignature() error {
	rs.rwMutex.Lock()
	defer rs.rwMutex.Unlock()

	blsSignature := rs.proposalBlock.BlsSignature()
	sigHash := blsSignature.Hash()
	signatures := make([]*bls.Signature, 0)
	pubKeys := make([]*bls.PublicKey, 0)
	pubKeyRaws := make([]string, 0)
	invalid_keys := make([]string, 0)

	// aggregate the signatures of client nodes
	for pubKeyRaw, commit := range rs.commitSignatures {
		pubKey, err := utils.HexToBlsPubKey(pubKeyRaw)
		if err != nil {
			// it is a critical error if the public key is not valid because it is from the database
			return err
		}
		sig, err := utils.HexToBlsSignature(commit.BlsSignature)
		if err != nil {
			logger.Errorf("failed to deserialize signature %v: %v", commit, err)
			invalid_keys = append(invalid_keys, pubKeyRaw)
			continue
		}
		signatures = append(signatures, sig)
		pubKeys = append(pubKeys, pubKey)
		pubKeyRaws = append(pubKeyRaws, pubKeyRaw)
	}

	aggSig := bls.AggregateSignatures(signatures)
	verified, err := aggSig.FastAggregateVerify(pubKeys, sigHash)
	if verified && len(invalid_keys) == 0 {
		rs.proposalBlock.AggSignature = utils.BlsSignatureToHex(aggSig)
		rs.proposalBlock.PubKeys = pubKeyRaws
		return nil
	}

	if err != nil {
		logger.Errorf("failed to verify the aggregated signature: %v", err)
	}

	// find the invalid signature
	for i, pubKeyRaw := range pubKeyRaws {
		commit := rs.commitSignatures[pubKeyRaw]
		commitHash := commit.Hash()
		if !bytes.Equal(commitHash, sigHash) {
			logger.Errorf("wrong commit message: %v, original: %v", common.Bytes2Hex(commitHash), common.Bytes2Hex(sigHash))
			invalid_keys = append(invalid_keys, pubKeyRaw)
			continue
		}
		verified, err := signatures[i].VerifyByte(pubKeys[i], commitHash)
		if err != nil {
			return err
		}
		if !verified {
			logger.Errorf("invalid signature: %v", commit)
			invalid_keys = append(invalid_keys, pubKeyRaw)
		}
	}

	// add invalid signatures to evidences
	for _, key := range invalid_keys {
		rs.evidences = append(rs.evidences, rs.commitSignatures[key])
		delete(rs.commitSignatures, key)
	}

	return ErrInvalidAggregativeSignature
}

// GetEvidences returns the evidences.
func (rs *RoundState) GetEvidences() ([]*Evidence, error) {
	rs.rwMutex.RLock()
	defer rs.rwMutex.RUnlock()
	var evidences []*Evidence

	for _, req := range rs.evidences {
		evidence, err := GetEvidence(req, rs.proposalBlock.BlockHash(), rs.proposalBlock.CurrentCommittee(), rs.proposalBlock.NextCommittee())
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

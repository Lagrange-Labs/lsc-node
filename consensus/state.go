package consensus

import (
	"bytes"
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/Lagrange-Labs/lagrange-node/consensus/types"
	"github.com/Lagrange-Labs/lagrange-node/crypto"
	"github.com/Lagrange-Labs/lagrange-node/logger"
	sequencerv2types "github.com/Lagrange-Labs/lagrange-node/sequencer/types/v2"
	storetypes "github.com/Lagrange-Labs/lagrange-node/store/types"
	"github.com/Lagrange-Labs/lagrange-node/utils"
)

const CheckInterval = 1 * time.Second

// State handles the consensus process.
type State struct {
	validators *types.ValidatorSet

	round           *types.RoundState
	previousBatch   *sequencerv2types.Batch
	lastBatchNumber uint64
	rwMutex         *sync.RWMutex
	blsScheme       crypto.BLSScheme

	proposerPrivKey  []byte
	proposerPubKey   string // hex string
	storage          storageInterface
	roundLimit       time.Duration
	roundInterval    time.Duration
	batchSize        uint32
	chainID          uint32
	lastCommittee    *sequencerv2types.CommitteeRoot
	blockedOperators map[string]struct{}

	chStop chan struct{}
}

// NewState returns a new State.
func NewState(cfg *Config, storage storageInterface, chainID uint32) *State {
	privKey := utils.Hex2Bytes(cfg.ProposerPrivateKey)
	blsScheme := crypto.NewBLSScheme(crypto.BLSCurve(cfg.BLSCurve))
	pubKey, err := blsScheme.GetPublicKey(privKey, true)
	if err != nil {
		logger.Fatalf("failed to get the public key: %v", err)
	}

	chStop := make(chan struct{})

	return &State{
		blsScheme:        blsScheme,
		proposerPrivKey:  privKey,
		proposerPubKey:   utils.Bytes2Hex(pubKey),
		storage:          storage,
		roundLimit:       time.Duration(cfg.RoundLimit),
		roundInterval:    time.Duration(cfg.RoundInterval),
		chainID:          chainID,
		batchSize:        cfg.BatchSize,
		chStop:           chStop,
		rwMutex:          &sync.RWMutex{},
		blockedOperators: make(map[string]struct{}),
	}
}

// GetBLSScheme returns the BLS scheme.
func (s *State) GetBLSScheme() crypto.BLSScheme {
	return s.blsScheme
}

// OnStart loads the first unverified block and starts the round.
func (s *State) OnStart() {
	logger.Info("Consensus process is started with the batch size: ", s.batchSize)

	for {
		// check if chStop is triggered
		select {
		case <-s.chStop:
			return
		default:
		}

		lastBatchNumber, err := s.storage.GetLastFinalizedBatchNumber(context.Background(), s.chainID)
		if err != nil && err != storetypes.ErrBatchNotFound {
			logger.Errorf("failed to get the last finalized batch number: %v", err)
			return
		}
		if err == nil {
			batch, err := s.storage.GetBatch(context.Background(), s.chainID, lastBatchNumber)
			if err == storetypes.ErrBatchNotFound {
				logger.Infof("the last finalized batch %d is not found", s.lastBatchNumber)
				s.previousBatch = nil
				s.lastBatchNumber = lastBatchNumber + 1
			} else if err != nil {
				logger.Errorf("failed to get the last finalized batch %d: %v", lastBatchNumber, err)
				return
			} else {
				s.previousBatch = batch
				s.lastBatchNumber = batch.BatchHeader.ToBlockNumber() + 1
			}
			break
		}

		logger.Info("waiting for the first block")
		time.Sleep(CheckInterval)
	}

	for {
		// check if chStop is triggered
		select {
		case <-s.chStop:
			return
		default:
		}

		logger.Infof("start the round with batch number %v", s.lastBatchNumber)
		if err := s.startRound(s.lastBatchNumber); err != nil {
			logger.Errorf("failed to start the round: %v", err)
			time.Sleep(s.roundInterval)
			continue
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Duration(s.roundLimit))
		defer cancel()
		isVoted := s.processRound(ctx)
		if !isVoted {
			logger.Error("the current batch is not finalized within the round limit")
		}

		// update the finalized batch
		updateFinalizedBatch := func() {
			batch := s.round.GetCurrentBatch()
			batch.FinalizedTime = fmt.Sprintf("%d", time.Now().UnixMicro())
			if err := s.storage.UpdateBatch(context.Background(), batch); err != nil {
				logger.Errorf("failed to update the batch %d: %v", s.lastBatchNumber, err)
				return
			}
			logger.Infof("the batch %d is finalized", s.lastBatchNumber)
		}

		if s.round.IsFinalized() {
			updateFinalizedBatch()
		} else {
			// TODO: handle the case when the batch is not finalized, now it will be run forever
			logger.Error("the infinite loop is started!")
			_ = s.processRound(context.Background())
			updateFinalizedBatch()
		}

		// store the evidences
		evidences, err := s.round.GetEvidences()
		if err != nil {
			logger.Errorf("failed to get the evidences: %v", err)
			continue
		}
		if len(evidences) > 0 {
			if err := s.storage.AddEvidences(ctx, evidences); err != nil {
				logger.Errorf("failed to add the evidences: %v", err)
				continue
			}

			// block the operators
			for _, evidence := range evidences {
				s.blockedOperators[evidence.Operator] = struct{}{}
			}
		}

		s.rwMutex.Lock()
		s.previousBatch = s.round.GetCurrentBatch()
		s.lastBatchNumber = s.previousBatch.BatchHeader.ToBlockNumber() + 1
		s.rwMutex.Unlock()
	}
}

// OnStop stops the consensus process.
func (s *State) OnStop() {
	logger.Infof("OnStop() called")
	s.chStop <- struct{}{}
	close(s.chStop)
}

// GetOpenBatch returns the batch of the current round.
func (s *State) GetOpenBatch(batchNumber uint64) *sequencerv2types.Batch {
	return s.round.GetCurrentBatch()
}

// GetOpenBatchNumber returns the batch number of the current round.
func (s *State) GetOpenBatchNumber() (uint64, uint64) {
	batch := s.round.GetCurrentBatch()
	prevL1BlockNumber := batch.L1BlockNumber()
	if s.previousBatch != nil {
		prevL1BlockNumber = s.previousBatch.L1BlockNumber()
	}
	return batch.BatchNumber(), prevL1BlockNumber
}

// AddBatchCommit adds the commit to the round state.
func (s *State) AddBatchCommit(commit *sequencerv2types.BlsSignature, stakeAddr string) error {
	s.rwMutex.Lock()
	defer s.rwMutex.Unlock()

	if s.validators == nil {
		return fmt.Errorf("the validator set is not initialized")
	}

	// check if the operator is blocked
	if _, ok := s.blockedOperators[stakeAddr]; ok {
		return fmt.Errorf("the operator %s is blocked", stakeAddr)
	}

	if s.validators.GetVotingPower(stakeAddr) == 0 {
		return fmt.Errorf("the stake address %s is not registered", stakeAddr)
	}

	if s.round.GetCurrentBatchNumber() != commit.BatchNumber() {
		return fmt.Errorf("the batch number %d is not matched with the current batch number %d", commit.BatchNumber(), s.round.GetCurrentBatchNumber())
	}

	return s.round.AddCommit(commit, s.validators.GetPublicKey(stakeAddr), stakeAddr)
}

// CheckCommitteeMember checks if the operator is a committee member.
func (s *State) CheckCommitteeMember(stakeAddr string, pubKey []byte) bool {
	if s.validators == nil {
		return false
	}
	return bytes.Equal(s.validators.GetPublicKey(stakeAddr), pubKey)
}

// IsFinalized returns true if the current batch is finalized.
func (s *State) IsFinalized(batchNumber uint64) bool {
	s.rwMutex.RLock()
	defer s.rwMutex.RUnlock()

	if s.lastBatchNumber > batchNumber {
		return true
	}

	return s.round.IsFinalized()
}

// startRound loads the next batch and initializes the round state.
func (s *State) startRound(batchNumber uint64) error {
	s.rwMutex.Lock()
	defer s.rwMutex.Unlock()

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(s.roundLimit))
	defer cancel()
	batch, err := s.getNextBatch(ctx, batchNumber)
	if err != nil {
		return fmt.Errorf("getting the next batch %d is failed: %v", batchNumber, err)
	}

	logger.Infof("the batch %d is loaded with L2 blocks from %d to %d", batch.BatchNumber(), batch.BatchHeader.FromBlockNumber(), batch.BatchHeader.ToBlockNumber())

	batch.CommitteeHeader = &sequencerv2types.CommitteeHeader{}
	// load the committee root
	var currentCommittee *sequencerv2types.CommitteeRoot
	if s.lastCommittee == nil || batch.L1BlockNumber() > s.lastCommittee.EpochEndBlockNumber {
		logger.Infof("the next committee root is loading: %v", batch.L1BlockNumber())
		nextCommittee, err := s.storage.GetCommitteeRoot(context.Background(), s.chainID, batch.L1BlockNumber())
		if err != nil {
			logger.Errorf("failed to get the committee root for the block number %d: %v", batch.L1BlockNumber(), err)
			return err
		}
		if s.lastCommittee == nil && s.previousBatch != nil {
			prevCommittee, err := s.storage.GetCommitteeRoot(context.Background(), s.chainID, s.previousBatch.L1BlockNumber())
			if err != nil {
				logger.Errorf("failed to get the previous committee root: %v", err)
				return err
			}
			if prevCommittee.EpochEndBlockNumber < nextCommittee.EpochEndBlockNumber {
				s.lastCommittee = prevCommittee
			}
		}
		if s.lastCommittee != nil {
			currentCommittee = s.lastCommittee
		} else {
			currentCommittee = nextCommittee
		}
		batch.CommitteeHeader.CurrentCommittee = currentCommittee.CurrentCommitteeRoot
		batch.CommitteeHeader.NextCommittee = nextCommittee.CurrentCommitteeRoot
		s.lastCommittee = nextCommittee
	} else {
		currentCommittee = s.lastCommittee
		batch.CommitteeHeader.CurrentCommittee = currentCommittee.CurrentCommitteeRoot
		batch.CommitteeHeader.NextCommittee = currentCommittee.CurrentCommitteeRoot
	}

	batch.CommitteeHeader.TotalVotingPower = currentCommittee.TotalVotingPower
	s.validators = types.NewValidatorSet(currentCommittee.Operators, currentCommittee.TotalVotingPower)

	s.round = types.NewEmptyRoundState(s.blsScheme)

	// generate a proposer signature
	blsSigHash := batch.BlsSignature().Hash()
	signature, err := s.blsScheme.Sign(s.proposerPrivKey, blsSigHash)
	if err != nil {
		logger.Errorf("failed to sign the batch %d: %v", batch.BatchNumber(), err)
		return err
	}
	batch.ProposerSignature = utils.Bytes2Hex(signature)
	batch.ProposerPubKey = s.proposerPubKey

	s.round.UpdateRoundState(batch)

	return nil
}

// getNextBatch returns the next batch from the storage.
func (s *State) getNextBatch(ctx context.Context, batchNumber uint64) (*sequencerv2types.Batch, error) {
	getBatch := func(batchNumber uint64) (*sequencerv2types.Batch, error) {
		batch, err := s.storage.GetBatch(ctx, uint32(s.chainID), batchNumber)
		if err == storetypes.ErrBatchNotFound {
			return nil, nil
		} else if err != nil {
			logger.Errorf("failed to get the next batch %d: %v", batchNumber, err)
			return nil, err
		}
		return batch, nil
	}

	batch, err := getBatch(batchNumber)
	if err != nil {
		return nil, err
	}
	if batch != nil {
		return batch, nil
	}
	// in case of batch not found, wait for it to be added from the sequencer
	ticker := time.NewTicker(CheckInterval)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		case <-ticker.C:
			batch, err := getBatch(batchNumber)
			if err != nil {
				return nil, err
			}
			if batch != nil {
				return batch, nil
			}
		}
	}
}

// processRound processes the round.
func (s *State) processRound(ctx context.Context) bool {
	checkCommit := func(round *types.RoundState) (bool, error) {
		if round.CheckEnoughVotingPower(s.validators) {
			round.BlockCommit()
			err := round.CheckAggregatedSignature()
			if err != nil {
				round.UnblockCommit()
				if err == types.ErrInvalidAggregativeSignature {
					logger.Warnf("the aggregated signature is invalid for the batch %d", round.GetCurrentBatchNumber())
					return false, nil
				}
				logger.Errorf("failed to check the aggregated signature for the batch %d: %v", round.GetCurrentBatchNumber(), err)
				return false, err
			}
			return true, nil
		}
		return false, nil
	}

	ticker := time.NewTicker(s.roundInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			break
		case <-ticker.C:
			isFinalized, err := checkCommit(s.round)
			if err != nil {
				logger.Errorf("failed to check the commit for the batch %d: %v", s.lastBatchNumber, err)
				return false
			}
			if isFinalized {
				return true
			}
		}
	}
}

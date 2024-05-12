package consensus

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/Lagrange-Labs/lagrange-node/consensus/types"
	"github.com/Lagrange-Labs/lagrange-node/crypto"
	"github.com/Lagrange-Labs/lagrange-node/logger"
	sequencerv2types "github.com/Lagrange-Labs/lagrange-node/sequencer/types/v2"
	storetypes "github.com/Lagrange-Labs/lagrange-node/store/types"
	"github.com/Lagrange-Labs/lagrange-node/telemetry"
	"github.com/Lagrange-Labs/lagrange-node/utils"
)

const CheckInterval = 1 * time.Second

// State handles the consensus process.
type State struct {
	validators *types.ValidatorSet

	round           *types.RoundState
	previousBatch   *sequencerv2types.Batch
	fromBatchNumber uint64
	rwMutex         *sync.RWMutex
	blsScheme       crypto.BLSScheme

	proposerPrivKey  []byte
	proposerPubKey   string // hex string
	storage          storageInterface
	roundLimit       time.Duration
	roundInterval    time.Duration
	chainID          uint32
	lastCommittee    *sequencerv2types.CommitteeRoot
	blockedOperators map[string]struct{}

	ctx    context.Context
	cancel context.CancelFunc
}

// NewState returns a new State.
func NewState(cfg *Config, storage storageInterface, chainID uint32) *State {
	if len(cfg.ProposerBLSKeystorePasswordPath) > 0 {
		var err error
		cfg.ProposerBLSKeystorePasswordPath, err = crypto.ReadKeystorePasswordFromFile(cfg.ProposerBLSKeystorePasswordPath)
		if err != nil {
			logger.Fatalf("failed to read the bls keystore password from %s: %v", cfg.ProposerBLSKeystorePasswordPath, err)
		}
	}
	privKey, err := crypto.LoadPrivateKey(crypto.CryptoCurve(cfg.BLSCurve), cfg.ProposerBLSKeystorePassword, cfg.ProposerBLSKeystorePath)
	if err != nil {
		logger.Fatalf("failed to load the bls keystore from %s: %v", cfg.ProposerBLSKeystorePath, err)
	}
	blsScheme := crypto.NewBLSScheme(crypto.BLSCurve(cfg.BLSCurve))
	pubKey, err := blsScheme.GetPublicKey(privKey, true)
	if err != nil {
		logger.Fatalf("failed to get the public key: %v", err)
	}

	return &State{
		blsScheme:        blsScheme,
		proposerPrivKey:  privKey,
		proposerPubKey:   utils.Bytes2Hex(pubKey),
		storage:          storage,
		roundLimit:       time.Duration(cfg.RoundLimit),
		roundInterval:    time.Duration(cfg.RoundInterval),
		chainID:          chainID,
		rwMutex:          &sync.RWMutex{},
		blockedOperators: make(map[string]struct{}),
	}
}

// GetBLSScheme returns the BLS scheme.
func (s *State) GetBLSScheme() crypto.BLSScheme {
	return s.blsScheme
}

// GetRoundInterval returns the round interval.
func (s *State) GetRoundInterval() time.Duration {
	return s.roundInterval + 500*time.Millisecond
}

// GetOpenBatch returns the batch of the current round.
func (s *State) GetOpenBatch() *sequencerv2types.Batch {
	return s.round.GetCurrentBatch()
}

// GetPrevBatch returns the previous batch.
func (s *State) GetPrevBatch() *sequencerv2types.Batch {
	if s.previousBatch == nil {
		return s.round.GetCurrentBatch()
	}

	return s.previousBatch
}

// OnStart loads the first unverified block and starts the round.
func (s *State) OnStart() {
	logger.Info("Consensus process is started")
	s.ctx, s.cancel = context.WithCancel(context.Background())

	for {
		// check if OnStop is triggered
		select {
		case <-s.ctx.Done():
			s.ctx = nil
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
				logger.Infof("the last finalized batch %d is not found", lastBatchNumber)
				s.previousBatch = nil
				s.fromBatchNumber = lastBatchNumber + 1
			} else if err != nil {
				logger.Errorf("failed to get the last finalized batch %d: %v", lastBatchNumber, err)
				return
			} else {
				s.previousBatch = batch
				s.fromBatchNumber = batch.BatchNumber() + 1
			}
			break
		}

		logger.Info("waiting for the first block")
		time.Sleep(CheckInterval)
	}

	for {
		// check if OnStop is triggered
		select {
		case <-s.ctx.Done():
			s.ctx = nil
			return
		default:
		}
		ti := time.Now()
		logger.Infof("start the round with batch number %v", s.fromBatchNumber)
		if err := s.startRound(s.fromBatchNumber); err != nil {
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
				logger.Errorf("failed to update the batch %d: %v", s.fromBatchNumber, err)
				return
			}
			telemetry.SetGauge(float64(len(batch.PubKeys)), "consensus", "committed_node_count")
			logger.Infof("the batch number %d, L1 block number %d, upto L2 block number %d  is finalized", batch.BatchNumber(), batch.L1BlockNumber(), batch.BatchHeader.ToBlockNumber())
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
		telemetry.SetGauge(float64(len(evidences)), "consensus", "evidence_count")

		s.rwMutex.Lock()
		telemetry.SetGauge(float64(s.fromBatchNumber), "consensus", "finalized_batch_number")
		s.previousBatch = s.round.GetCurrentBatch()
		s.fromBatchNumber++
		s.rwMutex.Unlock()
		telemetry.MeasureSince(ti, "consensus", "round_duration")
	}
}

// OnStop stops the consensus process.
func (s *State) OnStop() {
	logger.Infof("OnStop() called")
	if s != nil && s.ctx != nil {
		s.cancel()
	}
}

// IsStopped returns true if the consensus process is stopped.
func (s *State) IsStopped() bool {
	return s.ctx == nil
}

// AddBatchCommit adds the commit to the round state.
func (s *State) AddBatchCommit(commit *sequencerv2types.BlsSignature, stakeAddr, pubKey string) error {
	s.rwMutex.Lock()
	defer s.rwMutex.Unlock()

	if s.validators == nil {
		return fmt.Errorf("the validator set is not initialized")
	}

	// check if the operator is blocked
	if _, ok := s.blockedOperators[stakeAddr]; ok {
		return fmt.Errorf("the operator %s is blocked", stakeAddr)
	}

	if s.validators.GetVotingPower(stakeAddr, pubKey) == 0 {
		return fmt.Errorf("the operator address %s is not registered", stakeAddr)
	}

	if s.round.GetCurrentBatchNumber() != commit.BatchNumber() {
		return fmt.Errorf("the batch number %d is not matched with the current batch number %d", commit.BatchNumber(), s.round.GetCurrentBatchNumber())
	}

	return s.round.AddCommit(commit, pubKey, stakeAddr)
}

// CheckCommitteeMember checks if the operator is a committee member.
func (s *State) CheckCommitteeMember(stakeAddr, pubKey string) (bool, error) {
	if s.validators == nil {
		return false, fmt.Errorf("the validator set is not initialized")
	}
	return s.validators.GetVotingPower(stakeAddr, pubKey) > 0, nil
}

// CheckSignAddress checks if the sign address is valid.
func (s *State) CheckSignAddress(stakeAddr, signAddr string) bool {
	if s.validators == nil {
		return false
	}
	return s.validators.GetSignAddress(stakeAddr) == signAddr
}

// IsFinalized returns true if the current batch is finalized.
func (s *State) IsFinalized(batchNumber uint64) bool {
	s.rwMutex.RLock()
	defer s.rwMutex.RUnlock()

	if s.fromBatchNumber > batchNumber {
		return true
	}

	return s.round.IsFinalized()
}

// startRound loads the next batch and initializes the round state.
func (s *State) startRound(batchNumber uint64) error {
	ti := time.Now()
	defer telemetry.MeasureSince(ti, "consensus", "start_round")

	s.rwMutex.Lock()
	s.round = types.NewEmptyRoundState(s.blsScheme)
	s.rwMutex.Unlock()

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
	telemetry.SetGauge(float64(len(currentCommittee.Operators)), "consensus", "committee_size")
	telemetry.SetGauge(float64(currentCommittee.TotalVotingPower), "consensus", "committee_voting_power")

	// generate a proposer signature
	blsSigHash := batch.BlsSignature().Hash()
	signature, err := s.blsScheme.Sign(s.proposerPrivKey, blsSigHash)
	if err != nil {
		logger.Errorf("failed to sign the batch %d: %v", batch.BatchNumber(), err)
		return err
	}
	batch.ProposerSignature = utils.Bytes2Hex(signature)
	batch.ProposerPubKey = s.proposerPubKey

	s.rwMutex.Lock()
	defer s.rwMutex.Unlock()
	s.round.UpdateRoundState(batch)

	return nil
}

// getNextBatch returns the next batch from the storage.
func (s *State) getNextBatch(ctx context.Context, batchNumber uint64) (*sequencerv2types.Batch, error) {
	ti := time.Now()
	defer telemetry.MeasureSince(ti, "consensus", "get_next_batch")

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
	ti := time.Now()
	defer telemetry.MeasureSince(ti, "consensus", "process_round")

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
				logger.Errorf("failed to check the commit for the batch %d: %v", s.fromBatchNumber, err)
				return false
			}
			if isFinalized {
				return true
			}
		}
	}
}

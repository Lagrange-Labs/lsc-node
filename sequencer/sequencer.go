package sequencer

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/Lagrange-Labs/lagrange-node/logger"
	networktypes "github.com/Lagrange-Labs/lagrange-node/network/types"
	"github.com/Lagrange-Labs/lagrange-node/rpcclient"
	rpctypes "github.com/Lagrange-Labs/lagrange-node/rpcclient/types"
	"github.com/Lagrange-Labs/lagrange-node/scinterface/committee"
	v2types "github.com/Lagrange-Labs/lagrange-node/sequencer/types/v2"
	storetypes "github.com/Lagrange-Labs/lagrange-node/store/types"
	"github.com/Lagrange-Labs/lagrange-node/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	// SyncInterval is the interval between two block syncs after fully synced.
	SyncInterval = 1 * time.Second
)

// CommitteeParams is the committee parameters.
type CommitteeParams struct {
	StartBlock     uint64
	Duration       uint64
	FreezeDuration uint64
}

// Sequencer is the main component of the lagrange node.
// - It is responsible for fetching batch headers from the given L2 chain.
// - It is responsible for fetching the operator information details from the committee smart contract.
type Sequencer struct {
	storage         storageInterface
	rpcClient       rpctypes.RpcClient
	chainID         uint32
	lastBlockNumber uint64

	stakingInterval    time.Duration
	updatedEpochNumber uint64
	currentEpochNumber uint64
	committeeParams    *CommitteeParams
	committeeSC        *committee.Committee
	etherClient        *ethclient.Client

	// Operators sync is a heavy operation, so we do it only once
	isOpertorsSynced bool
	operators        []networktypes.ClientNode

	ctx    context.Context
	cancel context.CancelFunc
}

// NewSequencer creates a new sequencer instance.
func NewSequencer(cfg *Config, rpcCfg *rpcclient.Config, storage storageInterface) (*Sequencer, error) {
	logger.Infof("Creating sequencer with config: %+v", cfg)

	client, err := ethclient.Dial(cfg.EthereumURL)
	if err != nil {
		logger.Errorf("failed to connect to ethereum: %v", err)
		return nil, err
	}

	committeeSC, err := committee.NewCommittee(common.HexToAddress(cfg.CommitteeSCAddress), client)
	if err != nil {
		logger.Errorf("failed to create committee contract: %v", err)
		return nil, err
	}

	rpcClient, err := rpcclient.NewClient(cfg.Chain, rpcCfg)
	if err != nil {
		logger.Errorf("failed to create rpc client: %v", err)
		return nil, err
	}

	chainID, err := rpcClient.GetChainID()
	if err != nil {
		logger.Errorf("failed to get chain ID: %v", err)
		return nil, err
	}

	updatedEpochNumber, err := committeeSC.UpdatedEpoch(nil, chainID)
	if err != nil {
		logger.Errorf("failed to get updated epoch number: %d err: %v", chainID, err)
		return nil, err
	}

	lastEpochNumber, err := storage.GetLastCommitteeEpochNumber(context.Background(), chainID)
	if err != nil {
		logger.Errorf("failed to get last committee epoch number: %v", err)
		return nil, err
	}

	committeeParams, err := committeeSC.CommitteeParams(nil, chainID)
	if err != nil {
		return nil, err
	}

	lastBlockNumber := uint64(0)
	batchNumber, err := storage.GetLastBatchNumber(context.Background(), chainID)
	if err != nil {
		if errors.Is(err, storetypes.ErrBatchNotFound) {
			logger.Infof("no batch found")
			rpcClient.SetBeginBlockNumber(cfg.FromL1BlockNumber)
		} else {
			logger.Errorf("failed to get last batch number: %v", err)
			return nil, err
		}
	} else {
		batch, err := storage.GetBatch(context.Background(), chainID, batchNumber)
		if err != nil {
			logger.Errorf("failed to get batch for batch number: %d error : %v", batchNumber, err)
			return nil, err
		}
		lastBlockNumber = batch.BatchHeader.ToBlockNumber()
		rpcClient.SetBeginBlockNumber(batch.L1BlockNumber())
	}
	if cfg.FromL2BlockNumber > lastBlockNumber {
		lastBlockNumber = cfg.FromL2BlockNumber - 1
	}

	ctx, cancel := context.WithCancel(context.Background())
	return &Sequencer{
		storage:         storage,
		rpcClient:       rpcClient,
		lastBlockNumber: lastBlockNumber,
		chainID:         uint32(chainID),

		etherClient:        client,
		committeeSC:        committeeSC,
		stakingInterval:    time.Duration(cfg.StakingCheckInterval),
		updatedEpochNumber: updatedEpochNumber.Uint64(),
		currentEpochNumber: lastEpochNumber,
		committeeParams: &CommitteeParams{
			StartBlock:     committeeParams.StartBlock.Uint64(),
			Duration:       committeeParams.Duration.Uint64(),
			FreezeDuration: committeeParams.FreezeDuration.Uint64(),
		},

		ctx:    ctx,
		cancel: cancel,
	}, nil
}

// GetChainID returns the chain ID.
func (s *Sequencer) GetChainID() uint32 {
	return s.chainID
}

// Start starts the sequencer.
func (s *Sequencer) Start() error {
	// start the committee update process
	go func() {
		if err := s.updateCommittee(); err != nil {
			logger.Errorf("failed to update committee root: %w", err)
		}

		ticker := time.NewTicker(s.stakingInterval)
		defer ticker.Stop()

		for {
			select {
			case <-s.ctx.Done():
				return
			case <-ticker.C:
				if err := s.fetchOperatorInfos(); err != nil {
					logger.Errorf("failed to fetch operator infos: %w", err)
				}
				if err := s.updateCommittee(); err != nil {
					logger.Errorf("failed to update committee root: %w", err)
				}
			}
		}
	}()

	logger.Infof("Sequencer batch fetching started from %d", s.lastBlockNumber+1)

	for {
		select {
		case <-s.ctx.Done():
			return nil
		default:
			nextBlockNumber := s.lastBlockNumber + 1
			batchHeader, err := s.rpcClient.GetBatchHeaderByNumber(nextBlockNumber)
			if err != nil {
				if errors.Is(err, rpctypes.ErrBatchNotFound) {
					logger.Infof("block %d is not ready", nextBlockNumber)
					time.Sleep(SyncInterval)
					continue
				}
				logger.Errorf("failed to get batch header: %v", err)
				return err
			}

			if err := s.storage.AddBatch(context.Background(), &v2types.Batch{
				BatchHeader:   batchHeader,
				SequencedTime: fmt.Sprintf("%d", time.Now().UnixMicro()),
			}); err != nil {
				logger.Errorf("failed to add batch: %v", err)
				return err
			}

			s.lastBlockNumber = batchHeader.ToBlockNumber()
			logger.Infof("batch block sequenced up to %d", s.lastBlockNumber)
			time.Sleep(1 * time.Millisecond)
		}
	}
}

// fetch the operator information details from the committee smart contract.
func (s *Sequencer) fetchOperatorInfos() error {
	// check if the given epoch is locked
	isLocked, _, err := s.committeeSC.IsLocked(nil, s.chainID)
	if err != nil {
		logger.Errorf("failed to check if the given epoch is locked: %w", err)
		return err
	}
	if !isLocked {
		if s.isOpertorsSynced {
			logger.Infof("the given epoch is not locked and the operators are already synced")
		}
		s.isOpertorsSynced = false
		return nil
	}

	if s.isOpertorsSynced {
		return nil
	}
	logger.Info("start fetching operator infos")

	// get the leaf count
	epochEndBlockNumber := (s.updatedEpochNumber+1)*s.committeeParams.Duration + s.committeeParams.StartBlock - 1
	committeeData, err := s.committeeSC.GetCommittee(nil, s.chainID, big.NewInt(int64(epochEndBlockNumber)))
	if err != nil {
		logger.Errorf("failed to get the committee data: %w", err)
	}
	leafCount := committeeData.CurrentCommittee.LeafCount.Int64()

	// get the operator details
	operators := make([]networktypes.ClientNode, 0)
	for i := int64(0); i < leafCount; i++ {
		addr, err := s.committeeSC.CommitteeAddrs(nil, s.chainID, big.NewInt(i))
		if err != nil {
			return err
		}
		votingPower, err := s.committeeSC.GetOperatorVotingPower(nil, addr, s.chainID)
		if err != nil {
			return err
		}
		blsPubKey, err := s.committeeSC.GetBlsPubKey(nil, addr)
		if err != nil {
			return err
		}
		pubKey := make([]byte, 0)
		pubKey = append(pubKey, common.LeftPadBytes(blsPubKey[0].Bytes(), 32)...)
		pubKey = append(pubKey, common.LeftPadBytes(blsPubKey[1].Bytes(), 32)...)
		operators = append(operators, networktypes.ClientNode{
			StakeAddress: addr.String(),
			VotingPower:  votingPower.Uint64(),
			PublicKey:    pubKey,
		})
	}

	s.operators = operators
	s.isOpertorsSynced = true

	return nil
}

func (s *Sequencer) updateCommittee() error {
	logger.Infof("update committee is started, current epoch number %d, updated epoch number %d", s.currentEpochNumber, s.updatedEpochNumber)
	// check if there are any missing committee roots
	// NOTE: this is for only test scenario, it should not be happened in the live network
	for epochNumber := s.currentEpochNumber + 1; epochNumber <= s.updatedEpochNumber+1; epochNumber++ {
		committeeRoot, err := s.fetchCommitteeRoot(epochNumber)
		if err != nil {
			return err
		}
		if err := s.storage.UpdateCommitteeRoot(s.ctx, committeeRoot); err != nil {
			return err
		}
		s.currentEpochNumber = epochNumber
	}

	// check if the committee tree needs to be updated
	blockNumber, err := s.etherClient.BlockNumber(s.ctx)
	if err != nil {
		return err
	}
	currentEpochNumber := (blockNumber-s.committeeParams.StartBlock)/s.committeeParams.Duration + 1

	for epochNumber := s.updatedEpochNumber + 1; epochNumber <= currentEpochNumber; epochNumber++ {
		if epochNumber > s.currentEpochNumber {
			committeeRoot, err := s.fetchCommitteeRoot(epochNumber)
			if err != nil {
				return err
			}
			if err := s.storage.UpdateCommitteeRoot(context.Background(), committeeRoot); err != nil {
				return err
			}
			s.currentEpochNumber = epochNumber
		}
	}

	if currentEpochNumber > s.currentEpochNumber {
		return fmt.Errorf("missing committee roots")
	}

	return nil
}

// fetch the committee root from the committee smart contract.
func (s *Sequencer) fetchCommitteeRoot(epochNumber uint64) (*v2types.CommitteeRoot, error) {
	epochEndBlockNumber := epochNumber*s.committeeParams.Duration + s.committeeParams.StartBlock - 1
	committeeData, err := s.committeeSC.GetCommittee(nil, s.chainID, big.NewInt(int64(epochEndBlockNumber)))

	if err != nil {
		logger.Errorf("failed to get committee data for block number %d, epoch number %d: %w", epochEndBlockNumber, epochNumber, err)
		return nil, err
	}

	committeeRoot := &v2types.CommitteeRoot{
		ChainID:               s.chainID,
		CurrentCommitteeRoot:  utils.Bytes2Hex(committeeData.CurrentCommittee.Root[:]),
		TotalVotingPower:      committeeData.CurrentCommittee.TotalVotingPower.Uint64(),
		EpochStartBlockNumber: epochEndBlockNumber - s.committeeParams.Duration + 1,
		EpochEndBlockNumber:   epochEndBlockNumber,
		EpochNumber:           epochNumber,
		Operators:             s.operators,
	}

	tvl := uint64(0)
	for _, operator := range s.operators {
		tvl += operator.VotingPower
	}

	if committeeRoot.TotalVotingPower != tvl {
		logger.Errorf("total voting power mismatch, committee root %+v, tvl %d", committeeRoot, tvl)
		return nil, fmt.Errorf("total voting power mismatch")
	}

	logger.Infof("fetched committee root %+v", committeeRoot)

	return committeeRoot, nil
}

// Stop stops the sequencer.
func (s *Sequencer) Stop() {
	if s != nil {
		s.cancel()
	}
}

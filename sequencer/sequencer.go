package sequencer

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/Lagrange-Labs/lagrange-node/logger"
	networktypes "github.com/Lagrange-Labs/lagrange-node/network/types"
	"github.com/Lagrange-Labs/lagrange-node/rpcclient"
	rpctypes "github.com/Lagrange-Labs/lagrange-node/rpcclient/types"
	"github.com/Lagrange-Labs/lagrange-node/scinterface/committee"
	v2types "github.com/Lagrange-Labs/lagrange-node/sequencer/types/v2"
	storetypes "github.com/Lagrange-Labs/lagrange-node/store/types"
	"github.com/Lagrange-Labs/lagrange-node/telemetry"
	"github.com/Lagrange-Labs/lagrange-node/utils"
)

const (
	// SyncInterval is the interval between two block syncs after fully synced.
	SyncInterval = 1 * time.Second
)

// CommitteeParams is the committee parameters.
type CommitteeParams struct {
	StartBlock     uint64
	GenesisBlock   uint64
	L1Bias         int64
	Duration       uint64
	FreezeDuration uint64
}

// Sequencer is the main component of the lagrange node.
// - It is responsible for fetching batch headers from the given L2 chain.
// - It is responsible for fetching the operator information details from the committee smart contract.
type Sequencer struct {
	storage           storageInterface
	rpcClient         rpctypes.RpcClient
	chainID           uint32
	fromL1BlockNumber uint64
	fromL1TxIndex     uint32
	lastBatchNumber   uint64

	stakingInterval    time.Duration
	updatedEpochNumber uint64
	currentEpochNumber uint64
	committeeParams    *CommitteeParams
	committeeSC        *committee.Committee
	etherClient        *ethclient.Client

	ctx    context.Context
	cancel context.CancelFunc
}

// NewSequencer creates a new sequencer instance.
func NewSequencer(cfg *Config, rpcCfg *rpcclient.Config, storage storageInterface) (*Sequencer, error) {
	logger.Infof("Creating sequencer for chain: %s", cfg.Chain)

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

	fromL2BlockNumber := uint64(0)
	fromL1BlockNumber := uint64(0)
	fromL1TxIndex := uint32(0)
	batchNumber, err := storage.GetLastBatchNumber(context.Background(), chainID)
	if err != nil {
		if errors.Is(err, storetypes.ErrBatchNotFound) {
			logger.Infof("no batch found")
			fromL1BlockNumber = cfg.FromL1BlockNumber
			fromL2BlockNumber = cfg.FromL2BlockNumber
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
		fromL2BlockNumber = batch.BatchHeader.FromBlockNumber()
		fromL1BlockNumber = batch.L1BlockNumber()
		fromL1TxIndex = batch.BatchHeader.L1TxIndex
	}
	rpcClient.SetBeginBlockNumber(fromL1BlockNumber, fromL2BlockNumber)

	ctx, cancel := context.WithCancel(context.Background())
	return &Sequencer{
		storage:           storage,
		rpcClient:         rpcClient,
		fromL1BlockNumber: fromL1BlockNumber,
		fromL1TxIndex:     fromL1TxIndex,
		lastBatchNumber:   batchNumber,
		chainID:           uint32(chainID),

		etherClient:        client,
		committeeSC:        committeeSC,
		stakingInterval:    time.Duration(cfg.StakingCheckInterval),
		updatedEpochNumber: updatedEpochNumber.Uint64(),
		currentEpochNumber: lastEpochNumber,
		committeeParams: &CommitteeParams{
			StartBlock:     committeeParams.StartBlock.Uint64(),
			GenesisBlock:   committeeParams.GenesisBlock.Uint64(),
			L1Bias:         committeeParams.L1Bias.Int64(),
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
				if err := s.updateCommittee(); err != nil {
					logger.Errorf("failed to update committee root: %w", err)
				}
			}
		}
	}()

	logger.Infof("Sequencer batch fetching started from L1 block number %d, tx index %d, batch number %d", s.fromL1BlockNumber, s.fromL1TxIndex, s.lastBatchNumber+1)

	for {
		select {
		case <-s.ctx.Done():
			return nil
		default:
			ti := time.Now()
			batchHeader, err := s.rpcClient.NextBatch()
			if err != nil {
				logger.Errorf("failed to get batch header: %v", err)
				return err
			}
			if s.lastBatchNumber > 0 {
				if batchHeader.L1BlockNumber < s.fromL1BlockNumber || (batchHeader.L1BlockNumber == s.fromL1BlockNumber && batchHeader.L1TxIndex <= s.fromL1TxIndex) {
					logger.Infof("no new batch found L1 block number %d, tx index %d, waiting for the next batch", batchHeader.L1BlockNumber, batchHeader.L1TxIndex)
					time.Sleep(SyncInterval)
					continue
				}
			}
			telemetry.MeasureSince(ti, "sequencer", "next_batch")

			s.lastBatchNumber++
			batchHeader.BatchNumber = s.lastBatchNumber
			if err := s.storage.AddBatch(context.Background(), &v2types.Batch{
				BatchHeader:   batchHeader,
				SequencedTime: fmt.Sprintf("%d", time.Now().UnixMicro()),
			}); err != nil {
				logger.Errorf("failed to add batch: %v", err)
				return err
			}

			s.fromL1BlockNumber = batchHeader.L1BlockNumber
			s.fromL1TxIndex = batchHeader.L1TxIndex
			logger.Infof("batch block sequenced up to %d", batchHeader.ToBlockNumber())
			time.Sleep(1 * time.Millisecond)
		}
	}
}

// fetch the operator information details from the committee smart contract.
func (s *Sequencer) fetchOperatorInfos(blockNumber *big.Int, leafCount uint32) ([]networktypes.ClientNode, error) {
	logger.Info("start fetching operator infos")
	opts := &bind.CallOpts{
		BlockNumber: blockNumber,
	}
	// get the operator details
	operators := make([]networktypes.ClientNode, leafCount)
	operatorIndex := int64(0)
	leafIndex := uint32(0)
	for leafIndex < leafCount {
		addr, err := s.committeeSC.CommitteeAddrs(opts, s.chainID, big.NewInt(operatorIndex))
		if err != nil {
			return nil, err
		}
		if bytes.Equal(addr.Bytes(), common.Address{}.Bytes()) {
			return nil, fmt.Errorf("leafCount %d is not matched with leafIndex %d", leafCount, leafIndex)
		}

		operatorStatus, err := s.committeeSC.OperatorsStatus(opts, addr)
		if err != nil {
			return nil, err
		}
		blsPubKeys, err := s.committeeSC.GetBlsPubKeys(opts, addr)
		if err != nil {
			return nil, err
		}
		votingPowers, err := s.committeeSC.GetBlsPubKeyVotingPowers(opts, addr, s.chainID)
		if err != nil {
			return nil, err
		}
		for i, votingPower := range votingPowers {
			pubKey := make([]byte, 0)
			pubKey = append(pubKey, common.LeftPadBytes(blsPubKeys[i][0].Bytes(), 32)...)
			pubKey = append(pubKey, common.LeftPadBytes(blsPubKeys[i][1].Bytes(), 32)...)
			operators[leafIndex] = networktypes.ClientNode{
				StakeAddress: addr.Hex(),
				SignAddress:  operatorStatus.SignAddress.Hex(),
				VotingPower:  votingPower.Uint64(),
				PublicKey:    utils.Bytes2Hex(pubKey),
			}
			leafIndex++
		}
		operatorIndex++
	}

	return operators, nil
}

func (s *Sequencer) updateCommittee() error {
	logger.Infof("update committee is started, current epoch number %d, updated epoch number %d", s.currentEpochNumber, s.updatedEpochNumber)

	// check if the committee tree is updated
	updatedEpochNumber, err := s.committeeSC.UpdatedEpoch(nil, s.chainID)
	if err != nil {
		return fmt.Errorf("failed to get updated epoch number: %w", err)
	}
	if updatedEpochNumber.Uint64() > s.updatedEpochNumber {
		s.updatedEpochNumber = updatedEpochNumber.Uint64()
	}

	for epochNumber := s.currentEpochNumber + 1; epochNumber <= s.updatedEpochNumber; epochNumber++ {
		committeeRoot, err := s.fetchCommitteeRoot(epochNumber)
		if err != nil {
			return err
		}
		if err := s.storage.UpdateCommitteeRoot(s.ctx, committeeRoot); err != nil {
			return err
		}
		s.currentEpochNumber = epochNumber
	}

	return nil
}

// fetch the committee root from the committee smart contract.
func (s *Sequencer) fetchCommitteeRoot(epochNumber uint64) (*v2types.CommitteeRoot, error) {
	ti := time.Now()
	defer telemetry.MeasureSince(ti, "sequencer", "fetch_committee_root")

	epochEndBlockNumber := (epochNumber+1)*s.committeeParams.Duration + s.committeeParams.StartBlock - 1
	epochEndBlockNumber = uint64(int64(epochEndBlockNumber) - s.committeeParams.L1Bias)
	committeeData, err := s.committeeSC.GetCommittee(nil, s.chainID, big.NewInt(int64(epochEndBlockNumber)))
	if err != nil {
		logger.Errorf("failed to get committee data for block number %d, epoch number %d: %w", epochEndBlockNumber, epochNumber, err)
		return nil, err
	}
	if committeeData.LeafCount == 0 {
		logger.Warnf("no operator in the committee for block number %d, epoch number %d", epochEndBlockNumber, epochNumber)
		return nil, fmt.Errorf("no operator in the committee epoch number %d", epochNumber)
	}

	operators, err := s.fetchOperatorInfos(committeeData.UpdatedBlock, committeeData.LeafCount)
	if err != nil {
		logger.Errorf("failed to fetch operator infos: %w", err)
		return nil, err
	}
	tvl := uint64(0)
	for _, operator := range operators {
		tvl += operator.VotingPower
	}
	epochStart := epochEndBlockNumber - s.committeeParams.Duration + 1
	if epochNumber == uint64(1) {
		epochStart = uint64(int64(s.committeeParams.GenesisBlock) - s.committeeParams.L1Bias)
	}

	committeeRoot := &v2types.CommitteeRoot{
		ChainID:               s.chainID,
		CurrentCommitteeRoot:  utils.Bytes2Hex(committeeData.Root[:]),
		TotalVotingPower:      tvl,
		EpochStartBlockNumber: epochStart,
		EpochEndBlockNumber:   epochEndBlockNumber,
		EpochNumber:           epochNumber,
		Operators:             operators,
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

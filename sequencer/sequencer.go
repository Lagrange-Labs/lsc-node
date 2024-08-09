package sequencer

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/Lagrange-Labs/lagrange-node/core"
	"github.com/Lagrange-Labs/lagrange-node/core/logger"
	"github.com/Lagrange-Labs/lagrange-node/rpcclient"
	rpctypes "github.com/Lagrange-Labs/lagrange-node/rpcclient/types"
	"github.com/Lagrange-Labs/lagrange-node/scinterface/committee"
	"github.com/Lagrange-Labs/lagrange-node/scinterface/eigendm"
	"github.com/Lagrange-Labs/lagrange-node/scinterface/voteweigher"
	v2types "github.com/Lagrange-Labs/lagrange-node/sequencer/types/v2"
	servertypes "github.com/Lagrange-Labs/lagrange-node/server/types"
	storetypes "github.com/Lagrange-Labs/lagrange-node/store/types"
	"github.com/Lagrange-Labs/lagrange-node/telemetry"
)

const (
	// SyncInterval is the interval between two block syncs after fully synced.
	SyncInterval = 1 * time.Second

	maskBitLen = 112
)

var (
	operatorSharesIncreased = crypto.Keccak256Hash([]byte("OperatorSharesIncreased(address,address,address,uint256)"))
	operatorSharesDecreased = crypto.Keccak256Hash([]byte("OperatorSharesDecreased(address,address,address,uint256)"))
	updateCommittee         = crypto.Keccak256Hash([]byte("UpdateCommittee(uint256,uint256,bytes32)"))
	// This is the mask to get the block number from the UpdatedBlockNumber in Holesky network.
	l1BlockNumberMask = new(big.Int).Sub(new(big.Int).Lsh(big.NewInt(1), maskBitLen), big.NewInt(1))
)

// CommitteeParams is the committee parameters.
type CommitteeParams struct {
	StartBlock       uint64
	GenesisBlock     uint64
	L1Bias           int64
	Duration         uint64
	FreezeDuration   uint64
	QuorumNumber     uint8
	MinWeight        uint64
	MaxWeight        uint64
	WeightingDivisor *big.Int

	tokenMultipliers  map[common.Address]*big.Int
	committeeSCAddr   common.Address
	voteweigherSCAddr common.Address
	eigenDMSCAddr     common.Address
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
	voteweigherSC      *voteweigher.Voteweigher
	eigendmSC          *eigendm.Eigendm
	etherClient        *ethclient.Client

	ctx    context.Context
	cancel context.CancelFunc
}

// NewSequencer creates a new sequencer instance.
func NewSequencer(cfg *Config, rpcCfg *rpcclient.Config, storage storageInterface) (*Sequencer, error) {
	logger.Infof("Creating sequencer for chain: %s", cfg.Chain)

	rpcClient, err := rpcclient.NewClient(cfg.Chain, rpcCfg, false)
	if err != nil {
		logger.Errorf("failed to create rpc client: %v", err)
		return nil, err
	}
	chainID, err := rpcClient.GetChainID()
	if err != nil {
		logger.Errorf("failed to get chain ID: %v", err)
		return nil, err
	}

	client, err := ethclient.Dial(cfg.EthereumURL)
	if err != nil {
		logger.Errorf("failed to connect to ethereum: %v", err)
		return nil, err
	}

	committeeSCAddr := common.HexToAddress(cfg.CommitteeSCAddress)
	committeeSC, err := committee.NewCommittee(committeeSCAddr, client)
	if err != nil {
		logger.Errorf("failed to create committee contract: %v", err)
		return nil, err
	}
	committeeParams, err := committeeSC.CommitteeParams(nil, chainID)
	if err != nil {
		return nil, err
	}

	voteweigherSCAddr, err := committeeSC.VoteWeigher(nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get VoteWeigher address: %v", err)
	}
	voteweigherSC, err := voteweigher.NewVoteweigher(voteweigherSCAddr, client)
	if err != nil {
		return nil, fmt.Errorf("failed to create a new eigenDM SC: %v", err)
	}
	weightingDivisor, err := voteweigherSC.WEIGHTINGDIVISOR(nil)
	if err != nil {
		return nil, fmt.Errorf("failed to get weighting divisor: %v", err)
	}
	tokenMultipliers := make(map[common.Address]*big.Int)
	for i := uint8(0); i < 255; i++ {
		tokenMultiplier, err := voteweigherSC.QuorumMultipliers(nil, committeeParams.QuorumNumber, big.NewInt(int64(i)))
		if err != nil {
			logger.Warnf("failed to get token multiplier: %v", err)
			break
		}
		tokenMultipliers[tokenMultiplier.Token] = tokenMultiplier.Multiplier
	}

	eigendmSCAddr := common.HexToAddress(cfg.EigenDMSCAddress)
	eigendmSC, err := eigendm.NewEigendm(eigendmSCAddr, client)
	if err != nil {
		return nil, fmt.Errorf("failed to create a new eigenDM SC: %v", err)
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

	fromL1BlockNumber := uint64(0)
	fromL1TxIndex := uint32(0)
	batchNumber, err := storage.GetLastBatchNumber(context.Background(), chainID)
	if err != nil {
		if errors.Is(err, storetypes.ErrBatchNotFound) {
			logger.Infof("no batch found")
			fromL1BlockNumber = cfg.FromL1BlockNumber
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
		fromL1BlockNumber = batch.L1BlockNumber()
		fromL1TxIndex = batch.BatchHeader.L1TxIndex
	}
	rpcClient.SetBeginBlockNumber(fromL1BlockNumber)

	ctx, cancel := context.WithCancel(context.Background())
	return &Sequencer{
		storage:           storage,
		rpcClient:         rpcClient,
		fromL1BlockNumber: fromL1BlockNumber,
		fromL1TxIndex:     fromL1TxIndex,
		lastBatchNumber:   batchNumber,
		chainID:           chainID,

		etherClient:        client,
		committeeSC:        committeeSC,
		voteweigherSC:      voteweigherSC,
		eigendmSC:          eigendmSC,
		stakingInterval:    time.Duration(cfg.StakingCheckInterval),
		updatedEpochNumber: updatedEpochNumber.Uint64(),
		currentEpochNumber: lastEpochNumber,
		committeeParams: &CommitteeParams{
			StartBlock:        committeeParams.StartBlock.Uint64(),
			GenesisBlock:      committeeParams.GenesisBlock.Uint64(),
			L1Bias:            committeeParams.L1Bias.Int64(),
			Duration:          committeeParams.Duration.Uint64(),
			FreezeDuration:    committeeParams.FreezeDuration.Uint64(),
			QuorumNumber:      committeeParams.QuorumNumber,
			MinWeight:         committeeParams.MinWeight.Uint64(),
			MaxWeight:         committeeParams.MaxWeight.Uint64(),
			WeightingDivisor:  weightingDivisor,
			tokenMultipliers:  tokenMultipliers,
			voteweigherSCAddr: voteweigherSCAddr,
			committeeSCAddr:   committeeSCAddr,
			eigenDMSCAddr:     eigendmSCAddr,
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
			batchHeader.L2FromBlockNumber = batchHeader.FromBlockNumber()
			batchHeader.L2ToBlockNumber = batchHeader.ToBlockNumber()
			if err := s.storage.AddBatch(context.Background(), &v2types.Batch{
				BatchHeader:   batchHeader,
				SequencedTime: fmt.Sprintf("%d", time.Now().UnixMicro()),
			}); err != nil {
				logger.Errorf("failed to add batch: %v", err)
				return err
			}

			s.fromL1BlockNumber = batchHeader.L1BlockNumber
			s.fromL1TxIndex = batchHeader.L1TxIndex
			logger.Infof("batch L2 block sequenced up to %d, count: %d", batchHeader.ToBlockNumber(), batchHeader.ToBlockNumber()-batchHeader.FromBlockNumber()+1)
			time.Sleep(1 * time.Millisecond)
		}
	}
}

// fetch the operator information details from the committee smart contract.
func (s *Sequencer) fetchOperatorInfos(blockNumber *big.Int, leafCount uint32) ([]servertypes.ClientNode, error) {
	logger.Info("start fetching operator infos")
	opts := &bind.CallOpts{
		BlockNumber: blockNumber,
	}
	voteWeightChanges, err := s.getVoteWeightChanges(blockNumber)
	if err != nil {
		logger.Errorf("failed to get vote weight changes: %v", err)
		return nil, err
	}
	// get the operator details
	operators := make([]servertypes.ClientNode, leafCount)
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
		var votingPowers []*big.Int

		votingPowers, err = s.committeeSC.GetBlsPubKeyVotingPowers(opts, addr, s.chainID)
		if err != nil {
			return nil, err
		}

		if voteWeightChanges[addr] != 0 {
			voteWeightOrg, err := s.voteweigherSC.WeightOfOperator(opts, s.committeeParams.QuorumNumber, addr)
			if err != nil {
				return nil, err
			}
			voteWeight := voteWeightOrg.Int64() - voteWeightChanges[addr]
			if voteWeight < 0 {
				return nil, fmt.Errorf("vote weight %d is less than 0", voteWeight)
			}
			blsKeysCnt := len(blsPubKeys)
			votingPowers, err = s.distributeVotingPowers(uint64(voteWeight), uint64(blsKeysCnt), s.committeeParams.MinWeight, s.committeeParams.MaxWeight)
			if err != nil {
				return nil, err
			}
		}
		for i, votingPower := range votingPowers {
			pubKey := make([]byte, 0)
			pubKey = append(pubKey, common.LeftPadBytes(blsPubKeys[i][0].Bytes(), 32)...)
			pubKey = append(pubKey, common.LeftPadBytes(blsPubKeys[i][1].Bytes(), 32)...)
			operators[leafIndex] = servertypes.ClientNode{
				StakeAddress: addr.Hex(),
				SignAddress:  operatorStatus.SignAddress.Hex(),
				VotingPower:  votingPower.Uint64(),
				PublicKey:    core.Bytes2Hex(pubKey),
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
		if err := committeeRoot.Verify(); err != nil {
			logger.Errorf("failed to verify committee root: %v", err)
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

	committeeData, err := s.committeeSC.Committees(nil, s.chainID, big.NewInt(int64(epochNumber)))
	if err != nil {
		logger.Errorf("failed to get committee data for epoch number %d: %w", epochNumber, err)
		return nil, err
	}
	if committeeData.LeafCount == 0 {
		logger.Warnf("no operator in the committee for epoch number %d", epochNumber)
		return nil, fmt.Errorf("no operator in the committee epoch number %d", epochNumber)
	}

	epochStartBlockNumber := big.NewInt(0).Rsh(committeeData.UpdatedBlock, maskBitLen).Uint64()
	updatedBlockNumber := big.NewInt(0).And(committeeData.UpdatedBlock, l1BlockNumberMask)
	if epochStartBlockNumber == 0 {
		epochStartBlockNumber = uint64(updatedBlockNumber.Int64() - s.committeeParams.L1Bias)
	}
	operators, err := s.fetchOperatorInfos(updatedBlockNumber, committeeData.LeafCount)
	if err != nil {
		logger.Errorf("failed to fetch operator infos: %w", err)
		return nil, err
	}
	tvl := uint64(0)
	for _, operator := range operators {
		tvl += operator.VotingPower
	}
	if epochNumber == uint64(1) {
		epochStartBlockNumber = uint64(int64(s.committeeParams.GenesisBlock) - s.committeeParams.L1Bias)
	}

	committeeRoot := &v2types.CommitteeRoot{
		ChainID:               s.chainID,
		CurrentCommitteeRoot:  core.Bytes2Hex(committeeData.Root[:]),
		TotalVotingPower:      tvl,
		EpochStartBlockNumber: epochStartBlockNumber,
		EpochNumber:           epochNumber,
		Operators:             operators,
	}

	logger.Infof("fetched committee root %+v", committeeRoot)

	return committeeRoot, nil
}

func (s *Sequencer) getVoteWeightChanges(blockNumber *big.Int) (map[common.Address]int64, error) {
	var voteWeightChanges map[common.Address]int64
	queryFilter := ethereum.FilterQuery{
		FromBlock: blockNumber,
		ToBlock:   blockNumber,
		Addresses: []common.Address{s.committeeParams.committeeSCAddr},
		Topics:    [][]common.Hash{{updateCommittee}, {common.HexToHash(fmt.Sprintf("%x", s.chainID))}},
	}
	logs, err := s.etherClient.FilterLogs(s.ctx, queryFilter)
	if err != nil {
		logger.Errorf("failed to filter logs: %v", err)
		return nil, err
	}
	if len(logs) != 1 {
		return nil, fmt.Errorf("invalid logs count: %d", len(logs))
	}
	updateLogIndex := logs[0].Index

	voteWeightChanges = make(map[common.Address]int64)

	queryFilter = ethereum.FilterQuery{
		FromBlock: blockNumber,
		ToBlock:   blockNumber,
		Addresses: []common.Address{s.committeeParams.eigenDMSCAddr},
		Topics:    [][]common.Hash{{operatorSharesIncreased, operatorSharesDecreased}},
	}

	logs, err = s.etherClient.FilterLogs(s.ctx, queryFilter)
	if err != nil {
		logger.Errorf("failed to filter logs: %v", err)
		return nil, err
	}

	operatorTokenShareChanges := make(map[common.Address]map[common.Address]*big.Int)
	for _, vLog := range logs {
		if vLog.Index < updateLogIndex {
			continue
		}
		if vLog.Topics[0] != operatorSharesIncreased && vLog.Topics[0] != operatorSharesDecreased {
			continue
		}
		var (
			operator common.Address
			token    common.Address
			shares   *big.Int
		)
		if vLog.Topics[0] == operatorSharesIncreased {
			event, err := s.eigendmSC.ParseOperatorSharesIncreased(vLog)
			if err != nil {
				logger.Errorf("failed to parse operator shares increased: %v", err)
				return nil, err
			}
			operator = event.Operator
			token = event.Strategy
			shares = event.Shares
		} else if vLog.Topics[0] == operatorSharesDecreased {
			event, err := s.eigendmSC.ParseOperatorSharesDecreased(vLog)
			if err != nil {
				logger.Errorf("failed to parse operator shares decreased: %v", err)
				return nil, err
			}
			operator = event.Operator
			token = event.Strategy
			shares = big.NewInt(0).Neg(event.Shares)
		}

		if _, ok := s.committeeParams.tokenMultipliers[token]; ok {
			if _, ok := operatorTokenShareChanges[operator]; !ok {
				operatorTokenShareChanges[operator] = make(map[common.Address]*big.Int)
			}
			if _, ok := operatorTokenShareChanges[operator][token]; !ok {
				operatorTokenShareChanges[operator][token] = big.NewInt(0)
			}
			operatorTokenShareChanges[operator][token] = big.NewInt(0).Add(operatorTokenShareChanges[operator][token], shares)
		}
	}

	opts := &bind.CallOpts{
		BlockNumber: blockNumber,
	}

	for operator, tokenShareChanges := range operatorTokenShareChanges {
		voteWeightChanges[operator] = 0
		for token, shareChange := range tokenShareChanges {
			multiplier := s.committeeParams.tokenMultipliers[token]
			orgShares, err := s.eigendmSC.OperatorShares(opts, operator, token)
			if err != nil {
				return nil, err
			}
			voteWeightChanges[operator] += s.calcVotingPower(orgShares, multiplier) - s.calcVotingPower(big.NewInt(0).Sub(orgShares, shareChange), multiplier)
		}
	}

	return voteWeightChanges, nil
}

func (s *Sequencer) calcVotingPower(share *big.Int, multiplier *big.Int) int64 {
	return big.NewInt(0).Div(big.NewInt(0).Mul(share, multiplier), s.committeeParams.WeightingDivisor).Int64()
}

func (s *Sequencer) distributeVotingPowers(voteWeight uint64, blsKeysCnt uint64, minWeight uint64, maxWeight uint64) ([]*big.Int, error) {
	if voteWeight < minWeight {
		return make([]*big.Int, 0), nil
	}

	countLimit := ((voteWeight - 1) / maxWeight) + 1
	if countLimit < blsKeysCnt {
		blsKeysCnt = countLimit
	}

	var votingPowers []*big.Int = make([]*big.Int, blsKeysCnt)

	amountLimit := maxWeight * blsKeysCnt

	if voteWeight > amountLimit {
		voteWeight = amountLimit
	}

	index := 0
	remained := voteWeight
	for remained >= maxWeight+minWeight {
		votingPowers[index] = big.NewInt(int64(maxWeight))
		index++
		remained -= maxWeight
	}
	if remained > maxWeight {
		votingPowers[index] = big.NewInt(int64(minWeight))
		index++
		votingPowers[index] = big.NewInt(int64(remained - minWeight))
	} else {
		votingPowers[index] = big.NewInt(int64(remained))
	}

	return votingPowers, nil
}

// Stop stops the sequencer.
func (s *Sequencer) Stop() {
	if s != nil && s.ctx != nil {
		s.cancel()
	}
}

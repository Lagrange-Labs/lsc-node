package governance

import (
	"context"
	"fmt"
	"math/big"
	"time"

	"github.com/Lagrange-Labs/lagrange-node/crypto"
	"github.com/Lagrange-Labs/lagrange-node/governance/types"
	"github.com/Lagrange-Labs/lagrange-node/logger"
	networktypes "github.com/Lagrange-Labs/lagrange-node/network/types"
	"github.com/Lagrange-Labs/lagrange-node/scinterface/committee"
	"github.com/Lagrange-Labs/lagrange-node/utils"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// CommitteeParams is the committee parameters.
type CommitteeParams struct {
	StartBlock     uint64
	Duration       uint64
	FreezeDuration uint64
}

// Governance is the module which is responsible for the staking and slashing.
type Governance struct {
	stakingInterval    time.Duration
	blsScheme          crypto.BLSScheme
	chainID            uint32
	updatedEpochNumber uint64
	currentEpochNumber uint64
	committeeParams    *CommitteeParams

	committeeSC *committee.Committee
	storage     storageInterface
	etherClient *ethclient.Client
	auth        *bind.TransactOpts

	// Operators sync is a heavy operation, so we do it only once
	isOpertorsSynced bool
	operators        []networktypes.ClientNode

	ctx    context.Context
	cancel context.CancelFunc
}

// NewGovernance creates a new Governance instance.
func NewGovernance(cfg *types.Config, blsCurve crypto.BLSCurve, chainID uint32, storage storageInterface) (*Governance, error) {
	logger.Infof("Creating governance with config: %+v", cfg)

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

	auth, err := utils.GetSigner(context.Background(), client, cfg.PrivateKey)
	if err != nil {
		logger.Errorf("failed to get signer: %v", err)
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

	ctx, cancel := context.WithCancel(context.Background())
	return &Governance{
		stakingInterval:    time.Duration(cfg.StakingCheckInterval),
		committeeSC:        committeeSC,
		blsScheme:          crypto.NewBLSScheme(blsCurve),
		chainID:            chainID,
		storage:            storage,
		etherClient:        client,
		auth:               auth,
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

// Start starts the governance process.
func (g *Governance) Start() {
	if err := g.updateCommittee(); err != nil {
		logger.Errorf("failed to update committee root: %w", err)
	}

	ticker := time.NewTicker(g.stakingInterval)
	defer ticker.Stop()

	for {
		select {
		case <-g.ctx.Done():
			return
		case <-ticker.C:
			if err := g.fetchOperatorInfos(); err != nil {
				logger.Errorf("failed to fetch operator infos: %w", err)
			}
			if err := g.updateCommittee(); err != nil {
				logger.Errorf("failed to update committee root: %w", err)
			}
		}
	}
}

// Stop stops the governance process.
func (g *Governance) Stop() {
	g.cancel()
}

func (g *Governance) updateCommittee() error {
	logger.Infof("update committee is started, current epoch number %d, updated epoch number %d", g.currentEpochNumber, g.updatedEpochNumber)
	// check if there are any missing committee roots
	// NOTE: this is for only test scenario, it should not be happened in the live network
	for epochNumber := g.currentEpochNumber + 1; epochNumber <= g.updatedEpochNumber+1; epochNumber++ {
		committeeRoot, err := g.fetchCommitteeRoot(epochNumber)
		if err != nil {
			return err
		}
		if err := g.storage.UpdateCommitteeRoot(g.ctx, committeeRoot); err != nil {
			return err
		}
		g.currentEpochNumber = epochNumber
	}

	// check if the committee tree needs to be updated
	blockNumber, err := g.etherClient.BlockNumber(g.ctx)
	if err != nil {
		return err
	}
	currentEpochNumber := (blockNumber-g.committeeParams.StartBlock)/g.committeeParams.Duration + 1

	for epochNumber := g.updatedEpochNumber + 1; epochNumber <= currentEpochNumber; epochNumber++ {
		if epochNumber > g.currentEpochNumber {
			committeeRoot, err := g.fetchCommitteeRoot(epochNumber)
			if err != nil {
				return err
			}
			if err := g.storage.UpdateCommitteeRoot(context.Background(), committeeRoot); err != nil {
				return err
			}
			g.currentEpochNumber = epochNumber
		}

		// check if the committee tree needs to be updated
		updatable, err := g.committeeSC.IsUpdatable(nil, g.chainID, big.NewInt(int64(epochNumber)))
		if err != nil {
			logger.Errorf("failed to check if the committee tree is updatable: %w", err)
			return err
		}
		if updatable {
			logger.Infof("updating committee tree for epoch %d", epochNumber)
			tx, err := g.committeeSC.Update(g.auth, g.chainID, big.NewInt(int64(epochNumber)))
			if err != nil {
				logger.Errorf("failed to update committee tree: %w", err)
				return err
			}
			// wait for the transaction to be mined
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()
			receipt, err := bind.WaitMined(ctx, g.etherClient, tx)
			if err != nil {
				logger.Errorf("failed to wait for transaction to be mined: %w", err)
				return fmt.Errorf("failed to wait for transaction to be mined: %w", err)
			}
			if receipt.Status != 1 {
				logger.Errorf("transaction failed: %v", receipt)
				return fmt.Errorf("transaction failed: %v", receipt)
			}
			g.updatedEpochNumber = epochNumber
		}
	}

	if currentEpochNumber > g.currentEpochNumber {
		return fmt.Errorf("missing committee roots")
	}

	return nil
}

// fetch the operator information details from the committee smart contract.
func (g *Governance) fetchOperatorInfos() error {
	// check if the given epoch is locked
	isLocked, _, err := g.committeeSC.IsLocked(nil, g.chainID)
	if err != nil {
		logger.Errorf("failed to check if the given epoch is locked: %w", err)
		return err
	}
	if !isLocked {
		if g.isOpertorsSynced {
			logger.Infof("the given epoch is not locked and the operators are already synced")
		}
		g.isOpertorsSynced = false
		return nil
	}

	if g.isOpertorsSynced {
		return nil
	}
	logger.Info("start fetching operator infos")

	// get the leaf count
	epochEndBlockNumber := (g.updatedEpochNumber+1)*g.committeeParams.Duration + g.committeeParams.StartBlock - 1
	committeeData, err := g.committeeSC.GetCommittee(nil, g.chainID, big.NewInt(int64(epochEndBlockNumber)))
	if err != nil {
		logger.Errorf("failed to get the committee data: %w", err)
	}
	leafCount := committeeData.CurrentCommittee.LeafCount.Int64()

	// get the operator details
	operators := make([]networktypes.ClientNode, 0)
	for i := int64(0); i < leafCount; i++ {
		addr, err := g.committeeSC.CommitteeAddrs(nil, g.chainID, big.NewInt(i))
		if err != nil {
			return err
		}
		votingPower, err := g.committeeSC.GetOperatorVotingPower(nil, addr, g.chainID)
		if err != nil {
			return err
		}
		blsPubKey, err := g.committeeSC.GetBlsPubKey(nil, addr)
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

	g.operators = operators
	g.isOpertorsSynced = true

	return nil
}

// fetch the committee root from the committee smart contract.
func (g *Governance) fetchCommitteeRoot(epochNumber uint64) (*types.CommitteeRoot, error) {
	epochEndBlockNumber := epochNumber*g.committeeParams.Duration + g.committeeParams.StartBlock - 1
	committeeData, err := g.committeeSC.GetCommittee(nil, g.chainID, big.NewInt(int64(epochEndBlockNumber)))

	if err != nil {
		logger.Errorf("failed to get committee data for block number %d, epoch number %d: %w", epochEndBlockNumber, epochNumber, err)
		return nil, err
	}

	committeeRoot := &types.CommitteeRoot{
		ChainID:               g.chainID,
		CurrentCommitteeRoot:  utils.Bytes2Hex(committeeData.CurrentCommittee.Root[:]),
		TotalVotingPower:      committeeData.CurrentCommittee.TotalVotingPower.Uint64(),
		EpochStartBlockNumber: epochEndBlockNumber - g.committeeParams.Duration + 1,
		EpochEndBlockNumber:   epochEndBlockNumber,
		EpochNumber:           epochNumber,
		Operators:             g.operators,
	}

	tvl := uint64(0)
	for _, operator := range g.operators {
		tvl += operator.VotingPower
	}

	if committeeRoot.TotalVotingPower != tvl {
		logger.Errorf("total voting power mismatch, committee root %+v, tvl %d", committeeRoot, tvl)
		return nil, fmt.Errorf("total voting power mismatch")
	}

	logger.Infof("fetched committee root %+v", committeeRoot)

	return committeeRoot, nil
}

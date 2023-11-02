package governance

import (
	"context"
	"fmt"
	"math/big"
	"time"

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
	stakingInterval time.Duration

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
func NewGovernance(cfg *types.Config, chainID uint32, storage storageInterface) (*Governance, error) {
	client, err := ethclient.Dial(cfg.EthereumURL)
	if err != nil {
		return nil, err
	}

	committeeSC, err := committee.NewCommittee(common.HexToAddress(cfg.CommitteeSCAddress), client)
	if err != nil {
		return nil, err
	}

	auth, err := utils.GetSigner(context.Background(), client, cfg.PrivateKey)
	if err != nil {
		return nil, err
	}

	updatedEpochNumber, err := committeeSC.UpdatedEpoch(nil, chainID)
	if err != nil {
		return nil, err
	}

	lastEpochNumber, err := storage.GetLastCommitteeEpochNumber(context.Background(), chainID)
	if err != nil {
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
		logger.Fatalf("failed to update committee root: %w", err)
	}

	ticker := time.NewTicker(g.stakingInterval)
	defer ticker.Stop()

	for {
		select {
		case <-g.ctx.Done():
			return
		case <-ticker.C:
			if err := g.updateCommittee(); err != nil {
				logger.Errorf("failed to update committee root: %w", err)
			}
			if err := g.updateNodeStatuses(); err != nil {
				logger.Errorf("failed to update node status: %w", err)
			}
		}
	}
}

// Stop stops the governance process.
func (g *Governance) Stop() {
	g.cancel()
}

func (g *Governance) updateNodeStatuses() error {
	nodeStatuses := []networktypes.NodeStatus{networktypes.NodeJoined}

	nodes, err := g.storage.GetNodesByStatuses(g.ctx, nodeStatuses, g.chainID)
	if err != nil {
		return err
	}
	if len(nodes) > 0 {
		logger.Infof("updating nodes %v", nodes)
	}
	for _, node := range nodes {
		sNode, err := g.committeeSC.Operators(nil, common.HexToAddress(node.StakeAddress))
		if err != nil {
			logger.Warnf("failed to get node %s: %v", node.StakeAddress, err)
			continue
		}
		logger.Infof("node found %v", sNode)
		node.VotingPower = uint64(sNode.Amount.Int64())
		if node.VotingPower == 0 {
			logger.Infof("node %s has 0 voting power", node.StakeAddress)
			node.Status = networktypes.NodeUnstaked
		} else if sNode.Slashed {
			node.Status = networktypes.NodeSlashed
		} else {
			node.Status = networktypes.NodeRegistered
		}
		if err := g.storage.UpdateNode(g.ctx, &node); err != nil {
			return err
		}
	}

	return nil
}

func (g *Governance) updateCommittee() error {
	logger.Infof("update committee is started, current epoch number %d, updated epoch number %d", g.currentEpochNumber, g.updatedEpochNumber)
	g.isOpertorsSynced = false
	// check if there are any missing committee roots
	for epochNumber := g.currentEpochNumber + 1; epochNumber <= g.updatedEpochNumber+1; epochNumber++ {
		epochEndBlockNumber := epochNumber*g.committeeParams.Duration + g.committeeParams.StartBlock - 1

		committeeRoot, err := g.fetchCommitteeRoot(epochEndBlockNumber, epochNumber)
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

	currentEpochNumber, err := g.committeeSC.GetEpochNumber(nil, g.chainID, big.NewInt(int64(blockNumber)))
	if err != nil {
		return err
	}

	for epochNumber := g.updatedEpochNumber + 1; epochNumber <= currentEpochNumber.Uint64(); epochNumber++ {
		if epochNumber > g.currentEpochNumber {
			epochEndBlockNumber := epochNumber*g.committeeParams.Duration + g.committeeParams.StartBlock - 1

			committeeRoot, err := g.fetchCommitteeRoot(epochEndBlockNumber, epochNumber)
			if err != nil {
				return err
			}

			if err := g.storage.UpdateCommitteeRoot(g.ctx, committeeRoot); err != nil {
				return err
			}

			g.currentEpochNumber = epochNumber
		}

		// check if the committee tree needs to be updated
		updatable, err := g.committeeSC.IsUpdatable(nil, g.chainID, big.NewInt(int64(epochNumber)))
		if err != nil {
			return err
		}
		if updatable {
			logger.Infof("updating committee tree for epoch %d", epochNumber)
			tx, err := g.committeeSC.Update(g.auth, g.chainID, big.NewInt(int64(epochNumber)))
			if err != nil {
				return err
			}
			// wait for the transaction to be mined
			receipt, err := bind.WaitMined(g.ctx, g.etherClient, tx)
			if err != nil {
				return fmt.Errorf("failed to wait for transaction to be mined: %w", err)
			}
			if receipt.Status != 1 {
				return fmt.Errorf("transaction failed: %v", receipt)
			}

			g.updatedEpochNumber = epochNumber
		}
	}

	if currentEpochNumber.Uint64() > g.currentEpochNumber {
		return fmt.Errorf("missing committee roots")
	}

	return nil
}

// fetch the committee root from the smart contract
func (g *Governance) fetchCommitteeRoot(blockNumber, epochNumber uint64) (*types.CommitteeRoot, error) {
	committeeData, err := g.committeeSC.GetCommittee(nil, g.chainID, big.NewInt(int64(blockNumber)))

	if err != nil {
		return nil, err
	}

	committeeRoot := &types.CommitteeRoot{
		ChainID:              g.chainID,
		CurrentCommitteeRoot: common.Bytes2Hex(committeeData.CurrentCommittee.Root.Bytes()),
		TotalVotingPower:     committeeData.CurrentCommittee.TotalVotingPower.Uint64(),
		EpochBlockNumber:     blockNumber,
		EpochNumber:          epochNumber,
	}

	if committeeRoot.TotalVotingPower == 0 {
		logger.Errorf("total voting power is 0, committee root %v, epoch number %d", committeeRoot, epochNumber)
		return nil, fmt.Errorf("total voting power is 0")
	}

	if !g.isOpertorsSynced {
		operators := make([]networktypes.ClientNode, 0)
		for i := int64(0); i < committeeData.CurrentCommittee.Height.Int64(); i++ {
			addr, err := g.committeeSC.CommitteeAddrs(nil, g.chainID, big.NewInt(i))
			if err != nil {
				return nil, err
			}
			operator, err := g.committeeSC.Operators(nil, addr)
			if err != nil {
				return nil, err
			}
			operators = append(operators, networktypes.ClientNode{
				StakeAddress: addr.String(),
				VotingPower:  uint64(operator.Amount.Int64()),
				PublicKey:    common.Bytes2Hex(operator.BlsPubKey),
			})
		}

		g.operators = operators
		g.isOpertorsSynced = true
	}

	committeeRoot.Operators = g.operators

	return committeeRoot, nil
}

package governance

import (
	"context"
	"fmt"
	"math/big"
	"time"

	contypes "github.com/Lagrange-Labs/lagrange-node/consensus/types"
	"github.com/Lagrange-Labs/lagrange-node/governance/types"
	"github.com/Lagrange-Labs/lagrange-node/logger"
	networktypes "github.com/Lagrange-Labs/lagrange-node/network/types"
	"github.com/Lagrange-Labs/lagrange-node/scinterface/committee"
	"github.com/Lagrange-Labs/lagrange-node/scinterface/lagrange"
	"github.com/Lagrange-Labs/lagrange-node/utils"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Governance is the module which is responsible for the staking and slashing.
type Governance struct {
	stakingInterval    time.Duration
	evidenceInterval   time.Duration
	lagrangeSC         *lagrange.Lagrange
	committeeSC        *committee.Committee
	chainID            uint32
	updatedEpochNumber uint64
	currentEpochNumber uint64

	storage     storageInterface
	etherClient *ethclient.Client
	auth        *bind.TransactOpts

	ctx    context.Context
	cancel context.CancelFunc
}

// NewGovernance creates a new Governance instance.
func NewGovernance(cfg *types.Config, chainID uint32, storage storageInterface) (*Governance, error) {
	client, err := ethclient.Dial(cfg.EthereumURL)
	if err != nil {
		return nil, err
	}
	lagrangeSC, err := lagrange.NewLagrange(common.HexToAddress(cfg.StakingSCAddress), client)
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

	ctx, cancel := context.WithCancel(context.Background())
	return &Governance{
		stakingInterval:    time.Duration(cfg.StakingCheckInterval),
		evidenceInterval:   time.Duration(cfg.EvidenceUploadInterval),
		lagrangeSC:         lagrangeSC,
		committeeSC:        committeeSC,
		chainID:            chainID,
		storage:            storage,
		etherClient:        client,
		auth:               auth,
		updatedEpochNumber: updatedEpochNumber.Uint64(),
		currentEpochNumber: lastEpochNumber,
		ctx:                ctx,
		cancel:             cancel,
	}, nil
}

// Start starts the governance process.
func (g *Governance) Start() {
	if err := g.updateCommittee(); err != nil {
		logger.Fatalf("failed to update committee root: %w", err)
	}

	ticker := time.NewTicker(g.evidenceInterval)
	defer ticker.Stop()

	for {
		select {
		case <-g.ctx.Done():
			return
		case <-ticker.C:
			go func() {
				if err := g.uploadEvidences(); err != nil {
					logger.Fatalf("failed to upload evidences: %w", err)
				}
			}()
		case <-time.After(g.stakingInterval):
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

func (g *Governance) uploadEvidences() error {
	evidences, err := g.storage.GetEvidences(g.ctx)
	if err != nil {
		return err
	}
	for _, evidence := range evidences {
		_, err := g.lagrangeSC.UploadEvidence(g.auth, contypes.GetLagrangeServiceEvidence(evidence))
		if err != nil {
			return err
		}
		evidence.Status = true
		if err := g.storage.UpdateEvidence(g.ctx, evidence); err != nil {
			return err
		}
	}
	return nil
}

func (g *Governance) updateNodeStatuses() error {
	err := g.updateCommittee()
	if err != nil {
		return fmt.Errorf("failed to update committee root: %w", err)
	}

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
	blockNumber, err := g.etherClient.BlockNumber(g.ctx)
	if err != nil {
		return err
	}

	epochNumber, err := g.committeeSC.GetEpochNumber(nil, g.chainID, big.NewInt(int64(blockNumber)))
	if err != nil {
		return err
	}
	if epochNumber.Uint64() > g.updatedEpochNumber {
		updatable, err := g.committeeSC.IsUpdatable(nil, g.chainID, epochNumber)
		if err != nil {
			return err
		}
		isSkipped := (epochNumber.Uint64() > g.updatedEpochNumber+1)
		if updatable || isSkipped {
			// rotate the committee tree
			updateEpochNumber := big.NewInt(int64(epochNumber.Int64()))
			if isSkipped {
				updateEpochNumber = updateEpochNumber.Sub(epochNumber, big.NewInt(1))
			}
			logger.Infof("updating committee tree for epoch %d", updateEpochNumber.Int64())
			tx, err := g.committeeSC.Update(g.auth, g.chainID, updateEpochNumber)
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
			g.updatedEpochNumber = updateEpochNumber.Uint64()
		}
	}

	if epochNumber.Uint64() > g.currentEpochNumber {
		// fetch the committee root from the smart contract
		committeeData, err := g.committeeSC.GetCommittee(nil, g.chainID, big.NewInt(int64(blockNumber)))

		if err != nil {
			return err
		}

		committeeRoot := &types.CommitteeRoot{
			ChainID:              g.chainID,
			CurrentCommitteeRoot: common.Bytes2Hex(committeeData.CurrentCommittee.Root.Bytes()),
			TotalVotingPower:     committeeData.CurrentCommittee.TotalVotingPower.Uint64(),
			EpochBlockNumber:     blockNumber,
			EpochNumber:          epochNumber.Uint64(),
		}

		if committeeRoot.TotalVotingPower == 0 {
			logger.Errorf("total voting power is 0, committee root %v, epoch number %d", committeeRoot, epochNumber.Int64())
			return fmt.Errorf("total voting power is 0")
		}
		operators := make([]networktypes.ClientNode, 0)
		for i := int64(0); i < committeeData.CurrentCommittee.Height.Int64(); i++ {
			addr, err := g.committeeSC.CommitteeAddrs(nil, g.chainID, big.NewInt(i))
			if err != nil {
				return err
			}
			operator, err := g.committeeSC.Operators(nil, addr)
			if err != nil {
				return err
			}
			operators = append(operators, networktypes.ClientNode{
				StakeAddress: addr.String(),
				VotingPower:  uint64(operator.Amount.Int64()),
			})
		}

		committeeRoot.Operators = operators
		if err := g.storage.UpdateCommitteeRoot(g.ctx, committeeRoot); err != nil {
			return err
		}
		g.currentEpochNumber = epochNumber.Uint64()
	}

	return nil
}

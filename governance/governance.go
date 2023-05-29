package governance

import (
	"context"
	"fmt"
	"time"

	contypes "github.com/Lagrange-Labs/lagrange-node/consensus/types"
	"github.com/Lagrange-Labs/lagrange-node/logger"
	networktypes "github.com/Lagrange-Labs/lagrange-node/network/types"
	"github.com/Lagrange-Labs/lagrange-node/scinterface/lagrange"
	"github.com/Lagrange-Labs/lagrange-node/utils"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// Governance is the module which is responsible for the staking and slashing.
type Governance struct {
	stakingInterval  time.Duration
	evidenceInterval time.Duration
	lagrangeSC       *lagrange.Lagrange
	storage          storageInterface
	auth             *bind.TransactOpts

	ctx    context.Context
	cancel context.CancelFunc
}

// NewGovernance creates a new Governance instance.
func NewGovernance(cfg *Config, storage storageInterface) (*Governance, error) {
	client, err := ethclient.Dial(cfg.EthereumURL)
	if err != nil {
		return nil, err
	}
	lagrangeSC, err := lagrange.NewLagrange(common.HexToAddress(cfg.StakingSCAddress), client)
	if err != nil {
		return nil, err
	}

	auth, err := utils.GetSigner(context.Background(), client, cfg.PrivateKey)
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithCancel(context.Background())
	return &Governance{
		stakingInterval:  time.Duration(cfg.StakingCheckInterval),
		evidenceInterval: time.Duration(cfg.EvidenceUploadInterval),
		lagrangeSC:       lagrangeSC,
		storage:          storage,
		auth:             auth,
		ctx:              ctx,
		cancel:           cancel,
	}, nil
}

// Start starts the governance process.
func (g *Governance) Start() {
	ticker := time.NewTicker(g.stakingInterval)
	for {
		select {
		case <-g.ctx.Done():
			return
		case <-time.After(g.evidenceInterval):
			go func() {
				if err := g.uploadEvidences(); err != nil {
					panic(fmt.Errorf("failed to upload evidences: %w", err))
				}
			}()
		case <-ticker.C:
			if err := g.updateNodeStatuses(); err != nil {
				panic(fmt.Errorf("failed to update node status: %w", err))
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
	nodes, err := g.storage.GetNodesByStatuses(g.ctx, []networktypes.NodeStatus{networktypes.NodeJoined})
	logger.Infof("updating nodes %v", nodes)
	if err != nil {
		return err
	}
	for _, node := range nodes {
		sNode, err := g.lagrangeSC.Operators(nil, common.HexToAddress(node.StakeAddress))
		if err != nil {
			return err
		}
		logger.Infof("node found %v", sNode)
		node.VotingPower = uint64(sNode.Amount.Int64())
		if node.VotingPower == 0 {
			logger.Errorf("node %s has 0 voting power", node.StakeAddress)
			continue
		}

		if sNode.Slashed {
			node.Status = networktypes.NodeSlashed
			if err := g.storage.UpdateNode(g.ctx, &node); err != nil {
				return err
			}
		} else {
			node.Status = networktypes.NodeRegistered
			if err := g.storage.UpdateNode(g.ctx, &node); err != nil {
				return err
			}
		}
	}

	return nil
}

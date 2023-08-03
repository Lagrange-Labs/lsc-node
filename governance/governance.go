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
	updatedEpochNumber int64

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

	updatedEpochNumber, err := committeeSC.UpdatedEpoch(nil, big.NewInt(int64(chainID)))
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
		updatedEpochNumber: updatedEpochNumber.Int64(),
		ctx:                ctx,
		cancel:             cancel,
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
	nodes, err := g.storage.GetNodesByStatuses(g.ctx, []networktypes.NodeStatus{networktypes.NodeJoined}, g.chainID)
	logger.Infof("updating nodes %v", nodes)
	if err != nil {
		return err
	}
	// TODO - we need to update node status in case of new adding and removing in real time
	for _, node := range nodes {
		sNode, err := g.committeeSC.Operators(nil, common.HexToAddress(node.StakeAddress))
		if err != nil {
			logger.Warnf("failed to get node %s: %v", node.StakeAddress, err)
			continue
		}
		logger.Infof("node found %v", sNode)
		node.VotingPower = uint64(sNode.Amount.Int64())
		if node.VotingPower == 0 {
			logger.Errorf("node %s has 0 voting power", node.StakeAddress)
			continue
		}

		if node.PublicKey != common.Bytes2Hex(sNode.BlsPubKey) {
			logger.Errorf("node %s has different public key", node.StakeAddress)
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

	return g.updateCommittee()
}

func (g *Governance) updateCommittee() error {
	blockNumber, err := g.etherClient.BlockNumber(g.ctx)
	if err != nil {
		return err
	}

	epochNumber, err := g.committeeSC.GetEpochNumber(nil, big.NewInt(int64(g.chainID)), big.NewInt(int64(blockNumber)))
	if err != nil {
		return err
	}
	if epochNumber.Int64() > g.updatedEpochNumber {
		updatable, err := g.committeeSC.IsUpdatable(nil, epochNumber, big.NewInt(int64(g.chainID)))
		if err != nil {
			return err
		}
		if !updatable && epochNumber.Int64() > g.updatedEpochNumber+1 {
			epochNumber = big.NewInt(epochNumber.Int64() - 1)
			updatable = true
		}
		if updatable {
			// rotate the committee tree
			logger.Infof("updating committee tree for epoch %d", epochNumber.Int64())
			tx, err := g.committeeSC.Update(g.auth, big.NewInt(int64(g.chainID)), epochNumber)
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
			g.updatedEpochNumber = epochNumber.Int64()
		}
	}

	committeeData, err := g.committeeSC.GetCommittee(nil, big.NewInt(int64(g.chainID)), big.NewInt(int64(blockNumber)))
	if err != nil {
		return err
	}

	committeeRoot := &types.CommitteeRoot{
		ChainID:              g.chainID,
		CurrentCommitteeRoot: common.Bytes2Hex(committeeData.CurrentCommittee.Root.Bytes()),
		NextCommitteeRoot:    common.Bytes2Hex(committeeData.NextRoot.Bytes()),
		TotalVotingPower:     committeeData.CurrentCommittee.TotalVotingPower.Uint64(),
		EpochBlockNumber:     blockNumber,
	}

	if committeeRoot.TotalVotingPower == 0 {
		logger.Errorf("total voting power is 0, committee root %v, epoch number %d", committeeRoot, epochNumber.Int64())
		return fmt.Errorf("total voting power is 0")
	}

	return g.storage.UpdateCommitteeRoot(g.ctx, committeeRoot)
}

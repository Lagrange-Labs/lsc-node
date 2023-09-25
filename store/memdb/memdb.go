package memdb

import (
	"context"
	"time"

	contypes "github.com/Lagrange-Labs/lagrange-node/consensus/types"
	govtypes "github.com/Lagrange-Labs/lagrange-node/governance/types"
	"github.com/Lagrange-Labs/lagrange-node/logger"
	networktypes "github.com/Lagrange-Labs/lagrange-node/network/types"
	sequencertypes "github.com/Lagrange-Labs/lagrange-node/sequencer/types"
	"github.com/Lagrange-Labs/lagrange-node/store/types"
	"github.com/Lagrange-Labs/lagrange-node/utils"
)

const KeyLen = 32

var _ types.Storage = (*MemDB)(nil)

// DB is an in-memory database.
type MemDB struct {
	nodes          map[string]networktypes.ClientNode
	blocks         []*sequencertypes.Block
	evidences      []*contypes.Evidence
	committeeRoots []*govtypes.CommitteeRoot
}

// NewMemDB creates a new in-memory database.
func NewMemDB() (*MemDB, error) {
	nodes := make(map[string]networktypes.ClientNode, 0)
	db := &MemDB{nodes: nodes, blocks: []*sequencertypes.Block{}}
	go db.updateBlock(10 * time.Second)
	return db, nil
}

// AddNode adds a client node to the network.
func (d *MemDB) AddNode(ctx context.Context, node *networktypes.ClientNode) error {
	node.Status = networktypes.NodeRegistered
	node.VotingPower = 1
	d.nodes[node.PublicKey] = *node
	return nil
}

// GetNodeByStakeAddr returns the node with the given IP address.
func (d *MemDB) GetNodeByStakeAddr(ctx context.Context, stakeAddress string, chainID uint32) (*networktypes.ClientNode, error) {
	for _, node := range d.nodes {
		if node.StakeAddress == stakeAddress && node.ChainID == chainID {
			return &node, nil
		}
	}
	return nil, nil
}

// GetLastBlock returns the last block that was submitted to the network.
func (d *MemDB) GetLastFinalizedBlock(ctx context.Context, chainID uint32) (*sequencertypes.Block, error) {
	return d.blocks[len(d.blocks)-1], nil
}

// GetBlock returns the block for the given block number.
func (d *MemDB) GetBlock(ctx context.Context, chainID uint32, blockNumber uint64) (*sequencertypes.Block, error) {
	if blockNumber > uint64(len(d.blocks)) {
		return nil, types.ErrBlockNotFound
	}
	return d.blocks[blockNumber-1], nil
}

// GetBlocks returns the `count` blocks starting from `fromBlockNumber`.
func (d *MemDB) GetBlocks(ctx context.Context, chainID uint32, fromBlockNumber uint64, count uint32) ([]*sequencertypes.Block, error) {
	if fromBlockNumber > uint64(len(d.blocks)) {
		return nil, types.ErrBlockNotFound
	}
	if fromBlockNumber+uint64(count) > uint64(len(d.blocks)) {
		return d.blocks[fromBlockNumber-1:], nil
	}
	return d.blocks[fromBlockNumber-1 : fromBlockNumber+uint64(count)-1], nil
}

// AddBlock adds a new block to the database.
func (d *MemDB) AddBlock(ctx context.Context, block *sequencertypes.Block) error {
	blockNumber := uint64(len(d.blocks)) + 1
	lastBlock := &sequencertypes.Block{
		ChainHeader: &sequencertypes.ChainHeader{
			BlockNumber: blockNumber,
			BlockHash:   utils.RandomHex(32),
			ChainId:     1,
		},
		BlockHeader: &sequencertypes.BlockHeader{
			CurrentCommittee: utils.RandomHex(32),
			NextCommittee:    utils.RandomHex(32),
		},
	}

	d.blocks = append(d.blocks, lastBlock)
	return nil
}

// UpdateBlock updates the block in the database.
func (d *MemDB) UpdateBlock(ctx context.Context, block *sequencertypes.Block) error {
	for i := 0; i < len(d.blocks); i++ {
		if d.blocks[i].BlockNumber() == block.BlockNumber() {
			d.blocks[i] = block
		}
	}

	return nil
}

// GetLastFinalizedBlockNumber returns the last finalized block number.
func (d *MemDB) GetLastFinalizedBlockNumber(ctx context.Context, chainID uint32) (uint64, bool, error) {
	for i := len(d.blocks) - 1; i >= 0; i-- {
		if len(d.blocks[i].AggSignature) != 0 {
			return d.blocks[i].BlockNumber(), true, nil
		}
	}
	if len(d.blocks) > 0 {
		return d.blocks[0].BlockNumber(), false, nil
	}
	return 0, false, nil
}

// UpdateNode updates the node status in the database.
func (d *MemDB) UpdateNode(ctx context.Context, node *networktypes.ClientNode) error {
	d.nodes[node.PublicKey] = *node
	return nil
}

// GetNodesByStatuses returns the nodes with the given statuses.
func (d *MemDB) GetNodesByStatuses(ctx context.Context, statuses []networktypes.NodeStatus, chainID uint32) ([]networktypes.ClientNode, error) {
	res := make([]networktypes.ClientNode, 0)
	for _, node := range d.nodes {
		isBelonged := false
		for _, status := range statuses {
			if node.Status == status {
				isBelonged = true
				break
			}
		}
		if isBelonged {
			res = append(res, node)
		}
	}

	return res, nil
}

// GetLastBlockNumber returns the last block number.
func (d *MemDB) GetLastBlockNumber(ctx context.Context, chainID uint32) (uint64, error) {
	return uint64(len(d.blocks)), nil
}

func (d *MemDB) updateBlock(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		<-ticker.C
		if err := d.AddBlock(context.Background(), nil); err != nil {
			panic(err)
		}
		logger.Infof("new block added: %d", len(d.blocks))
	}
}

// AddEvidences adds the given evidences to the database.
func (d *MemDB) AddEvidences(ctx context.Context, evidences []*contypes.Evidence) error {
	d.evidences = append(d.evidences, evidences...)
	return nil
}

// UpdateEvidence updates the given evidence in the database.
func (d *MemDB) UpdateEvidence(ctx context.Context, evidence *contypes.Evidence) error {
	for i := 0; i < len(d.evidences); i++ {
		if d.evidences[i].BlockHash == evidence.BlockHash && d.evidences[i].Operator == evidence.Operator {
			d.evidences[i].Status = true
		}
	}
	return nil
}

// GetEvidences returns the pending evidences.
func (d *MemDB) GetEvidences(ctx context.Context) ([]*contypes.Evidence, error) {
	evidences := make([]*contypes.Evidence, 0)
	for _, evidence := range d.evidences {
		if !evidence.Status {
			evidences = append(evidences, evidence)
		}
	}
	return evidences, nil
}

// UpdateCommitteeRoot updates the committee root in the database.
func (d *MemDB) UpdateCommitteeRoot(ctx context.Context, committeeRoot *govtypes.CommitteeRoot) error {
	d.committeeRoots = append(d.committeeRoots, committeeRoot)
	return nil
}

// GetLastCommitteeRoot returns the last committee root for the given chainID.
func (d *MemDB) GetLastCommitteeRoot(ctx context.Context, chainID uint32, isFinalized bool) (*govtypes.CommitteeRoot, error) {
	for i := len(d.committeeRoots) - 1; i >= 0; i-- {
		if d.committeeRoots[i].ChainID == chainID && d.committeeRoots[i].IsFinalized == isFinalized {
			return d.committeeRoots[i], nil
		}
	}
	return nil, nil
}

// GetLastCommitteeEpochNumber returns the last committee epoch number for the given chainID.
func (d *MemDB) GetLastCommitteeEpochNumber(ctx context.Context, chainID uint32) (uint64, error) {
	for i := len(d.committeeRoots) - 1; i >= 0; i-- {
		if d.committeeRoots[i].ChainID == chainID {
			return d.committeeRoots[i].EpochNumber, nil
		}
	}
	return 0, nil
}

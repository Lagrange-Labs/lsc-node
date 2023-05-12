package store

import (
	"context"
	"time"

	"github.com/Lagrange-Labs/lagrange-node/logger"
	sequencertypes "github.com/Lagrange-Labs/lagrange-node/sequencer/types"
	"github.com/Lagrange-Labs/lagrange-node/store/types"
	"github.com/Lagrange-Labs/lagrange-node/utils"
)

const KeyLen = 32

var _ Storage = (*MemDB)(nil)

// DB is an in-memory database.
type MemDB struct {
	nodes  map[string]sequencertypes.ClientNode
	blocks []*sequencertypes.Block
}

// NewMemDB creates a new in-memory database.
func NewMemDB() (*MemDB, error) {
	nodes := make(map[string]sequencertypes.ClientNode, 0)
	db := &MemDB{nodes: nodes, blocks: []*sequencertypes.Block{}}
	go db.updateBlock(10 * time.Second)
	return db, nil
}

// AddNode adds a client node to the network.
func (d *MemDB) AddNode(ctx context.Context, node *sequencertypes.ClientNode) error {
	node.Status = sequencertypes.NodeRegistered
	node.VotingPower = 1
	d.nodes[node.PublicKey] = *node
	return nil
}

// GetNode returns the node with the given IP address.
func (d *MemDB) GetNode(ctx context.Context, ip string) (*sequencertypes.ClientNode, error) {
	for _, node := range d.nodes {
		if node.IPAddress == ip {
			return &node, nil
		}
	}
	return nil, nil
}

// GetNodeCount returns the number of nodes in the network.
func (d *MemDB) GetNodeCount(ctx context.Context) (uint16, error) {
	return uint16(len(d.nodes)), nil
}

// GetLastBlock returns the last block that was submitted to the network.
func (d *MemDB) GetLastBlock(ctx context.Context) (*sequencertypes.Block, error) {
	return d.blocks[len(d.blocks)-1], nil
}

// GetBlock returns the block for the given block number.
func (d *MemDB) GetBlock(ctx context.Context, blockNumber uint64) (*sequencertypes.Block, error) {
	if blockNumber > uint64(len(d.blocks)) {
		return nil, types.ErrBlockNotFound
	}
	return d.blocks[blockNumber-1], nil
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
func (d *MemDB) GetLastFinalizedBlockNumber(ctx context.Context) (uint64, error) {
	for i := len(d.blocks) - 1; i >= 0; i-- {
		if len(d.blocks[i].AggSignature) != 0 {
			return d.blocks[i].BlockNumber(), nil
		}
	}

	return 0, nil
}

// UpdateNode updates the node status in the database.
func (d *MemDB) UpdateNode(ctx context.Context, node *sequencertypes.ClientNode) error {
	d.nodes[node.PublicKey] = *node
	return nil
}

// GetNodesByStatuses returns the nodes with the given statuses.
func (d *MemDB) GetNodesByStatuses(ctx context.Context, statuses []sequencertypes.NodeStatus) ([]sequencertypes.ClientNode, error) {
	res := make([]sequencertypes.ClientNode, 0)
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
func (d *MemDB) GetLastBlockNumber(ctx context.Context) (uint64, error) {
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

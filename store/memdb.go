package store

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/Lagrange-Labs/lagrange-node/logger"
	sequencertypes "github.com/Lagrange-Labs/lagrange-node/sequencer/types"
	synctypes "github.com/Lagrange-Labs/lagrange-node/synchronizer/types"
	"github.com/Lagrange-Labs/lagrange-node/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/umbracle/go-eth-consensus/bls"
	"google.golang.org/protobuf/proto"
)

const KeyLen = 32

var _ Storage = (*MemDB)(nil)

// DB is an in-memory database.
type MemDB struct {
	nodes   map[string]sequencertypes.ClientNode
	blocks  []*sequencertypes.Block
	privKey *bls.SecretKey
	pubKey  string
}

// NewMemDB creates a new in-memory database.
func NewMemDB() (*MemDB, error) {
	nodes := make(map[string]sequencertypes.ClientNode, 0)
	privKey, pubKey := utils.RandomBlsKey()
	db := &MemDB{nodes: nodes, blocks: []*sequencertypes.Block{}, privKey: privKey, pubKey: pubKey}
	go db.updateBlock(10 * time.Second)
	return db, nil
}

// AddNode adds a client node to the network.
func (d *MemDB) AddNode(ctx context.Context, node *sequencertypes.ClientNode) error {
	d.nodes[node.IPAddress] = *node
	return nil
}

// GetNode returns the node with the given IP address.
func (d *MemDB) GetNode(ctx context.Context, ip string) (*sequencertypes.ClientNode, error) {
	node, ok := d.nodes[ip]
	if !ok {
		return nil, nil
	}
	return &node, nil
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
		return nil, fmt.Errorf("the block %d is not ready", blockNumber)
	}
	return d.blocks[blockNumber-1], nil
}

// AddBlock adds a new block to the database.
func (d *MemDB) AddBlock(ctx context.Context, block *sequencertypes.Block) error {
	blockNumber := uint64(len(d.blocks)) + 1
	parentHash := common.Hash{}.Hex()
	if blockNumber > 1 {
		parentHash = d.blocks[len(d.blocks)-1].Header.BlockHash
	}
	lastBlock := &sequencertypes.Block{
		ChainHeader: &synctypes.ChainHeader{
			BlockNumber: blockNumber,
			StateRoot:   randomHex(32),
			Chain:       "test",
		},
		Header: &sequencertypes.BlockHeader{
			BlockNumber:    blockNumber,
			ParentHash:     parentHash,
			ProposerPubKey: d.pubKey,
		},
	}
	blockMsg, err := proto.Marshal(lastBlock)
	if err != nil {
		panic(err)
	}
	blockHash := utils.Hash(blockMsg)
	lastBlock.Header.BlockHash = common.Bytes2Hex(blockHash)

	sig, err := d.privKey.Sign(blockHash)
	if err != nil {
		panic(err)
	}
	sigMsg := sig.Serialize()
	lastBlock.Header.ProposerSignature = common.Bytes2Hex(sigMsg[:])

	d.blocks = append(d.blocks, lastBlock)
	return nil
}

// UpdateBlock updates the block in the database.
func (d *MemDB) UpdateBlock(ctx context.Context, block *sequencertypes.Block) error {
	for i := 0; i < len(d.blocks); i++ {
		if d.blocks[i].Header.BlockNumber == block.Header.BlockNumber {
			d.blocks[i] = block
		}
	}

	return nil
}

// GetLastFinalizedBlockNumber returns the last finalized block number.
func (d *MemDB) GetLastFinalizedBlockNumber(ctx context.Context) (uint64, error) {
	for i := len(d.blocks) - 1; i >= 0; i-- {
		if len(d.blocks[i].AggSignature) != 0 {
			return d.blocks[i].Header.BlockNumber, nil
		}
	}

	return 0, nil
}

// UpdateNode updates the node status in the database.
func (d *MemDB) UpdateNode(ctx context.Context, node *sequencertypes.ClientNode) error {
	d.nodes[node.IPAddress] = *node
	return nil
}

// GetNodesByStatuses returns the nodes with the given statuses.
func (d *MemDB) GetNodesByStatuses(ctx context.Context, statuses []sequencertypes.NodeStatus) ([]*sequencertypes.ClientNode, error) {
	res := make([]*sequencertypes.ClientNode, 0)
	for _, node := range d.nodes {
		isBelonged := false
		for _, status := range statuses {
			if node.Status == status {
				isBelonged = true
				break
			}
		}
		if isBelonged {
			res = append(res, &node)
		}
	}

	return res, nil
}

// GetLastBlockNumber returns the last block number.
func (d *MemDB) GetLastBlockNumber(ctx context.Context) (uint64, error) {
	return uint64(len(d.blocks)), nil
}

func randomHex(n int) string {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		panic(err)
	}
	return hex.EncodeToString(bytes)
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

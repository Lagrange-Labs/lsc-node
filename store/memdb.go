package store

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"time"

	sequencertypes "github.com/Lagrange-Labs/Lagrange-Node/sequencer/types"
	synchronizertypes "github.com/Lagrange-Labs/Lagrange-Node/synchronizer/types"
	"github.com/Lagrange-Labs/Lagrange-Node/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/umbracle/go-eth-consensus/bls"
	"google.golang.org/protobuf/proto"
)

const KeyLen = 32

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
	priv := new(bls.SecretKey)
	if err := priv.Unmarshal(common.FromHex("0x0642cf177a12c962938366d7c2d286f49806625831aaed8e861405bfdd1f652c")); err != nil {
		panic(err)
	}
	pubKeyMsg := priv.GetPublicKey().Serialize()
	db := &MemDB{nodes: nodes, blocks: []*sequencertypes.Block{}, privKey: priv, pubKey: common.Bytes2Hex(pubKeyMsg[:])}
	go db.updateBlock(5 * time.Second)
	return db, nil
}

// AddNode adds a client node to the network.
func (d *MemDB) AddNode(ctx context.Context, stakeAdr, pubKey, ipAdr string) error {
	d.nodes[ipAdr] = sequencertypes.ClientNode{
		StakeAddress: stakeAdr,
		PublicKey:    pubKey,
		IPAddress:    ipAdr,
	}
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
	if blockNumber >= uint64(len(d.blocks)) {
		return nil, fmt.Errorf("the block %d is not ready", blockNumber)
	}
	return d.blocks[blockNumber], nil
}

// AddBlock adds a new block to the database.
func (d *MemDB) AddBlock(ctx context.Context, block *sequencertypes.Block) error {
	return nil
}

// UpdateNode updates the node status in the database.
func (d *MemDB) UpdateNode(ctx context.Context, node *sequencertypes.ClientNode) error {
	return nil
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
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	for {
		<-ticker.C
		blockNumber := uint64(len(d.blocks))
		parentHash := common.Hash{}.Hex()
		if blockNumber > 0 {
			parentHash = d.blocks[len(d.blocks)-1].Header.BlockHash
		}
		lastBlock := &sequencertypes.Block{
			Delta: &synchronizertypes.BlockDelta{
				BlockNumber: blockNumber,
				StateRoot:   randomHex(32),
				Chain:       "ethereum",
				Delta: []*synchronizertypes.DeltaItem{
					{
						Address: randomHex(32),
						Key:     "balance",
						Value:   &synchronizertypes.DeltaItem_StringValue{StringValue: "100000"},
					},
					{
						Address: randomHex(32),
						Key:     "storage",
						Value: &synchronizertypes.DeltaItem_StorageValue{
							StorageValue: &synchronizertypes.StorageItemList{
								Items: []*synchronizertypes.StorageItem{
									{
										Skey:   randomHex(32),
										Svalue: randomHex(32),
									},
									{
										Skey:   randomHex(32),
										Svalue: randomHex(32),
									},
									{
										Skey:   randomHex(32),
										Svalue: randomHex(32),
									},
								},
							},
						},
					},
				},
			},
			Proof: randomHex(32),
			Header: &sequencertypes.BlockHeader{
				BlockNumber:    blockNumber,
				ParentHash:     parentHash,
				ProposerPubKey: d.pubKey,
			},
		}
		deltaMsg, err := proto.Marshal(lastBlock.Delta)
		if err != nil {
			panic(err)
		}
		deltaHash := utils.Hash(deltaMsg)
		lastBlock.Delta.DeltaHash = common.Bytes2Hex(deltaHash)
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
	}
}

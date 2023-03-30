package store

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/Lagrange-Labs/Lagrange-Node/network"
	"github.com/Lagrange-Labs/Lagrange-Node/network/types"
	"github.com/Lagrange-Labs/Lagrange-Node/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/umbracle/go-eth-consensus/bls"
	"google.golang.org/protobuf/proto"
)

const KeyLen = 32

// DB is an in-memory database.
type MemDB struct {
	nodes   map[string]network.ClientNode
	blocks  []*types.Block
	privKey *bls.SecretKey
	pubKey  string
}

// NewMemDB creates a new in-memory database.
func NewMemDB() (*MemDB, error) {
	nodes := make(map[string]network.ClientNode, 0)
	priv := new(bls.SecretKey)
	if err := priv.Unmarshal(common.FromHex("0x0642cf177a12c962938366d7c2d286f49806625831aaed8e861405bfdd1f652c")); err != nil {
		panic(err)
	}
	pubKeyMsg := priv.GetPublicKey().Serialize()
	db := &MemDB{nodes: nodes, blocks: []*types.Block{}, privKey: priv, pubKey: common.Bytes2Hex(pubKeyMsg[:])}
	go db.updateBlock(5 * time.Second)
	return db, nil
}

// AddNode adds a client node to the network.
func (d *MemDB) AddNode(ctx context.Context, stakeAdr, pubKey, ipAdr string) error {
	pk := new(bls.PublicKey)
	err := pk.Deserialize(common.FromHex(pubKey))
	if err != nil {
		return err
	}
	d.nodes[ipAdr] = network.ClientNode{
		StakeAddress: stakeAdr,
		PublicKey:    pk,
		IPAddress:    ipAdr,
	}
	return nil
}

// GetNode returns the node with the given IP address.
func (d *MemDB) GetNode(ctx context.Context, ip string) (*network.ClientNode, error) {
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
func (d *MemDB) GetLastBlock(ctx context.Context) (*types.Block, error) {
	return d.blocks[len(d.blocks)-1], nil
}

// GetBlock returns the block for the given block number.
func (d *MemDB) GetBlock(ctx context.Context, blockNumber uint64) (*types.Block, error) {
	if blockNumber >= uint64(len(d.blocks)) {
		return nil, fmt.Errorf("the block %d is not ready", blockNumber)
	}
	return d.blocks[blockNumber], nil
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
		lastBlock := &types.Block{
			Delta: &types.BlockDelta{
				BlockNumber: blockNumber,
				StateRoot:   randomHex(32),
				Chain:       "ethereum",
				Delta: []*types.DeltaItem{
					{
						Address: randomHex(32),
						Key:     "balance",
						Value:   &types.DeltaItem_StringValue{StringValue: "100000"},
					},
					{
						Address: randomHex(32),
						Key:     "storage",
						Value: &types.DeltaItem_StorageValue{
							StorageValue: &types.StorageItemList{
								Items: []*types.StorageItem{
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
			Header: &types.BlockHeader{
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

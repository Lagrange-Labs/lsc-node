package store

import (
	"context"
	"crypto/rand"
	"encoding/base64"
	"time"

	"github.com/Lagrange-Labs/Lagrange-Node/network"
	"github.com/Lagrange-Labs/Lagrange-Node/network/pb"
	"github.com/ethereum/go-ethereum/common"
	"github.com/umbracle/go-eth-consensus/bls"
)

// DB is an in-memory database.
type MemDB struct {
	nodes     map[string]network.ClientNode
	lastProof *pb.ProofMessage
}

// NewMemDB creates a new in-memory database.
func NewMemDB() (*MemDB, error) {
	nodes := make(map[string]network.ClientNode, 0)
	db := &MemDB{nodes: nodes}
	go db.updateProof(5 * time.Second)
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

// GetLastProof returns the last proof that was submitted to the network.
func (d *MemDB) GetLastProof(ctx context.Context) (*pb.ProofMessage, error) {
	return d.lastProof, nil
}

func randomString(n int) string {
	b := make([]byte, n)
	_, _ = rand.Read(b)
	return base64.URLEncoding.EncodeToString(b)[:n]
}

func (d *MemDB) updateProof(interval time.Duration) {
	ticker := time.NewTicker(10 * time.Second)
	defer ticker.Stop()

	d.lastProof = &pb.ProofMessage{
		Message: randomString(90),
		Proof:   randomString(32),
		ProofId: 1,
	}

	for {
		<-ticker.C
		d.lastProof = &pb.ProofMessage{
			Message: randomString(90),
			Proof:   randomString(32),
			ProofId: d.lastProof.ProofId + 1,
		}
	}
}

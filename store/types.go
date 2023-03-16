package store

import (
	"context"

	"github.com/Lagrange-Labs/Lagrange-Node/network"
	"github.com/Lagrange-Labs/Lagrange-Node/network/pb"
)

type DB interface {
	// AddNode adds a new node to the database.
	AddNode(ctx context.Context, stakeAdr, pubKey, ipAdr string) error
	// GetNode returns the node with the given IP address.
	GetNode(ctx context.Context, ip string) (*network.ClientNode, error)
	// GetLastProof returns the last proof that was submitted to the network.
	GetLastProof(ctx context.Context) (*pb.ProofMessage, error)
	// GetNodeCount returns the number of nodes in the network.
	GetNodeCount(ctx context.Context) (uint16, error)
}

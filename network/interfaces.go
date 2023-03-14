package network

import (
	"context"

	"github.com/Lagrange-Labs/Lagrange-Node/network/pb"
)

type storageInterface interface {
	RegisterNode(ctx context.Context, stakeAdr, pubKey, ipAdr string) error
	GetNode(ctx context.Context, ip string) (*ClientNode, error)
	GetLastProof(ctx context.Context) (*pb.ProofMessage, error)
	GetNodeCount(ctx context.Context) (uint16, error)
}

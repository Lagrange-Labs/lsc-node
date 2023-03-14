package network

import (
	"context"
	"time"

	"github.com/umbracle/go-eth-consensus/bls"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/proto"

	"github.com/Lagrange-Labs/Lagrange-Node/network/pb"
)

// ClientNode is a struct to store the information of a node.
type ClientNode struct {
	// PublicKey is the public key of the node.
	PublicKey *bls.PublicKey
	// IPAddress is the IP address of the client node.
	IPAddress string
	// StakeAddress is the ethereum address of the staking.
	StakeAddress string
}

// Client is a gRPC client to join the network
type Client struct {
	pb.NetworkServiceClient
	ctx          context.Context
	cancelFunc   context.CancelFunc
	privateKey   *bls.SecretKey
	pullInterval time.Duration
}

// NewClient creates a new client.
func NewClient(cfg ClientConfig) (*Client, error) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	conn, err := grpc.DialContext(ctx, cfg.GrpcURL, opts...)
	if err != nil {
		return nil, err
	}
	priv := new(bls.SecretKey)
	if err := priv.Unmarshal([]byte(cfg.PrivateKey)); err != nil {
		return nil, err
	}

	return &Client{
		NetworkServiceClient: pb.NewNetworkServiceClient(conn),
		privateKey:           priv,
		pullInterval:         cfg.PullInterval,
		ctx:                  ctx,
		cancelFunc:           cancel,
	}, nil
}

func (c *Client) Start() {
	for {
		select {
		case <-c.ctx.Done():
			return
		case <-time.After(c.pullInterval):
			// TODO logging error
			res, err := c.GetLastProof(c.ctx, &pb.GetLastProofRequest{ProofId: 0}) // TODO track the proof id
			if err != nil {
				continue
			}
			msg, err := proto.Marshal(res.Proof)
			if err != nil {
				continue
			}
			sig, err := c.privateKey.Sign(msg)
			if err != nil {
				continue
			}
			sigMsg := sig.Serialize()
			resS, err := c.UploadSignature(c.ctx, &pb.UploadSignatureRequest{
				ProofId:   0,
				Signature: string(sigMsg[:]),
			})
			if err != nil {
				continue
			}
			if !resS.Result {
				continue
			}
		}
	}
}

// Stop function stops the client node.
func (c *Client) Stop() {
	c.cancelFunc()
}

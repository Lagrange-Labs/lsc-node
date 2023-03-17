package network

import (
	"context"
	"fmt"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/umbracle/go-eth-consensus/bls"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/health/grpc_health_v1"
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
	stakeAddress string
	lastProofID  uint64
	pullInterval time.Duration
}

// NewClient creates a new client.
func NewClient(cfg ClientConfig) (*Client, error) {
	ctx, cancel := context.WithCancel(context.Background())

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	conn, err := grpc.DialContext(ctx, cfg.GrpcURL, opts...)
	if err != nil {
		panic(err)
	}

	healthClient := grpc_health_v1.NewHealthClient(conn)
	hctx, hcancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer hcancel()

	response, err := healthClient.Check(hctx, &grpc_health_v1.HealthCheckRequest{})
	if err != nil {
		fmt.Println("Failed to check gRPC health:", err)
		panic(err)
	}

	if response.Status == grpc_health_v1.HealthCheckResponse_SERVING {
		fmt.Println("gRPC server is healthy")
	} else {
		fmt.Println("gRPC server is not healthy")
	}

	priv := new(bls.SecretKey)
	if err := priv.Unmarshal(common.FromHex(cfg.PrivateKey)); err != nil {
		panic(err)
	}

	return &Client{
		NetworkServiceClient: pb.NewNetworkServiceClient(conn),
		privateKey:           priv,
		stakeAddress:         cfg.StakeAddress,
		pullInterval:         time.Duration(cfg.PullInterval),
		ctx:                  ctx,
		cancelFunc:           cancel,
	}, nil
}

// Start starts the connection loop.
func (c *Client) Start() {
	pk := c.privateKey.GetPublicKey().Serialize()
	req := &pb.JoinNetworkRequest{
		PublicKey:    common.Bytes2Hex(pk[:]),
		StakeAddress: c.stakeAddress,
	}
	reqMsg, err := proto.Marshal(req)
	if err != nil {
		panic(err)
	}
	sig, err := c.privateKey.Sign(reqMsg)
	if err != nil {
		panic(err)
	}
	sigMsg := sig.Serialize()
	req.Signature = common.Bytes2Hex(sigMsg[:])
	res, err := c.NetworkServiceClient.JoinNetwork(context.Background(), req)
	if err != nil {
		panic(err)
	}

	if !res.Result {
		panic(fmt.Errorf("failed to join the network: %s", res.Message))
	}

	fmt.Printf("joined the network: %v\n", req)

	for {
		select {
		case <-c.ctx.Done():
			return
		case <-time.After(c.pullInterval):
			// TODO logging error
			res, err := c.GetLastProof(context.Background(), &pb.GetLastProofRequest{ProofId: c.lastProofID}) // TODO track the proof id
			if err != nil {
				fmt.Printf("failed to get the last proof: %v", err)
				continue
			}
			// TODO proof validation
			c.lastProofID = res.Proof.ProofId

			fmt.Printf("got the current proof: %v\n", res.Proof)

			msg, err := proto.Marshal(res.Proof)
			if err != nil {
				fmt.Printf("failed to marshal the proof: %v", err)
				continue
			}
			sig, err := c.privateKey.Sign(msg)
			if err != nil {
				fmt.Printf("failed to sign the proof: %v", err)
				continue
			}
			sigMsg := sig.Serialize()
			resS, err := c.UploadSignature(c.ctx, &pb.UploadSignatureRequest{
				ProofId:   c.lastProofID,
				Signature: common.Bytes2Hex(sigMsg[:]),
			})
			if err != nil {
				fmt.Printf("failed to upload signature: %v", err)
				continue
			}
			if !resS.Result {
				fmt.Printf("failed to upload signature: %s", resS.Message)
				continue
			}

			fmt.Printf("uploaded the signature: %v\n", resS)

		}
	}
}

// Stop function stops the client node.
func (c *Client) Stop() {
	c.cancelFunc()
}

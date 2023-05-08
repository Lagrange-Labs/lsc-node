package network

import (
	"context"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/umbracle/go-eth-consensus/bls"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/protobuf/proto"

	"github.com/Lagrange-Labs/lagrange-node/logger"
	"github.com/Lagrange-Labs/lagrange-node/network/types"
	"github.com/Lagrange-Labs/lagrange-node/utils"
)

// Client is a gRPC client to join the network
type Client struct {
	types.NetworkServiceClient
	ctx             context.Context
	cancelFunc      context.CancelFunc
	privateKey      *bls.SecretKey
	stakeAddress    string
	lastBlockNumber uint64
	pullInterval    time.Duration
}

// NewClient creates a new client.
func NewClient(cfg *ClientConfig) (*Client, error) {
	ctx, cancel := context.WithCancel(context.Background())

	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	conn, err := grpc.DialContext(ctx, cfg.GrpcURL, opts...)
	if err != nil {
		cancel()
		return nil, err
	}

	healthClient := grpc_health_v1.NewHealthClient(conn)
	hctx, hcancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer hcancel()

	watcher, err := healthClient.Watch(hctx, &grpc_health_v1.HealthCheckRequest{})
	if err != nil {
		logger.Error("Failed to check gRPC health:", err)
		cancel()
		return nil, err
	}

	for {
		response, err := watcher.Recv()
		if err != nil {
			logger.Info("Failed to get gRPC health response:", err)
		}
		if response.Status == grpc_health_v1.HealthCheckResponse_SERVING {
			logger.Info("gRPC server is healthy")
			break
		} else {
			logger.Info("gRPC server is not healthy")
		}
	}

	priv, err := utils.HexToBlsPrivKey(cfg.PrivateKey)
	if err != nil {
		panic(err)
	}

	return &Client{
		NetworkServiceClient: types.NewNetworkServiceClient(conn),
		privateKey:           priv,
		stakeAddress:         cfg.StakeAddress,
		pullInterval:         time.Duration(cfg.PullInterval),
		ctx:                  ctx,
		cancelFunc:           cancel,
		lastBlockNumber:      1,
	}, nil
}

// Start starts the connection loop.
func (c *Client) Start() {
	pubkey := utils.BlsPubKeyToHex(c.privateKey.GetPublicKey())
	req := &types.JoinNetworkRequest{
		PublicKey:    pubkey,
		StakeAddress: c.stakeAddress,
	}
	reqMsg, err := proto.Marshal(req)
	if err != nil {
		logger.Fatalf("failed to marshal the request: %v\n", err)
	}

	sig, err := c.privateKey.Sign(reqMsg)
	if err != nil {
		logger.Fatalf("failed to sign the request: %v\n", err)
	}

	req.Signature = utils.BlsSignatureToHex(sig)
	res, err := c.NetworkServiceClient.JoinNetwork(context.Background(), req)
	if err != nil {
		logger.Fatalf("failed to join the network: %v\n", err)
	}

	if !res.Result {
		logger.Panicf("failed to join the network: %s", res.Message)
	}

	logger.Infof("joined the network: %v\n", req)

	for {
		select {
		case <-c.ctx.Done():
			return
		case <-time.After(c.pullInterval):
			res, err := c.GetBlock(context.Background(), &types.GetBlockRequest{BlockNumber: c.lastBlockNumber})
			if err != nil {
				logger.Errorf("failed to get the last block: %v\n", err)
				continue
			}

			logger.Infof("got the current block: %v\n", res.Block)

			// verify the proposer signature
			if len(res.Block.ProposerPubKey()) == 0 {
				logger.Warnf("the block %d is not opened yet", res.Block.BlockNumber())
				continue
			}

			// generate the BLS signature
			blsSignature := res.Block.BlsSignature()
			blsSigMsg, err := proto.Marshal(blsSignature)
			if err != nil {
				logger.Errorf("failed to marshal the BLS signature: %v\n", err)
				continue
			}

			verified, err := utils.VerifySignature(common.FromHex(res.Block.ProposerPubKey()), blsSigMsg, common.FromHex(res.Block.ProposerSignature()))
			if err != nil || !verified {
				logger.Errorf("failed to verify the proposer signature: %v\n", err)
				continue
			}

			blsSig, err := c.privateKey.Sign(blsSigMsg)
			if err != nil {
				logger.Errorf("failed to sign the BLS signature: %v\n", err)
			}
			blsSignature.Signature = utils.BlsSignatureToHex(blsSig)

			req := &types.CommitBlockRequest{
				BlsSignature: blsSignature,
				PubKey:       pubkey,
			}
			msg, err := proto.Marshal(req)
			if err != nil {
				logger.Errorf("failed to marshal the block: %v\n", err)
				continue
			}
			sig, err := c.privateKey.Sign(msg)
			if err != nil {
				logger.Errorf("failed to sign the block: %v\n", err)
				continue
			}
			req.Signature = utils.BlsSignatureToHex(sig)
			resS, err := c.CommitBlock(c.ctx, req)
			if err != nil {
				logger.Errorf("failed to upload signature: %v\n", err)
				continue
			}
			if !resS.Result {
				logger.Infof("failed to upload signature: %s\n", resS.Message)
				continue
			}

			c.lastBlockNumber += 1
			logger.Infof("uploaded the signature: %v\n", resS)
		}
	}
}

// Stop function stops the client node.
func (c *Client) Stop() {
	c.cancelFunc()
}

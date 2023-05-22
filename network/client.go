package network

import (
	"context"
	"crypto/ecdsa"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/umbracle/go-eth-consensus/bls"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/protobuf/proto"

	contypes "github.com/Lagrange-Labs/lagrange-node/consensus/types"
	"github.com/Lagrange-Labs/lagrange-node/logger"
	"github.com/Lagrange-Labs/lagrange-node/network/types"
	"github.com/Lagrange-Labs/lagrange-node/rpcclient"
	"github.com/Lagrange-Labs/lagrange-node/utils"
)

// Client is a gRPC client to join the network
type Client struct {
	types.NetworkServiceClient
	rpcClient       rpcclient.RpcClient
	chainID         int32
	blsPrivateKey   *bls.SecretKey
	ecdsaPrivateKey *ecdsa.PrivateKey
	stakeAddress    string
	lastBlockNumber uint64
	pullInterval    time.Duration

	ctx        context.Context
	cancelFunc context.CancelFunc
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

	blsPriv, err := utils.HexToBlsPrivKey(cfg.BLSPrivateKey)
	if err != nil {
		panic(err)
	}
	ecdsaPriv, err := crypto.HexToECDSA(strings.TrimPrefix(cfg.ECDSAPrivateKey, "0x"))
	if err != nil {
		panic(err)
	}
	stakeAddress := crypto.PubkeyToAddress(ecdsaPriv.PublicKey).Hex()

	rpcClient, err := rpcclient.CreateRPCClient(cfg.Chain, cfg.RPCEndpoint)
	if err != nil {
		panic(err)
	}
	chainID, err := rpcClient.GetChainID()
	if err != nil {
		panic(err)
	}

	return &Client{
		NetworkServiceClient: types.NewNetworkServiceClient(conn),
		blsPrivateKey:        blsPriv,
		ecdsaPrivateKey:      ecdsaPriv,
		stakeAddress:         stakeAddress,
		pullInterval:         time.Duration(cfg.PullInterval),
		rpcClient:            rpcClient,
		chainID:              chainID,
		lastBlockNumber:      1,

		ctx:        ctx,
		cancelFunc: cancel,
	}, nil
}

// Start starts the connection loop.
func (c *Client) Start() {
	pubkey := utils.BlsPubKeyToHex(c.blsPrivateKey.GetPublicKey())
	req := &types.JoinNetworkRequest{
		PublicKey:    pubkey,
		StakeAddress: c.stakeAddress,
	}
	reqMsg, err := proto.Marshal(req)
	if err != nil {
		logger.Fatalf("failed to marshal the request: %v\n", err)
	}

	sig, err := c.blsPrivateKey.Sign(reqMsg)
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
			res, err := c.GetBlock(context.Background(), &types.GetBlockRequest{BlockNumber: c.lastBlockNumber, ChainId: c.chainID})
			if err != nil {
				logger.Errorf("failed to get the last block: %v\n", err)
				continue
			}
			if res.CurrentBlockNumber == 0 {
				logger.Warnf("the current block is not ready\n")
				continue
			}

			logger.Infof("got the current block: %v\n", res.Block)
			if res.CurrentBlockNumber != c.lastBlockNumber {
				// TODO determine how to handle the sync
				logger.Warnf("the current block number %d is not equal to the last block number %d\n", res.CurrentBlockNumber, c.lastBlockNumber)
				c.lastBlockNumber = res.CurrentBlockNumber
				continue
			}

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
			// verify the proposer signature
			verified, err := utils.VerifySignature(common.FromHex(res.Block.ProposerPubKey()), blsSigMsg, common.FromHex(res.Block.ProposerSignature()))
			if err != nil || !verified {
				logger.Errorf("failed to verify the proposer signature: %v\n", err)
				continue
			}
			// verify if the block hash is correct
			blockHash := res.Block.BlockHash()
			rBlockHash, err := c.rpcClient.GetBlockHashByNumber(c.lastBlockNumber)
			if err != nil {
				logger.Errorf("failed to get the block hash by number: %v\n", err)
				continue
			}
			if blockHash != rBlockHash {
				logger.Errorf("the block hash %s is not equal to the rpc block hash %s", blockHash, rBlockHash)
				continue
			}

			blsSig, err := c.blsPrivateKey.Sign(blsSigMsg)
			if err != nil {
				logger.Errorf("failed to sign the BLS signature: %v\n", err)
			}
			blsSignature.Signature = utils.BlsSignatureToHex(blsSig)

			req := &types.CommitBlockRequest{
				BlsSignature: blsSignature,
				EpochNumber:  res.Block.EpochNumber(),
				PubKey:       pubkey,
			}
			// generate the ECDSA signature
			msg := contypes.GetCommitRequestHash(req)
			sig, err := crypto.Sign(msg, c.ecdsaPrivateKey)
			if err != nil {
				logger.Errorf("failed to ecdsa sign the block: %v\n", err)
				continue
			}
			req.Signature = common.Bytes2Hex(sig)
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

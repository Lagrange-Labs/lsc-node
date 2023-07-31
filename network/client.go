package network

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/umbracle/go-eth-consensus/bls"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/protobuf/proto"

	contypes "github.com/Lagrange-Labs/lagrange-node/consensus/types"
	"github.com/Lagrange-Labs/lagrange-node/logger"
	"github.com/Lagrange-Labs/lagrange-node/network/types"
	"github.com/Lagrange-Labs/lagrange-node/rpcclient"
	"github.com/Lagrange-Labs/lagrange-node/scinterface/committee"
	sequencertypes "github.com/Lagrange-Labs/lagrange-node/sequencer/types"
	"github.com/Lagrange-Labs/lagrange-node/utils"
)

// Client is a gRPC client to join the network
type Client struct {
	types.NetworkServiceClient
	rpcClient   rpcclient.RpcClient
	committeeSC *committee.Committee

	chainID         uint32
	blsPrivateKey   *bls.SecretKey
	blsPublicKey    string
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
	pubkey := utils.BlsPubKeyToHex(blsPriv.GetPublicKey())
	ecdsaPriv, err := crypto.HexToECDSA(strings.TrimPrefix(cfg.ECDSAPrivateKey, "0x"))
	if err != nil {
		panic(err)
	}
	stakeAddress := crypto.PubkeyToAddress(ecdsaPriv.PublicKey).Hex()

	rpcClient, err := rpcclient.CreateRPCClient(cfg.Chain, cfg.RPCEndpoint)
	if err != nil {
		panic(err)
	}
	etherClient, err := ethclient.Dial(cfg.EthereumURL)
	if err != nil {
		panic(err)
	}
	committeeSC, err := committee.NewCommittee(common.HexToAddress(cfg.CommitteeSCAddress), etherClient)
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
		blsPublicKey:         pubkey,
		ecdsaPrivateKey:      ecdsaPriv,
		stakeAddress:         stakeAddress,
		pullInterval:         time.Duration(cfg.PullInterval),
		rpcClient:            rpcClient,
		committeeSC:          committeeSC,
		chainID:              chainID,
		lastBlockNumber:      1,

		ctx:        ctx,
		cancelFunc: cancel,
	}, nil
}

// GetStakeAddress returns the stake address.
func (c *Client) GetStakeAddress() string {
	return c.stakeAddress
}

// Start starts the connection loop.
func (c *Client) Start() {
	err := c.TryJoinNetwork()
	if err != nil {
		panic(fmt.Errorf("failed to join the network: %v", err))
	}

	logger.Infof("joined the network: %v\n", c.stakeAddress)

	for {
		select {
		case <-c.ctx.Done():
			return
		default:
			block, err := c.TryGetCurrentBlock()
			if err != nil {
				logger.Errorf("failed to get the current block: %v\n", err)
				continue
			}

			if err := c.TryCommitBlock(block); err != nil {
				logger.Errorf("failed to commit the block: %v\n", err)
				continue
			}

			c.lastBlockNumber += 1
			logger.Info("uploaded the signature")
			time.Sleep(c.pullInterval)
		}
	}
}

// TryJoinNetwork tries to join the network.
func (c *Client) TryJoinNetwork() error {
	req := &types.JoinNetworkRequest{
		PublicKey:    c.blsPublicKey,
		StakeAddress: c.stakeAddress,
	}
	reqMsg, err := proto.Marshal(req)
	if err != nil {
		return err
	}
	sig, err := c.blsPrivateKey.Sign(reqMsg)
	if err != nil {
		return err
	}
	req.Signature = utils.BlsSignatureToHex(sig)
	res, err := c.JoinNetwork(c.ctx, req)
	if err != nil {
		return err
	}
	if !res.Result {
		return fmt.Errorf("failed to join the network: %s", res.Message)
	}
	return nil
}

// TryGetCurrentBlock tries to get the block from the network.
func (c *Client) TryGetCurrentBlock() (*sequencertypes.Block, error) {
	stream, err := c.GetCurrentBlock(c.ctx, &types.GetBlockRequest{BlockNumber: c.lastBlockNumber, StakeAddress: c.stakeAddress})
	if err != nil {
		return nil, fmt.Errorf("failed to get the stream: %v", err)
	}

	// receive the notification from the server via the stream
	res, err := stream.Recv()
	if err != nil {
		return nil, fmt.Errorf("failed to receive the block: %v", err)
	}
	c.lastBlockNumber = res.Block.BlockNumber()

	// verify the proposer signature
	if len(res.Block.ProposerPubKey()) == 0 {
		return nil, fmt.Errorf("the block %d proposer key is empty", res.Block.BlockNumber())
	}
	blsSigHash := res.Block.BlsSignature().Hash()
	if err != nil {
		return nil, fmt.Errorf("failed to marshal the BLS signature: %v", err)
	}
	// verify the proposer signature
	verified, err := utils.VerifySignature(common.FromHex(res.Block.ProposerPubKey()), blsSigHash, common.FromHex(res.Block.ProposerSignature()))
	if err != nil || !verified {
		return nil, fmt.Errorf("failed to verify the proposer signature: %v", err)
	}
	// verify if the block hash is correct
	blockHash := res.Block.BlockHash()
	rBlockHash, err := c.rpcClient.GetBlockHashByNumber(c.lastBlockNumber)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch the block hash by number: %v", err)
	}
	if blockHash != rBlockHash {
		return nil, fmt.Errorf("the block hash %s is not equal to the rpc block hash %s", blockHash, rBlockHash)
	}
	// verify the committee root
	committeeData, err := c.committeeSC.GetCommittee(nil, big.NewInt(int64(c.chainID)), big.NewInt(int64(res.Block.EpochBlockNumber())))
	if err != nil {
		return nil, fmt.Errorf("failed to get the committee root: %v", err)
	}
	if res.Block.CurrentCommittee() != common.Bytes2Hex(committeeData.CurrentCommittee.Root.Bytes()) {
		return nil, fmt.Errorf("the block committee root %s is not equal to the current root %v", res.Block.CurrentCommittee(), committeeData)
	}
	if res.Block.NextCommittee() != common.Bytes2Hex(committeeData.NextRoot.Bytes()) {
		return nil, fmt.Errorf("the block committee root %s is not equal to the next root %v", res.Block.NextCommittee(), committeeData)
	}

	return res.Block, nil
}

// TryCommitBlock tries to commit the signature to the network.
func (c *Client) TryCommitBlock(block *sequencertypes.Block) error {
	blsSignature := block.BlsSignature()
	blsSig, err := c.blsPrivateKey.Sign(blsSignature.Hash())
	if err != nil {
		return fmt.Errorf("failed to sign the BLS signature: %v", err)
	}
	blsSignature.Signature = utils.BlsSignatureToHex(blsSig)

	req := &types.CommitBlockRequest{
		BlsSignature:     blsSignature,
		EpochBlockNumber: block.EpochBlockNumber(),
		PubKey:           c.blsPublicKey,
	}
	// generate the ECDSA signature
	msg := contypes.GetCommitRequestHash(req)
	sig, err := crypto.Sign(msg, c.ecdsaPrivateKey)
	if err != nil {
		return fmt.Errorf("failed to ecdsa sign the block: %v", err)
	}
	req.Signature = common.Bytes2Hex(sig)
	resS, err := c.CommitBlock(c.ctx, req)
	if err != nil {
		return fmt.Errorf("failed to upload signature: %v", err)
	}
	if !resS.Result {
		return fmt.Errorf("failed to upload signature with message : %s", resS.Message)
	}

	return nil
}

// Stop function stops the client node.
func (c *Client) Stop() {
	c.cancelFunc()
}

var (
	ErrBlockNotReady = fmt.Errorf("the block is not ready")
)

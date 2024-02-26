package network

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	ecrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/protobuf/proto"

	contypes "github.com/Lagrange-Labs/lagrange-node/consensus/types"
	"github.com/Lagrange-Labs/lagrange-node/crypto"
	"github.com/Lagrange-Labs/lagrange-node/logger"
	"github.com/Lagrange-Labs/lagrange-node/network/types"
	"github.com/Lagrange-Labs/lagrange-node/rpcclient"
	rpctypes "github.com/Lagrange-Labs/lagrange-node/rpcclient/types"
	"github.com/Lagrange-Labs/lagrange-node/scinterface/committee"
	sequencertypes "github.com/Lagrange-Labs/lagrange-node/sequencer/types"
	"github.com/Lagrange-Labs/lagrange-node/utils"
)

// Client is a gRPC client to join the network
type Client struct {
	types.NetworkServiceClient
	rpcClient   rpctypes.RpcClient
	committeeSC *committee.Committee
	blsScheme   crypto.BLSScheme

	nextCommitteeRoot string
	chainID           uint32
	blsPrivateKey     []byte
	blsPublicKey      string
	ecdsaPrivateKey   *ecdsa.PrivateKey
	stakeAddress      string
	lastBlockNumber   uint64
	pullInterval      time.Duration

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

	blsScheme := crypto.NewBLSScheme(crypto.BLSCurve(cfg.BLSCurve))
	blsPriv := utils.Hex2Bytes(cfg.BLSPrivateKey)
	pubkey, err := blsScheme.GetPublicKey(blsPriv, true)
	if err != nil {
		logger.Fatalf("failed to get the bls public key: %v", err)
	}

	ecdsaPriv, err := ecrypto.HexToECDSA(strings.TrimPrefix(cfg.ECDSAPrivateKey, "0x"))
	if err != nil {
		logger.Fatalf("failed to get the ecdsa private key: %v", err)
	}
	stakeAddress := ecrypto.PubkeyToAddress(ecdsaPriv.PublicKey).Hex()

	rpcClient, err := rpcclient.NewClient(cfg.Chain, cfg.RPCEndpoint, cfg.EthereumURL, "")
	if err != nil {
		logger.Fatalf("failed to create the rpc client: %v", err)
	}
	etherClient, err := ethclient.Dial(cfg.EthereumURL)
	if err != nil {
		logger.Fatalf("failed to create the ethereum client: %v", err)
	}
	committeeSC, err := committee.NewCommittee(common.HexToAddress(cfg.CommitteeSCAddress), etherClient)
	if err != nil {
		logger.Fatalf("failed to create the committee smart contract: %v", err)
	}

	chainID, err := rpcClient.GetChainID()
	if err != nil {
		logger.Fatalf("failed to get the chain ID: %v", err)
	}

	return &Client{
		NetworkServiceClient: types.NewNetworkServiceClient(conn),
		blsScheme:            blsScheme,
		blsPrivateKey:        blsPriv,
		blsPublicKey:         utils.Bytes2Hex(pubkey),
		ecdsaPrivateKey:      ecdsaPriv,
		stakeAddress:         stakeAddress,
		pullInterval:         time.Duration(cfg.PullInterval),
		rpcClient:            rpcClient,
		committeeSC:          committeeSC,
		chainID:              chainID,
		lastBlockNumber:      1,
		nextCommitteeRoot:    "",

		ctx:        ctx,
		cancelFunc: cancel,
	}, nil
}

// GetStakeAddress returns the stake address.
func (c *Client) GetStakeAddress() string {
	return c.stakeAddress
}

// GetChainID returns the chain ID.
func (c *Client) GetChainID() uint32 {
	return c.chainID
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
		case <-time.After(c.pullInterval):
			blocks, err := c.TryGetBlocks()
			if err != nil {
				logger.Errorf("failed to get the current block: %v\n", err)
				continue
			}

			if err := c.TryCommitBlocks(blocks); err != nil {
				logger.Errorf("failed to commit the block: %v\n", err)
				continue
			}

			c.lastBlockNumber = blocks[len(blocks)-1].BlockNumber() + 1
			logger.Infof("uploaded the signature up to block %d\n", c.lastBlockNumber-1)
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
	sig, err := c.blsScheme.Sign(c.blsPrivateKey, reqMsg)
	if err != nil {
		return err
	}
	req.Signature = utils.Bytes2Hex(sig)
	res, err := c.NetworkServiceClient.JoinNetwork(context.Background(), req)
	if err != nil {
		return err
	}
	if !res.Result {
		return fmt.Errorf("failed to join the network: %s", res.Message)
	}
	return nil
}

// TryGetBlocks tries to get the block batch from the network.
func (c *Client) TryGetBlocks() ([]*sequencertypes.Block, error) {
	res, err := c.GetBatch(context.Background(), &types.GetBatchRequest{BlockNumber: c.lastBlockNumber, StakeAddress: c.stakeAddress})
	if err != nil {
		return nil, err
	}

	if len(res.Batch) == 0 {
		return nil, ErrBlockNotReady
	}

	logger.Infof("got the block batch: %d, %d\n", res.Batch[0].BlockNumber(), res.Batch[len(res.Batch)-1].BlockNumber())

	if c.lastBlockNumber < res.Batch[0].BlockNumber() {
		c.nextCommitteeRoot = ""
	}

	// verify the committee root
	if len(c.nextCommitteeRoot) > 0 && res.Batch[0].CurrentCommittee() != c.nextCommitteeRoot {
		return nil, fmt.Errorf("the block committee root %s is not equal to the previous batch's next committee root %s", res.Batch[0].CurrentCommittee(), c.nextCommitteeRoot)
	}
	committeeData, err := c.committeeSC.GetCommittee(nil, c.chainID, big.NewInt(int64(res.Batch[0].EpochBlockNumber())))
	if err != nil {
		return nil, fmt.Errorf("failed to get the committee data: %v chainID: %d Batch: %v", err, c.chainID, res.Batch)
	}
	for i := range res.Batch {
		if res.Batch[i].EpochBlockNumber() != res.Batch[0].EpochBlockNumber() {
			return nil, fmt.Errorf("the epoch block number is not equal: %d, %d", res.Batch[i].EpochBlockNumber(), res.Batch[0].EpochBlockNumber())
		}
		if res.Batch[i].CurrentCommittee() != common.Bytes2Hex(committeeData.CurrentCommittee.Root[:]) {
			return nil, fmt.Errorf("the block committee root %s is not equal to the current root %v", res.Batch[i].CurrentCommittee(), committeeData)
		}
		if i > 0 && res.Batch[i-1].NextCommittee() != res.Batch[i].CurrentCommittee() {
			return nil, fmt.Errorf("the block next committee root %s is not equal to the next block's committee root %s", res.Batch[i-1].NextCommittee(), res.Batch[i].CurrentCommittee())
		}
	}
	c.nextCommitteeRoot = res.Batch[len(res.Batch)-1].NextCommittee()

	wg := sync.WaitGroup{}
	wg.Add(len(res.Batch))
	chError := make(chan error, len(res.Batch))

	for _, block := range res.Batch {
		go func(block *sequencertypes.Block) {
			defer wg.Done()
			// verify the proposer signature
			if len(block.ProposerPubKey()) == 0 {
				chError <- fmt.Errorf("the block %d proposer key is empty", block.BlockNumber())
				return
			}
			blsSigHash := block.BlsSignature().Hash()
			// verify the proposer signature
			verified, err := c.blsScheme.VerifySignature(common.FromHex(block.ProposerPubKey()), blsSigHash, common.FromHex(block.ProposerSignature()))
			if err != nil || !verified {
				chError <- fmt.Errorf("failed to verify the proposer signature: %v", err)
				return
			}
			// verify if the block hash is correct
			blockHash := block.BlockHash()
			rBlockHash, err := c.rpcClient.GetBlockHashByNumber(block.BlockNumber())
			if err != nil {
				chError <- fmt.Errorf("failed to fetch the block hash by number: %v", err)
				return
			}
			if blockHash != rBlockHash {
				chError <- fmt.Errorf("the block hash %s is not equal to the rpc block hash %s", blockHash, rBlockHash)
				return
			}
		}(block)
	}
	wg.Wait()

	close(chError)
	for err := range chError {
		logger.Errorf("failed to verify the block: %v", err)
		return nil, err
	}

	return res.Batch, nil
}

// TryCommitBlocks tries to commit the signature to the network.
func (c *Client) TryCommitBlocks(blocks []*sequencertypes.Block) error {
	wg := sync.WaitGroup{}
	wg.Add(len(blocks))
	chError := make(chan error, len(blocks))

	sigs := make(chan *sequencertypes.BlsSignature, len(blocks))
	for _, block := range blocks {
		go func(block *sequencertypes.Block) {
			defer wg.Done()
			blsSignature := block.BlsSignature()
			blsSig, err := c.blsScheme.Sign(c.blsPrivateKey, blsSignature.Hash())
			if err != nil {
				chError <- fmt.Errorf("failed to sign the BLS signature: %v", err)
				return
			}
			blsSignature.BlsSignature = utils.Bytes2Hex(blsSig)

			// generate the ECDSA signature
			msg := contypes.GetCommitRequestHash(blsSignature)
			sig, err := ecrypto.Sign(msg, c.ecdsaPrivateKey)
			if err != nil {
				chError <- fmt.Errorf("failed to ecdsa sign the block: %v", err)
				return
			}
			blsSignature.EcdsaSignature = common.Bytes2Hex(sig)
			sigs <- blsSignature
		}(block)
	}

	wg.Wait()
	close(chError)
	close(sigs)
	for err := range chError {
		logger.Errorf("failed to sign the block: %v", err)
		return err
	}

	// upload the signature
	blsSignatures := make([]*sequencertypes.BlsSignature, 0, len(blocks))
	for blsSignature := range sigs {
		blsSignatures = append(blsSignatures, blsSignature)
	}

	req := &types.CommitBatchRequest{
		BlsSignatures: blsSignatures,
		StakeAddress:  c.stakeAddress,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	stream, err := c.CommitBatch(ctx, req)
	if err != nil {
		return fmt.Errorf("failed to upload signature: %v", err)
	}

	res, err := stream.Recv()
	if err != nil {
		return fmt.Errorf("failed to get the response from the stream: %v", err)
	}
	if !res.Result {
		return fmt.Errorf("the current batch is not finalized yet")
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

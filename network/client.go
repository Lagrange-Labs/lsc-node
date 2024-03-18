package network

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/lru"
	ecrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/protobuf/proto"

	"github.com/Lagrange-Labs/lagrange-node/crypto"
	"github.com/Lagrange-Labs/lagrange-node/logger"
	networkv2types "github.com/Lagrange-Labs/lagrange-node/network/types/v2"
	"github.com/Lagrange-Labs/lagrange-node/rpcclient"
	rpctypes "github.com/Lagrange-Labs/lagrange-node/rpcclient/types"
	"github.com/Lagrange-Labs/lagrange-node/scinterface/committee"
	sequencerv2types "github.com/Lagrange-Labs/lagrange-node/sequencer/types/v2"
	"github.com/Lagrange-Labs/lagrange-node/utils"
)

const CommitteeCacheSize = 10

type PreviousBatchInfo struct {
	NextCommitteeRoot string
	L1BlockNumber     uint64
}

// Client is a gRPC client to join the network
type Client struct {
	networkv2types.NetworkServiceClient
	rpcClient   rpctypes.RpcClient
	committeeSC *committee.Committee
	blsScheme   crypto.BLSScheme

	chainID         uint32
	blsPrivateKey   []byte
	blsPublicKey    string
	ecdsaPrivateKey *ecdsa.PrivateKey
	jwToken         string
	stakeAddress    string
	openBatchNumber uint64
	pullInterval    time.Duration
	committeeCache  *lru.Cache[uint64, *committee.ILagrangeCommitteeCommitteeData]

	ctx        context.Context
	cancelFunc context.CancelFunc
}

// NewClient creates a new client.
func NewClient(cfg *ClientConfig, rpcCfg *rpcclient.Config) (*Client, error) {
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

	rpcClient, err := rpcclient.NewClient(cfg.Chain, rpcCfg)
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
		NetworkServiceClient: networkv2types.NewNetworkServiceClient(conn),
		blsScheme:            blsScheme,
		blsPrivateKey:        blsPriv,
		blsPublicKey:         utils.Bytes2Hex(pubkey),
		ecdsaPrivateKey:      ecdsaPriv,
		stakeAddress:         stakeAddress,
		pullInterval:         time.Duration(cfg.PullInterval),
		rpcClient:            rpcClient,
		committeeSC:          committeeSC,
		chainID:              chainID,
		committeeCache:       lru.NewCache[uint64, *committee.ILagrangeCommitteeCommitteeData](CommitteeCacheSize),

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
	c.TryJoinNetwork()

	for {
		select {
		case <-c.ctx.Done():
			return
		case <-time.After(c.pullInterval):
			batch, err := c.TryGetBatch()
			if err != nil {
				if errors.Is(err, ErrInvalidToken) {
					c.TryJoinNetwork()
					continue
				}
				logger.Errorf("failed to get the current block: %v\n", err)
				continue
			}

			if err := c.TryCommitBatch(batch); err != nil {
				if errors.Is(err, ErrInvalidToken) {
					c.TryJoinNetwork()
					continue
				}
				logger.Errorf("failed to commit the block: %v\n", err)
				continue
			}

			c.openBatchNumber = batch.BatchNumber() + 1
			logger.Infof("uploaded the signature up to block %d\n", batch.BatchHeader.ToBlockNumber())
		}
	}
}

// TryJoinNetwork tries to join the network.
func (c *Client) TryJoinNetwork() {
	for {
		if err := c.joinNetwork(); err != nil {
			logger.Errorf("failed to join the network: %v", err)
			time.Sleep(5 * time.Second)
			continue
		}
		logger.Infof("joined the network with the new token: %s\n", c.jwToken)
		break
	}
}

func (c *Client) joinNetwork() error {
	req := &networkv2types.JoinNetworkRequest{
		PublicKey:    c.blsPublicKey,
		StakeAddress: c.stakeAddress,
	}
	reqMsg, err := proto.Marshal(req)
	if err != nil {
		return fmt.Errorf("failed to marshal the request: %v", err)
	}
	sig, err := c.blsScheme.Sign(c.blsPrivateKey, reqMsg)
	if err != nil {
		return fmt.Errorf("failed to sign the request: %v", err)
	}
	req.Signature = utils.Bytes2Hex(sig)
	res, err := c.NetworkServiceClient.JoinNetwork(context.Background(), req)
	if err != nil {
		return fmt.Errorf("failed to join the network: %v", err)
	}
	if len(res.Token) == 0 {
		return fmt.Errorf("the token is empty")
	}

	c.jwToken = res.Token
	c.openBatchNumber = res.OpenBatchNumber
	c.rpcClient.SetBeginBlockNumber(res.L1BlockNumber)

	return nil
}

// TryGetBatch tries to get the batch from the network.
func (c *Client) TryGetBatch() (*sequencerv2types.Batch, error) {
	res, err := c.GetBatch(context.Background(), &networkv2types.GetBatchRequest{
		BatchNumber: c.openBatchNumber, StakeAddress: c.stakeAddress, Token: c.jwToken})
	if err != nil {
		if strings.Contains(err.Error(), ErrInvalidToken.Error()) {
			return nil, ErrInvalidToken
		}
		return nil, err
	}

	if res.Batch == nil {
		return nil, ErrBlockNotReady
	}
	batch := res.Batch
	logger.Infof("got the batch the block number from %d to %d\n", res.FromBlockNumber, res.ToBlockNumber)

	// verify the L1 block number
	batchHeader, err := c.rpcClient.GetBatchHeaderByNumber(res.FromBlockNumber)
	if err != nil {
		return nil, fmt.Errorf("failed to get the batch header by number: %v", err)
	}
	if batch.L1BlockNumber() != batchHeader.L1BlockNumber {
		return nil, fmt.Errorf("the batch L1 block number %d is not equal to the rpc L1 block number %d", res.Batch.L1BlockNumber(), batchHeader.L1BlockNumber)
	}
	if res.FromBlockNumber != batchHeader.FromBlockNumber() {
		return nil, fmt.Errorf("the batch from block number %d is not equal to the rpc from block number %d", res.FromBlockNumber, batchHeader.FromBlockNumber())
	}
	if res.ToBlockNumber != batchHeader.ToBlockNumber() {
		return nil, fmt.Errorf("the batch to block number %d is not equal to the rpc to block number %d", res.ToBlockNumber, batchHeader.ToBlockNumber())
	}
	// verify the committee root
	if err := c.verifyCommitteeRoot(batch); err != nil {
		return nil, fmt.Errorf("failed to verify the committee root: %v", err)
	}

	// verify the proposer signature
	if len(batch.ProposerPubKey) == 0 {
		return nil, fmt.Errorf("the block %d proposer key is empty", batch.BatchNumber())
	}
	blsSigHash := batch.BlsSignature().Hash()
	// verify the proposer signature
	verified, err := c.blsScheme.VerifySignature(common.FromHex(batch.ProposerPubKey), blsSigHash, common.FromHex(batch.ProposerSignature))
	if err != nil || !verified {
		return nil, fmt.Errorf("failed to verify the proposer signature: %v", err)
	}
	// verify if the batch hash is correct
	batchHash := batch.BatchHash()
	bhHash := batchHeader.Hash()
	if !bytes.Equal(common.FromHex(batchHash), bhHash) {
		return nil, fmt.Errorf("the batch hash %s is not equal to the batch header hash %s", batchHash, utils.Bytes2Hex(bhHash))
	}

	return batch, nil
}

func (c *Client) getCommitteeRoot(blockNumber uint64) (*committee.ILagrangeCommitteeCommitteeData, error) { //nolint
	if committeeData, ok := c.committeeCache.Get(blockNumber); ok {
		return committeeData, nil
	}

	committeeData, err := c.committeeSC.GetCommittee(nil, c.chainID, big.NewInt(int64(blockNumber)))
	if err != nil {
		return nil, fmt.Errorf("failed to get the committee data: %v", err)
	}
	c.committeeCache.Add(blockNumber, &committeeData.CurrentCommittee)

	return &committeeData.CurrentCommittee, nil
}

// TODO: verify the committee root
func (c *Client) verifyCommitteeRoot(batch *sequencerv2types.Batch) error {
	// initialize the previous batch info
	// if len(c.prevBatchInfo.NextCommitteeRoot) == 0 {
	// 	previousBlock, err := c.GetBlock(context.Background(), &types.GetBlockRequest{BlockNumber: batch[0].BlockNumber() - 1, StakeAddress: c.stakeAddress, Token: c.jwToken})
	// 	if err != nil && !strings.Contains(err.Error(), "block not found") {
	// 		// TODO: handle the block not found error
	// 		return fmt.Errorf("failed to get the previous block: %v", err)
	// 	}
	// 	if previousBlock == nil || previousBlock.Block == nil {
	// 		c.prevBatchInfo = PreviousBatchInfo{
	// 			NextCommitteeRoot: batch[0].CurrentCommittee(),
	// 			L1BlockNumber:     batch[0].L1BlockNumber(),
	// 		}
	// 	} else {
	// 		blockHeader, err := c.rpcClient.GetBlockHeaderByNumber(previousBlock.Block.BlockNumber(), previousBlock.Block.L1TxHash())
	// 		if err != nil {
	// 			return fmt.Errorf("failed to get the block header by number: %v", err)
	// 		}
	// 		if previousBlock.Block.L1BlockNumber() != blockHeader.L1BlockNumber {
	// 			return fmt.Errorf("the previous block L1 block number %d is not equal to the rpc L1 block number %d", previousBlock.Block.L1BlockNumber(), blockHeader.L1BlockNumber)
	// 		}

	// 		previousCommitteeData, err := c.getCommitteeRoot(blockHeader.L1BlockNumber)
	// 		if err != nil {
	// 			return fmt.Errorf("failed to get the previous committee root: %v", err)
	// 		}
	// 		if previousBlock.Block.NextCommittee() != utils.Bytes2Hex(previousCommitteeData.Root[:]) {
	// 			return fmt.Errorf("the previous block next committee root %s is not equal to the epoch committee root %s", previousBlock.Block.NextCommittee(), utils.Bytes2Hex(previousCommitteeData.Root[:]))
	// 		}

	// 		c.prevBatchInfo = PreviousBatchInfo{
	// 			NextCommitteeRoot: previousBlock.Block.NextCommittee(),
	// 			L1BlockNumber:     previousBlock.Block.L1BlockNumber(),
	// 		}
	// 	}
	// }
	// // verify the previous block's next committee root
	// if len(c.prevBatchInfo.NextCommitteeRoot) > 0 && batch[0].CurrentCommittee() != c.prevBatchInfo.NextCommitteeRoot {
	// 	return fmt.Errorf("the block committee root %s is not equal to the previous batch's next committee root %s", batch[0].CurrentCommittee(), c.prevBatchInfo.NextCommitteeRoot)
	// }
	// for i, block := range batch {
	// 	if i > 0 && block.CurrentCommittee() != batch[i-1].NextCommittee() {
	// 		return fmt.Errorf("the block %d committee root %s is not equal to the previous block's next committee root %s", i, block.CurrentCommittee(), batch[i-1].NextCommittee())
	// 	}
	// }

	// lastBlock := batch[len(batch)-1]
	// previousBlockNumber := c.prevBatchInfo.L1BlockNumber
	// if len(batch) > 1 {
	// 	previousBlockNumber = batch[len(batch)-2].L1BlockNumber()
	// }
	// previousCommitteeData, err := c.getCommitteeRoot(previousBlockNumber)
	// if err != nil {
	// 	return fmt.Errorf("failed to get the previous committee root: %v", err)
	// }
	// for _, block := range batch {
	// 	if block.CurrentCommittee() != utils.Bytes2Hex(previousCommitteeData.Root[:]) {
	// 		return fmt.Errorf("the block %d committee root %s is not equal to the epoch committee root %s", block.BlockNumber(), block.CurrentCommittee(), utils.Bytes2Hex(previousCommitteeData.Root[:]))
	// 	}
	// }
	// // still can verify the next committee root even if the committee epoch rotates
	// committeeData, err := c.getCommitteeRoot(lastBlock.L1BlockNumber())
	// if err != nil {
	// 	return fmt.Errorf("failed to get the committee root: %v", err)
	// }
	// if lastBlock.NextCommittee() != utils.Bytes2Hex(committeeData.Root[:]) {
	// 	return fmt.Errorf("the last block %d next committee root %s is not equal to the epoch committee root %s", lastBlock.BlockNumber(), lastBlock.NextCommittee(), utils.Bytes2Hex(committeeData.Root[:]))
	// }

	// c.prevBatchInfo = PreviousBatchInfo{
	// 	NextCommitteeRoot: lastBlock.NextCommittee(),
	// 	L1BlockNumber:     lastBlock.L1BlockNumber(),
	// }

	return nil
}

// TryCommitBatch tries to commit the signature to the network.
func (c *Client) TryCommitBatch(batch *sequencerv2types.Batch) error {
	blsSignature := batch.BlsSignature()
	blsSig, err := c.blsScheme.Sign(c.blsPrivateKey, blsSignature.Hash())
	if err != nil {
		return fmt.Errorf("failed to sign the BLS signature: %v", err)
	}
	blsSignature.BlsSignature = utils.Bytes2Hex(blsSig)

	// generate the ECDSA signature
	msg := blsSignature.CommitHash()
	sig, err := ecrypto.Sign(msg, c.ecdsaPrivateKey)
	if err != nil {
		return fmt.Errorf("failed to ecdsa sign the block: %v", err)
	}
	blsSignature.EcdsaSignature = common.Bytes2Hex(sig)

	req := &networkv2types.CommitBatchRequest{
		BlsSignature: blsSignature,
		StakeAddress: c.stakeAddress,
		Token:        c.jwToken,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	stream, err := c.CommitBatch(ctx, req)
	if err != nil {
		if strings.Contains(err.Error(), ErrInvalidToken.Error()) {
			return ErrInvalidToken
		}
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
	if c != nil {
		c.cancelFunc()
	}
}

var (
	ErrBlockNotReady = fmt.Errorf("the block is not ready")
)

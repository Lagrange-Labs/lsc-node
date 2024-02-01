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
	"github.com/ethereum/go-ethereum/common/lru"
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

const CommitteeCacheSize = 10

type NextBlockInfo struct {
	NextCommitteeRoot string
	L1BlockNumber     uint64
}

// Client is a gRPC client to join the network
type Client struct {
	types.NetworkServiceClient
	rpcClient   rpctypes.RpcClient
	committeeSC *committee.Committee
	blsScheme   crypto.BLSScheme

	chainID         uint32
	blsPrivateKey   []byte
	blsPublicKey    string
	ecdsaPrivateKey *ecdsa.PrivateKey
	stakeAddress    string
	lastBlockNumber uint64
	pullInterval    time.Duration
	nextBlockInfo   NextBlockInfo
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
		logger.Errorf("failed to marshal the request: %v", err)
		return err
	}
	sig, err := c.blsScheme.Sign(c.blsPrivateKey, reqMsg)
	if err != nil {
		logger.Errorf("failed to sign the request: %v", err)
		return err
	}
	req.Signature = utils.Bytes2Hex(sig)
	res, err := c.NetworkServiceClient.JoinNetwork(context.Background(), req)
	if err != nil {
		logger.Errorf("failed to join the network: %v", err)
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

	logger.Infof("got the batch the block number from %d to %d\n", res.Batch[0].BlockNumber(), res.Batch[len(res.Batch)-1].BlockNumber())

	wg := sync.WaitGroup{}
	wg.Add(len(res.Batch))
	chError := make(chan error, len(res.Batch))

	blockHeaders := make(map[uint64]*rpctypes.L2BlockHeader)
	for _, block := range res.Batch {
		blockHeader, err := c.rpcClient.GetBlockHeaderByNumber(block.BlockNumber(), block.L1TxHash())
		if err != nil {
			return nil, fmt.Errorf("failed to get the block header by number: %v", err)
		}
		blockHeaders[block.BlockNumber()] = blockHeader
	}

	// verify the L1 block number
	for i, block := range res.Batch {
		l1BlockNumber := blockHeaders[block.BlockNumber()].L1BlockNumber
		if l1BlockNumber != block.L1BlockNumber() {
			return nil, fmt.Errorf("the L1 block number %d is not equal to the rpc L1 block number %d", block.L1BlockNumber(), l1BlockNumber)
		}
		if i > 0 && block.BlockNumber() != res.Batch[i-1].BlockNumber()+1 {
			return nil, fmt.Errorf("the batch blocks order is not sorted at index: %d the current block number: %d the previous block number: %d", i, block.BlockNumber(), res.Batch[i-1].BlockNumber())
		}
	}

	if err := c.verifyCommitteeRoot(res.Batch); err != nil {
		return nil, fmt.Errorf("failed to verify the committee root: %v", err)
	}

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
			rBlockHash := blockHeaders[block.BlockNumber()].L2BlockHash.Hex()
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

func (c *Client) getCommitteeRoot(blockNumber uint64) (*committee.ILagrangeCommitteeCommitteeData, error) {
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

func (c *Client) verifyCommitteeRoot(batch []*sequencertypes.Block) error {
	// initialize the next block info
	if len(c.nextBlockInfo.NextCommitteeRoot) == 0 {
		previousBlock, err := c.GetBlock(context.Background(), &types.GetBlockRequest{BlockNumber: batch[0].BlockNumber() - 1})
		if err != nil {
			return fmt.Errorf("failed to get the previous block: %v", err)
		}
		if previousBlock.Block == nil {
			c.nextBlockInfo = NextBlockInfo{
				NextCommitteeRoot: batch[0].CurrentCommittee(),
				L1BlockNumber:     batch[0].L1BlockNumber(),
			}
		} else {
			blockHeader, err := c.rpcClient.GetBlockHeaderByNumber(previousBlock.Block.BlockNumber(), previousBlock.Block.L1TxHash())
			if err != nil {
				return fmt.Errorf("failed to get the block header by number: %v", err)
			}
			if previousBlock.Block.L1BlockNumber() != blockHeader.L1BlockNumber {
				return fmt.Errorf("the previous block L1 block number %d is not equal to the rpc L1 block number %d", previousBlock.Block.L1BlockNumber(), blockHeader.L1BlockNumber)
			}

			previousCommitteeData, err := c.getCommitteeRoot(blockHeader.L1BlockNumber)
			if err != nil {
				return fmt.Errorf("failed to get the previous committee root: %v", err)
			}
			if previousBlock.Block.NextCommittee() != common.Bytes2Hex(previousCommitteeData.Root.Bytes()) {
				return fmt.Errorf("the previous block next committee root %s is not equal to the epoch committee root %s", previousBlock.Block.NextCommittee(), common.Bytes2Hex(previousCommitteeData.Root.Bytes()))
			}

			c.nextBlockInfo = NextBlockInfo{
				NextCommitteeRoot: previousBlock.Block.NextCommittee(),
				L1BlockNumber:     previousBlock.Block.L1BlockNumber(),
			}
		}
	}
	// verify the previous block's next committee root
	if len(c.nextBlockInfo.NextCommitteeRoot) > 0 && batch[0].CurrentCommittee() != c.nextBlockInfo.NextCommitteeRoot {
		return fmt.Errorf("the block committee root %s is not equal to the previous batch's next committee root %s", batch[0].CurrentCommittee(), c.nextBlockInfo.NextCommitteeRoot)
	}
	for i, block := range batch {
		if i > 0 && block.CurrentCommittee() != batch[i-1].NextCommittee() {
			return fmt.Errorf("the block %d committee root %s is not equal to the previous block's next committee root %s", i, block.CurrentCommittee(), batch[i-1].NextCommittee())
		}
	}

	lastBlock := batch[len(batch)-1]
	previousBlockNumber := c.nextBlockInfo.L1BlockNumber
	if len(batch) > 1 {
		previousBlockNumber = batch[len(batch)-2].L1BlockNumber()
	}
	previousCommitteeData, err := c.getCommitteeRoot(previousBlockNumber)
	if err != nil {
		return fmt.Errorf("failed to get the previous committee root: %v", err)
	}
	for _, block := range batch {
		if block.CurrentCommittee() != common.Bytes2Hex(previousCommitteeData.Root.Bytes()) {
			return fmt.Errorf("the block %d committee root %s is not equal to the epoch committee root %s", block.BlockNumber(), block.CurrentCommittee(), common.Bytes2Hex(previousCommitteeData.Root.Bytes()))
		}
	}
	// still can verify the next committee root even if the committee epoch rotates
	committeeData, err := c.getCommitteeRoot(lastBlock.L1BlockNumber())
	if err != nil {
		return fmt.Errorf("failed to get the committee root: %v", err)
	}
	if lastBlock.NextCommittee() != common.Bytes2Hex(committeeData.Root.Bytes()) {
		return fmt.Errorf("the last block %d next committee root %s is not equal to the epoch committee root %s", lastBlock.BlockNumber(), lastBlock.NextCommittee(), common.Bytes2Hex(committeeData.Root.Bytes()))
	}

	c.nextBlockInfo = NextBlockInfo{
		NextCommitteeRoot: lastBlock.NextCommittee(),
		L1BlockNumber:     lastBlock.L1BlockNumber(),
	}

	return nil
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
	if c != nil {
		c.cancelFunc()
	}
}

var (
	ErrBlockNotReady = fmt.Errorf("the block is not ready")
)

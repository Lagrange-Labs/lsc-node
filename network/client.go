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

	chainID           uint32
	blsPrivateKey     *bls.SecretKey
	blsPublicKey      string
	ecdsaPrivateKey   *ecdsa.PrivateKey
	stakeAddress      string
	lastBlockNumber   uint64
	pullInterval      time.Duration
	nextCommitteeRoot string

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

	rpcClient, err := rpcclient.NewClient(cfg.Chain, cfg.RPCEndpoint, cfg.EthereumURL, "")
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
	sig, err := c.blsPrivateKey.Sign(reqMsg)
	if err != nil {
		return err
	}
	req.Signature = utils.BlsSignatureToHex(sig)
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

	logger.Info("got the batch: ", res.Batch)

	wg := sync.WaitGroup{}
	wg.Add(len(res.Batch))
	chError := make(chan error, len(res.Batch))

	// verify the L1 block number
	for i, block := range res.Batch {
		l1BlockNumber, err := c.rpcClient.GetL1BlockNumber(block.BlockNumber())
		if err != nil {
			return nil, fmt.Errorf("failed to get the L1 block number: %v", err)
		}
		if l1BlockNumber != block.L1BlockNumber() {
			return nil, fmt.Errorf("the L1 block number %d is not equal to the rpc L1 block number %d", block.L1BlockNumber(), l1BlockNumber)
		}
		if i > 0 && block.L1BlockNumber() < res.Batch[i-1].L1BlockNumber() {
			return nil, fmt.Errorf("the batch blocks order is not sorted")
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
			verified, err := utils.VerifySignature(common.FromHex(block.ProposerPubKey()), blsSigHash, common.FromHex(block.ProposerSignature()))
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

func (c *Client) verifyCommitteeRoot(batch []*sequencertypes.Block) error {
	// verify the committee root
	if len(c.nextCommitteeRoot) > 0 && batch[0].CurrentCommittee() != c.nextCommitteeRoot {
		return fmt.Errorf("the block committee root %s is not equal to the previous batch's next committee root %s", batch[0].CurrentCommittee(), c.nextCommitteeRoot)
	}
	lastBlock := batch[len(batch)-1]
	firstEpochNumber, err := c.committeeSC.GetEpochNumber(nil, c.chainID, big.NewInt(int64(batch[0].L1BlockNumber())))
	if err != nil {
		return fmt.Errorf("failed to get the epoch number from SC: %v", err)
	}
	lastEpochNumber, err := c.committeeSC.GetEpochNumber(nil, c.chainID, big.NewInt(int64(lastBlock.L1BlockNumber())))
	if err != nil {
		return fmt.Errorf("failed to get the epoch number from SC: %v", err)
	}
	isRotated := false
	if firstEpochNumber.Uint64() != lastEpochNumber.Uint64() || len(batch) == 1 {
		// check the committee root rotation
		l1BlockNumber, err := c.rpcClient.GetL1BlockNumber(lastBlock.BlockNumber() - 1)
		if err != nil {
			return fmt.Errorf("failed to get the L1 block number: %v", err)
		}
		if l1BlockNumber >= lastBlock.L1BlockNumber() {
			if len(batch) > 1 {
				return fmt.Errorf("the committee rotation is detected but no L1 block number changes")
			}
		} else {
			prevEpochNumber, err := c.committeeSC.GetEpochNumber(nil, c.chainID, big.NewInt(int64(l1BlockNumber)))
			if err != nil {
				return fmt.Errorf("failed to get the epoch number from SC: %v", err)
			}
			if prevEpochNumber.Uint64() >= lastEpochNumber.Uint64() {
				if len(batch) > 1 {
					return fmt.Errorf("the committee rotation is detected but no epoch number changes")
				}
			} else {
				isRotated = true
				prevCommitteeData, err := c.committeeSC.GetCommittee(nil, c.chainID, big.NewInt(int64(l1BlockNumber)))
				if err != nil {
					return fmt.Errorf("failed to get the previous committee data")
				}
				currentCommitteeData, err := c.committeeSC.GetCommittee(nil, c.chainID, big.NewInt(int64(lastBlock.L1BlockNumber())))
				if err != nil {
					return fmt.Errorf("failed to get the current committee data")
				}

				if lastBlock.CurrentCommittee() != common.Bytes2Hex(prevCommitteeData.CurrentCommittee.Root.Bytes()) {
					return fmt.Errorf("the current committee root %s of the next epoch beginning block is not equal to the prev committee %v", lastBlock.CurrentCommittee(), prevCommitteeData)
				}
				if lastBlock.NextCommittee() != common.Bytes2Hex(currentCommitteeData.CurrentCommittee.Root.Bytes()) {
					return fmt.Errorf("the next committee root %s of the next epoch beginning block is not equal to the current committee %v", lastBlock.NextCommittee(), currentCommitteeData)
				}
			}
		}
	}
	committeeData, err := c.committeeSC.GetCommittee(nil, c.chainID, big.NewInt(int64(batch[0].L1BlockNumber())))
	if err != nil {
		return fmt.Errorf("failed to get the committee data: %v chainID: %d Batch: %v", err, c.chainID, batch)
	}
	currentCommitteeRoot := common.Bytes2Hex(committeeData.CurrentCommittee.Root.Bytes())
	for i := range batch {
		if batch[i].CurrentCommittee() != currentCommitteeRoot {
			if i != len(batch)-1 || !isRotated {
				return fmt.Errorf("the block committee root %s is not equal to the current root %v", batch[i].CurrentCommittee(), committeeData)
			}
		}
		if i > 0 && batch[i].NextCommittee() != batch[i-1].CurrentCommittee() {
			return fmt.Errorf("the block next committee root %s is not equal to the next block's committee root %s", batch[i].NextCommittee(), batch[i-1].CurrentCommittee())
		}
	}

	c.nextCommitteeRoot = lastBlock.CurrentCommittee()

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
			blsSig, err := c.blsPrivateKey.Sign(blsSignature.Hash())
			if err != nil {
				chError <- fmt.Errorf("failed to sign the BLS signature: %v", err)
				return
			}
			blsSignature.BlsSignature = utils.BlsSignatureToHex(blsSig)

			// generate the ECDSA signature
			msg := contypes.GetCommitRequestHash(blsSignature)
			sig, err := crypto.Sign(msg, c.ecdsaPrivateKey)
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

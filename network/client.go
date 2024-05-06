package network

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/binary"
	"errors"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"strings"
	"sync/atomic"
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
	"github.com/Lagrange-Labs/lagrange-node/store/goleveldb"
	"github.com/Lagrange-Labs/lagrange-node/telemetry"
	"github.com/Lagrange-Labs/lagrange-node/utils"
)

const (
	CommitteeCacheSize = 10
	ClientDBPath       = ".lagrange/db/"
	PruningBlocks      = 1000
)

var (
	// ErrBatchNotFinalized is returned when the current batch is not finalized yet.
	ErrBatchNotFinalized = errors.New("the current batch is not finalized yet")
)

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

	chainID               uint32
	blsPrivateKey         []byte
	blsPublicKey          string
	signerECDSAPrivateKey *ecdsa.PrivateKey
	jwToken               string
	stakeAddress          string
	genesisBlockNumber    uint64
	pullInterval          time.Duration
	committeeCache        *lru.Cache[uint64, *committee.ILagrangeCommitteeCommitteeData]
	openL1BlockNumber     atomic.Uint64

	db         *goleveldb.DB
	ctx        context.Context
	cancelFunc context.CancelFunc
	chErr      chan error
}

// NewClient creates a new client.
func NewClient(cfg *ClientConfig, rpcCfg *rpcclient.Config) (*Client, error) {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	conn, err := grpc.NewClient(cfg.GrpcURL, opts...)
	if err != nil {
		return nil, err
	}

	healthClient := grpc_health_v1.NewHealthClient(conn)
	hctx, hcancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer hcancel()

	watcher, err := healthClient.Watch(hctx, &grpc_health_v1.HealthCheckRequest{})
	if err != nil {
		logger.Error("Failed to check gRPC health:", err)
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

	if len(cfg.BLSKeystorePasswordPath) > 0 {
		cfg.BLSKeystorePassword, err = crypto.ReadKeystorePasswordFromFile(cfg.BLSKeystorePasswordPath)
		if err != nil {
			logger.Fatalf("failed to read the bls keystore password from %s: %v", cfg.BLSKeystorePasswordPath, err)
		}
	}
	blsPriv, err := crypto.LoadPrivateKey(crypto.CryptoCurve(cfg.BLSCurve), cfg.BLSKeystorePassword, cfg.BLSKeystorePath)
	if err != nil {
		logger.Fatalf("failed to load the bls keystore from %s: %v", cfg.BLSKeystorePath, err)
	}
	blsScheme := crypto.NewBLSScheme(crypto.BLSCurve(cfg.BLSCurve))
	pubkey, err := blsScheme.GetPublicKey(blsPriv, false)
	if err != nil {
		logger.Fatalf("failed to get the bls public key: %v", err)
	}

	if len(cfg.SignerECDSAKeystorePasswordPath) > 0 {
		cfg.SignerECDSAKeystorePassword, err = crypto.ReadKeystorePasswordFromFile(cfg.SignerECDSAKeystorePasswordPath)
		if err != nil {
			logger.Fatalf("failed to read the ecdsa keystore password from %s: %v", cfg.SignerECDSAKeystorePasswordPath, err)
		}
	}
	ecdsaPrivKey, err := crypto.LoadPrivateKey(crypto.CryptoCurve("ECDSA"), cfg.SignerECDSAKeystorePassword, cfg.SignerECDSAKeystorePath)
	if err != nil {
		logger.Fatalf("failed to load the ecdsa keystore from %s: %v", cfg.SignerECDSAKeystorePath, err)
	}
	ecdsaPriv, err := ecrypto.ToECDSA(ecdsaPrivKey)
	if err != nil {
		logger.Fatalf("failed to get the ecdsa private key: %v", err)
	}

	rpcClient, err := rpcclient.NewClient(cfg.Chain, rpcCfg)
	if err != nil {
		logger.Fatalf("failed to create the rpc client: %v, please check the chain name, the chain name should look like 'optimism', 'base'", err)
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

	params, err := committeeSC.CommitteeParams(nil, chainID)
	if err != nil {
		logger.Fatalf("failed to get the committee params: %v", err)
	}

	homePath, err := os.UserHomeDir()
	if err != nil {
		logger.Fatalf("failed to get the home directory: %v", err)
	}
	dbPath := filepath.Clean(filepath.Join(homePath, ClientDBPath))
	if err := os.MkdirAll(dbPath, os.ModePerm); err != nil {
		logger.Fatalf("failed to create the database directory: %v", err)
	}
	pubKey, err := blsScheme.GetPublicKey(blsPriv, false)
	if err != nil {
		logger.Fatalf("failed to get the public key: %v", err)
	}
	dbPath = filepath.Join(dbPath, fmt.Sprintf("client_%d_%x.db", chainID, pubKey))
	db, err := goleveldb.NewDB(dbPath)
	if err != nil {
		logger.Fatalf("failed to create the database: %v", err)
	}
	ctx, cancel := context.WithCancel(context.Background())

	c := &Client{
		NetworkServiceClient:  networkv2types.NewNetworkServiceClient(conn),
		blsScheme:             blsScheme,
		blsPrivateKey:         blsPriv,
		blsPublicKey:          utils.Bytes2Hex(pubkey),
		signerECDSAPrivateKey: ecdsaPriv,
		stakeAddress:          cfg.OperatorAddress,
		pullInterval:          time.Duration(cfg.PullInterval),
		rpcClient:             rpcClient,
		committeeSC:           committeeSC,
		chainID:               chainID,
		genesisBlockNumber:    uint64(params.GenesisBlock.Int64() - params.L1Bias.Int64()),
		committeeCache:        lru.NewCache[uint64, *committee.ILagrangeCommitteeCommitteeData](CommitteeCacheSize),

		db:         db,
		ctx:        ctx,
		cancelFunc: cancel,
		chErr:      make(chan error),
	}
	go c.startBatchFetching()

	return c, nil
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
func (c *Client) Start() error {
	c.TryJoinNetwork()

	for {
		select {
		case <-c.ctx.Done():
			return errors.New("the client is stopped")
		case err := <-c.chErr:
			return err
		case <-time.After(c.pullInterval):
			batch, err := c.TryGetBatch()
			if err != nil {
				if errors.Is(err, ErrInvalidToken) {
					c.TryJoinNetwork()
					continue
				}
				if errors.Is(err, ErrBatchNotFound) {
					// if the batch is not found, set the begin block number to the L1 block number
					logger.Infof("the batch is not found, set the begin block number to the L1 block number %d", batch.L1BlockNumber())
					c.rpcClient.SetBeginBlockNumber(batch.L1BlockNumber(), batch.BatchHeader.FromBlockNumber())
					continue
				}
				if errors.Is(err, ErrBatchNotReady) {
					logger.Infof("the batch is not ready yet")
					continue
				}

				logger.Errorf("failed to get the current block: %v", err)
				continue
			}

			if err := c.TryCommitBatch(batch); err != nil {
				if errors.Is(err, ErrInvalidToken) {
					c.TryJoinNetwork()
				} else if errors.Is(err, ErrBatchNotFinalized) {
					logger.Infof("NOTE: the current batch is not finalized yet due to a lack of voting power. Please wait until getting enough voting power.")
				} else {
					logger.Errorf("failed to commit the batch: %v", err)
				}
				continue
			}

			logger.Infof("uploaded the signature up to block %d", batch.BatchHeader.ToBlockNumber())
		}
	}
}

// TryJoinNetwork tries to join the network.
func (c *Client) TryJoinNetwork() {
	for {
		if err := c.joinNetwork(); err != nil {
			logger.Infof("failed to join the network: %v", err)
			if strings.Contains(err.Error(), ErrNotCommitteeMember.Error()) {
				logger.Warn("NOTE: If you just joined the network, please wait for the next committee rotation. If you have been observing this message for a long time, please check if the BLS public key and the operator address are set correctly in the config file or contact the Largrange team.")
			} else if strings.Contains(err.Error(), ErrCheckCommitteeMember.Error()) {
				logger.Warn("NOTE: The given round is not initialized yet. It may be because the sequencer is waiting for the next batch since it is almost caught up with the current block. Please wait for the next batch.")
			}
			time.Sleep(5 * time.Second)
			continue
		}
		logger.Info("joined the network with the new token")
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
	ti := time.Now()
	res, err := c.NetworkServiceClient.JoinNetwork(context.Background(), req)
	if err != nil {
		return fmt.Errorf("failed to join the network: %v", err)
	}
	if len(res.Token) == 0 {
		return fmt.Errorf("the token is empty")
	}
	telemetry.MeasureSince(ti, "client", "join_network_request")

	c.jwToken = res.Token

	c.rpcClient.SetBeginBlockNumber(res.PrevL1BlockNumber, res.PrevL2BlockNumber)
	return c.verifyPrevBatch(res.PrevL1BlockNumber, res.PrevL2BlockNumber)
}

// startBatchFetching starts the batch fetching loop.
func (c *Client) startBatchFetching() {
	for {
		batch, err := c.rpcClient.NextBatch()
		if err != nil {
			logger.Errorf("failed to get the next batch: %v", err)
			c.chErr <- err
			return
		}
		// block the writeBatchHeader if the batch is too far from the current block
		for openBlockNumber := c.openL1BlockNumber.Load(); openBlockNumber > 0 && openBlockNumber+PruningBlocks/2 < batch.L1BlockNumber; openBlockNumber = c.openL1BlockNumber.Load() {
			time.Sleep(1 * time.Second)
		}
		if err := c.writeBatchHeader(batch); err != nil {
			logger.Errorf("failed to write the batch header: %v", err)
			c.chErr <- err
			return
		}
		if batch.L1BlockNumber > PruningBlocks {
			prunedBlockNumber := batch.L1BlockNumber - PruningBlocks
			prefix := make([]byte, 8)
			binary.BigEndian.PutUint64(prefix, prunedBlockNumber)
			if err := c.db.Prune(prefix); err != nil {
				logger.Errorf("failed to prune the database: %v", err)
			}
		}
		time.Sleep(100 * time.Millisecond)
	}
}

// writeBatchHeader writes the batch header to the database.
func (c *Client) writeBatchHeader(batchHeader *sequencerv2types.BatchHeader) error {
	key := make([]byte, 12)
	binary.BigEndian.PutUint64(key, batchHeader.L1BlockNumber)
	binary.BigEndian.PutUint32(key[8:], batchHeader.L1TxIndex)
	value, err := proto.Marshal(batchHeader)
	if err != nil {
		return fmt.Errorf("failed to marshal the batch header: %v", err)
	}

	return c.db.Put(key, value)
}

// getPrevBatchL1Number gets the previous batch L1 number from the database.
func (c *Client) getPrevBatchL1Number(l1BlockNumber uint64, l1TxIndex uint32) (uint64, error) {
	key := make([]byte, 12)
	binary.BigEndian.PutUint64(key, l1BlockNumber)
	binary.BigEndian.PutUint32(key[8:], l1TxIndex)

	prevKey, _, err := c.db.Prev(key)
	if err != nil {
		return 0, fmt.Errorf("failed to get the previous key: %v", err)
	}
	var prevL1BlockNumber uint64
	if prevKey != nil {
		prevL1BlockNumber = binary.BigEndian.Uint64(prevKey[:8])
	}

	return prevL1BlockNumber, nil
}

// getBatchHeader gets the batch header from the database.
func (c *Client) getBatchHeader(l1BlockNumber, l2BlockNumber uint64) (*sequencerv2types.BatchHeader, error) {
	ti := time.Now()
	defer telemetry.MeasureSince(ti, "client", "get_batch_header")

	prefix := make([]byte, 8)
	binary.BigEndian.PutUint64(prefix, l1BlockNumber)

	var res *sequencerv2types.BatchHeader
	if err := c.db.Iterate(prefix, func(key, value []byte) error {
		var batchHeader sequencerv2types.BatchHeader
		if err := proto.Unmarshal(value, &batchHeader); err != nil {
			return fmt.Errorf("failed to unmarshal the batch header: %v", err)
		}
		if batchHeader.FromBlockNumber() == l2BlockNumber {
			res = &batchHeader
			return nil
		}
		return fmt.Errorf("the batch header is not found for the L1 block number %d, L2 block number %d", l1BlockNumber, l2BlockNumber)
	}); err != nil {
		return nil, fmt.Errorf("failed to iterate the database: %v", err)
	}

	return res, nil
}

// verifyPrevBatch verifies the previous batch.
func (c *Client) verifyPrevBatch(l1BlockNumber, l2BlockNumber uint64) error {
	batchHeader, err := c.getBatchHeader(l1BlockNumber, l2BlockNumber)
	if err != nil {
		return fmt.Errorf("failed to get the batch header: %v", err)
	}

	if batchHeader == nil {
		return fmt.Errorf("the batch header is not found for L1 block number %d, L2 block number %d", l1BlockNumber, l2BlockNumber)
	}

	return nil
}

// TryGetBatch tries to get the batch from the network.
func (c *Client) TryGetBatch() (*sequencerv2types.Batch, error) {
	ti := time.Now()
	res, err := c.GetBatch(context.Background(), &networkv2types.GetBatchRequest{StakeAddress: c.stakeAddress, Token: c.jwToken})
	if err != nil {
		if strings.Contains(err.Error(), ErrInvalidToken.Error()) {
			return nil, ErrInvalidToken
		}
		return nil, err
	}
	if res.Batch == nil {
		return nil, ErrBatchNotReady
	}
	telemetry.MeasureSince(ti, "client", "get_batch_request")

	batch := res.Batch
	fromBlockNumber := batch.BatchHeader.FromBlockNumber()
	toBlockNumber := batch.BatchHeader.ToBlockNumber()
	c.openL1BlockNumber.Store(batch.L1BlockNumber())
	logger.Infof("got the batch the block number from %d to %d", fromBlockNumber, toBlockNumber)

	// verify the L1 block number
	batchHeader, err := c.getBatchHeader(batch.L1BlockNumber(), fromBlockNumber)
	if err != nil || batchHeader == nil {
		logger.Errorf("failed to get the batch header: %v", err)
		return batch, ErrBatchNotFound
	}
	if batch.L1BlockNumber() != batchHeader.L1BlockNumber {
		return nil, fmt.Errorf("the batch L1 block number %d is not equal to the rpc L1 block number %d", res.Batch.L1BlockNumber(), batchHeader.L1BlockNumber)
	}
	if fromBlockNumber != batchHeader.FromBlockNumber() {
		return nil, fmt.Errorf("the batch from block number %d is not equal to the rpc from block number %d", fromBlockNumber, batchHeader.FromBlockNumber())
	}
	if toBlockNumber != batchHeader.ToBlockNumber() {
		return nil, fmt.Errorf("the batch to block number %d is not equal to the rpc to block number %d", toBlockNumber, batchHeader.ToBlockNumber())
	}

	// verify the committee root
	if err := c.verifyCommitteeRoot(batch); err != nil {
		return nil, fmt.Errorf("failed to verify the committee root: %v", err)
	}

	// verify if the batch hash is correct
	batchHash := batch.BatchHeader.Hash()
	bhHash := batchHeader.Hash()
	if !bytes.Equal(batchHash, bhHash) {
		return nil, fmt.Errorf("the batch hash %s is not equal to the batch header hash %s", batchHash, utils.Bytes2Hex(bhHash))
	}

	// verify the proposer signature
	if len(batch.ProposerPubKey) == 0 {
		return nil, fmt.Errorf("the block %d proposer key is empty", batch.BatchNumber())
	}
	blsSigHash := batch.BlsSignature().Hash()
	verified, err := c.blsScheme.VerifySignature(common.FromHex(batch.ProposerPubKey), blsSigHash, common.FromHex(batch.ProposerSignature))
	if err != nil || !verified {
		return nil, fmt.Errorf("failed to verify the proposer signature: %v", err)
	}

	telemetry.SetGauge(float64(batch.BatchNumber()), "client", "current_batch_number")

	return batch, nil
}

func (c *Client) getCommitteeRoot(blockNumber uint64) (*committee.ILagrangeCommitteeCommitteeData, error) { //nolint
	if committeeData, ok := c.committeeCache.Get(blockNumber); ok {
		return committeeData, nil
	}

	ti := time.Now()
	defer telemetry.MeasureSince(ti, "client", "get_committee")

	committeeData, err := c.committeeSC.GetCommittee(nil, c.chainID, big.NewInt(int64(blockNumber)))
	if err != nil {
		return nil, fmt.Errorf("failed to get the committee data: %v", err)
	}
	c.committeeCache.Add(blockNumber, &committeeData)

	return &committeeData, nil
}

func (c *Client) verifyCommitteeRoot(batch *sequencerv2types.Batch) error {
	blockNumber := batch.L1BlockNumber()
	prevBatchL1Number := batch.L1BlockNumber()
	isGenesis := c.genesisBlockNumber == blockNumber
	// verify the previous batch's next committee root
	if !isGenesis {
		var err error
		prevBatchL1Number, err = c.getPrevBatchL1Number(batch.L1BlockNumber(), batch.BatchHeader.L1TxIndex)
		if err != nil {
			return fmt.Errorf("failed to get the previous batch L1 number: %v", err)
		}
	}
	prevCommitteeData, err := c.getCommitteeRoot(prevBatchL1Number)
	if err != nil {
		return fmt.Errorf("failed to get the previous committee root: %v", err)
	}
	if !bytes.Equal(utils.Hex2Bytes(batch.CurrentCommittee()), prevCommitteeData.Root[:]) {
		return fmt.Errorf("the current batch committee root %s is not equal to the previous batch next committee root %s", batch.CurrentCommittee(), utils.Bytes2Hex(prevCommitteeData.Root[:]))
	}

	// verify the current batch's next committee root
	curCommitteeData, err := c.getCommitteeRoot(blockNumber)
	if err != nil {
		return fmt.Errorf("failed to get the current committee root: %v", err)
	}
	if !bytes.Equal(utils.Hex2Bytes(batch.NextCommittee()), curCommitteeData.Root[:]) {
		return fmt.Errorf("the current batch next committee root %s is not equal to the current committee root %s", batch.NextCommittee(), utils.Bytes2Hex(curCommitteeData.Root[:]))
	}

	return nil
}

// TryCommitBatch tries to commit the signature to the network.
func (c *Client) TryCommitBatch(batch *sequencerv2types.Batch) error {
	ti := time.Now()
	defer telemetry.MeasureSince(ti, "client", "try_commit_batch")

	blsSignature := batch.BlsSignature()
	blsSig, err := c.blsScheme.Sign(c.blsPrivateKey, blsSignature.Hash())
	if err != nil {
		return fmt.Errorf("failed to sign the BLS signature: %v", err)
	}
	blsSignature.BlsSignature = utils.Bytes2Hex(blsSig)

	// generate the ECDSA signature
	msg := blsSignature.CommitHash()
	sig, err := ecrypto.Sign(msg, c.signerECDSAPrivateKey)
	if err != nil {
		return fmt.Errorf("failed to ecdsa sign the block: %v", err)
	}
	blsSignature.EcdsaSignature = common.Bytes2Hex(sig)

	req := &networkv2types.CommitBatchRequest{
		BlsSignature: blsSignature,
		StakeAddress: c.stakeAddress,
		PublicKey:    c.blsPublicKey,
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
		return ErrBatchNotFinalized
	}

	telemetry.SetGauge(float64(batch.BatchNumber()), "client", "commit_batch_number")
	telemetry.AddSample(float32(batch.BatchNumber()), "client", "commit_batch_number_sample")

	return nil
}

// Stop function stops the client node.
func (c *Client) Stop() {
	if c != nil {
		c.cancelFunc()
	}
}

var (
	ErrBatchNotReady = fmt.Errorf("the batch is not ready")
	ErrBatchNotFound = fmt.Errorf("the batch is not found")
)

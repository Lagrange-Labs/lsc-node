package client

import (
	"bytes"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	ecrypto "github.com/ethereum/go-ethereum/crypto"
	"google.golang.org/protobuf/proto"

	"github.com/Lagrange-Labs/lagrange-node/crypto"
	"github.com/Lagrange-Labs/lagrange-node/logger"
	"github.com/Lagrange-Labs/lagrange-node/rpcclient"
	sequencerv2types "github.com/Lagrange-Labs/lagrange-node/sequencer/types/v2"
	"github.com/Lagrange-Labs/lagrange-node/server"
	serverv2types "github.com/Lagrange-Labs/lagrange-node/server/types/v2"
	"github.com/Lagrange-Labs/lagrange-node/telemetry"
	"github.com/Lagrange-Labs/lagrange-node/utils"
)

const (
	CommitteeCacheSize = 10
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
	serverv2types.NetworkServiceClient
	healthMgr *healthManager
	blsScheme crypto.BLSScheme
	adapter   *rpcAdapter

	blsPrivateKey         []byte
	blsPublicKey          string
	signerECDSAPrivateKey *ecdsa.PrivateKey
	jwToken               string
	stakeAddress          string
	pullInterval          time.Duration

	chErr chan error
}

// NewClient creates a new client.
func NewClient(cfg *Config, rpcCfg *rpcclient.Config) (*Client, error) {
	if len(cfg.BLSKeystorePasswordPath) > 0 {
		var err error
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

	adapter, err := newRpcAdapter(rpcCfg, cfg, pubkey)
	if err != nil {
		return nil, fmt.Errorf("failed to create the rpc adapter: %v", err)
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

	healthMgr, err := newHealthManager(cfg.GrpcURLs)
	if err != nil {
		logger.Fatalf("failed to create the health manager: %v", err)
	}
	healthClient, err := healthMgr.getHealthClient()
	if err != nil {
		logger.Fatalf("failed to get the health client: %v", err)
	}

	c := &Client{
		NetworkServiceClient:  healthClient,
		healthMgr:             healthMgr,
		adapter:               adapter,
		blsScheme:             blsScheme,
		blsPrivateKey:         blsPriv,
		blsPublicKey:          utils.Bytes2Hex(pubkey),
		signerECDSAPrivateKey: ecdsaPriv,
		stakeAddress:          cfg.OperatorAddress,
		pullInterval:          time.Duration(cfg.PullInterval),

		chErr: make(chan error, CommitteeCacheSize),
	}
	go adapter.startBatchFetching(c.chErr)

	return c, nil
}

// GetStakeAddress returns the stake address.
func (c *Client) GetStakeAddress() string {
	return c.stakeAddress
}

// GetChainID returns the chain ID.
func (c *Client) GetChainID() uint32 {
	return c.adapter.chainID
}

// Start starts the connection loop.
func (c *Client) Start() error {
	if err := c.TryJoinNetwork(); err != nil {
		return err
	}

	for {
		select {
		case err := <-c.healthMgr.chErr:
			if errors.Is(err, ErrCurrentServerNotServing) {
				c.NetworkServiceClient, err = c.healthMgr.getHealthClient()
				if err != nil {
					return err
				}
				continue
			}
			return err
		case err := <-c.chErr:
			return err
		case <-time.After(c.pullInterval):
			batch, err := c.TryGetBatch()
			if err != nil {
				if errors.Is(err, server.ErrInvalidToken) {
					if err := c.TryJoinNetwork(); err != nil {
						return err
					}
					continue
				}
				if errors.Is(err, ErrBatchNotFound) {
					logger.Warnf("The batch is not found, please check the metrics for the RPC provider. There may be a delay or a performance issue.")
					if err := c.adapter.initBeginBlockNumber(batch.L1BlockNumber()); err != nil {
						logger.Errorf("failed to initialize the begin block number: %v", err)
					}
					continue
				}
				if errors.Is(err, ErrBatchNotReady) {
					logger.Warn("NOTE: The given round is not initialized yet. It may be because the sequencer is waiting for the next batch since it is almost caught up with the current block. Please wait for the next batch.")
					continue
				}

				logger.Errorf("failed to get the current block: %v", err)
				continue
			}

			if err := c.TryCommitBatch(batch); err != nil {
				if errors.Is(err, server.ErrInvalidToken) {
					if err := c.TryJoinNetwork(); err != nil {
						return err
					}
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

// TryJoinNetwork tries to join the server.
func (c *Client) TryJoinNetwork() error {
	for {
		select {
		case err := <-c.healthMgr.chErr:
			if errors.Is(err, ErrCurrentServerNotServing) {
				c.NetworkServiceClient, err = c.healthMgr.getHealthClient()
				if err != nil {
					return err
				}
				continue
			}
			return err
		case err := <-c.chErr:
			return err
		default:
			if err := c.joinNetwork(); err != nil {
				logger.Infof("failed to join the network: %v", err)
				if strings.Contains(err.Error(), server.ErrNotCommitteeMember.Error()) {
					logger.Warn("NOTE: If you just joined the network, please wait for the next committee rotation. If you have been observing this message for a long time, please check if the BLS public key and the operator address are set correctly in the config file or contact the Largrange team.")
				} else if strings.Contains(err.Error(), server.ErrCheckCommitteeMember.Error()) {
					logger.Warn("NOTE: The given round is not initialized yet. It may be because the sequencer is waiting for the next batch since it is almost caught up with the current block. Please wait for the next batch.")
				}
				time.Sleep(5 * time.Second)
				continue
			}
			logger.Info("joined the network with the new token")
			return nil
		}
	}
}

func (c *Client) joinNetwork() error {
	req := &serverv2types.JoinNetworkRequest{
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
	res, err := c.NetworkServiceClient.JoinNetwork(utils.GetContext(), req)
	if err != nil {
		return fmt.Errorf("failed to join the network: %v", err)
	}
	if len(res.Token) == 0 {
		return fmt.Errorf("the token is empty")
	}
	telemetry.MeasureSince(ti, "client", "join_network_request")

	c.jwToken = res.Token

	if err := c.adapter.initBeginBlockNumber(res.PrevL1BlockNumber); err != nil {
		return fmt.Errorf("failed to initialize the begin block number: %v", err)
	}

	return c.verifyPrevBatch(res.PrevL1BlockNumber, res.PrevL2BlockNumber)
}

// verifyPrevBatch verifies the previous batch.
func (c *Client) verifyPrevBatch(l1BlockNumber, l2BlockNumber uint64) error {
	batchHeader, err := c.adapter.getBatchHeader(l1BlockNumber, l2BlockNumber, 0)
	if err != nil {
		return fmt.Errorf("failed to get the previous batch header for L1 block number %d, L2 block number %d: %v", l1BlockNumber, l2BlockNumber, err)
	}

	if batchHeader == nil {
		return fmt.Errorf("the batch header is not found for L1 block number %d, L2 block number %d", l1BlockNumber, l2BlockNumber)
	}

	return nil
}

// TryGetBatch tries to get the batch from the server.
func (c *Client) TryGetBatch() (*sequencerv2types.Batch, error) {
	ti := time.Now()
	res, err := c.GetBatch(utils.GetContext(), &serverv2types.GetBatchRequest{StakeAddress: c.stakeAddress, Token: c.jwToken})
	if err != nil {
		if strings.Contains(err.Error(), server.ErrInvalidToken.Error()) {
			return nil, server.ErrInvalidToken
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
	c.adapter.setOpenL1BlockNumber(batch.L1BlockNumber())
	logger.Infof("get the batch with L1 block number %d with L2 block number from %d to %d", batch.L1BlockNumber(), fromBlockNumber, toBlockNumber)

	// verify the L1 block number
	batchHeader, err := c.adapter.getBatchHeader(batch.L1BlockNumber(), fromBlockNumber, batch.BatchHeader.L1TxIndex)
	if err != nil || batchHeader == nil {
		logger.Errorf("failed to get the batch header for L1BlockNumber %d, L2FromBlockNumber %d, L1TxIndex %d: %v", batch.L1BlockNumber(), fromBlockNumber, batch.BatchHeader.L1TxIndex, err)
		return batch, ErrBatchNotFound
	}
	if batch.L1BlockNumber() != batchHeader.L1BlockNumber {
		return batch, fmt.Errorf("the batch L1 block number %d is not equal to the rpc L1 block number %d", res.Batch.L1BlockNumber(), batchHeader.L1BlockNumber)
	}
	if fromBlockNumber != batchHeader.FromBlockNumber() {
		return batch, fmt.Errorf("the batch from block number %d is not equal to the rpc from block number %d", fromBlockNumber, batchHeader.FromBlockNumber())
	}
	if toBlockNumber != batchHeader.ToBlockNumber() {
		return batch, fmt.Errorf("the batch to block number %d is not equal to the rpc to block number %d", toBlockNumber, batchHeader.ToBlockNumber())
	}

	// verify the committee root
	if err := c.verifyCommitteeRoot(batch); err != nil {
		logger.Warnf("failed to verify the committee root: %v", err)
		return batch, err
	}

	// verify if the batch hash is correct
	batchHash := batch.BatchHeader.Hash()
	bhHash := batchHeader.Hash()
	if !bytes.Equal(batchHash, bhHash) {
		return batch, fmt.Errorf("the batch hash %s is not equal to the batch header hash %s", batchHash, utils.Bytes2Hex(bhHash))
	}

	// verify the proposer signature
	if len(batch.ProposerPubKey) == 0 {
		return batch, fmt.Errorf("the block %d proposer key is empty", batch.BatchNumber())
	}
	blsSigHash := batch.BlsSignature().Hash()
	verified, err := c.blsScheme.VerifySignature(common.FromHex(batch.ProposerPubKey), blsSigHash, common.FromHex(batch.ProposerSignature))
	if err != nil || !verified {
		return batch, fmt.Errorf("failed to verify the proposer signature: %v", err)
	}

	telemetry.SetGauge(float64(batch.BatchNumber()), "client", "current_batch_number")

	return batch, nil
}

func (c *Client) verifyCommitteeRoot(batch *sequencerv2types.Batch) error {
	blockNumber := batch.L1BlockNumber()
	prevBatchL1Number := uint64(0)
	isGenesis := c.adapter.genesisBlockNumber == blockNumber
	// verify the previous batch's next committee root
	if !isGenesis {
		var err error
		prevBatchL1Number, err = c.adapter.getPrevBatchL1Number(batch.L1BlockNumber(), batch.BatchHeader.L1TxIndex)
		if err != nil {
			return fmt.Errorf("failed to get the previous batch L1 number: %v", err)
		}
		if prevBatchL1Number == 0 {
			return ErrBatchNotFound
		}
	}
	prevCommitteeData, err := c.adapter.getCommitteeRoot(prevBatchL1Number)
	if err != nil {
		return fmt.Errorf("failed to get the previous committee root: %v", err)
	}
	if !bytes.Equal(utils.Hex2Bytes(batch.CurrentCommittee()), prevCommitteeData.Root[:]) {
		return fmt.Errorf("the current batch committee root %s is not equal to the previous batch next committee root %s", batch.CurrentCommittee(), utils.Bytes2Hex(prevCommitteeData.Root[:]))
	}

	// verify the current batch's next committee root
	curCommitteeData, err := c.adapter.getCommitteeRoot(blockNumber)
	if err != nil {
		return fmt.Errorf("failed to get the current committee root: %v", err)
	}
	if !bytes.Equal(utils.Hex2Bytes(batch.NextCommittee()), curCommitteeData.Root[:]) {
		return fmt.Errorf("the current batch next committee root %s is not equal to the current committee root %s", batch.NextCommittee(), utils.Bytes2Hex(curCommitteeData.Root[:]))
	}

	return nil
}

// TryCommitBatch tries to commit the signature to the server.
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

	req := &serverv2types.CommitBatchRequest{
		BlsSignature: blsSignature,
		StakeAddress: c.stakeAddress,
		PublicKey:    c.blsPublicKey,
		Token:        c.jwToken,
	}

	ctx, cancel := utils.GetContextWithCancel()
	defer cancel()

	stream, err := c.CommitBatch(ctx, req)
	if err != nil {
		if strings.Contains(err.Error(), server.ErrInvalidToken.Error()) {
			return server.ErrInvalidToken
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
		if c.healthMgr != nil {
			c.healthMgr.cancel()
		}
	}
}

var (
	ErrBatchNotReady = fmt.Errorf("the batch is not ready")
	ErrBatchNotFound = fmt.Errorf("the batch is not found")
)

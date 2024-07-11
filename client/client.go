package client

import (
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
	adapter   *RpcAdapter
	verifier  *Verifier

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

	adapter, chainID, err := newRpcAdapter(rpcCfg, cfg, pubkey)
	if err != nil {
		return nil, fmt.Errorf("failed to create the rpc adapter: %v", err)
	}

	verifier, err := newVerifier(cfg, adapter, chainID)
	if err != nil {
		return nil, fmt.Errorf("failed to create the verifier: %v", err)
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
		verifier:              verifier,
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

	return c.verifier.VerifyPrevBatch(res.PrevL1BlockNumber, res.PrevL2BlockNumber)
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
	c.adapter.setOpenL1BlockNumber(batch.L1BlockNumber())
	logger.Infof("get the batch with L1 block number %d with L2 block number from %d to %d", batch.L1BlockNumber(), batch.BatchHeader.FromBlockNumber(), batch.BatchHeader.ToBlockNumber())

	// verify the batch
	if err := c.verifier.VerifyBatch(batch); err != nil {
		logger.Warnf("failed to verify the batch: %v", err)
		return batch, err
	}

	return batch, nil
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
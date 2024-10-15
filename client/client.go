package client

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"google.golang.org/protobuf/proto"

	"github.com/Lagrange-Labs/lagrange-node/core"
	"github.com/Lagrange-Labs/lagrange-node/core/logger"
	"github.com/Lagrange-Labs/lagrange-node/core/telemetry"
	"github.com/Lagrange-Labs/lagrange-node/rpcclient"
	sequencerv2types "github.com/Lagrange-Labs/lagrange-node/sequencer/types/v2"
	"github.com/Lagrange-Labs/lagrange-node/server"
	serverv2types "github.com/Lagrange-Labs/lagrange-node/server/types/v2"
)

var (
	// ErrBatchNotFinalized is returned when the current batch is not finalized yet.
	ErrBatchNotFinalized = errors.New("the current batch is not finalized yet")

	// NodeVersion is used to check the compatibility between the client and the server.
	NodeVersion = "v1.1.0"
)

// VerifierCaller is the interface to verify the batch.
type VerifierCaller interface {
	VerifyBatch(*sequencerv2types.Batch) error
	VerifyPrevBatch(uint64, uint64) error
}

// SignerCaller is the interface to sign the batch.
type SignerCaller interface {
	Sign(keyType string, msg []byte) ([]byte, error)
	GetPublicKey(keyType string) (string, error)
}

// StatusMessage is the struct to represent the node status message.
type StatusMessage struct {
	NodeStatus serverv2types.ClientNodeStatus
	Message    string
}

// Client is a gRPC client to join the network
type Client struct {
	serverv2types.NetworkServiceClient
	healthMgr *healthManager
	verifier  VerifierCaller
	signer    SignerCaller

	blsPublicKey       string
	jwToken            string
	stakeAddress       string
	pullInterval       time.Duration
	isUploadNodeStatus bool

	chErr        chan error
	chNodeStatus chan StatusMessage
}

// NewClient creates a new client.
func NewClient(cfg *Config, rpcCfg *rpcclient.Config) (*Client, error) {
	signer, err := NewSignerClient(cfg)
	if err != nil {
		return nil, fmt.Errorf("failed to create the signer client: %v", err)
	}

	blsPublicKey, err := signer.GetPublicKey("BLS")
	if err != nil {
		return nil, fmt.Errorf("failed to get the BLS public key: %v", err)
	}

	chNodeStatus := make(chan StatusMessage, 10)
	adapter, chainID, err := newRpcAdapter(rpcCfg, cfg, cfg.BLSKeyAccountID, chNodeStatus)
	if err != nil {
		return nil, fmt.Errorf("failed to create the rpc adapter: %v", err)
	}

	verifier, err := newVerifier(cfg, adapter, chainID)
	if err != nil {
		return nil, fmt.Errorf("failed to create the verifier: %v", err)
	}

	healthMgr, err := newHealthManager(cfg.GrpcURLs)
	if err != nil {
		logger.Fatalf("failed to create the health manager: %v", err)
	}
	healthClient, err := healthMgr.getHealthClient()
	if err != nil {
		logger.Fatalf("failed to get the health client: %v", err)
	}

	return &Client{
		NetworkServiceClient: healthClient,
		signer:               signer,
		healthMgr:            healthMgr,
		verifier:             verifier,
		blsPublicKey:         blsPublicKey,
		stakeAddress:         cfg.OperatorAddress,
		pullInterval:         time.Duration(cfg.PullInterval),
		isUploadNodeStatus:   cfg.IsUploadNodeStatus,

		chErr:        make(chan error, 10),
		chNodeStatus: chNodeStatus,
	}, nil
}

// GetStakeAddress returns the stake address.
func (c *Client) GetStakeAddress() string {
	return c.stakeAddress
}

func (c *Client) uploadStatus() {
	for msg := range c.chNodeStatus {
		if !c.isUploadNodeStatus {
			continue
		}
		req := &serverv2types.UploadStatusRequest{
			Status:       msg.NodeStatus,
			Message:      msg.Message,
			StakeAddress: c.stakeAddress,
			PublicKey:    c.blsPublicKey,
			Token:        c.jwToken,
		}
		_, err := c.UploadStatus(core.GetContext(), req)
		if err != nil {
			c.chErr <- fmt.Errorf("failed to upload the status: %v", err)
		}
	}
}

// Start starts the connection loop.
func (c *Client) Start() error {
	go c.uploadStatus()

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
		NodeVersion:  NodeVersion,
	}
	reqMsg, err := proto.Marshal(req)
	if err != nil {
		return fmt.Errorf("failed to marshal the request: %v", err)
	}
	sig, err := c.signer.Sign("BLS", reqMsg)
	if err != nil {
		c.chNodeStatus <- StatusMessage{NodeStatus: serverv2types.ClientNodeStatus_BLS_KEY_ISSUE, Message: "failed to sign the join network request"}
		return fmt.Errorf("failed to sign the request: %v", err)
	}
	req.Signature = core.Bytes2Hex(sig)
	ti := time.Now()
	res, err := c.NetworkServiceClient.JoinNetwork(core.GetContext(), req)
	if err != nil {
		return fmt.Errorf("failed to join the network: %v", err)
	}
	if len(res.Token) == 0 {
		return fmt.Errorf("the token is empty")
	}
	telemetry.MeasureSince(ti, "client", "join_network_request")

	c.jwToken = res.Token

	return c.verifier.VerifyPrevBatch(res.PrevL1BlockNumber, res.PrevL2BlockNumber)
}

// TryGetBatch tries to get the batch from the server.
func (c *Client) TryGetBatch() (*sequencerv2types.Batch, error) {
	ti := time.Now()
	res, err := c.GetBatch(core.GetContext(), &serverv2types.GetBatchRequest{StakeAddress: c.stakeAddress, Token: c.jwToken})
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
	logger.Infof("get the batch with L1 block number %d with L2 block number from %d to %d", batch.L1BlockNumber(), batch.BatchHeader.FromBlockNumber(), batch.BatchHeader.ToBlockNumber())

	// verify the batch
	if err := c.verifier.VerifyBatch(batch); err != nil {
		logger.Warnf("failed to verify the batch: %v", err)
		return batch, err
	}
	telemetry.SetGauge(float64(batch.BatchNumber()), "client", "current_batch_number")
	return batch, nil
}

// TryCommitBatch tries to commit the signature to the server.
func (c *Client) TryCommitBatch(batch *sequencerv2types.Batch) error {
	ti := time.Now()
	defer telemetry.MeasureSince(ti, "client", "try_commit_batch")

	blsSignature := batch.BlsSignature()
	blsSig, err := c.signer.Sign("BLS", blsSignature.Hash())
	if err != nil {
		c.chNodeStatus <- StatusMessage{NodeStatus: serverv2types.ClientNodeStatus_BLS_KEY_ISSUE, Message: "failed to sign the BLS signature"}
		return fmt.Errorf("failed to sign the BLS signature: %v", err)
	}
	blsSignature.BlsSignature = core.Bytes2Hex(blsSig)

	// generate the ECDSA signature
	msg := blsSignature.CommitHash()
	sig, err := c.signer.Sign("ECDSA", msg)
	if err != nil {
		c.chNodeStatus <- StatusMessage{NodeStatus: serverv2types.ClientNodeStatus_SIGNER_KEY_ISSUE, Message: "failed to ecdsa sign the batch"}
		return fmt.Errorf("failed to ecdsa sign the batch: %v", err)
	}
	blsSignature.EcdsaSignature = core.Bytes2Hex(sig)

	req := &serverv2types.CommitBatchRequest{
		BlsSignature: blsSignature,
		StakeAddress: c.stakeAddress,
		PublicKey:    c.blsPublicKey,
		Token:        c.jwToken,
	}

	ctx, cancel := core.GetContextWithCancel()
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
		c.chNodeStatus <- StatusMessage{NodeStatus: serverv2types.ClientNodeStatus_SERVER_ISSUE, Message: "failed to get the response from the stream"}
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

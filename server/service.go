package server

import (
	context "context"
	"errors"
	"fmt"
	"strings"
	"time"

	"google.golang.org/grpc/peer"
	"google.golang.org/protobuf/proto"

	"github.com/Lagrange-Labs/lagrange-node/core"
	"github.com/Lagrange-Labs/lagrange-node/core/crypto"
	"github.com/Lagrange-Labs/lagrange-node/core/logger"
	"github.com/Lagrange-Labs/lagrange-node/core/telemetry"
	"github.com/Lagrange-Labs/lagrange-node/server/types"
	v2types "github.com/Lagrange-Labs/lagrange-node/server/types/v2"
)

var (
	// ErrWrongBlockNumber is returned when the block number is not matched.
	ErrWrongBlockNumber = errors.New("the block number is not matched")
	// ErrInvalidToken is returned when the token is invalid.
	ErrInvalidToken = errors.New("the token is invalid")
	// ErrNotCommitteeMember is returned when the operator is not a committee member.
	ErrNotCommitteeMember = errors.New("the given operator is not a member of the current committee")
	// ErrCheckCommitteeMember is returned when the check committee member failed.
	ErrCheckCommitteeMember = errors.New("failed to check the committee member")

	// MinCompatibleVersion is the minimum compatible version.
	MinCompatibleVersion = core.Version{Major: 1, Minor: 1, Patch: 0}
	// ExpectedVersion is the expected version.
	ExpectedVersion = core.Version{Major: 1, Minor: 1, Patch: 2}
)

type sequencerService struct {
	storage   storageInterface
	consensus consensusInterface
	v2types.UnimplementedNetworkServiceServer

	blsScheme crypto.BLSScheme
	chainID   uint32
}

// NewSequencerService creates the sequencer service.
func NewSequencerService(storage storageInterface, consensus consensusInterface, blsScheme crypto.BLSScheme, chainID uint32) (v2types.NetworkServiceServer, error) {
	return &sequencerService{
		storage:   storage,
		consensus: consensus,
		blsScheme: blsScheme,
		chainID:   chainID,
	}, nil
}

// JoinNetwork is a method to join the attestation network.
func (s *sequencerService) JoinNetwork(ctx context.Context, req *v2types.JoinNetworkRequest) (*v2types.JoinNetworkResponse, error) {
	// Check if the consensus is initialized
	s.consensus.Start()

	logger.Infof("JoinNetwork request: %+v", req)
	ti := time.Now()
	defer telemetry.MeasureSince(ti, "server", "join_network")

	// Verify the node version
	if len(req.NodeVersion) == 0 {
		logger.Warnf("The node version is empty, expected version: %s", ExpectedVersion.String())
	} else {
		nv, err := core.GetVersion(req.NodeVersion)
		if err != nil {
			return nil, fmt.Errorf("failed to parse the node version: %v", err)
		}
		if MinCompatibleVersion.Compare(nv) > 0 {
			logger.Warnf("The node version is not compatible, the minimum compatible version: %s", MinCompatibleVersion.String())
			return nil, fmt.Errorf("the node version is not compatible, the minimum compatible version: %s", MinCompatibleVersion.String())
		}
		if ExpectedVersion.Compare(nv) > 0 {
			logger.Warnf("The node version is not expected, the expected version: %s", ExpectedVersion)
		}
	}

	// Verify signature
	sigMessage := req.Signature
	req.Signature = ""
	msg, err := proto.Marshal(req)
	if err != nil {
		return nil, err
	}
	verified, err := s.blsScheme.VerifySignature(core.Hex2Bytes(req.PublicKey), msg, core.Hex2Bytes(sigMessage), true)
	if err != nil || !verified {
		logger.Warnf("BLS signature verification failed: %v", err)
		return nil, fmt.Errorf("BLS signature verification failed: %v", err)
	}
	// Check if the operator is a committee member
	isMember, err := s.consensus.CheckCommitteeMember(core.GetValidAddress(req.StakeAddress), strings.TrimPrefix(req.PublicKey, "0x"))
	if err != nil {
		return nil, errors.Join(ErrCheckCommitteeMember, err)
	}
	if !isMember {
		logger.Warnf("The operator %s is not a committee member", req.StakeAddress)
		return &v2types.JoinNetworkResponse{}, ErrNotCommitteeMember
	}
	// Register node
	ip, err := getIPAddress(ctx)
	if err != nil {
		return nil, err
	}
	if err := s.storage.AddNode(ctx,
		&types.ClientNode{
			StakeAddress: core.GetValidAddress(req.StakeAddress),
			PublicKey:    req.PublicKey,
			IPAddress:    ip,
			ChainID:      s.chainID,
			JoinedAt:     time.Now().UnixMilli(),
			Status:       types.NodeJoined,
		}); err != nil {
		return nil, err
	}

	token, err := GenerateToken(req.StakeAddress)
	if err != nil {
		return nil, err
	}

	logger.Infof("New node %v joined the network", req.StakeAddress)
	prevBatch := s.consensus.GetPrevBatch()
	return &v2types.JoinNetworkResponse{
		Token:             token,
		PrevL2BlockNumber: prevBatch.BatchHeader.FromBlockNumber(),
		PrevL1BlockNumber: prevBatch.L1BlockNumber(),
	}, nil
}

// GetBatch is a method to get the proposed batch.
func (s *sequencerService) GetBatch(ctx context.Context, req *v2types.GetBatchRequest) (*v2types.GetBatchResponse, error) {
	ti := time.Now()
	defer telemetry.MeasureSince(ti, "server", "get_batch")

	valid, err := ValidateToken(req.Token, req.StakeAddress)
	if err != nil || !valid {
		logger.Warnf("Failed to validate the token: %v", err)
		return nil, ErrInvalidToken
	}

	logger.Infof("GetBatch request from %v, %d", req.StakeAddress, req.BatchNumber)

	return &v2types.GetBatchResponse{
		Batch: s.consensus.GetOpenBatch(),
	}, nil
}

// CommitBatch is a method to commit the proposed batch.
func (s *sequencerService) CommitBatch(req *v2types.CommitBatchRequest, stream v2types.NetworkService_CommitBatchServer) error {
	ti := time.Now()
	defer telemetry.MeasureSince(ti, "server", "commit_batch")

	valid, err := ValidateToken(req.Token, req.StakeAddress)
	if err != nil || !valid {
		logger.Warnf("Failed to validate the token: %v", err)
		return ErrInvalidToken
	}

	signature := req.BlsSignature
	batchNumber := signature.BatchNumber()
	logger.Infof("CommitBatch request from %v, %d", req.StakeAddress, batchNumber)

	// verify the peer signature
	reqHash := signature.CommitHash()
	isVerified, addr, err := crypto.VerifyECDSASignature(reqHash, core.Hex2Bytes(signature.EcdsaSignature))
	if err != nil || !isVerified {
		logger.Errorf("failed to verify the ECDSA signature: %v, %v", err, isVerified)
		return fmt.Errorf("failed to verify the ECDSA signature: %v, %v", err, isVerified)
	}

	if !s.consensus.CheckSignAddress(core.GetValidAddress(req.StakeAddress), addr.Hex()) {
		logger.Errorf("the sign address is not matched in ECDSA signature: %v, %v", addr, req.StakeAddress)
		return fmt.Errorf("the sign address is not matched in ECDSA signature: %v, %v", addr, req.StakeAddress)
	}

	// upload the commit to the consensus layer
	if err := s.consensus.AddBatchCommit(signature, core.GetValidAddress(req.StakeAddress), strings.TrimPrefix(req.PublicKey, "0x")); err != nil {
		logger.Errorf("failed to add the commit to the consensus layer: %v", err)
		return err
	}

	timeoutCtx, cancel := context.WithTimeout(stream.Context(), s.consensus.GetRoundInterval())
	defer cancel()

	for {
		select {
		case <-timeoutCtx.Done():
			logger.Warnf("Failed to commit the batch: %v err: %v", req.StakeAddress, timeoutCtx.Err())
			return stream.Send(&v2types.CommitBatchResponse{
				Result: false,
			})
		default:
			if s.consensus.IsFinalized(batchNumber) {
				return stream.Send(&v2types.CommitBatchResponse{
					Result: true,
				})
			}
			time.Sleep(200 * time.Millisecond)
		}
	}
}

func getIPAddress(ctx context.Context) (string, error) {
	// Get the client IP address from the gRPC StreamInfo
	pr, ok := peer.FromContext(ctx)
	if !ok {
		return "", fmt.Errorf("failed to get peer from context")
	}

	return strings.Split(pr.Addr.String(), ":")[0], nil
}

package network

import (
	context "context"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"google.golang.org/grpc/peer"
	"google.golang.org/protobuf/proto"

	"github.com/Lagrange-Labs/lagrange-node/crypto"
	"github.com/Lagrange-Labs/lagrange-node/logger"
	"github.com/Lagrange-Labs/lagrange-node/network/types"
	networkv2types "github.com/Lagrange-Labs/lagrange-node/network/types/v2"
	"github.com/Lagrange-Labs/lagrange-node/utils"
)

var (
	// ErrWrongBlockNumber is returned when the block number is not matched.
	ErrWrongBlockNumber = errors.New("the block number is not matched")
	// ErrInvalidToken is returned when the token is invalid.
	ErrInvalidToken = errors.New("the token is invalid")
)

type sequencerService struct {
	storage   storageInterface
	consensus consensusInterface
	networkv2types.UnimplementedNetworkServiceServer

	blsScheme crypto.BLSScheme
	chainID   uint32
}

// NewSequencerService creates the sequencer service.
func NewSequencerService(storage storageInterface, consensus consensusInterface, blsScheme crypto.BLSScheme, chainID uint32) (networkv2types.NetworkServiceServer, error) {
	return &sequencerService{
		storage:   storage,
		consensus: consensus,
		blsScheme: blsScheme,
		chainID:   chainID,
	}, nil
}

// JoinNetwork is a method to join the attestation network.
func (s *sequencerService) JoinNetwork(ctx context.Context, req *networkv2types.JoinNetworkRequest) (*networkv2types.JoinNetworkResponse, error) {
	logger.Infof("JoinNetwork request: %+v\n", req)

	// Verify signature
	sigMessage := req.Signature
	req.Signature = ""
	msg, err := proto.Marshal(req)
	if err != nil {
		return nil, err
	}
	verified, err := s.blsScheme.VerifySignature(utils.Hex2Bytes(req.PublicKey), msg, utils.Hex2Bytes(sigMessage))
	if err != nil || !verified {
		logger.Warnf("BLS signature verification failed: %v", err)
		return nil, fmt.Errorf("BLS signature verification failed: %v", err)
	}
	// Check if the operator is a committee member
	isMember, err := s.consensus.CheckCommitteeMember(req.StakeAddress, req.PublicKey)
	if err != nil {
		return nil, err
	}
	if !isMember {
		logger.Warnf("The operator is not a committee member")
		return &networkv2types.JoinNetworkResponse{}, fmt.Errorf("the operator is not a committee member")
	}
	// Register node
	ip, err := getIPAddress(ctx)
	if err != nil {
		return nil, err
	}
	if err := s.storage.AddNode(ctx,
		&types.ClientNode{
			StakeAddress: req.StakeAddress,
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

	logger.Infof("New node %v joined the network\n", req.StakeAddress)
	prevBatch := s.consensus.GetPrevBatch()
	return &networkv2types.JoinNetworkResponse{
		Token:             token,
		PrevL2BlockNumber: prevBatch.BatchHeader.FromBlockNumber(),
		PrevL1BlockNumber: prevBatch.L1BlockNumber(),
	}, nil
}

// GetBatch is a method to get the proposed batch.
func (s *sequencerService) GetBatch(ctx context.Context, req *networkv2types.GetBatchRequest) (*networkv2types.GetBatchResponse, error) {
	valid, err := ValidateToken(req.Token, req.StakeAddress)
	if err != nil || !valid {
		logger.Warnf("Failed to validate the token: %v", err)
		return nil, ErrInvalidToken
	}

	logger.Infof("GetBatch request from %v, %d", req.StakeAddress, req.BatchNumber)

	return &networkv2types.GetBatchResponse{
		Batch: s.consensus.GetOpenBatch(),
	}, nil
}

// CommitBatch is a method to commit the proposed batch.
func (s *sequencerService) CommitBatch(req *networkv2types.CommitBatchRequest, stream networkv2types.NetworkService_CommitBatchServer) error {
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
	isVerified, addr, err := utils.VerifyECDSASignature(reqHash, common.FromHex(signature.EcdsaSignature))
	if err != nil || !isVerified {
		logger.Errorf("failed to verify the ECDSA signature: %v, %v", err, isVerified)
		return fmt.Errorf("failed to verify the ECDSA signature: %v, %v", err, isVerified)
	}
	if !s.consensus.CheckSignAddress(req.StakeAddress, addr.Hex()) {
		logger.Errorf("the sign address is not matched in ECDSA signature: %v, %v", addr, req.StakeAddress)
		return fmt.Errorf("the sign address is not matched in ECDSA signature: %v, %v", addr, req.StakeAddress)
	}

	// upload the commit to the consensus layer
	if err := s.consensus.AddBatchCommit(signature, req.StakeAddress, req.PublicKey); err != nil {
		logger.Errorf("failed to add the commit to the consensus layer: %v", err)
		return err
	}

	timeoutCtx, cancel := context.WithTimeout(stream.Context(), s.consensus.GetRoundInterval())
	defer cancel()

	for {
		select {
		case <-timeoutCtx.Done():
			logger.Warnf("Failed to commit the batch: %v err: %v", req.StakeAddress, timeoutCtx.Err())
			return fmt.Errorf("failed to commit the batch: %v", timeoutCtx.Err())
		default:
			if s.consensus.IsFinalized(batchNumber) {
				return stream.Send(&networkv2types.CommitBatchResponse{
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

package network

import (
	context "context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"google.golang.org/grpc/peer"
	"google.golang.org/protobuf/proto"

	"github.com/Lagrange-Labs/lagrange-node/logger"
	"github.com/Lagrange-Labs/lagrange-node/network/types"
	"github.com/Lagrange-Labs/lagrange-node/utils"
)

var (
	// ErrWrongBlockNumber is returned when the block number is not matched.
	ErrWrongBlockNumber = fmt.Errorf("the block number is not matched")
)

type sequencerService struct {
	storage   storageInterface
	consensus consensusInterface
	types.UnimplementedNetworkServiceServer
}

// NewSequencerService creates the sequencer service.
func NewSequencerService(storage storageInterface, consensus consensusInterface) (types.NetworkServiceServer, error) {
	return &sequencerService{
		storage:   storage,
		consensus: consensus,
	}, nil
}

// JoinNetwork is a method to join the attestation network.
func (s *sequencerService) JoinNetwork(ctx context.Context, req *types.JoinNetworkRequest) (*types.JoinNetworkResponse, error) {
	logger.Infof("JoinNetwork request: %v\n", req)

	// Verify signature
	sigMessage := req.Signature
	req.Signature = ""
	msg, err := proto.Marshal(req)
	if err != nil {
		return nil, err
	}
	verified, err := utils.VerifySignature(common.FromHex(req.PublicKey), msg, common.FromHex(sigMessage))
	if err != nil || !verified {
		return &types.JoinNetworkResponse{
			Result:  false,
			Message: fmt.Sprintf("Signature verification failed: %v", err),
		}, nil
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
		}); err != nil {
		return nil, err
	}

	logger.Infof("New node %v joined the network\n", req)

	return &types.JoinNetworkResponse{
		Result:  true,
		Message: "Joined successfully",
	}, nil
}

// GetBlock is a method to get the last block with a proof.
func (s *sequencerService) GetBlock(ctx context.Context, req *types.GetBlockRequest) (*types.GetBlockResponse, error) {
	logger.Infof("GetBlock request: %v\n", req)

	// verify the registered node
	ip, err := getIPAddress(ctx)
	if err != nil {
		return nil, err
	}
	_, err = s.storage.GetNode(ctx, ip)
	if err != nil {
		return nil, err
	}

	block := s.consensus.GetCurrentBlock()
	if block == nil || block.BlockNumber() != req.BlockNumber {
		block, err := s.storage.GetBlock(ctx, req.BlockNumber)
		return &types.GetBlockResponse{
			Block: block,
		}, err
	}

	return &types.GetBlockResponse{
		Block: block,
	}, nil
}

// CommitBlock is a method to commit a block.
func (s *sequencerService) CommitBlock(ctx context.Context, req *types.CommitBlockRequest) (*types.CommitBlockResponse, error) {
	logger.Infof("CommitBlock request: %v\n", req)

	// verify the peer signature
	signature := req.Signature
	req.Signature = ""
	reqMsg, err := proto.Marshal(req)
	if err != nil {
		logger.Errorf("Failed to marshal the request: %v", err)
		return nil, err
	}

	isVerified, err := utils.VerifySignature(common.FromHex(req.PubKey), reqMsg, common.FromHex(signature))
	if err != nil || !isVerified {
		return &types.CommitBlockResponse{
			Result:  false,
			Message: fmt.Sprintf("Failed to verify the signature: %v", err),
		}, nil
	}

	// check if the block number is matched
	blockNumber := s.consensus.GetCurrentBlockNumber()
	if blockNumber != req.BlsSignature.BlockNumber() {
		return &types.CommitBlockResponse{
			Result:  false,
			Message: fmt.Sprintf("The block number is not matched: %v", blockNumber),
		}, nil
	}

	// upload the commit to the consensus layer
	s.consensus.AddCommit(req)

	return &types.CommitBlockResponse{
		Result:  true,
		Message: "Uploaded successfully",
	}, nil
}

func getIPAddress(ctx context.Context) (string, error) {
	// Get the client IP address from the gRPC StreamInfo
	pr, ok := peer.FromContext(ctx)
	if !ok {
		return "", fmt.Errorf("failed to get peer from context")
	}

	return pr.Addr.String(), nil
}

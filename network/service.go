package network

import (
	context "context"
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"google.golang.org/grpc/peer"
	"google.golang.org/protobuf/proto"

	contypes "github.com/Lagrange-Labs/lagrange-node/consensus/types"
	"github.com/Lagrange-Labs/lagrange-node/logger"
	"github.com/Lagrange-Labs/lagrange-node/network/types"
	storetypes "github.com/Lagrange-Labs/lagrange-node/store/types"
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

	chainID uint32
}

// NewSequencerService creates the sequencer service.
func NewSequencerService(storage storageInterface, consensus consensusInterface, chainID uint32) (types.NetworkServiceServer, error) {
	return &sequencerService{
		storage:   storage,
		consensus: consensus,
		chainID:   chainID,
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
			ChainID:      s.chainID,
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
	// verify the registered node
	ip, err := getIPAddress(ctx)
	if err != nil {
		return nil, err
	}
	node, err := s.storage.GetNodeByStakeAddr(ctx, req.StakeAddress)
	if err != nil {
		return nil, err
	}
	if node.IPAddress != ip {
		logger.Warnf("The IP address is not matched: %v, %v\n", node.IPAddress, ip)
	}

	sBlock, err := s.storage.GetBlock(ctx, s.chainID, req.BlockNumber)
	if err == storetypes.ErrBlockNotFound {
		err = nil
	}

	return &types.GetBlockResponse{
		Block: sBlock,
	}, err
}

// GetCurrentBlock is a method to get the current proposed block.
func (s *sequencerService) GetCurrentBlock(req *types.GetBlockRequest, stream types.NetworkService_GetCurrentBlockServer) error {
	// verify the registered node
	ip, err := getIPAddress(stream.Context())
	if err != nil {
		return err
	}
	node, err := s.storage.GetNodeByStakeAddr(context.Background(), req.StakeAddress)
	if err != nil {
		return err
	}
	if node.IPAddress != ip {
		logger.Warnf("The IP address is not matched: %v, %v\n", node.IPAddress, ip)
	}
	if node.Status != types.NodeRegistered {
		return fmt.Errorf("node is not registered: %v", node)
	}

	block := s.consensus.GetCurrentBlock()
	if block != nil && req.BlockNumber <= block.BlockNumber() {
		return stream.Send(&types.GetBlockResponse{Block: block})
	}

	for {
		select {
		case <-stream.Context().Done():
			return stream.Context().Err()
		case block := <-s.consensus.GetNextBlock():
			return stream.Send(&types.GetBlockResponse{Block: block})
		}
	}
}

// CommitBlock is a method to commit a block.
func (s *sequencerService) CommitBlock(ctx context.Context, req *types.CommitBlockRequest) (*types.CommitBlockResponse, error) {
	// verify the peer signature
	logger.Infof("CommitBlock request: %v\n", req)
	signature := req.Signature
	reqHash := contypes.GetCommitRequestHash(req)
	isVerified, addr, err := utils.VerifyECDSASignature(reqHash, common.FromHex(signature))
	if err != nil || !isVerified {
		return &types.CommitBlockResponse{
			Result:  false,
			Message: fmt.Sprintf("Failed to verify the signature: %v", err),
		}, nil
	}

	// verify the registered node
	ip, err := getIPAddress(ctx)
	if err != nil {
		return nil, err
	}
	node, err := s.storage.GetNodeByStakeAddr(ctx, addr.Hex())
	if err != nil {
		return nil, err
	}
	if node.IPAddress != ip {
		logger.Warnf("The IP address is not matched: %v, %v\n", node.IPAddress, ip)
	}
	if node.Status != types.NodeRegistered {
		return nil, fmt.Errorf("the node is not registered: %v", node)
	}

	// check if the block number is matched
	blockNumber := s.consensus.GetCurrentBlockNumber()
	if blockNumber != req.BlsSignature.BlockNumber() {
		return &types.CommitBlockResponse{
			Result:  false,
			Message: fmt.Sprintf("The block number is not matched: %v", blockNumber),
		}, nil
	}

	// check if the epoch number is matched
	epochNumber := s.consensus.GetCurrentEpochBlockNumber()
	if epochNumber != req.EpochBlockNumber {
		return &types.CommitBlockResponse{
			Result:  false,
			Message: fmt.Sprintf("The epoch number is not matched: %v", epochNumber),
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

	return strings.Split(pr.Addr.String(), ":")[0], nil
}

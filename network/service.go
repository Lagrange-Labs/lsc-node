package network

import (
	context "context"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"google.golang.org/grpc/peer"
	"google.golang.org/protobuf/proto"

	contypes "github.com/Lagrange-Labs/lagrange-node/consensus/types"
	"github.com/Lagrange-Labs/lagrange-node/logger"
	"github.com/Lagrange-Labs/lagrange-node/network/types"
	sequencertypes "github.com/Lagrange-Labs/lagrange-node/sequencer/types"
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

// GetBlock is a method to get the block.
func (s *sequencerService) GetBlock(ctx context.Context, req *types.GetBlockRequest) (*types.GetBlockResponse, error) {
	block, err := s.storage.GetBlock(ctx, s.chainID, req.BlockNumber)
	if err != nil {
		if err == storetypes.ErrBlockNotFound {
			return &types.GetBlockResponse{
				Block: nil,
			}, nil
		}
		return nil, err
	}

	return &types.GetBlockResponse{
		Block: block,
	}, nil
}

// GetBatch is a method to get the proposed batch.
func (s *sequencerService) GetBatch(ctx context.Context, req *types.GetBatchRequest) (*types.GetBatchResponse, error) {
	// verify the registered node
	ip, err := getIPAddress(ctx)
	if err != nil {
		return nil, err
	}
	node, err := s.storage.GetNodeByStakeAddr(ctx, req.StakeAddress, s.chainID)
	if err != nil {
		return nil, err
	}
	if node.IPAddress != ip {
		logger.Warnf("The IP address is not matched: %v, %v\n", node.IPAddress, ip)
	}

	blocks := s.consensus.GetOpenRoundBlocks(req.BlockNumber)

	return &types.GetBatchResponse{
		Batch: blocks,
	}, nil
}

// CommitBatch is a method to commit the proposed batch.
func (s *sequencerService) CommitBatch(req *types.CommitBatchRequest, stream types.NetworkService_CommitBatchServer) error {
	logger.Infof("CommitBatch request from %v\n", req.StakeAddress)
	// verify the registered node
	ip, err := getIPAddress(stream.Context())
	if err != nil {
		return err
	}
	node, err := s.storage.GetNodeByStakeAddr(context.Background(), req.StakeAddress, s.chainID)
	if err != nil {
		return err
	}
	if node.IPAddress != ip {
		logger.Warnf("The IP address is not matched: %v, %v\n", node.IPAddress, ip)
	}
	if node.Status != types.NodeRegistered {
		return fmt.Errorf("the node is not registered: %v", node.Status)
	}

	wg := sync.WaitGroup{}
	wg.Add(len(req.BlsSignatures))
	chError := make(chan error, len(req.BlsSignatures))
	lastBlockNumber := uint64(0)

	for _, signature := range req.BlsSignatures {
		if signature.BlockNumber() > lastBlockNumber {
			lastBlockNumber = signature.BlockNumber()
		}
		go func(signature *sequencertypes.BlsSignature) {
			defer wg.Done()
			// verify the peer signature
			reqHash := contypes.GetCommitRequestHash(signature)
			isVerified, addr, err := utils.VerifyECDSASignature(reqHash, common.FromHex(signature.EcdsaSignature))
			if err != nil || !isVerified {
				chError <- fmt.Errorf("failed to verify the ECDSA signature: %v, %v", err, isVerified)
				return
			}
			if addr != common.HexToAddress(node.StakeAddress) {
				chError <- fmt.Errorf("the stake address is not matched in ECDSA signature: %v, %v", addr, node.StakeAddress)
				return
			}

			// upload the commit to the consensus layer
			err = s.consensus.AddCommit(signature, node.PublicKey, node.StakeAddress)
			if err != nil {
				chError <- err
			}
		}(signature)
	}

	wg.Wait()
	close(chError)

	for err := range chError {
		return err
	}

	timeoutCtx, cancel := context.WithTimeout(stream.Context(), 5*time.Second)
	defer cancel()

	for {
		select {
		case <-timeoutCtx.Done():
			return fmt.Errorf("failed to commit the batch: %v", timeoutCtx.Err())
		default:
			if s.consensus.IsFinalized(lastBlockNumber) {
				return stream.Send(&types.CommitBatchResponse{
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

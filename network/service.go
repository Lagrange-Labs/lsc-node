package network

import (
	"bytes"
	context "context"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"google.golang.org/grpc/peer"
	"google.golang.org/protobuf/proto"

	contypes "github.com/Lagrange-Labs/lagrange-node/consensus/types"
	"github.com/Lagrange-Labs/lagrange-node/crypto"
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

	blsScheme crypto.BLSScheme
	chainID   uint32
}

// NewSequencerService creates the sequencer service.
func NewSequencerService(storage storageInterface, consensus consensusInterface, blsScheme crypto.BLSScheme, chainID uint32) (types.NetworkServiceServer, error) {
	return &sequencerService{
		storage:   storage,
		consensus: consensus,
		blsScheme: blsScheme,
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
	verified, err := s.blsScheme.VerifySignature(utils.Hex2Bytes(req.PublicKey), msg, utils.Hex2Bytes(sigMessage))
	if err != nil || !verified {
		return &types.JoinNetworkResponse{
			Result:  false,
			Message: fmt.Sprintf("BLS signature verification failed: %v", err),
		}, nil
	}
	// Register node
	ip, err := getIPAddress(ctx)
	if err != nil {
		return nil, err
	}
	// Check if the node is already registered
	node, err := s.storage.GetNodeByStakeAddr(ctx, req.StakeAddress, s.chainID)
	if err != nil {
		if err == storetypes.ErrNodeNotFound {
			if err := s.storage.AddNode(ctx,
				&types.ClientNode{
					StakeAddress: req.StakeAddress,
					PublicKey:    utils.Hex2Bytes(req.PublicKey),
					IPAddress:    ip,
					ChainID:      s.chainID,
				}); err != nil {
				return nil, err
			}
		} else {
			return nil, err
		}
	} else if node != nil {
		if node.Status == types.NodeRegistered {
			if !bytes.Equal(node.PublicKey, utils.Hex2Bytes(req.PublicKey)) {
				logger.Warnf("The public key is not matched: %v, %v", node.PublicKey, utils.Hex2Bytes(req.PublicKey))
				return &types.JoinNetworkResponse{
					Result:  false,
					Message: "The public key is not matched",
				}, nil
			}
		}
		if node.IPAddress != ip {
			node.IPAddress = ip
			if err := s.storage.UpdateNode(ctx, node); err != nil {
				return nil, err
			}
		}
	}

	logger.Infof("New node %v joined the network\n", req)

	return &types.JoinNetworkResponse{
		Result:  true,
		Message: "Joined successfully",
	}, nil
}

// GetBatch is a method to get the proposed batch.
func (s *sequencerService) GetBatch(ctx context.Context, req *types.GetBatchRequest) (*types.GetBatchResponse, error) {
	logger.Infof("GetBatch request from %v, %d", req.StakeAddress, req.BlockNumber)
	// verify the registered node
	ip, err := getIPAddress(ctx)
	if err != nil {
		logger.Warnf("Failed to get IP address: %v", err)
		return nil, err
	}
	node, err := s.storage.GetNodeByStakeAddr(ctx, req.StakeAddress, s.chainID)
	if err != nil {
		logger.Warnf("Failed to get the node: %v err: %v", req.StakeAddress, err)
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
	logger.Infof("CommitBatch request from %v, %d", req.StakeAddress, req.BlsSignatures[0].BlockNumber())
	// verify the registered node
	ip, err := getIPAddress(stream.Context())
	if err != nil {
		logger.Warnf("Failed to get IP address: %v", err)
		return err
	}
	node, err := s.storage.GetNodeByStakeAddr(context.Background(), req.StakeAddress, s.chainID)
	if err != nil {
		logger.Warnf("Failed to get the node: %v err: %v", req.StakeAddress, err)
		return err
	}
	if node.IPAddress != ip {
		logger.Warnf("The IP address is not matched: %v, %v", node.IPAddress, ip)
	}
	if node.Status != types.NodeRegistered {
		logger.Warnf("The node is not registered: %v", node.Status)
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
				logger.Errorf("failed to verify the ECDSA signature: %v, %v", err, isVerified)
				chError <- fmt.Errorf("failed to verify the ECDSA signature: %v, %v", err, isVerified)
				return
			}
			if addr != common.HexToAddress(node.StakeAddress) {
				logger.Errorf("the stake address is not matched in ECDSA signature: %v, %v", addr, node.StakeAddress)
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
		logger.Warnf("Failed to commit the batch: %v err: %v", req.StakeAddress, err)
		return err
	}

	timeoutCtx, cancel := context.WithTimeout(stream.Context(), 5*time.Second)
	defer cancel()

	for {
		select {
		case <-timeoutCtx.Done():
			logger.Warnf("Failed to commit the batch: %v err: %v", req.StakeAddress, timeoutCtx.Err())
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

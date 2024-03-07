package network

import (
	context "context"
	"errors"
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
	pubKey := utils.Hex2Bytes(req.PublicKey)
	verified, err := s.blsScheme.VerifySignature(pubKey, msg, utils.Hex2Bytes(sigMessage))
	if err != nil || !verified {
		return &types.JoinNetworkResponse{
			Message: fmt.Sprintf("BLS signature verification failed: %v", err),
		}, nil
	}
	// Check if the operator is a committee member
	rawPubKey, err := s.blsScheme.ConvertPublicKey(pubKey, false)
	if err != nil {
		return nil, err
	}
	if !s.consensus.CheckCommitteeMember(req.StakeAddress, rawPubKey) {
		return &types.JoinNetworkResponse{
			Message: "The operator is not a committee member",
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
			PublicKey:    utils.Hex2Bytes(req.PublicKey),
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

	logger.Infof("New node %v joined the network\n", req)
	return &types.JoinNetworkResponse{
		Token:   token,
		Message: "Joined successfully",
	}, nil
}

// GetBatch is a method to get the proposed batch.
func (s *sequencerService) GetBatch(ctx context.Context, req *types.GetBatchRequest) (*types.GetBatchResponse, error) {
	valid, err := ValidateToken(req.Token, req.StakeAddress)
	if err != nil || !valid {
		logger.Warnf("Failed to validate the token: %v", err)
		return nil, ErrInvalidToken
	}

	logger.Infof("GetBatch request from %v, %d", req.StakeAddress, req.BlockNumber)
	blocks := s.consensus.GetOpenRoundBlocks(req.BlockNumber)

	return &types.GetBatchResponse{
		Batch: blocks,
	}, nil
}

// GetBlock is a method to get the proposed block.
func (s *sequencerService) GetBlock(ctx context.Context, req *types.GetBlockRequest) (*types.GetBlockResponse, error) {
	valid, err := ValidateToken(req.Token, req.StakeAddress)
	if err != nil || !valid {
		logger.Warnf("Failed to validate the token: %v", err)
		return nil, ErrInvalidToken
	}

	logger.Infof("GetBlock request from %v, %d", req.StakeAddress, req.BlockNumber)
	block, err := s.storage.GetBlock(ctx, s.chainID, req.BlockNumber)
	if err != nil {
		return nil, err
	}

	return &types.GetBlockResponse{
		Block: block,
	}, nil
}

// CommitBatch is a method to commit the proposed batch.
func (s *sequencerService) CommitBatch(req *types.CommitBatchRequest, stream types.NetworkService_CommitBatchServer) error {
	valid, err := ValidateToken(req.Token, req.StakeAddress)
	if err != nil || !valid {
		logger.Warnf("Failed to validate the token: %v", err)
		return ErrInvalidToken
	}

	logger.Infof("CommitBatch request from %v, %d", req.StakeAddress, req.BlsSignatures[0].BlockNumber())

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
			if addr != common.HexToAddress(req.StakeAddress) {
				logger.Errorf("the stake address is not matched in ECDSA signature: %v, %v", addr, req.StakeAddress)
				chError <- fmt.Errorf("the stake address is not matched in ECDSA signature: %v, %v", addr, req.StakeAddress)
				return
			}

			// upload the commit to the consensus layer
			err = s.consensus.AddCommit(signature, req.StakeAddress)
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

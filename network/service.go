package network

import (
	context "context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/umbracle/go-eth-consensus/bls"
	"google.golang.org/grpc/peer"
	"google.golang.org/protobuf/proto"

	"github.com/Lagrange-Labs/lagrange-node/logger"
	"github.com/Lagrange-Labs/lagrange-node/network/types"
	"github.com/Lagrange-Labs/lagrange-node/utils"
)

var (
	// ErrWrongBlockNumber is returned when the block number is not latest.
	ErrWrongBlockNumber = fmt.Errorf("the block number is not latest")
)

type sequencerService struct {
	threshold    uint16
	storage      storageInterface
	commitStatus map[string]bool
	publicKeys   []*bls.PublicKey
	signatures   []*bls.Signature
	types.UnimplementedNetworkServiceServer
}

// NewSequencerService creates the sequencer service.
func NewSequencerService(storage storageInterface) (types.NetworkServiceServer, error) {
	ctx := context.Background()

	count, err := storage.GetNodeCount(ctx)
	if err != nil {
		return nil, err
	}

	return &sequencerService{
		storage:      storage,
		threshold:    count * 2 / 3,
		commitStatus: map[string]bool{},
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
	if err != nil {
		return nil, err
	}
	if !verified {
		return &types.JoinNetworkResponse{
			Result:  false,
			Message: "Signature verification failed",
		}, nil
	}
	// Register node
	ip, err := getIPAddress(ctx)
	if err != nil {
		return nil, err
	}
	if err := s.storage.AddNode(ctx, req.StakeAddress, req.PublicKey, ip); err != nil {
		return nil, err
	}
	count, err := s.storage.GetNodeCount(ctx)
	if err != nil {
		return nil, err
	}
	s.threshold = count * 2 / 3

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

	block, err := s.storage.GetBlock(ctx, req.BlockNumber)
	if err != nil {
		return nil, err
	}

	return &types.GetBlockResponse{
		Block: block,
	}, nil
}

// CommitBlock is a method to commit a block.
func (s *sequencerService) CommitBlock(ctx context.Context, req *types.CommitBlockRequest) (*types.CommitBlockResponse, error) {
	logger.Infof("CommitBlock request: %v\n", req)

	ip, err := getIPAddress(ctx)
	if err != nil {
		return nil, err
	}
	node, err := s.storage.GetNode(ctx, ip)
	if err != nil {
		return nil, err
	}

	block, err := s.storage.GetLastBlock(ctx)
	if err != nil {
		return nil, err
	}

	if block.Header.BlockNumber != req.BlockNumber {
		s.commitStatus[node.StakeAddress] = false
		return &types.CommitBlockResponse{
			Result:  false,
			Message: fmt.Sprintf("The wrong block number: %d", block.Header.BlockNumber),
		}, nil
	}

	// check the commit status
	if s.commitStatus[node.StakeAddress] {
		return &types.CommitBlockResponse{
			Result:  false,
			Message: "Already committed",
		}, nil
	}
	s.commitStatus[node.StakeAddress] = true

	pk := new(bls.PublicKey)
	if err := pk.Deserialize(common.FromHex(node.PublicKey)); err != nil {
		return nil, err
	}
	s.publicKeys = append(s.publicKeys, pk)
	sig := new(bls.Signature)
	if err := sig.Deserialize(common.FromHex(req.Signature)); err != nil {
		return nil, err
	}
	s.signatures = append(s.signatures, sig)

	if len(s.signatures) > int(s.threshold) {
		// TODO next generation of the proof
		msg, err := proto.Marshal(block)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal the proof: %v", err)
		}
		aggSig := bls.AggregateSignatures(s.signatures)
		verified, err := aggSig.FastAggregateVerify(s.publicKeys, msg)
		if err != nil {
			return nil, err
		}
		if !verified {
			// TODO punishing mechanism

			logger.Errorf("The current proof is not verifed %v %v", s.signatures, s.publicKeys)

			return &types.CommitBlockResponse{
				Result:  false,
				Message: "Signature verification failed",
			}, nil
		}
		s.signatures = []*bls.Signature{}
		s.publicKeys = []*bls.PublicKey{}
		s.commitStatus = map[string]bool{}
		aggSigMsg := aggSig.Serialize()
		block.AggSignature = common.Bytes2Hex(aggSigMsg[:])
		// TODO store the block
		logger.Info("The current proof is verifed successfully")
	}

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

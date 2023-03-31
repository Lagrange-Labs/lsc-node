package network

import (
	context "context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/umbracle/go-eth-consensus/bls"
	"google.golang.org/grpc/peer"
	"google.golang.org/protobuf/proto"

	"github.com/Lagrange-Labs/Lagrange-Node/network/types"
)

type sequencerService struct {
	threshold  uint16
	storage    storageInterface
	publicKeys []*bls.PublicKey
	signatures []*bls.Signature
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
		storage:   storage,
		threshold: count * 2 / 3,
	}, nil
}

// JoinNetwork is a method to join the attestation network.
func (s *sequencerService) JoinNetwork(ctx context.Context, req *types.JoinNetworkRequest) (*types.JoinNetworkResponse, error) {
	fmt.Printf("JoinNetwork request: %v\n", req)

	// Verify signature
	sigMessage := req.Signature
	req.Signature = ""
	msg, err := proto.Marshal(req)
	if err != nil {
		return nil, err
	}
	sig := new(bls.Signature)
	pub := new(bls.PublicKey)
	if err := pub.Deserialize(common.FromHex(req.PublicKey)); err != nil {
		return nil, err
	}
	if err := sig.Deserialize(common.FromHex(sigMessage)); err != nil {
		return nil, err
	}
	verified, err := sig.VerifyByte(pub, msg)
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

	fmt.Printf("New node %v joined the network\n", req)

	return &types.JoinNetworkResponse{
		Result:  true,
		Message: "Joined successfully",
	}, nil
}

// GetBlock is a method to get the last block with a proof.
func (s *sequencerService) GetBlock(ctx context.Context, req *types.GetBlockRequest) (*types.GetBlockResponse, error) {
	fmt.Printf("GetBlock request: %v\n", req)

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
	fmt.Printf("CommitBlock request: %v\n", req)

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
		return nil, fmt.Errorf("the proof id is not correct")
	}
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

			fmt.Printf("The current proof is verifed\n")

			return &types.CommitBlockResponse{
				Result:  false,
				Message: "Signature verification failed",
			}, nil
		}
		s.signatures = []*bls.Signature{}
		s.publicKeys = []*bls.PublicKey{}
		aggSigMsg := aggSig.Serialize()
		block.Signature = common.Bytes2Hex(aggSigMsg[:])
		// TODO store the block
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

package network

import (
	context "context"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/umbracle/go-eth-consensus/bls"
	"google.golang.org/grpc/peer"
	"google.golang.org/protobuf/proto"

	"github.com/Lagrange-Labs/Lagrange-Node/network/pb"
)

type sequencerService struct {
	threshold  uint16
	storage    storageInterface
	publicKeys []*bls.PublicKey
	signatures []*bls.Signature
	pb.UnimplementedNetworkServiceServer
}

// NewSequencer creates the sequencer service.
func NewSequencer(storage storageInterface) (pb.NetworkServiceServer, error) {
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
func (s *sequencerService) JoinNetwork(ctx context.Context, req *pb.JoinNetworkRequest) (*pb.JoinNetworkResponse, error) {
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
		return &pb.JoinNetworkResponse{
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

	return &pb.JoinNetworkResponse{
		Result:  true,
		Message: "Joined successfully",
	}, nil
}

// GetLastProof is a method to get the last proof.
func (s *sequencerService) GetLastProof(ctx context.Context, req *pb.GetLastProofRequest) (*pb.GetLastProofResponse, error) {
	fmt.Printf("GetLastProof request: %v\n", req)

	ip, err := getIPAddress(ctx)
	if err != nil {
		return nil, err
	}
	_, err = s.storage.GetNode(ctx, ip)
	if err != nil {
		return nil, err
	}
	proof, err := s.storage.GetLastProof(ctx)
	if err != nil {
		return nil, err
	}

	if proof.ProofId <= req.ProofId {
		return nil, fmt.Errorf("the current proof is not ready yet")
	}

	return &pb.GetLastProofResponse{
		Proof: proof,
	}, nil
}

// UploadSignature is a method to interact with the uploading signature in consensus.
func (s *sequencerService) UploadSignature(ctx context.Context, req *pb.UploadSignatureRequest) (*pb.UploadSignatureResponse, error) {
	fmt.Printf("UploadSignature request: %v\n", req)

	ip, err := getIPAddress(ctx)
	if err != nil {
		return nil, err
	}
	node, err := s.storage.GetNode(ctx, ip)
	if err != nil {
		return nil, err
	}
	s.publicKeys = append(s.publicKeys, node.PublicKey)

	proof, err := s.storage.GetLastProof(ctx)
	if err != nil {
		return nil, err
	}

	if proof.ProofId != req.ProofId {
		return nil, fmt.Errorf("the proof id is not correct")
	}

	sig := new(bls.Signature)
	if err := sig.Deserialize(common.FromHex(req.Signature)); err != nil {
		return nil, err
	}
	s.signatures = append(s.signatures, sig)

	if len(s.signatures) > int(s.threshold) {
		// TODO next generation of the proof
		msg, err := proto.Marshal(proof)
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

			return &pb.UploadSignatureResponse{
				Result:  false,
				Message: "Signature verification failed",
			}, nil
		}
		s.signatures = []*bls.Signature{}
		s.publicKeys = []*bls.PublicKey{}
	}

	return &pb.UploadSignatureResponse{
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

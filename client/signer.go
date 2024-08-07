package client

import (
	"context"
	"errors"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/Lagrange-Labs/lagrange-node/signer"
	"github.com/Lagrange-Labs/lagrange-node/signer/types"
	"github.com/Lagrange-Labs/lagrange-node/utils"
)

var _ SignerCaller = (*SignerClient)(nil)

// Signer is the struct to generate the signature.
type SignerClient struct {
	types.SignerServiceClient

	blsKeyAccountID    string
	signerKeyAccountID string
}

// NewSignerClient creates a new SignerClient.
func NewSignerClient(blsKeyAccountID, signerKeyAccountID, signerServerURL string) (*SignerClient, error) {
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}
	conn, err := grpc.NewClient(signerServerURL, opts...)
	if err != nil {
		return nil, err
	}

	signerClient := types.NewSignerServiceClient(conn)
	return &SignerClient{
		SignerServiceClient: signerClient,
		blsKeyAccountID:     blsKeyAccountID,
		signerKeyAccountID:  signerKeyAccountID,
	}, nil
}

// Sign signs the message with the given key type.
func (sc *SignerClient) Sign(keyType string, msg []byte) ([]byte, error) {
	req := &types.SignRequest{
		Message: utils.Bytes2Hex(msg),
	}
	switch keyType {
	case "BLS":
		req.AccountId = sc.blsKeyAccountID
	case "ECDSA":
		req.AccountId = sc.signerKeyAccountID
	default:
		return nil, errors.New("invalid key type")
	}
	resp, err := sc.SignerServiceClient.Sign(context.Background(), req)
	if err != nil {
		return nil, err
	}
	return utils.Hex2Bytes(resp.Signature), nil
}

// GetPublicKey gets the public key with the given key type.
func (sc *SignerClient) GetPublicKey(keyType string) (string, error) {
	req := &types.SignRequest{
		SignMethod: signer.PublicKeyMethod,
	}
	switch keyType {
	case "BLS":
		req.AccountId = sc.blsKeyAccountID
	case "ECDSA":
		req.AccountId = sc.signerKeyAccountID
	default:
		return "", errors.New("invalid key type")
	}
	resp, err := sc.SignerServiceClient.Sign(context.Background(), req)
	if err != nil {
		return "", err
	}
	return resp.Signature, nil
}

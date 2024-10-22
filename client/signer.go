package client

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"

	ecrypto "github.com/ethereum/go-ethereum/crypto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/Lagrange-Labs/lsc-node/core"
	"github.com/Lagrange-Labs/lsc-node/core/crypto"
	"github.com/Lagrange-Labs/lsc-node/core/logger"
	"github.com/Lagrange-Labs/lsc-node/signer"
	"github.com/Lagrange-Labs/lsc-node/signer/types"
)

var _ SignerCaller = (*SignerClient)(nil)

// Signer is the struct to generate the signature.
type SignerClient struct {
	types.SignerServiceClient

	isRemote              bool
	blsScheme             crypto.BLSScheme
	blsPrivateKey         []byte
	blsPublicKey          string
	signerECDSAPrivateKey *ecdsa.PrivateKey

	blsKeyAccountID    string
	signerKeyAccountID string
}

// NewSignerClient creates a new SignerClient.
func NewSignerClient(cfg *Config) (*SignerClient, error) {
	// Remote Signer
	if len(cfg.SignerServerURL) > 0 {
		opts := []grpc.DialOption{}
		if cfg.TLSConfig != nil {
			creds, err := core.LoadTLS(cfg.TLSConfig, false)
			if err != nil {
				return nil, err
			}
			tlsCredentials := credentials.NewTLS(creds)
			opts = append(opts, grpc.WithTransportCredentials(tlsCredentials))
		} else {
			opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
		}

		conn, err := grpc.NewClient(cfg.SignerServerURL, opts...)
		if err != nil {
			return nil, err
		}

		signerClient := types.NewSignerServiceClient(conn)
		return &SignerClient{
			isRemote:            true,
			SignerServiceClient: signerClient,
			blsKeyAccountID:     cfg.BLSKeyAccountID,
			signerKeyAccountID:  cfg.SignerKeyAccountID,
		}, nil
	}

	if len(cfg.BLSKeystorePasswordPath) > 0 {
		var err error
		cfg.BLSKeystorePassword, err = crypto.ReadKeystorePasswordFromFile(cfg.BLSKeystorePasswordPath)
		if err != nil {
			return nil, fmt.Errorf("failed to read the bls keystore password from %s: %v", cfg.BLSKeystorePasswordPath, err)
		}
	}
	blsPriv, err := crypto.LoadPrivateKey(crypto.CryptoCurve(cfg.BLSCurve), cfg.BLSKeystorePassword, cfg.BLSKeystorePath)
	if err != nil {
		return nil, fmt.Errorf("failed to load the bls keystore from %s: %v", cfg.BLSKeystorePath, err)
	}
	blsScheme := crypto.NewBLSScheme(crypto.BLSCurve(cfg.BLSCurve))
	pubkey, err := blsScheme.GetPublicKey(blsPriv, false, true)
	if err != nil {
		return nil, fmt.Errorf("failed to get the bls public key: %v", err)
	}
	if len(cfg.SignerECDSAKeystorePasswordPath) > 0 {
		cfg.SignerECDSAKeystorePassword, err = crypto.ReadKeystorePasswordFromFile(cfg.SignerECDSAKeystorePasswordPath)
		if err != nil {
			return nil, fmt.Errorf("failed to read the ecdsa keystore password from %s: %v", cfg.SignerECDSAKeystorePasswordPath, err)
		}
	}
	ecdsaPrivKey, err := crypto.LoadPrivateKey(crypto.CryptoCurve("ECDSA"), cfg.SignerECDSAKeystorePassword, cfg.SignerECDSAKeystorePath)
	if err != nil {
		logger.Fatalf("failed to load the ecdsa keystore from %s: %v", cfg.SignerECDSAKeystorePath, err)
	}
	ecdsaPriv, err := ecrypto.ToECDSA(ecdsaPrivKey)
	if err != nil {
		logger.Fatalf("failed to get the ecdsa private key: %v", err)
	}

	return &SignerClient{
		isRemote:              false,
		blsScheme:             blsScheme,
		blsPrivateKey:         blsPriv,
		blsPublicKey:          core.Bytes2Hex(pubkey),
		signerECDSAPrivateKey: ecdsaPriv,
	}, nil
}

// Sign signs the message with the given key type.
func (sc *SignerClient) Sign(keyType string, msg []byte) ([]byte, error) {
	if !sc.isRemote {
		switch keyType {
		case "BLS":
			return sc.blsScheme.Sign(sc.blsPrivateKey, msg, true)
		case "ECDSA":
			return ecrypto.Sign(msg, sc.signerECDSAPrivateKey)
		default:
			return nil, errors.New("invalid key type")
		}
	}

	req := &types.SignRequest{
		Message: core.Bytes2Hex(msg),
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
	return core.Hex2Bytes(resp.Signature), nil

}

// GetPublicKey gets the public key with the given key type.
func (sc *SignerClient) GetPublicKey(keyType string) (string, error) {
	if !sc.isRemote {
		switch keyType {
		case "BLS":
			return sc.blsPublicKey, nil
		case "ECDSA":
			return core.Bytes2Hex(ecrypto.FromECDSAPub(&sc.signerECDSAPrivateKey.PublicKey)), nil
		default:
			return "", errors.New("invalid key type")
		}
	}

	req := &types.SignRequest{
		SignMethod: signer.PublicKeyMethodG1,
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

package local

import (
	"crypto/ecdsa"
	"errors"

	ecrypto "github.com/ethereum/go-ethereum/crypto"

	"github.com/Lagrange-Labs/lsc-node/core/crypto"
	"github.com/Lagrange-Labs/lsc-node/signer"
)

var _ signer.Signer = (*provider)(nil)

type provider struct {
	keyType         string
	privateKey      []byte
	ecdsaPrivateKey *ecdsa.PrivateKey
	scheme          crypto.BLSScheme
}

// NewProvider creates a new local provider.
func NewProvider(cfg *signer.LocalConfig) (*provider, error) {
	password, err := crypto.ReadKeystorePasswordFromFile(cfg.PasswordKeyPath)
	if err != nil {
		return nil, err
	}
	privateKey, err := crypto.LoadPrivateKey(crypto.CryptoCurve(cfg.KeyType), password, cfg.PrivateKeyPath)
	if err != nil {
		return nil, err
	}

	p := &provider{
		keyType:    cfg.KeyType,
		privateKey: privateKey,
	}

	switch cfg.KeyType {
	case string(crypto.BLS12381):
		p.scheme = crypto.NewBLSScheme(crypto.BLS12381)
	case string(crypto.BN254):
		p.scheme = crypto.NewBLSScheme(crypto.BN254)
	case "ECDSA":
		p.ecdsaPrivateKey, err = ecrypto.ToECDSA(privateKey)
	default:
		return nil, errors.New("invalid curve")
	}

	return p, err
}

// Sign signs the message.
func (p *provider) Sign(msg []byte, isG1 bool) ([]byte, error) {
	switch p.keyType {
	case string(crypto.BLS12381), string(crypto.BN254):
		return p.scheme.Sign(p.privateKey, msg, isG1)
	case "ECDSA":
		return ecrypto.Sign(msg, p.ecdsaPrivateKey)
	default:
		return nil, errors.New("invalid curve")
	}
}

// GetPubKey gets the public key.
func (p *provider) GetPubKey(isG1 bool) ([]byte, error) {
	switch p.keyType {
	case string(crypto.BLS12381), string(crypto.BN254):
		return p.scheme.GetPublicKey(p.privateKey, false, isG1)
	case "ECDSA":
		addr := ecrypto.PubkeyToAddress(p.ecdsaPrivateKey.PublicKey)
		return addr.Bytes(), nil
	default:
		return nil, errors.New("invalid curve")
	}
}

// Verify verifies the signature.
func (p *provider) Verify(pubKey, msg, sig []byte, isG1 bool) (bool, error) {
	switch p.keyType {
	case string(crypto.BLS12381), string(crypto.BN254):
		return p.scheme.VerifySignature(pubKey, msg, sig, isG1)
	case "ECDSA":
		res, _, err := crypto.VerifyECDSASignature(msg, sig)
		return res, err
	default:
		return false, errors.New("invalid curve")
	}
}

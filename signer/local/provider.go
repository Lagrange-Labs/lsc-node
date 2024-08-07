package local

import (
	"crypto/ecdsa"
	"errors"

	ecrypto "github.com/ethereum/go-ethereum/crypto"

	"github.com/Lagrange-Labs/lagrange-node/crypto"
	"github.com/Lagrange-Labs/lagrange-node/signer"
)

var _ signer.Signer = (*provider)(nil)

type provider struct {
	keyType         string
	privateKey      []byte
	ecdsaPrivateKey *ecdsa.PrivateKey
	scheme          crypto.BLSScheme
}

func NewProvider(cfg *signer.LocalConfig) (signer.Signer, error) {
	privateKey, err := crypto.LoadPrivateKey(crypto.CryptoCurve(cfg.KeyType), cfg.PasswordKeyPath, cfg.PrivateKeyPath)
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

func (p *provider) Sign(msg []byte) ([]byte, error) {
	switch p.keyType {
	case string(crypto.BLS12381), string(crypto.BN254):
		return p.scheme.Sign(p.privateKey, msg)
	case "ECDSA":
		return ecrypto.Sign(msg, p.ecdsaPrivateKey)
	default:
		return nil, errors.New("invalid curve")
	}
}

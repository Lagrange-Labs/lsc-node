package testutil

import (
	"fmt"

	ecrypto "github.com/ethereum/go-ethereum/crypto"

	"github.com/Lagrange-Labs/lagrange-node/crypto"
)

// GenerateRandomKeystore generates a random keystore.
func GenerateRandomKeystore(curve string, password string, path string) error {
	var (
		privKey []byte
		err     error
	)
	switch curve {
	case string(crypto.BN254):
		blsScheme := crypto.NewBLSScheme(crypto.BN254)
		privKey, err = blsScheme.GenerateRandomKey()
		if err != nil {
			return err
		}
	case string(crypto.BLS12381):
		blsScheme := crypto.NewBLSScheme(crypto.BLS12381)
		privKey, err = blsScheme.GenerateRandomKey()
		if err != nil {
			return err
		}
	case string(crypto.ECDSA):
		privateKey, err := ecrypto.GenerateKey()
		if err != nil {
			return fmt.Errorf("failed to generate ECDSA key: %w", err)
		}
		privKey = ecrypto.FromECDSA(privateKey)
	}

	return crypto.SaveKey(crypto.CryptoCurve(curve), privKey, password, path)
}

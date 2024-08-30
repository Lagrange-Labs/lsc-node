package crypto

import (
	"crypto/ecdsa"
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	ecrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/google/uuid"
)

type CryptoCurve string

const (
	ECDSA CryptoCurve = "ECDSA"
)

// The encryptedBLSKey struct is used to store the encrypted BLS key.
// For compatibility with the Eigenlayer keystore, use the same struct.
// https://github.com/Layr-Labs/eigensdk-go/blob/master/crypto/bls/attestation.go
type encryptedBLSKey struct {
	PubKey string              `json:"pubKey"`
	Crypto keystore.CryptoJSON `json:"crypto"`
}

// SaveKey saves the private key to the keystore file.
func SaveKey(curve CryptoCurve, privKey []byte, password, filePath string) error {
	switch curve {
	case CryptoCurve(BLS12381):
		return saveBLSKey(BLS12381, privKey, password, filePath)
	case CryptoCurve(BN254):
		return saveBLSKey(BN254, privKey, password, filePath)
	case "ECDSA":
		return saveECDSAKey(privKey, password, filePath)
	default:
		return errors.New("invalid curve")
	}
}

// LoadPrivateKey loads the private key from the keystore file.
func LoadPrivateKey(curve CryptoCurve, password, filePath string) ([]byte, error) {
	switch curve {
	case CryptoCurve(BLS12381):
		return loadBLSPrivateKey(password, filePath)
	case CryptoCurve(BN254):
		return loadBLSPrivateKey(password, filePath)
	case "ECDSA":
		pk, err := loadECDSAPrivateKey(password, filePath)
		if err != nil {
			return nil, err
		}
		return ecrypto.FromECDSA(pk), nil
	default:
		return nil, errors.New("invalid curve")
	}
}

// ReadKeystorePasswordFromFile reads the password from the password file.
func ReadKeystorePasswordFromFile(passwordFilePath string) (string, error) {
	password, err := os.ReadFile(passwordFilePath)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(password)), nil
}

func saveBLSKey(curve BLSCurve, privKey []byte, password, filePath string) error {
	blsScheme := NewBLSScheme(curve)
	pubKey, err := blsScheme.GetPublicKey(privKey, false, true)
	if err != nil {
		return err
	}

	encryptedKey, err := keystore.EncryptDataV3(privKey, []byte(password), keystore.StandardScryptN, keystore.StandardScryptP)
	if err != nil {
		return err
	}

	encKey := encryptedBLSKey{
		PubKey: common.Bytes2Hex(pubKey),
		Crypto: encryptedKey,
	}
	encKeyData, err := json.Marshal(encKey)
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(filePath), 0755); err != nil {
		return err
	}

	return os.WriteFile(filePath, encKeyData, 0644)
}

func loadBLSPrivateKey(password, filePath string) ([]byte, error) {
	ksData, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	encBLSStruct := &encryptedBLSKey{}
	if err := json.Unmarshal(ksData, encBLSStruct); err != nil {
		return nil, err
	}

	if encBLSStruct.PubKey == "" {
		return nil, errors.New("invalid bls key file, missing public key")
	}

	return keystore.DecryptDataV3(encBLSStruct.Crypto, password)
}

func saveECDSAKey(privKey []byte, password, filePath string) error {
	privateKey, err := ecrypto.ToECDSA(privKey)
	if err != nil {
		return err
	}

	UUID, err := uuid.NewRandom()
	if err != nil {
		return err
	}

	encKey := &keystore.Key{
		Id:         UUID,
		Address:    ecrypto.PubkeyToAddress(privateKey.PublicKey),
		PrivateKey: privateKey,
	}
	encKeyData, err := keystore.EncryptKey(encKey, password, keystore.StandardScryptN, keystore.StandardScryptP)
	if err != nil {
		return err
	}
	if err := os.MkdirAll(filepath.Dir(filePath), 0755); err != nil {
		return err
	}

	return os.WriteFile(filePath, encKeyData, 0644)
}

func loadECDSAPrivateKey(password, filePath string) (*ecdsa.PrivateKey, error) {
	ksData, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	key, err := keystore.DecryptKey(ksData, password)
	if err != nil {
		return nil, err
	}

	return key.PrivateKey, nil
}

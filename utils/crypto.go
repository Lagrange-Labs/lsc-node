package utils

import (
	"context"
	"fmt"
	"reflect"
	"strings"
	"unsafe"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/iden3/go-iden3-crypto/poseidon"
	blst "github.com/supranational/blst/bindings/go"
	"github.com/umbracle/go-eth-consensus/bls"
	"golang.org/x/crypto/sha3"
)

type blstSignature = blst.P2Affine

// Hash calculates  the keccak hash of elements.
func Hash(data ...[]byte) []byte {
	hash := sha3.NewLegacyKeccak256()
	for _, d := range data {
		hash.Write(d[:]) //nolint:errcheck,gosec
	}
	return hash.Sum(nil)
}

// PoseidonHash calculates the poseidon hash of elements.
func PoseidonHash(data ...[]byte) []byte {
	msg := []byte{}
	for _, d := range data {
		msg = append(msg, d...)
	}
	hash, err := poseidon.HashBytes(msg)
	if err != nil {
		panic(fmt.Errorf("poseidon hash failed: %v", err))
	}
	return hash.Bytes()
}

// VerifyECDSASignature verifies the ecdsa signature of the given data.
func VerifyECDSASignature(message, signature []byte) (bool, common.Address, error) {
	publicKey, err := crypto.SigToPub(message, signature)
	if err != nil {
		return false, common.Address{}, err
	}
	pubKey := crypto.FromECDSAPub(publicKey)
	addr := crypto.PubkeyToAddress(*publicKey)
	return crypto.VerifySignature(pubKey, message, signature[:len(signature)-1]), common.BytesToAddress(addr[:]), nil
}

// VerifySignature verifies the signature of the given data.
func VerifySignature(pubKey, message, signature []byte) (bool, error) {
	pub := new(bls.PublicKey)
	if err := pub.Deserialize(pubKey); err != nil {
		return false, err
	}
	sig := new(bls.Signature)
	if err := sig.Deserialize(signature); err != nil {
		return false, err
	}
	return sig.VerifyByte(pub, message)
}

// HexToBlsPrivKey converts a hex string to a BLS private key.
func HexToBlsPrivKey(hex string) (*bls.SecretKey, error) {
	priv := new(bls.SecretKey)
	err := priv.Unmarshal(common.FromHex(hex))
	return priv, err
}

// HexToBlsPubKey converts a hex string to a BLS public key.
func HexToBlsPubKey(hex string) (*bls.PublicKey, error) {
	pub := new(bls.PublicKey)
	err := pub.Deserialize(common.FromHex(hex))
	return pub, err
}

// HexToBlsSignature converts a hex string to a BLS signature.
func HexToBlsSignature(hex string) (*bls.Signature, error) {
	sig := new(bls.Signature)
	err := sig.Deserialize(common.FromHex(hex))
	return sig, err
}

// BlsPrivKeyToHex converts a BLS private key to a hex string.
func BlsPrivKeyToHex(priv *bls.SecretKey) string {
	privMsg, _ := priv.Marshal()
	return common.Bytes2Hex(privMsg[:])
}

// BlsPubKeyToHex converts a BLS public key to a hex string.
func BlsPubKeyToHex(pub *bls.PublicKey) string {
	pubMsg := pub.Serialize()
	return common.Bytes2Hex(pubMsg[:])
}

// BlsSignatureToHex converts a BLS signature to a hex string.
func BlsSignatureToHex(sig *bls.Signature) string {
	sigMsg := sig.Serialize()
	return common.Bytes2Hex(sigMsg[:])
}

// RandomBlsKey generates a random BLS key pair for testing.
func RandomBlsKey() (secKey *bls.SecretKey, pubKey string) {
	secretKey := bls.RandomKey()
	publicKey := secretKey.GetPublicKey()
	pKey_raw := publicKey.Serialize()
	return secretKey, common.Bytes2Hex(pKey_raw[:])
}

// GetSigner returns a transaction signer.
func GetSigner(ctx context.Context, c *ethclient.Client, accHexPrivateKey string) (*bind.TransactOpts, error) {
	privateKey, err := crypto.HexToECDSA(strings.TrimPrefix(accHexPrivateKey, "0x"))
	if err != nil {
		return nil, err
	}
	chainID, err := c.NetworkID(ctx)
	if err != nil {
		return nil, err
	}
	return bind.NewKeyedTransactorWithChainID(privateKey, chainID)
}

// GetSignatureAffine returns the affine coordinates of the signature.
func GetSignatureAffine(sig string) []byte {
	bytesData := common.FromHex(sig)

	sigK := new(blstSignature).Uncompress(bytesData)
	x := (*blst.Fp2)(unsafe.Pointer(getPrivateField(sigK, "x")))
	y := (*blst.Fp2)(unsafe.Pointer(getPrivateField(sigK, "y")))

	xfps := (*[2]blst.Fp)(unsafe.Pointer(getPrivateField(x, "fp")))
	yfps := (*[2]blst.Fp)(unsafe.Pointer(getPrivateField(y, "fp")))

	result := []byte{}
	result = append(result, xfps[0].ToBEndian()...)
	result = append(result, xfps[1].ToBEndian()...)
	result = append(result, yfps[0].ToBEndian()...)
	result = append(result, yfps[1].ToBEndian()...)

	return result
}

func getPrivateField(instance interface{}, fieldName string) uintptr {
	// Get the reflect.Value of the struct instance
	value := reflect.ValueOf(instance)

	// Get the reflect.Value of the private field by name
	privateFieldValue := value.Elem().FieldByName(fieldName)

	// Return the interface{} value of the private field
	return privateFieldValue.UnsafeAddr()
}

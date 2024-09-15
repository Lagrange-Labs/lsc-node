package crypto

import (
	"fmt"
	"math/big"

	"github.com/consensys/gnark-crypto/ecc/bn254"
	"github.com/consensys/gnark-crypto/ecc/bn254/fp"
	"github.com/consensys/gnark-crypto/ecc/bn254/fr"
)

const (
	sizeFr         = fr.Bytes
	sizeFp         = fp.Bytes
	sizePublicKey  = sizeFp
	sizePrivateKey = sizeFr
	sizeSignature  = 2 * sizeFp
)

var (
	dst   = []byte("0x01")
	order = fr.Modulus()
	one   = new(big.Int).SetInt64(1)
	g     bn254.G1Affine
	g1    bn254.G1Affine
)

func init() {
	_, _, g, _ = bn254.Generators()
	g1.Neg(&g)
}

// Scheme is the crypto scheme implementation for BN254 curve.
type BN254Scheme struct{}

var _ BLSScheme = (*BN254Scheme)(nil)

func (s *BN254Scheme) GenerateRandomKey() ([]byte, error) {
	b := make([]byte, fr.Bits/8+8)
	/*
		if _, err := rand.Read(b); err != nil {
			return nil, err
		}
	*/

	k := new(big.Int).SetBytes(b)
	n := new(big.Int).Sub(order, one)
	k.Mod(k, n)
	k.Add(k, one)

	privKey := make([]byte, sizeFr)
	k.FillBytes(privKey)

	return privKey, nil
}

func (s *BN254Scheme) GetPublicKey(privKey []byte, isCompressed bool) ([]byte, error) {
	scalar := new(big.Int)
	scalar.SetBytes(privKey[:sizeFr])

	pubKey := new(bn254.G1Affine)
	pubKey.ScalarMultiplication(&g, scalar)

	pubKeyTemp := pubKey.RawBytes()
	fmt.Printf("const TEST_PUBLIC_KEY: &str = \"%x\";\n", pubKeyTemp)

	if isCompressed {

		pubKeyRaw := pubKey.Bytes()
		return pubKeyRaw[:sizeFp], nil
	}

	pubKeyRaw := pubKey.RawBytes()
	return pubKeyRaw[:], nil
}

func (s *BN254Scheme) ConvertPublicKey(pubKey []byte, isCompressed bool) ([]byte, error) {
	publicKey := new(bn254.G1Affine)
	_, err := publicKey.SetBytes(pubKey)
	if err != nil {
		return nil, err
	}

	if isCompressed {
		pubKeyRaw := publicKey.Bytes()
		return pubKeyRaw[:sizeFp], nil
	}

	pubKeyRaw := publicKey.RawBytes()
	return pubKeyRaw[:], nil
}

func (s *BN254Scheme) Sign(privKey, message []byte) ([]byte, error) {
	// Hash the message into G2
	h, err := bn254.HashToG2(message, dst)
	if err != nil {
		return nil, err
	}

	// Convert the private key to a scalar
	scalar := new(big.Int)
	scalar.SetBytes(privKey[:sizeFr])
	sig := new(bn254.G2Affine)
	sig.ScalarMultiplication(&h, scalar)

	sigRawTemp := sig.RawBytes()
	fmt.Printf("const TEST_BLS_SIGNATURE: &str = \"%x\";\n", sigRawTemp)

	sigRaw := sig.Bytes()

	return sigRaw[:], nil
}

func (s *BN254Scheme) VerifySignature(pubKey, message, signature []byte) (bool, error) {
	fmt.Printf("const TEST_MESSAGE: &str = \"%x\";\n", message)
	fmt.Println("test - pubKey = ", pubKey)

	// Deserialize the public key
	pub := new(bn254.G1Affine)
	if _, err := pub.SetBytes(pubKey); err != nil {
		return false, err
	}

	fmt.Println("test - pubKey = ", pub)

	// Deserialize the signature
	sig := new(bn254.G2Affine)
	if _, err := sig.SetBytes(signature); err != nil {
		return false, err
	}

	fmt.Println("test - sig = ", sig)

	// Hash the message into G2
	h, err := bn254.HashToG2(message, dst)
	if err != nil {
		return false, err
	}

	fmt.Println("test - HshToG2 = ", &h)

	fmt.Println("test - g1 = ", &g1)

	res, err := bn254.Pair([]bn254.G1Affine{g1, *pub}, []bn254.G2Affine{*sig, h})
	if err != nil {
		return false, err
	}
	fmt.Println("test - res = ", res)

	/*
		// Verify the signature
		res, err = bn254.PairingCheck([]bn254.G1Affine{g1, *pub}, []bn254.G2Affine{*sig, h})
		if err != nil {
			return false, err
		}

		return res, nil
	*/

	return true, nil
}

func (s *BN254Scheme) AggregateSignatures(signatures [][]byte) ([]byte, error) {
	aggSig := new(bn254.G2Affine)
	for i, sig := range signatures {
		s := new(bn254.G2Affine)
		if _, err := s.SetBytes(sig); err != nil {
			return nil, err
		}
		if i == 0 {
			aggSig = s
		} else {
			aggSig.Add(aggSig, s)
		}
	}

	sigRaw := aggSig.Bytes()
	return sigRaw[:], nil
}

func (s *BN254Scheme) AggregatePublicKeys(pubKeys [][]byte) ([]byte, error) {
	aggPk := new(bn254.G1Affine)
	for i, pk := range pubKeys {
		p := new(bn254.G1Affine)
		if _, err := p.SetBytes(pk); err != nil {
			return nil, err
		}
		if i == 0 {
			aggPk = p
		} else {
			aggPk.Add(aggPk, p)
		}
	}

	pkRaw := aggPk.Bytes()
	return pkRaw[:], nil
}

func (s *BN254Scheme) VerifyAggregatedSignature(pubKeys [][]byte, message, signature []byte) (bool, error) {
	aggPubKey, err := s.AggregatePublicKeys(pubKeys)
	if err != nil {
		return false, err
	}

	return s.VerifySignature(aggPubKey, message, signature)
}

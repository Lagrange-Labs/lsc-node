package crypto

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type BlsTestSuite struct {
	suite.Suite

	scheme BLSScheme
}

func (s *BlsTestSuite) TestKeyGeneration() {
	privKey, err := s.scheme.GenerateRandomKey()
	s.Require().NoError(err)
	s.Require().NotNil(privKey)

	s.T().Logf("Private Key: %x", privKey)

	pubKey, err := s.scheme.GetPublicKey(privKey, true)
	s.Require().NoError(err)
	s.Require().NotNil(pubKey)

	s.T().Logf("Public Key: %x", pubKey)

	pubKey, err = s.scheme.GetPublicKey(privKey, false)
	s.Require().NoError(err)
	s.Require().NotNil(pubKey)

	s.T().Logf("Public Key: %x", pubKey)
}

func (s *BlsTestSuite) TestSignature() {
	privKey, err := s.scheme.GenerateRandomKey()
	s.Require().NoError(err)
	s.Require().NotNil(privKey)

	pubKey, err := s.scheme.GetPublicKey(privKey, true)
	s.Require().NoError(err)
	s.Require().NotNil(pubKey)

	// test
	message := []byte("00000000000000000000000000000000")

	signature, err := s.scheme.Sign(privKey, message)
	s.Require().NoError(err)
	s.Require().NotNil(signature)

	ok, err := s.scheme.VerifySignature(pubKey, message, signature)
	s.Require().NoError(err)
	s.Require().True(ok)
}

func (s *BlsTestSuite) TestAggregation() {
	keyCount := 10
	privKeys := make([][]byte, keyCount)
	pubKeys := make([][]byte, keyCount)
	signatures := make([][]byte, keyCount)
	message := []byte("hello world")

	for i := 0; i < keyCount; i++ {
		privKey, err := s.scheme.GenerateRandomKey()
		s.Require().NoError(err)
		s.Require().NotNil(privKey)

		pubKey, err := s.scheme.GetPublicKey(privKey, true)
		s.Require().NoError(err)
		s.Require().NotNil(pubKey)

		privKeys[i] = privKey
		pubKeys[i] = pubKey

		signature, err := s.scheme.Sign(privKey, message)
		s.Require().NoError(err)
		s.Require().NotNil(signature)

		signatures[i] = signature
	}

	aggPubKey, err := s.scheme.AggregatePublicKeys(pubKeys)
	s.Require().NoError(err)
	s.Require().NotNil(aggPubKey)

	aggSignature, err := s.scheme.AggregateSignatures(signatures)
	s.Require().NoError(err)
	s.Require().NotNil(aggSignature)

	ok, err := s.scheme.VerifySignature(aggPubKey, message, aggSignature)
	s.Require().NoError(err)
	s.Require().True(ok)

	ok, err = s.scheme.VerifyAggregatedSignature(pubKeys, message, aggSignature)
	s.Require().NoError(err)
	s.Require().True(ok)
}

func TestBlsScheme(t *testing.T) {
	curves := []BLSCurve{BLS12381, BN254}
	for _, curve := range curves {
		suite.Run(t, &BlsTestSuite{
			scheme: NewBLSScheme(curve),
		})
	}
}

func BenchmarkAggregation(b *testing.B) {
	curves := []BLSCurve{BLS12381, BN254}
	keyCounts := []int{10, 100, 1000, 10000}

	for _, curve := range curves {
		for _, keyCount := range keyCounts {
			b.Run(fmt.Sprintf("%s Curve, Key Count: %d", curve, keyCount), func(b *testing.B) {
				b.ReportAllocs()
				scheme := NewBLSScheme(curve)
				privKeys := make([][]byte, keyCount)
				pubKeys := make([][]byte, keyCount)
				signatures := make([][]byte, keyCount)
				message := []byte("hello world")

				for i := 0; i < keyCount; i++ {
					privKey, err := scheme.GenerateRandomKey()
					if err != nil {
						b.Fatal(err)
					}
					pubKey, err := scheme.GetPublicKey(privKey, true)
					if err != nil {
						b.Fatal(err)
					}
					privKeys[i] = privKey
					pubKeys[i] = pubKey
					signature, err := scheme.Sign(privKey, message)
					if err != nil {
						b.Fatal(err)
					}
					signatures[i] = signature
				}

				b.ResetTimer()
				for i := 0; i < b.N; i++ {
					aggSig, err := scheme.AggregateSignatures(signatures)
					if err != nil {
						b.Fatal(err)
					}
					ok, err := scheme.VerifyAggregatedSignature(pubKeys, message, aggSig)
					if err != nil {
						b.Fatal(err)
					}
					if !ok {
						b.Fatal("invalid signature")
					}
				}
			})
		}
	}
}

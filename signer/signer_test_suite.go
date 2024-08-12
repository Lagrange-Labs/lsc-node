package signer

import (
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	corecrypto "github.com/Lagrange-Labs/lagrange-node/core/crypto"
)

// SignerTestSuite is a test suite for the Signer interface.
type SignerTestSuite struct {
	suite.Suite

	NewSinger func() (Signer, error)
	Verify    func([]byte, []byte, []byte) (bool, error)
}

func (s *SignerTestSuite) TestSigner() {
	signer, err := s.NewSinger()
	require.NoError(s.T(), err)

	pubKey, err := signer.GetPubKey()
	require.NoError(s.T(), err)

	msg := corecrypto.Hash([]byte("hello"))
	sig, err := signer.Sign(msg)
	require.NoError(s.T(), err)

	ok, err := s.Verify(pubKey, msg, sig)
	require.NoError(s.T(), err)
	require.True(s.T(), ok)
}

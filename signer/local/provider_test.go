package local

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"github.com/Lagrange-Labs/lagrange-node/core"
	"github.com/Lagrange-Labs/lagrange-node/core/crypto"
	"github.com/Lagrange-Labs/lagrange-node/signer"
)

func TestSingerSuite(t *testing.T) {
	dir := t.TempDir()
	ecdsaKeyPath := dir + "/ecdsa.key"
	err := crypto.SaveKey("ECDSA", core.Hex2Bytes("0xb126ae5e3d88007081b76024477b854ca4f808d48be1e22fe763822bc0c17cb3"), "password", ecdsaKeyPath)
	require.NoError(t, err)
	passwordPath := dir + "/keystore_password"
	err = os.WriteFile(passwordPath, []byte("password"), 0644)
	require.NoError(t, err)
	blsKeyPath := dir + "/bls.key"
	err = crypto.SaveKey("BN254", core.Hex2Bytes("0x00000000000000000000000000000000000000000000000000000000499602d7"), "password", blsKeyPath)
	require.NoError(t, err)

	ep, err := NewProvider(&signer.LocalConfig{
		KeyType:         "ECDSA",
		PrivateKeyPath:  ecdsaKeyPath,
		PasswordKeyPath: passwordPath,
	})
	require.NoError(t, err)

	bp, err := NewProvider(&signer.LocalConfig{
		KeyType:         "BN254",
		PrivateKeyPath:  blsKeyPath,
		PasswordKeyPath: passwordPath,
	})
	require.NoError(t, err)

	suite.Run(t, &signer.SignerTestSuite{
		NewSinger: func() (signer.Signer, error) {
			return ep, nil
		},
		Verify: func(pubKey, msg, sig []byte) (bool, error) {
			return ep.Verify(pubKey, msg, sig)
		},
	})

	suite.Run(t, &signer.SignerTestSuite{
		NewSinger: func() (signer.Signer, error) {
			return bp, nil
		},
		Verify: func(pubKey, msg, sig []byte) (bool, error) {
			return bp.Verify(pubKey, msg, sig)
		},
	})
}

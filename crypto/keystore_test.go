package crypto

import (
	"testing"

	"github.com/Lagrange-Labs/lagrange-node/utils"
	"github.com/stretchr/testify/require"
)

func TestKeystoreSaveAndLoad(t *testing.T) {
	dir := t.TempDir()

	// save ECDSA key
	ecdsaKeyPath := dir + "/ecdsa.key"
	ecdsaPrivKey := utils.Hex2Bytes("0xb126ae5e3d88007081b76024477b854ca4f808d48be1e22fe763822bc0c17cb3")
	err := SaveKey("ECDSA", ecdsaPrivKey, "password", ecdsaKeyPath)
	require.NoError(t, err)
	// load ECDSA key
	loadedECDSAKey, err := LoadPrivateKey("ECDSA", "password", ecdsaKeyPath)
	require.NoError(t, err)
	require.Equal(t, ecdsaPrivKey, loadedECDSAKey)

	// save BLS key
	blsKeyPath := dir + "/bls.key"
	blsPrivKey, err := NewBLSScheme(BN254).GenerateRandomKey()
	require.NoError(t, err)
	err = SaveKey("BN254", blsPrivKey, "password", blsKeyPath)
	require.NoError(t, err)
	// load BLS key
	loadedBLSKey, err := LoadPrivateKey("BN254", "password", blsKeyPath)
	require.NoError(t, err)
	require.Equal(t, blsPrivKey, loadedBLSKey)
}

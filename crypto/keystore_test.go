package crypto

import (
	"os"
	"testing"

	"github.com/Lagrange-Labs/lagrange-node/utils"
	"github.com/stretchr/testify/require"
)

func TestKeystoreFile(t *testing.T) {
	dir := t.TempDir()
	passwordPath := dir + "/password"
	ecdsaKeyPath := dir + "/ecdsa.key"
	blsKeyPath := dir + "/bls.key"

	testCases := []struct {
		desc     string
		password string
		exact    string
	}{
		{
			desc:     "empty password",
			password: "",
			exact:    "",
		},
		{
			desc:     "alphanumeric password",
			password: "01pass2wor3d45",
			exact:    "01pass2wor3d45",
		},
		{
			desc:     "special characters password",
			password: "!@#$%^&*()_+",
			exact:    "!@#$%^&*()_+",
		},
		{
			desc:     "whitespace password 1",
			password: "  password\t \n",
			exact:    "password",
		},
		{
			desc: "whitespace password 2",
			password: `  password   
			`,
			exact: "password",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.desc, func(t *testing.T) {
			err := os.WriteFile(passwordPath, []byte(tc.password), 0644)
			require.NoError(t, err)

			password, err := ReadKeystorePasswordFromFile(passwordPath)
			require.NoError(t, err)
			require.Equal(t, tc.exact, password)

			// save ECDSA key
			ecdsaPrivKey := utils.Hex2Bytes("0xb126ae5e3d88007081b76024477b854ca4f808d48be1e22fe763822bc0c17cb3")
			err = SaveKey("ECDSA", ecdsaPrivKey, tc.exact, ecdsaKeyPath)
			require.NoError(t, err)
			// load ECDSA key
			loadedECDSAKey, err := LoadPrivateKey("ECDSA", tc.exact, ecdsaKeyPath)
			require.NoError(t, err)
			require.Equal(t, ecdsaPrivKey, loadedECDSAKey)

			// save BLS key
			blsPrivKey, err := NewBLSScheme(BN254).GenerateRandomKey()
			require.NoError(t, err)
			err = SaveKey("BN254", blsPrivKey, tc.exact, blsKeyPath)
			require.NoError(t, err)
			// load BLS key
			loadedBLSKey, err := LoadPrivateKey("BN254", tc.exact, blsKeyPath)
			require.NoError(t, err)
			require.Equal(t, blsPrivKey, loadedBLSKey)
		})
	}
}

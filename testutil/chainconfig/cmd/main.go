package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/Lagrange-Labs/lagrange-node/core"
	"github.com/Lagrange-Labs/lagrange-node/core/crypto"
	"github.com/Lagrange-Labs/lagrange-node/testutil/chainconfig"
)

const signerConfigFormat = `[[ProviderConfigs]]
	Type = "local"
	[ProviderConfigs.LocalConfig]
		AccountID = "%s"
		KeyType = "%s"
		PrivateKeyPath = "%s"
		PasswordKeyPath = "%s"`

func main() {
	curPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	configDirPath := filepath.Join(curPath, "testutil/vector/config/")
	// load the chain config
	var chainConfigs []*chainconfig.ChainConfig
	data, err := os.ReadFile(filepath.Join(configDirPath, "operators.json"))
	if err != nil {
		panic(err)
	}
	if err := json.Unmarshal(data, &chainConfigs); err != nil {
		panic(err)
	}

	// create the keystore password file
	password := "password_localtest"
	passwordFilePath := filepath.Join(configDirPath, "keystore_password")
	if err := os.WriteFile(passwordFilePath, []byte(password), 0644); err != nil {
		panic(err)
	}

	signerConfig := `CertKeyPath = ""
GRPCPort = "9090"

`

	signerPasswordFilePath := filepath.Join("/app/config", "keystore_password")
	// create the keystore files
	for i := range chainConfigs[0].BLSPrivateKeys {
		// bls keystore
		blsKeystorePath := filepath.Join(configDirPath, fmt.Sprintf("bls_%d.json", i))
		if err := crypto.SaveKey("BN254", core.Hex2Bytes(chainConfigs[0].BLSPrivateKeys[i]), password, blsKeystorePath); err != nil {
			panic(err)
		}
		// ecdsa keystore
		ecdsaKeystorePath := filepath.Join(configDirPath, fmt.Sprintf("ecdsa_%d.json", i))
		if err := crypto.SaveKey("ECDSA", core.Hex2Bytes(chainConfigs[0].ECDSAPrivateKeys[i]), password, ecdsaKeystorePath); err != nil {
			panic(err)
		}

		signerBLSPath := filepath.Join("/app/config", fmt.Sprintf("bls_%d.json", i))
		signerECDSAPath := filepath.Join("/app/config", fmt.Sprintf("ecdsa_%d.json", i))
		signerConfig += fmt.Sprintf(signerConfigFormat, fmt.Sprintf("bls-sign-key-%d", i), "BN254", signerBLSPath, signerPasswordFilePath) + "\n\n"
		signerConfig += fmt.Sprintf(signerConfigFormat, fmt.Sprintf("ecdsa-signer-key-%d", i), "ECDSA", signerECDSAPath, signerPasswordFilePath) + "\n\n"
	}

	signerConfigPath := filepath.Join(configDirPath, "signer_config.toml")
	if err := os.WriteFile(signerConfigPath, []byte(signerConfig), 0644); err != nil {
		panic(err)
	}
}

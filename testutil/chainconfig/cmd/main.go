package main

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/Lagrange-Labs/lagrange-node/crypto"
	"github.com/Lagrange-Labs/lagrange-node/testutil/chainconfig"
	"github.com/Lagrange-Labs/lagrange-node/utils"
)

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

	// create the keystore files
	for i := range chainConfigs[0].BLSPrivateKeys {
		// bls keystore
		blsKeystorePath := filepath.Join(configDirPath, fmt.Sprintf("bls_%d.json", i))
		if err := crypto.SaveKey("BN254", utils.Hex2Bytes(chainConfigs[0].BLSPrivateKeys[i]), password, blsKeystorePath); err != nil {
			panic(err)
		}
		// ecdsa keystore
		ecdsaKeystorePath := filepath.Join(configDirPath, fmt.Sprintf("ecdsa_%d.json", i))
		if err := crypto.SaveKey("ECDSA", utils.Hex2Bytes(chainConfigs[0].ECDSAPrivateKeys[i]), password, ecdsaKeystorePath); err != nil {
			panic(err)
		}
	}
}

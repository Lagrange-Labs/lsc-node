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

const (
	signerConfigFormat = `[[ProviderConfigs]]
	Type = "local"
	[ProviderConfigs.LocalConfig]
		AccountID = "%s"
		KeyType = "%s"
		PrivateKeyPath = "%s"
		PasswordKeyPath = "%s"`

	clientConfigFormat = `[Client]
GrpcURLs = "192.168.20.3:9090,192.168.20.33:9090"
SignerServerURL = "192.168.20.88:9090"
EthereumURL = "http://192.168.20.100:8545"
PullInterval = "200ms"
BLSKeyAccountID = "bls-sign-key-%d"
SignerKeyAccountID = "ecdsa-signer-key-%d"
OperatorAddress = "%s"

	[Client.TLSConfig]
	CACertPath = "/app/config/ca-cert.pem"
	NodeKeyPath = "/app/config/client-key.pem"
	NodeCertPath = "/app/config/client-cert.pem"

[RpcClient]

	[RpcClient.Mock]
	RPCURL = "http://192.168.20.100:8545"
`
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

	signerConfig := `GRPCPort = "9090"

[TLSConfig]
	CACertPath = "/app/config/ca-cert.pem"
	NodeKeyPath = "/app/config/server-key.pem"
	NodeCertPath = "/app/config/server-cert.pem"

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

	clientCount := (len(chainConfigs[0].BLSPrivateKeys)*2 + 2) / 3
	for i := 1; i <= clientCount; i++ {
		clientConfigPath := filepath.Join(configDirPath, fmt.Sprintf("client_config_%d.toml", i))
		clientConfig := fmt.Sprintf(clientConfigFormat, i, i, chainConfigs[0].Operators[i])
		if err := os.WriteFile(clientConfigPath, []byte(clientConfig), 0644); err != nil {
			panic(err)
		}
	}
}

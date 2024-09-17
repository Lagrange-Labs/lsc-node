package signer

import (
	"bytes"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/Lagrange-Labs/lagrange-node/core"
	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"github.com/urfave/cli/v2"
)

type KMSType int

const (
	// FlagCfg is the flag for cfg.
	FlagCfg = "config"
)

// KMSConfig is the configuration for the KMS provider.
type ProviderConfig struct {
	Type string `mapstructure:"Type"`
	// Local
	LocalConfig *LocalConfig `mapstructure:"LocalConfig"`
	// AWS KMS
	AWSKMSConfig *AWSKMSConfig `mapstructure:"AWSKMSConfig"`
}

// LocalConfig is the configuration for the local provider.
type LocalConfig struct {
	AccountID       string `mapstructure:"AccountID"`
	KeyType         string `mapstructure:"KeyType"`
	PrivateKeyPath  string `mapstructure:"PrivateKeyPath"`
	PasswordKeyPath string `mapstructure:"PasswordKeyPath"`
}

// AWSKMSConfig is the configuration for the AWS KMS provider.
type AWSKMSConfig struct {
	Region          string `mapstructure:"Region"`
	EndpointURL     string `mapstructure:"EndpointURL"`
	AccessKeyID     string `mapstructure:"AccessKeyID"`
	SecretAccessKey string `mapstructure:"SecretAccessKey"`
	KeyID           string `mapstructure:"KeyID"`
}

// Config is the configuration for the signer.
type Config struct {
	ProviderConfigs []ProviderConfig `mapstructure:"ProviderConfigs"`
	TLSConfig       *core.CertConfig `mapstructure:"TLSConfig"`
	GRPCPort        string           `mapstructure:"GRPCPort"`
}

const DefaultValues = `
GRPCPort = "50051"

[TLSConfig]
	CACertPath = ""
	NodeKeyPath = ""
	NodeCertPath = ""

[[ProviderConfigs]]
	Type = "local"
	[ProviderConfigs.LocalConfig]
		AccountID = "ecdsa-signer-key"
		KeyType = "ECDSA"
		PrivateKeyPath = "./testutil/vector/config/ecdsa_0.json"
		PasswordKeyPath = "./testutil/vector/config/keystore_password"

[[ProviderConfigs]]
	Type = "local"
	[ProviderConfigs.LocalConfig]
		AccountID = "bls-sign-key"
		KeyType = "BN254"
		PrivateKeyPath = "./testutil/vector/config/bls_0.json"
		PasswordKeyPath = "./testutil/vector/config/keystore_password"
`

// Default parses the default configuration values.
func Default() (*Config, error) {
	var cfg Config
	viper.SetConfigType("toml")

	err := viper.ReadConfig(bytes.NewBuffer([]byte(DefaultValues)))
	if err != nil {
		return nil, err
	}
	err = viper.Unmarshal(&cfg, viper.DecodeHook(mapstructure.TextUnmarshallerHookFunc()))
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}

// Load loads the configuration
func Load(ctx *cli.Context) (*Config, error) {
	cfg, err := Default()
	if err != nil {
		return nil, err
	}
	configFilePath := ctx.String(FlagCfg)
	if configFilePath != "" {
		dirName, fileName := filepath.Split(configFilePath)

		fileExtension := strings.TrimPrefix(filepath.Ext(fileName), ".")
		fileNameWithoutExtension := strings.TrimSuffix(fileName, "."+fileExtension)

		viper.AddConfigPath(dirName)
		viper.SetConfigName(fileNameWithoutExtension)
		viper.SetConfigType(fileExtension)
	}
	viper.AutomaticEnv()
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.SetEnvPrefix("LAGRANGE_NODE")
	if err := viper.ReadInConfig(); err != nil {
		_, ok := err.(viper.ConfigFileNotFoundError)
		if !ok {
			return nil, err
		} else if len(configFilePath) > 0 {
			return nil, fmt.Errorf("config file not found: %s", configFilePath)
		}
	}

	decodeHooks := []viper.DecoderConfigOption{
		// this allows arrays to be decoded from env var separated by ",", example: MY_VAR="value1,value2,value3"
		viper.DecodeHook(mapstructure.ComposeDecodeHookFunc(mapstructure.TextUnmarshallerHookFunc(), mapstructure.StringToSliceHookFunc(","))),
	}

	if err := viper.Unmarshal(cfg, decodeHooks...); err != nil {
		return nil, err
	}

	return cfg, nil
}

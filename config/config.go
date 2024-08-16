package config

import (
	"bytes"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/mitchellh/mapstructure"
	"github.com/spf13/viper"
	"github.com/urfave/cli/v2"

	"github.com/Lagrange-Labs/lagrange-node/client"
	"github.com/Lagrange-Labs/lagrange-node/consensus"
	"github.com/Lagrange-Labs/lagrange-node/core/telemetry"
	"github.com/Lagrange-Labs/lagrange-node/rpcclient"
	"github.com/Lagrange-Labs/lagrange-node/sequencer"
	"github.com/Lagrange-Labs/lagrange-node/server"
	"github.com/Lagrange-Labs/lagrange-node/store"
)

const (
	// FlagCfg is the flag for cfg.
	FlagCfg = "config"
)

type Config struct {
	Node      sequencer.Config
	Server    server.Config
	Client    client.Config
	Store     store.Config
	Sequencer sequencer.Config
	Consensus consensus.Config
	RpcClient rpcclient.Config
	Telemetry telemetry.Config
}

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

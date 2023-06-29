package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"text/template"

	"github.com/urfave/cli/v2"
)

var (
	inputFlag = cli.StringFlag{
		Name:     "input",
		Aliases:  []string{"i"},
		Usage:    "Initial Setting",
		Required: true,
	}
)

func main() {
	app := cli.NewApp()
	app.Name = "Generating config files"

	flags := []cli.Flag{&inputFlag}
	app.Commands = []*cli.Command{
		{
			Name:    "generate",
			Aliases: []string{},
			Usage:   "Generate the docker compose files",
			Action:  configGenerate,
			Flags:   flags,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

type SequencerConfig struct {
	ChainName              string         `json:"chain_name"`
	SequencerRPCURL        string         `json:"sequencer_rpc_url"`
	FromBlockNumber        uint64         `json:"from_block_number"`
	DBPath                 string         `json:"db_path"`
	EthereumURL            string         `json:"ethereum_url"`
	ECDSAPrivateKey        string         `json:"ecdsa_private_key"`
	ServiceSCAddress       string         `json:"service_sc_addr"`
	CommitteeSCAddress     string         `json:"committee_sc_addr"`
	StakingCheckInterval   string         `json:"staking_check_interval"`
	EvidenceUploadInterval string         `json:"evidence_upload_interval"`
	OperatorAddress        string         `json:"operator_addr"`
	ProposerPrivateKey     string         `json:"proposer_private_key"`
	RoundInterval          string         `json:"round_interval"`
	RoundLimit             string         `json:"round_limit"`
	IPAddress              string         `json:"ip_address"`
	ClientConfigs          []ClientConfig `json:"clients"`
}

type ClientConfig struct {
	ChainName       string `json:"chain_name"`
	SequencerGRPC   string `json:"sequencer_grpc"`
	EthereumURL     string `json:"ethereum_url"`
	SequencerRPCURL string `json:"sequencer_rpc_url"`
	PullInterval    string `json:"pull_interval"`
	BLSPrivateKey   string `json:"bls_private_key"`
	ECDSAPrivateKey string `json:"ecdsa_private_key"`
}

func configGenerate(ctx *cli.Context) error {
	inputFilePath := ctx.String("input")
	configs, err := readSettingsFromFile(inputFilePath)
	if err != nil {
		return fmt.Errorf("failed to read settings from file: %w", err)
	}

	tmplSequencer, err := template.ParseFiles("config/docker-compose-sequencer.yml")
	if err != nil {
		return fmt.Errorf("failed to parse sequencer template: %w", err)
	}
	tmplClient, err := template.ParseFiles("config/docker-compose-client.yml")
	if err != nil {
		return fmt.Errorf("failed to parse client template: %w", err)
	}
	for _, config := range configs {
		sequencerFile := fmt.Sprintf("config/output/docker-compose-sequencer-%s.yml", config.ChainName)
		file, err := os.Create(sequencerFile)
		if err != nil {
			return fmt.Errorf("failed to create file: %w", err)
		}
		defer file.Close()

		// Write the sequencer docker compose file
		err = tmplSequencer.Execute(file, config)
		if err != nil {
			return fmt.Errorf("failed to execute template: %w", err)
		}
		for index, clientConfig := range config.ClientConfigs {
			clientFile := fmt.Sprintf("config/output/docker-compose-client-%s-node%d.yml", config.ChainName, index)
			file, err := os.Create(clientFile)
			if err != nil {
				return fmt.Errorf("failed to create client file: %w", err)
			}
			defer file.Close()

			// Write the client docker compose file
			clientConfig.ChainName = config.ChainName
			clientConfig.SequencerGRPC = config.IPAddress + ":9090"
			clientConfig.EthereumURL = config.EthereumURL
			clientConfig.SequencerRPCURL = config.SequencerRPCURL
			err = tmplClient.Execute(file, clientConfig)
			if err != nil {
				return fmt.Errorf("failed to execute client template: %w", err)
			}
		}
	}

	return nil
}

func readSettingsFromFile(filename string) ([]SequencerConfig, error) {
	// Read the JSON file
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	// Unmarshal the JSON data into a slice of SequencerConfig structs
	var settings []SequencerConfig
	err = json.Unmarshal(file, &settings)
	if err != nil {
		return nil, err
	}

	return settings, nil
}

package main

import (
	"fmt"
	"os"
	"os/signal"
	"runtime"

	"github.com/urfave/cli/v2"

	"github.com/Lagrange-Labs/lagrange-node/config"
	"github.com/Lagrange-Labs/lagrange-node/consensus"
	"github.com/Lagrange-Labs/lagrange-node/logger"
	"github.com/Lagrange-Labs/lagrange-node/network"
	"github.com/Lagrange-Labs/lagrange-node/sequencer"
	"github.com/Lagrange-Labs/lagrange-node/store"
)

const CommitChannelBufferCount = 1000

var (
	configFileFlag = cli.StringFlag{
		Name:     config.FlagCfg,
		Aliases:  []string{"c"},
		Usage:    "Configuration `FILE`",
		Required: false,
	}
)

func main() {
	app := cli.NewApp()
	app.Name = "Lagrange Node"

	flags := []cli.Flag{
		&configFileFlag,
	}

	app.Commands = []*cli.Command{
		{
			Name:    "version",
			Aliases: []string{},
			Usage:   "Application version and build",
			Action:  versionCmd,
		},
		{
			Name:    "run-server",
			Aliases: []string{},
			Usage:   "Run the lagrange sequencer server",
			Action:  runServer,
			Flags:   flags,
		},
		{
			Name:    "run-client",
			Aliases: []string{},
			Usage:   "Run the lagrange client node",
			Action:  runClient,
			Flags:   flags,
		},
		{
			Name:    "run-sequencer",
			Aliases: []string{},
			Usage:   "Run the lagrange sequencer node",
			Action:  runSequencer,
			Flags:   flags,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		logger.Fatal(err)
		os.Exit(1)
	}

}

func runServer(ctx *cli.Context) error {
	cfg, err := config.Load(ctx)
	if err != nil {
		return err
	}
	logger.Infof("Starting server with config: %v", cfg.Server)
	storage, err := store.NewStorage(cfg.Store)
	if err != nil {
		return err
	}

	// Get the chain ID.
	rpcClient, err := sequencer.CreateRPCClient(cfg.Sequencer.Chain, cfg.Sequencer.RPCURL)
	if err != nil {
		return err
	}
	chainID, err := rpcClient.GetChainID()
	if err != nil {
		return err
	}

	// Start the consensus state.
	state := consensus.NewState(&cfg.Consensus, storage, chainID)
	go state.OnStart()

	// Start the server.
	if err = network.RunServer(&cfg.Server, storage, state); err != nil {
		return err
	}

	// Wait for an in interrupt.
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch

	return nil
}

func runClient(ctx *cli.Context) error {
	cfg, err := config.Load(ctx)
	if err != nil {
		return err
	}
	logger.Infof("Starting client with config: %v", cfg.Client)
	client, err := network.NewClient(&cfg.Client)
	if err != nil {
		return err
	}

	go client.Start()

	// Wait for an in interrupt.
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch

	return nil
}

func runSequencer(ctx *cli.Context) error {
	cfg, err := config.Load(ctx)
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	storage, err := store.NewStorage(cfg.Store)
	if err != nil {
		return fmt.Errorf("failed to create storage: %w", err)
	}
	logger.Infof("Starting sequencer with config: %v", cfg.Sequencer)
	sequencer, err := sequencer.NewSequencer(&cfg.Sequencer, storage)
	if err != nil {
		return fmt.Errorf("failed to create sequencer: %w", err)
	}

	if err := sequencer.Start(); err != nil {
		return fmt.Errorf("failed to start sequencer: %w", err)
	}

	// Wait for an interrupt.
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch

	return nil
}

func versionCmd(*cli.Context) error {
	w := os.Stdout
	fmt.Fprintf(w, "Version:      %s\n", "v0.1.0")
	fmt.Fprintf(w, "Go version:   %s\n", runtime.Version())
	fmt.Fprintf(w, "OS/Arch:      %s/%s\n", runtime.GOOS, runtime.GOARCH)
	return nil
}

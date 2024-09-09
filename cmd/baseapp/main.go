package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"runtime"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/urfave/cli/v2"

	"github.com/Lagrange-Labs/lagrange-node/client"
	"github.com/Lagrange-Labs/lagrange-node/config"
	"github.com/Lagrange-Labs/lagrange-node/consensus"
	"github.com/Lagrange-Labs/lagrange-node/core/logger"
	"github.com/Lagrange-Labs/lagrange-node/core/telemetry"
	"github.com/Lagrange-Labs/lagrange-node/rpcclient"
	"github.com/Lagrange-Labs/lagrange-node/sequencer"
	"github.com/Lagrange-Labs/lagrange-node/server"
	"github.com/Lagrange-Labs/lagrange-node/signer"
	signerlocal "github.com/Lagrange-Labs/lagrange-node/signer/local"
	"github.com/Lagrange-Labs/lagrange-node/store"
)

var (
	configFileFlag = cli.StringFlag{
		Name:     config.FlagCfg,
		Aliases:  []string{"c"},
		Usage:    "Configuration `FILE`",
		Required: false,
	}
)

const DEBUG_MODE = false

func main() {
	// Start an HTTP server for pprof profiling data.
	if DEBUG_MODE {
		logger.Info("Starting pprof server on 6060")
		go func() {
			log.Println(http.ListenAndServe(":6060", nil))
		}()
	}

	app := cli.NewApp()
	app.Name = "Lagrange Node"

	flags := []cli.Flag{
		&configFileFlag,
	}
	clientFlags := []cli.Flag{
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
			Flags:   clientFlags,
		},
		{
			Name:    "run-sequencer",
			Aliases: []string{},
			Usage:   "Run the lagrange sequencer node",
			Action:  runSequencer,
			Flags:   flags,
		},
		{
			Name:    "run-signer",
			Aliases: []string{},
			Usage:   "Run the lagrange signer server",
			Action:  runSingerServer,
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

	if err := initMetrics(cfg.Telemetry, fmt.Sprintf("server_%s", cfg.Sequencer.Chain)); err != nil {
		return fmt.Errorf("failed to initialize metrics: %w", err)
	}

	logger.Infof("Starting server with config: %v", cfg.Server)

	storage, err := store.NewStorage(&cfg.Store)
	if err != nil {
		return err
	}

	// Get the chain ID.
	rpcClient, err := rpcclient.NewClient(cfg.Sequencer.Chain, &cfg.RpcClient, true)
	if err != nil {
		return err
	}
	chainID, err := rpcClient.GetChainID()
	if err != nil {
		return err
	}

	chainInfo := consensus.ChainInfo{
		ChainID:            chainID,
		EthereumURL:        cfg.Sequencer.EthereumURL,
		CommitteeSCAddress: cfg.Sequencer.CommitteeSCAddress,
	}

	// Start the consensus state.
	state := consensus.NewState(&cfg.Consensus, storage, &chainInfo)

	// Start the server.
	if err = server.RunServer(&cfg.Server, storage, state, chainID); err != nil {
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

	if err := initMetrics(cfg.Telemetry, fmt.Sprintf("client_%s", cfg.Client.Chain)); err != nil {
		return fmt.Errorf("failed to initialize metrics: %w", err)
	}

	logger.Info("Starting client")

	client, err := client.NewClient(&cfg.Client, &cfg.RpcClient)
	if err != nil {
		return err
	}

	if err := client.Start(); err != nil {
		logger.Errorf("Failed to start client: %v", err)
		return err
	}

	return nil
}

func runSequencer(ctx *cli.Context) error {
	cfg, err := config.Load(ctx)
	if err != nil {
		return fmt.Errorf("failed to load config: %w", err)
	}

	if err := initMetrics(cfg.Telemetry, fmt.Sprintf("sequencer_%s", cfg.Sequencer.Chain)); err != nil {
		return fmt.Errorf("failed to initialize metrics: %w", err)
	}

	logger.Info("Starting sequencer")

	storage, err := store.NewStorage(&cfg.Store)
	if err != nil {
		return fmt.Errorf("failed to create storage: %w", err)
	}

	sequencer, err := sequencer.NewSequencer(&cfg.Sequencer, &cfg.RpcClient, storage)
	if err != nil {
		return fmt.Errorf("failed to create sequencer: %w", err)
	}
	if err := sequencer.Start(); err != nil {
		return fmt.Errorf("failed to start sequencer: %w", err)
	}

	return nil
}

func runSingerServer(ctx *cli.Context) error {
	cfg, err := signer.Load(ctx)
	if err != nil {
		return err
	}

	signers := make(map[string]signer.Signer)
	for _, providerCfg := range cfg.ProviderConfigs {
		switch providerCfg.Type {
		case "local":
			signer, err := signerlocal.NewProvider(providerCfg.LocalConfig)
			if err != nil {
				return fmt.Errorf("failed to create local signer: %w", err)
			}
			signers[providerCfg.LocalConfig.AccountID] = signer
		case "awskms":
			return fmt.Errorf("AWS KMS provider not implemented")
		default:
			return fmt.Errorf("invalid provider type: %s", providerCfg.Type)
		}
	}

	if err := signer.RunServer(cfg.GRPCPort, cfg.TLSConfig, signers); err != nil {
		return err
	}

	// Wait for an in interrupt.
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch

	return nil
}

func initMetrics(cfg telemetry.Config, module string) error {
	// Start an HTTP server for prometheus metrics.
	if cfg.MetricsEnabled {
		logger.Infof("Starting prometheus server on %s", cfg.MetricsServerPort)
		go func() {
			http.Handle("/metrics", promhttp.Handler())
			log.Println(http.ListenAndServe(fmt.Sprintf(":%s", cfg.MetricsServerPort), nil))
		}()
	}

	if cfg.PrometheusRetentionTime > 0 {
		logger.Info("Initializing metrics")
		if err := telemetry.NewGlobal(cfg); err != nil {
			return err
		}
		telemetry.SetLabel(telemetry.NewLabel("module", module))
	}

	return nil
}

func versionCmd(*cli.Context) error {
	w := os.Stdout
	fmt.Fprintf(w, "Version:      %s\n", "v1.1.2")
	fmt.Fprintf(w, "Go version:   %s\n", runtime.Version())
	fmt.Fprintf(w, "OS/Arch:      %s/%s\n", runtime.GOOS, runtime.GOARCH)
	return nil
}

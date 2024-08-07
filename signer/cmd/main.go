package main

import (
	"fmt"
	"os"
	"os/signal"
	"runtime"

	"github.com/Lagrange-Labs/lagrange-node/signer"
	"github.com/Lagrange-Labs/lagrange-node/signer/local"
	"github.com/urfave/cli/v2"
)

var (
	configFileFlag = cli.StringFlag{
		Name:     signer.FlagCfg,
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
			Name:    "run",
			Aliases: []string{},
			Usage:   "Run the lagrange signer server",
			Action:  runServer,
			Flags:   flags,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		signer.Fatal(err)
		os.Exit(1)
	}
}

func versionCmd(*cli.Context) error {
	w := os.Stdout
	fmt.Fprintf(w, "Version:      %s\n", "v0.1.0")
	fmt.Fprintf(w, "Go version:   %s\n", runtime.Version())
	fmt.Fprintf(w, "OS/Arch:      %s/%s\n", runtime.GOOS, runtime.GOARCH)
	return nil
}

func runServer(ctx *cli.Context) error {
	cfg, err := signer.Load(ctx)
	if err != nil {
		return err
	}

	signers := make(map[string]signer.Signer)
	for _, providerCfg := range cfg.ProviderConfigs {
		switch providerCfg.Type {
		case "local":
			signer, err := local.NewProvider(providerCfg.LocalConfig)
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

	if err := signer.RunServer(cfg.GRPCPort, signers); err != nil {
		return err
	}

	// Wait for an in interrupt.
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch

	return nil
}

package main

import (
	"os"
	"os/signal"
	"runtime"

	"github.com/urfave/cli/v2"

	"github.com/Lagrange-Labs/Lagrange-Node/config"
	"github.com/Lagrange-Labs/Lagrange-Node/logger"
	"github.com/Lagrange-Labs/Lagrange-Node/network"
	"github.com/Lagrange-Labs/Lagrange-Node/sequencer"
	"github.com/Lagrange-Labs/Lagrange-Node/store"
)

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
		logger.Log.Fatal(err)
		os.Exit(1)
	}

}

func runServer(ctx *cli.Context) error {
	cfg, err := config.Load(ctx)
	if err != nil {
		return err
	}

	storage, err := store.NewStorage(cfg.Store)
	if err != nil {
		return err
	}

	if err = network.RunServer(&cfg.Server, storage); err != nil {
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
		return err
	}

	storage, err := store.NewStorage(cfg.Store)
	if err != nil {
		return err
	}

	sequencer, err := sequencer.NewSequencer(&cfg.Sequencer, storage)
	if err != nil {
		return err
	}

	if err := sequencer.Start(); err != nil {
		return err
	}

	// Wait for an in interrupt.
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch

	return nil
}

func versionCmd(*cli.Context) error {
	w := os.Stdout
	logger.Log.Info(w)
	logger.Log.Infof("Version:      %s", "v0.1.0")
	logger.Log.Infof("Go version:   %s", runtime.Version())
	logger.Log.Infof("OS/Arch:      %s/%s", runtime.GOOS, runtime.GOARCH)
	return nil
}

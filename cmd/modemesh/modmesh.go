package main

import (
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/Lagrange-Labs/Lagrange-Node/config"
	"github.com/Lagrange-Labs/Lagrange-Node/node"

	"github.com/urfave/cli/v2"
)

// accounts "github.com/ethereum/go-ethereum/accounts"
// common "github.com/ethereum/go-ethereum/common"

var LOG_LEVEL int

var (
	configFileFlag = cli.StringFlag{
		Name:     config.FlagCfg,
		Aliases:  []string{"c"},
		Usage:    "Configuration `FILE`",
		Required: false,
	}
)

// Placeholder - Track staking listening to contract via rpc
var STAKE_STATE []string

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
			Usage:   "Run the lagrange-node",
			Action:  start,
			Flags:   flags,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

}

func start(ctx *cli.Context) error {
	lnode := node.NewLagrangeNode()

	cfg, err := config.Load(ctx)
	if err != nil {
		return err
	}

	logLevel := cfg.Node.LogLevel
	LOG_LEVEL = logLevel

	// Placeholder - Return first Hardhat private key for now
	PRIVATE_KEY_STRING := "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
	ADDRESS := "0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"

	lnode.SetWalletPath("./wallets")
	lnode.SetAddress(ADDRESS)
	lnode.LoadAccount()

	if !lnode.HasAccount(ADDRESS) {
		lnode.GenerateAccountFromPrivateKeyString(PRIVATE_KEY_STRING)
	} else {
		lnode.LoadAccount()
	}

	lnode.Start(&cfg.Node)

	return nil
}

func versionCmd(*cli.Context) error {
	w := os.Stdout
	fmt.Fprintf(w, "Version:      %s\n", "v0.1.0")
	fmt.Fprintf(w, "Go version:   %s\n", runtime.Version())
	fmt.Fprintf(w, "OS/Arch:      %s/%s\n", runtime.GOOS, runtime.GOARCH)
	return nil
}

package main

import (
	"log"
	"os"

	"github.com/Lagrange-Labs/lagrange-node/migrations"
	"github.com/urfave/cli/v2"
)

var (
	stepFlag = cli.UintFlag{
		Name:     "step",
		Aliases:  []string{"s"},
		Usage:    "Migration Step",
		Required: false,
	}
	uriFlag = cli.StringFlag{
		Name:     "uri",
		Aliases:  []string{"u"},
		Usage:    "Database URI",
		Required: true,
	}
)

func main() {
	app := cli.NewApp()
	app.Name = "Migration Tool"

	flags := []cli.Flag{&stepFlag, &uriFlag}
	app.Commands = []*cli.Command{
		{
			Name:    "up",
			Aliases: []string{},
			Usage:   "Run the migrations up to the given version",
			Action:  migrateUp,
			Flags:   flags,
		},
		{
			Name:    "down",
			Aliases: []string{},
			Usage:   "Run the migrations up to the given version",
			Action:  migrateDown,
			Flags:   flags,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func migrateUp(ctx *cli.Context) error {
	uri := ctx.String("uri")
	step := ctx.Uint("step")

	mm := migrations.RegisterDB(uri)
	return mm.MigrateUp(uint32(step))
}

func migrateDown(ctx *cli.Context) error {
	uri := ctx.String("uri")
	step := ctx.Uint("step")

	mm := migrations.RegisterDB(uri)
	return mm.MigrateDown(uint32(step))
}

package main

import (
	"fmt"
	"os"

	logging "github.com/ipfs/go-log"
	"github.com/urfave/cli/v2"
)

var log = logging.Logger("processing")

func main() {
	local := []*cli.Command{
		initCmd,
		runCmd,
	}

	app := &cli.App{
		Name:    "Processing",
		Usage:   "Public API server for lotus GUI server",
		Version: "0.0.1",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "repo",
				EnvVars: []string{"LOTUS_PROCESSING"},
				Value:   "~/.lotus-processing", // TODO: Consider XDG_DATA_HOME
			},
		},

		Commands: local,
	}
	app.Setup()

	if err := app.Run(os.Args); err != nil {
		log.Warnf("%+v", err)
		return
	}
}

var initCmd = &cli.Command{
	Name:  "init",
	Usage: "init processing server node",
	Action: func(cctx *cli.Context) error {
		log.Info("initial processing server node")
		return nil
	},
}

var runCmd = &cli.Command{
	Name:  "run",
	Usage: "start processing server",
	Action: func(cctx *cli.Context) error {
		fmt.Print("Processing server is running...")
		return nil
	},
}

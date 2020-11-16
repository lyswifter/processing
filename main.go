package main

import (
	"os"

	logging "github.com/ipfs/go-log"
	"github.com/urfave/cli/v2"

	"github.com/lyswifter/processing/cmd"
)

var log = logging.Logger("processing")

func main() {
	local := []*cli.Command{
		cmd.InitCmd,
		cmd.RunCmd,
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

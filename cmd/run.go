package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// RunCmd RunCmd
var RunCmd = &cli.Command{
	Name:  "run",
	Usage: "start processing server",
	Action: func(cctx *cli.Context) error {
		fmt.Print("Processing server is running...")
		return nil
	},
}

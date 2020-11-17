package cmd

import (
	"fmt"

	logging "github.com/ipfs/go-log"
	"github.com/urfave/cli/v2"

	"github.com/lyswifter/processing/repo"
)

var log = logging.Logger("processing")

var repoPath = "~/.processing"

// InitCmd InitCmd
var InitCmd = &cli.Command{
	Name:  "init",
	Usage: "init processing server node",
	Action: func(cctx *cli.Context) error {

		r, err := repo.NewFS(repoPath)
		if err != nil {
			return err
		}

		isExist, err := r.Exists()
		if isExist {
			return nil
		}

		err = r.Init()
		if err != nil {
			return err
		}

		fmt.Printf("Initial processing node at %s successfully\n", r.Path())

		return nil
	},
}

package cmd

import (
	"github.com/qiniu/x/log"
	"github.com/urfave/cli/v2"

	"github.com/lyswifter/processing/repo"
)

var repoPath = "processing"

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

		log.Info("Initial processing server node successfully")

		return nil
	},
}

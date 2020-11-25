package cmd

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/ipfs/go-datastore"
	"github.com/urfave/cli/v2"

	"github.com/lyswifter/processing/db"
	handler "github.com/lyswifter/processing/handlers"
)

const (
	dealDsNamespace        = "deal"
	powerDsNamespace       = "power"
	slaveDsNamespace       = "slave"
	uploaderDsNamespace    = "uploader"
	windowDsNamespace      = "window"
	winningDsNamespace     = "winning"
	posterDsNamespace      = "poster"
	ospProviderDsNamespace = "osp-provider"
	ospworkerDsNamespace   = "osp-worker"
)

var router *gin.Engine

// RunCmd RunCmd
var RunCmd = &cli.Command{
	Name:  "run",
	Usage: "start processing server",
	Action: func(cctx *cli.Context) error {

		// Set Gin to production mode
		gin.SetMode(gin.ReleaseMode)

		// Set the router as the default one provided by Gin
		router = gin.Default()

		// Process the templates at the start so that they don't have to be loaded
		// from the disk again. This makes serving HTML pages very fast.
		// router.LoadHTMLGlob("template/*")

		// Initialize the routes
		initializeRoutes()

		err := initDs()
		if err != nil {
			fmt.Printf("initDs: %s", err.Error())
			return err
		}

		fmt.Println("Processing server is running...")

		// Start serving the application
		router.Run(":9090")
		return nil
	},
}

func initializeRoutes() {
	router.GET("/", handler.HandleIndex)

	dealRoutes := router.Group("/deal")
	{
		dealRoutes.GET("/send", handler.HandleSendDeal)
	}

	// send
	powerRoutes := router.Group("/pow")
	{
		powerRoutes.GET("/send", handler.HandleSendPower)
	}
}

func initDs() error {
	ds, err := db.OpenDs(repoPath, datastore.NewKey(dealDsNamespace).String())
	if err != nil {
		return err
	}
	db.DealDs = ds

	ds, err = db.OpenDs(repoPath, datastore.NewKey(powerDsNamespace).String())
	if err != nil {
		return err
	}
	db.PowerDs = ds

	ds, err = db.OpenDs(repoPath, datastore.NewKey(windowDsNamespace).String())
	if err != nil {
		return err
	}
	db.WindowDs = ds

	ds, err = db.OpenDs(repoPath, datastore.NewKey(winningDsNamespace).String())
	if err != nil {
		return err
	}
	db.WinningDs = ds

	ds, err = db.OpenDs(repoPath, datastore.NewKey(slaveDsNamespace).String())
	if err != nil {
		return err
	}
	db.SlaveDs = ds

	//...

	return nil
}

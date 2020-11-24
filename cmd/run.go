package cmd

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/urfave/cli/v2"

	handler "github.com/lyswifter/processing/handlers"
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

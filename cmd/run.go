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
		// router.LoadHTMLGlob("templates/*")

		// Initialize the routes
		initializeSendPowerRoutes()

		fmt.Print("Processing server is running...")

		// Start serving the application
		router.Run(":9090")
		return nil
	},
}

func initializeSendPowerRoutes() {
	router.GET("/", handler.ShowIndexPage)

	userRoutes := router.Group("/pow")
	{
		userRoutes.GET("send", handler.HandleSendPower)
	}
}

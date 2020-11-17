package handler

import "github.com/gin-gonic/gin"

func ShowPowerPage(c *gin.Context) {
	// Call the render function with the name of the template to render
	render(c, gin.H{
		"title": "Power",
	}, "power.html")
}

func HandleSendPower(c *gin.Context) {
	render(c, gin.H{
		"title": "sendPower",
	}, "send.html")
}

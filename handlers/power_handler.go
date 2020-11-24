package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HandleSendPower HandleSendPower
func HandleSendPower(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "HandleSendPower"})
}

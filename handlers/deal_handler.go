package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HandleSendDeal HandleSendDeal
func HandleSendDeal(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "HandleSendDeal"})
}

package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// HandleIndex HandleIndex
func HandleIndex(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "HandleIndex", "code": "0", "msg": ""})
}

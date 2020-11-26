package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lyswifter/processing/model"
)

// HandleRecordEvent HandleRecordEvent
func HandleRecordEvent(c *gin.Context) {
	var evt model.Event

	err := c.ShouldBindJSON(&evt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "json invaild"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "ok"})
}

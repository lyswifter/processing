package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lyswifter/processing/db"
	"github.com/lyswifter/processing/model"
)

// HandleSectorRestart HandleSectorRestart
func HandleSectorRestart(c *gin.Context) {
	r, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	var info model.RestartUpInfo
	err = json.Unmarshal(r, &info)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	fmt.Printf("RestartUpInfo: %+v\n", info)

	st := model.SealingStateEvt{
		TimeStamp:    info.LastTime,
		SectorNumber: info.SectorNumber,
		SectorType:   info.SectorType,
		From:         "",
		After:        info.LastState,
		Error:        "",
	}

	db.SectorStore.Incoming <- &st

	fmt.Println("push sector ok")
	c.JSON(http.StatusOK, gin.H{"data": "HandleSectorRestart"})
}

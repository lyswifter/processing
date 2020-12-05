package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/gin-gonic/gin"
	"github.com/lyswifter/processing/db"
	"github.com/lyswifter/processing/model"
)

// HandleSectorInfo HandleSectorInfo
func HandleSectorInfo(c *gin.Context) {
	r, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	fmt.Printf("SealingStateEvt fir: %s\n", string(r))

	var evt model.Event
	err = json.Unmarshal(r, &evt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	item := evt.Data.(map[string]interface{})
	st := model.SealingStateEvt{
		TimeStamp:    evt.Timestamp.Unix(),
		SectorNumber: abi.SectorNumber(item["SectorNumber"].(float64)),
		SectorType:   abi.RegisteredSealProof(item["SectorType"].(float64)),
		From:         model.SectorState(item["From"].(string)),
		After:        model.SectorState(item["After"].(string)),
		Error:        item["Error"].(string),
	}

	db.SectorStore.Incoming <- &st

	fmt.Println("push sector ok")
	c.JSON(http.StatusOK, gin.H{"data": "HandleSectorInfo"})
}

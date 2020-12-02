package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"reflect"

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

	fmt.Printf("Event: %+v", evt)

	ty := reflect.TypeOf(evt.Data)
	va := reflect.ValueOf(evt.Data)

	fmt.Printf("type: %v value: %v\n", ty, va)

	item := evt.Data.(map[string]interface{})

	sectorNumberT := reflect.TypeOf(item["SectorNumber"])
	SectorTypeT := reflect.TypeOf(item["SectorType"])
	FromT := reflect.TypeOf(item["From"])
	AfterT := reflect.TypeOf(item["After"])
	errorT := reflect.TypeOf(item["Error"])
	fmt.Printf("sectorNumberT: %v SectorTypeT: %v FromT: %v AfterT: %v errorT: %v\n", sectorNumberT, SectorTypeT, FromT, AfterT, errorT)

	st := model.SealingStateEvt{
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

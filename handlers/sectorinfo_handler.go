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

// HandleSectorInfo HandleSectorInfo
func HandleSectorInfo(c *gin.Context) {

	r, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	log.Infof("SealingStateEvt: %s", string(r))

	var sinfo model.SealingStateEvt
	err = json.Unmarshal(r, &sinfo)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	// var evt *model.SealingStateInfoEvt
	// err = json.Unmarshal(r, evt)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	// 	return
	// }

	// var bSinfo model.SectorInfo
	// if err := cborutil.ReadCborRPC(bytes.NewBuffer(evt.BInfo), &bSinfo); err != nil {
	// 	if err.Error() != "EOF" {
	// 		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	// 		return
	// 	}
	// }

	// fmt.Printf("SectorBInfo: %+v", bSinfo)

	// var aSinfo model.SectorInfo
	// if err := cborutil.ReadCborRPC(bytes.NewBuffer(evt.AInfo), &bSinfo); err != nil {
	// 	if err.Error() != "EOF" {
	// 		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	// 		return
	// 	}
	// }

	// fmt.Printf("SectorAInfo: %+v", aSinfo)

	// var extInfo model.SectorInfoExt
	// if err := cborutil.ReadCborRPC(bytes.NewBuffer(evt.ExtInfo), &extInfo); err != nil {
	// 	if err.Error() != "EOF" {
	// 		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	// 		return
	// 	}
	// }

	// fmt.Printf("SectorInfoExt: %+v", extInfo)

	db.SectorStore.Incoming <- &sinfo

	fmt.Println("push sector ok")

	c.JSON(http.StatusOK, gin.H{"data": "HandleSectorInfo"})
}

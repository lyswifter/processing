package handler

import (
	"fmt"
	"io/ioutil"
	"net/http"

	cborutil "github.com/filecoin-project/go-cbor-util"
	"github.com/gin-gonic/gin"
	"github.com/ipfs/go-datastore"
	"github.com/lyswifter/processing/db"
	"github.com/lyswifter/processing/model"
)

// HandleSectorInfo HandleSectorInfo
func HandleSectorInfo(c *gin.Context) {
	var sinfo model.SectorInfo
	if err := cborutil.ReadCborRPC(c.Request.Body, &sinfo); err != nil {
		if err.Error() != "EOF" {
			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}
	}

	fmt.Printf("SectorInfo: %+v", sinfo)

	r, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	// save or update sectorinfo in database
	err = db.DealDs.Put(datastore.NewKey(sinfo.SectorNumber.String()), r)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	fmt.Println("Save ok")

	c.JSON(http.StatusOK, gin.H{"data": "HandleSectorInfo"})
}

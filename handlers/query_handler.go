package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/lyswifter/processing/db"
	"github.com/lyswifter/processing/model"
	"golang.org/x/xerrors"
)

// HandleQuerySector HandleQuerySector
func HandleQuerySector(c *gin.Context) {
	qStates := c.QueryArray("name")
	offset := c.DefaultQuery("offset", "0")
	size := c.DefaultQuery("size", "50")

	fmt.Printf("HandleQuerySector states: %v offset: %s size: %s\n", qStates, offset, size)

	var sInfos []*model.SealingStateEvt
	var allKeys []string
	for _, ss := range qStates {
		keys, ret, err := querySpecify(model.SectorState(ss), offset, size)
		if err != nil {
			fmt.Printf("querySpecify: %s\n", err.Error())
			continue
		}

		allKeys = append(allKeys, keys...)
		sInfos = append(sInfos, ret...)
	}

	if len(allKeys) == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "allkeys is empty"})
		return
	}

	if len(sInfos) == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "result is empty"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"total":   len(allKeys),
		"offset":  offset,
		"size":    size,
		"data":    sInfos,
		"message": "ok",
		"code":    0,
	})
}

func querySpecify(stat model.SectorState, offset string, pSize string) ([]string, []*model.SealingStateEvt, error) {
	off, err := strconv.ParseInt(offset, 10, 64)
	if err != nil {
		return nil, nil, err
	}

	size, err := strconv.ParseInt(pSize, 10, 64)
	if err != nil {
		return nil, nil, err
	}

	var sInfo []*model.SealingStateEvt
	var allKeys []string
	switch stat {
	case model.Empty:
		err := db.SectorStore.Empty.ListBy(int(off), int(size), &sInfo)
		if err != nil {
			return nil, nil, err
		}

		keyList, err := db.SectorStore.Empty.ListKey()
		if err != nil {
			return nil, nil, err
		}
		allKeys = keyList
	case model.WaitDeals:
		err := db.SectorStore.WaitDeals.ListBy(int(off), int(size), &sInfo)
		if err != nil {
			return nil, nil, err
		}

		keyList, err := db.SectorStore.WaitDeals.ListKey()
		if err != nil {
			return nil, nil, err
		}
		allKeys = keyList
	case model.Packing:
		err := db.SectorStore.Packing.ListBy(int(off), int(size), &sInfo)
		if err != nil {
			return nil, nil, err
		}

		keyList, err := db.SectorStore.Packing.ListKey()
		if err != nil {
			return nil, nil, err
		}
		allKeys = keyList
	case model.GetTicket:
		err := db.SectorStore.GetTicket.ListBy(int(off), int(size), &sInfo)
		if err != nil {
			return nil, nil, err
		}

		keyList, err := db.SectorStore.GetTicket.ListKey()
		if err != nil {
			return nil, nil, err
		}
		allKeys = keyList
	case model.PreCommit1:
		err := db.SectorStore.PreCommit1.ListBy(int(off), int(size), &sInfo)
		if err != nil {
			return nil, nil, err
		}

		keyList, err := db.SectorStore.PreCommit1.ListKey()
		if err != nil {
			return nil, nil, err
		}
		allKeys = keyList
	case model.PreCommit2:
		err := db.SectorStore.PreCommit2.ListBy(int(off), int(size), &sInfo)
		if err != nil {
			return nil, nil, err
		}

		keyList, err := db.SectorStore.PreCommit2.ListKey()
		if err != nil {
			return nil, nil, err
		}
		allKeys = keyList
	case model.PreCommitting:
		err := db.SectorStore.PreCommitting.ListBy(int(off), int(size), &sInfo)
		if err != nil {
			return nil, nil, err
		}

		keyList, err := db.SectorStore.PreCommitting.ListKey()
		if err != nil {
			return nil, nil, err
		}
		allKeys = keyList
	case model.PreCommitWait:
		err := db.SectorStore.PreCommitWait.ListBy(int(off), int(size), &sInfo)
		if err != nil {
			return nil, nil, err
		}

		keyList, err := db.SectorStore.PreCommitWait.ListKey()
		if err != nil {
			return nil, nil, err
		}
		allKeys = keyList
	case model.WaitSeed:
		err := db.SectorStore.WaitSeed.ListBy(int(off), int(size), &sInfo)
		if err != nil {
			return nil, nil, err
		}

		keyList, err := db.SectorStore.WaitSeed.ListKey()
		if err != nil {
			return nil, nil, err
		}
		allKeys = keyList
	case model.Committing:
		err := db.SectorStore.Committing.ListBy(int(off), int(size), &sInfo)
		if err != nil {
			return nil, nil, err
		}

		keyList, err := db.SectorStore.Committing.ListKey()
		if err != nil {
			return nil, nil, err
		}
		allKeys = keyList
	case model.CommitWait:
		err := db.SectorStore.CommitWait.ListBy(int(off), int(size), &sInfo)
		if err != nil {
			return nil, nil, err
		}

		keyList, err := db.SectorStore.CommitWait.ListKey()
		if err != nil {
			return nil, nil, err
		}
		allKeys = keyList
	case model.FinalizeSector:
		err := db.SectorStore.FinalizeSector.ListBy(int(off), int(size), &sInfo)
		if err != nil {
			return nil, nil, err
		}

		keyList, err := db.SectorStore.FinalizeSector.ListKey()
		if err != nil {
			return nil, nil, err
		}
		allKeys = keyList
	case model.Proving:
		err := db.SectorStore.Proving.ListBy(int(off), int(size), &sInfo)
		if err != nil {
			return nil, nil, err
		}

		keyList, err := db.SectorStore.Proving.ListKey()
		if err != nil {
			return nil, nil, err
		}
		allKeys = keyList
	case model.FailedUnrecoverable:
		err := db.SectorStore.FailedUnrecoverable.ListBy(int(off), int(size), &sInfo)
		if err != nil {
			return nil, nil, err
		}

		keyList, err := db.SectorStore.FailedUnrecoverable.ListKey()
		if err != nil {
			return nil, nil, err
		}
		allKeys = keyList
	case model.SealPreCommit1Failed:
		err := db.SectorStore.SealPreCommit1Failed.ListBy(int(off), int(size), &sInfo)
		if err != nil {
			return nil, nil, err
		}

		keyList, err := db.SectorStore.SealPreCommit1Failed.ListKey()
		if err != nil {
			return nil, nil, err
		}
		allKeys = keyList
	case model.SealPreCommit2Failed:
		err := db.SectorStore.SealPreCommit2Failed.ListBy(int(off), int(size), &sInfo)
		if err != nil {
			return nil, nil, err
		}

		keyList, err := db.SectorStore.SealPreCommit2Failed.ListKey()
		if err != nil {
			return nil, nil, err
		}
		allKeys = keyList
	case model.PreCommitFailed:
		err := db.SectorStore.PreCommitFailed.ListBy(int(off), int(size), &sInfo)
		if err != nil {
			return nil, nil, err
		}

		keyList, err := db.SectorStore.PreCommitFailed.ListKey()
		if err != nil {
			return nil, nil, err
		}
		allKeys = keyList
	case model.ComputeProofFailed:
		err := db.SectorStore.ComputeProofFailed.ListBy(int(off), int(size), &sInfo)
		if err != nil {
			return nil, nil, err
		}

		keyList, err := db.SectorStore.ComputeProofFailed.ListKey()
		if err != nil {
			return nil, nil, err
		}
		allKeys = keyList
	case model.CommitFailed:
		err := db.SectorStore.CommitFailed.ListBy(int(off), int(size), &sInfo)
		if err != nil {
			return nil, nil, err
		}

		keyList, err := db.SectorStore.CommitFailed.ListKey()
		if err != nil {
			return nil, nil, err
		}
		allKeys = keyList
	case model.PackingFailed:
		err := db.SectorStore.PackingFailed.ListBy(int(off), int(size), &sInfo)
		if err != nil {
			return nil, nil, err
		}

		keyList, err := db.SectorStore.PackingFailed.ListKey()
		if err != nil {
			return nil, nil, err
		}
		allKeys = keyList
	case model.FinalizeFailed:
		err := db.SectorStore.FinalizeFailed.ListBy(int(off), int(size), &sInfo)
		if err != nil {
			return nil, nil, err
		}

		keyList, err := db.SectorStore.FinalizeFailed.ListKey()
		if err != nil {
			return nil, nil, err
		}
		allKeys = keyList
	case model.DealsExpired:
		err := db.SectorStore.DealsExpired.ListBy(int(off), int(size), &sInfo)
		if err != nil {
			return nil, nil, err
		}

		keyList, err := db.SectorStore.DealsExpired.ListKey()
		if err != nil {
			return nil, nil, err
		}
		allKeys = keyList
	case model.RecoverDealIDs:
		err := db.SectorStore.RecoverDealIDs.ListBy(int(off), int(size), &sInfo)
		if err != nil {
			return nil, nil, err
		}

		keyList, err := db.SectorStore.RecoverDealIDs.ListKey()
		if err != nil {
			return nil, nil, err
		}
		allKeys = keyList
	case model.Faulty:
		err := db.SectorStore.Faulty.ListBy(int(off), int(size), &sInfo)
		if err != nil {
			return nil, nil, err
		}

		keyList, err := db.SectorStore.Faulty.ListKey()
		if err != nil {
			return nil, nil, err
		}
		allKeys = keyList
	case model.FaultReported:
		err := db.SectorStore.FaultReported.ListBy(int(off), int(size), &sInfo)
		if err != nil {
			return nil, nil, err
		}

		keyList, err := db.SectorStore.FaultReported.ListKey()
		if err != nil {
			return nil, nil, err
		}
		allKeys = keyList
	case model.FaultedFinal:
		err := db.SectorStore.FaultedFinal.ListBy(int(off), int(size), &sInfo)
		if err != nil {
			return nil, nil, err
		}

		keyList, err := db.SectorStore.FaultedFinal.ListKey()
		if err != nil {
			return nil, nil, err
		}
		allKeys = keyList
	case model.Removing:
		err := db.SectorStore.Removing.ListBy(int(off), int(size), &sInfo)
		if err != nil {
			return nil, nil, err
		}

		keyList, err := db.SectorStore.Removing.ListKey()
		if err != nil {
			return nil, nil, err
		}
		allKeys = keyList
	case model.RemoveFailed:
		err := db.SectorStore.RemoveFailed.ListBy(int(off), int(size), &sInfo)
		if err != nil {
			return nil, nil, err
		}

		keyList, err := db.SectorStore.RemoveFailed.ListKey()
		if err != nil {
			return nil, nil, err
		}
		allKeys = keyList
	case model.Removed:
		err := db.SectorStore.Removed.ListBy(int(off), int(size), &sInfo)
		if err != nil {
			return nil, nil, err
		}

		keyList, err := db.SectorStore.Removed.ListKey()
		if err != nil {
			return nil, nil, err
		}
		allKeys = keyList
	case model.SlaveCommittingFailed:
		err := db.SectorStore.SlaveCommittingFailed.ListBy(int(off), int(size), &sInfo)
		if err != nil {
			return nil, nil, err
		}

		keyList, err := db.SectorStore.SlaveCommittingFailed.ListKey()
		if err != nil {
			return nil, nil, err
		}
		allKeys = keyList
	case model.SlaveCommitFailed:
		err := db.SectorStore.SlaveCommitFailed.ListBy(int(off), int(size), &sInfo)
		if err != nil {
			return nil, nil, err
		}

		keyList, err := db.SectorStore.SlaveCommitFailed.ListKey()
		if err != nil {
			return nil, nil, err
		}
		allKeys = keyList
	case model.PreSealed:
		err := db.SectorStore.PreSealed.ListBy(int(off), int(size), &sInfo)
		if err != nil {
			return nil, nil, err
		}

		keyList, err := db.SectorStore.PreSealed.ListKey()
		if err != nil {
			return nil, nil, err
		}
		allKeys = keyList
	case model.ReSealing:
		err := db.SectorStore.ReSealing.ListBy(int(off), int(size), &sInfo)
		if err != nil {
			return nil, nil, err
		}

		keyList, err := db.SectorStore.ReSealing.ListKey()
		if err != nil {
			return nil, nil, err
		}
		allKeys = keyList
	case model.ReSealedFailed:
		err := db.SectorStore.ReSealedFailed.ListBy(int(off), int(size), &sInfo)
		if err != nil {
			return nil, nil, err
		}

		keyList, err := db.SectorStore.ReSealedFailed.ListKey()
		if err != nil {
			return nil, nil, err
		}
		allKeys = keyList
	case model.SlaveFinal:
		err := db.SectorStore.SlaveFinal.ListBy(int(off), int(size), &sInfo)
		if err != nil {
			return nil, nil, err
		}

		keyList, err := db.SectorStore.SlaveFinal.ListKey()
		if err != nil {
			return nil, nil, err
		}
		allKeys = keyList
	case model.SectorDownloading:
		err := db.SectorStore.SectorDownloading.ListBy(int(off), int(size), &sInfo)
		if err != nil {
			return nil, nil, err
		}

		keyList, err := db.SectorStore.SectorDownloading.ListKey()
		if err != nil {
			return nil, nil, err
		}
		allKeys = keyList
	case model.DownloadFailed:
		err := db.SectorStore.DownloadFailed.ListBy(int(off), int(size), &sInfo)
		if err != nil {
			return nil, nil, err
		}

		keyList, err := db.SectorStore.DownloadFailed.ListKey()
		if err != nil {
			return nil, nil, err
		}
		allKeys = keyList
	case model.ProxyCommitted:
		err := db.SectorStore.ProxyCommitted.ListBy(int(off), int(size), &sInfo)
		if err != nil {
			return nil, nil, err
		}

		keyList, err := db.SectorStore.ProxyCommitted.ListKey()
		if err != nil {
			return nil, nil, err
		}
		allKeys = keyList
	case model.ProxyFetchingSector:
		err := db.SectorStore.ProxyFetchingSector.ListBy(int(off), int(size), &sInfo)
		if err != nil {
			return nil, nil, err
		}

		keyList, err := db.SectorStore.ProxyFetchingSector.ListKey()
		if err != nil {
			return nil, nil, err
		}
		allKeys = keyList
	case model.OrphanFailed:
		err := db.SectorStore.OrphanFailed.ListBy(int(off), int(size), &sInfo)
		if err != nil {
			return nil, nil, err
		}

		keyList, err := db.SectorStore.OrphanFailed.ListKey()
		if err != nil {
			return nil, nil, err
		}
		allKeys = keyList
	case model.OspCommittingRequest:
		err := db.SectorStore.OspCommittingRequest.ListBy(int(off), int(size), &sInfo)
		if err != nil {
			return nil, nil, err
		}

		keyList, err := db.SectorStore.OspCommittingRequest.ListKey()
		if err != nil {
			return nil, nil, err
		}
		allKeys = keyList
	case model.OspCommittingRequested:
		err := db.SectorStore.OspCommittingRequested.ListBy(int(off), int(size), &sInfo)
		if err != nil {
			return nil, nil, err
		}

		keyList, err := db.SectorStore.OspCommittingRequested.ListKey()
		if err != nil {
			return nil, nil, err
		}
		allKeys = keyList
	case model.OspCommitFailed:
		err := db.SectorStore.OspCommitFailed.ListBy(int(off), int(size), &sInfo)
		if err != nil {
			return nil, nil, err
		}

		keyList, err := db.SectorStore.OspCommitFailed.ListKey()
		if err != nil {
			return nil, nil, err
		}
		allKeys = keyList
	case model.OfflineDealsReady:
		err := db.SectorStore.OfflineDealsReady.ListBy(int(off), int(size), &sInfo)
		if err != nil {
			return nil, nil, err
		}

		keyList, err := db.SectorStore.OfflineDealsReady.ListKey()
		if err != nil {
			return nil, nil, err
		}
		allKeys = keyList
	case model.OfflineDealsReadyFailed:
		err := db.SectorStore.OfflineDealsReadyFailed.ListBy(int(off), int(size), &sInfo)
		if err != nil {
			return nil, nil, err
		}

		keyList, err := db.SectorStore.OfflineDealsReadyFailed.ListKey()
		if err != nil {
			return nil, nil, err
		}
		allKeys = keyList
	default:
		return allKeys, sInfo, xerrors.Errorf("no records")
	}

	for _, info := range sInfo {
		delta := time.Since(time.Unix(info.TimeStamp, 0))
		secs := int64(delta.Seconds())
		color := model.GetStateColor(info.After, secs)

		info.Interval = fmt.Sprintf("%ss", strings.Split(delta.String(), ".")[0])
		info.Sec = secs
		info.Color = color
	}

	return allKeys, sInfo, nil
}

package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/lyswifter/processing/db"
	"github.com/lyswifter/processing/model"
	"golang.org/x/xerrors"
)

// HandleQuerySector HandleQuerySector
func HandleQuerySector(c *gin.Context) {
	qStates := c.QueryArray("name")
	fmt.Printf("HandleQuerySector: %v\n", qStates)

	var sInfos []model.SealingStateEvt
	for _, ss := range qStates {
		ret, err := querySpecify(model.SectorState(ss))
		if err != nil {
			fmt.Printf("querySpecify: %s\n", err.Error())
			continue
		}

		sInfos = append(sInfos, ret...)
	}

	if len(sInfos) == 0 {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "result is empty"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": sInfos})
}

func querySpecify(stat model.SectorState) ([]model.SealingStateEvt, error) {
	var sInfo []model.SealingStateEvt
	switch stat {
	case model.Empty:
		err := db.SectorStore.Empty.List(&sInfo)
		if err != nil {
			return nil, err
		}
	case model.WaitDeals:
		err := db.SectorStore.WaitDeals.List(&sInfo)
		if err != nil {
			return nil, err
		}
	case model.Packing:
		err := db.SectorStore.Packing.List(&sInfo)
		if err != nil {
			return nil, err
		}
	case model.GetTicket:
		err := db.SectorStore.GetTicket.List(&sInfo)
		if err != nil {
			return nil, err
		}
	case model.PreCommit1:
		err := db.SectorStore.PreCommit1.List(&sInfo)
		if err != nil {
			return nil, err
		}
	case model.PreCommit2:
		err := db.SectorStore.PreCommit2.List(&sInfo)
		if err != nil {
			return nil, err
		}
	case model.PreCommitting:
		err := db.SectorStore.PreCommitting.List(&sInfo)
		if err != nil {
			return nil, err
		}
	case model.PreCommitWait:
		err := db.SectorStore.PreCommitWait.List(&sInfo)
		if err != nil {
			return nil, err
		}
	case model.WaitSeed:
		err := db.SectorStore.WaitSeed.List(&sInfo)
		if err != nil {
			return nil, err
		}
	case model.Committing:
		err := db.SectorStore.Committing.List(&sInfo)
		if err != nil {
			return nil, err
		}
	case model.CommitWait:
		err := db.SectorStore.CommitWait.List(&sInfo)
		if err != nil {
			return nil, err
		}
	case model.FinalizeSector:
		err := db.SectorStore.FinalizeSector.List(&sInfo)
		if err != nil {
			return nil, err
		}
	case model.Proving:
		err := db.SectorStore.Proving.List(&sInfo)
		if err != nil {
			return nil, err
		}
	case model.FailedUnrecoverable:
		err := db.SectorStore.FailedUnrecoverable.List(&sInfo)
		if err != nil {
			return nil, err
		}
	case model.SealPreCommit1Failed:
		err := db.SectorStore.SealPreCommit1Failed.List(&sInfo)
		if err != nil {
			return nil, err
		}
	case model.SealPreCommit2Failed:
		err := db.SectorStore.SealPreCommit2Failed.List(&sInfo)
		if err != nil {
			return nil, err
		}
	case model.PreCommitFailed:
		err := db.SectorStore.PreCommitFailed.List(&sInfo)
		if err != nil {
			return nil, err
		}
	case model.ComputeProofFailed:
		err := db.SectorStore.ComputeProofFailed.List(&sInfo)
		if err != nil {
			return nil, err
		}
	case model.CommitFailed:
		err := db.SectorStore.CommitFailed.List(&sInfo)
		if err != nil {
			return nil, err
		}
	case model.PackingFailed:
		err := db.SectorStore.PackingFailed.List(&sInfo)
		if err != nil {
			return nil, err
		}
	case model.FinalizeFailed:
		err := db.SectorStore.FinalizeFailed.List(&sInfo)
		if err != nil {
			return nil, err
		}
	case model.DealsExpired:
		err := db.SectorStore.DealsExpired.List(&sInfo)
		if err != nil {
			return nil, err
		}
	case model.RecoverDealIDs:
		err := db.SectorStore.RecoverDealIDs.List(&sInfo)
		if err != nil {
			return nil, err
		}
	case model.Faulty:
		err := db.SectorStore.Faulty.List(&sInfo)
		if err != nil {
			return nil, err
		}
	case model.FaultReported:
		err := db.SectorStore.FaultReported.List(&sInfo)
		if err != nil {
			return nil, err
		}
	case model.FaultedFinal:
		err := db.SectorStore.FaultedFinal.List(&sInfo)
		if err != nil {
			return nil, err
		}
	case model.Removing:
		err := db.SectorStore.Removing.List(&sInfo)
		if err != nil {
			return nil, err
		}
	case model.RemoveFailed:
		err := db.SectorStore.RemoveFailed.List(&sInfo)
		if err != nil {
			return nil, err
		}
	case model.Removed:
		err := db.SectorStore.Removed.List(&sInfo)
		if err != nil {
			return nil, err
		}
	case model.SlaveCommittingFailed:
		err := db.SectorStore.SlaveCommittingFailed.List(&sInfo)
		if err != nil {
			return nil, err
		}
	case model.SlaveCommitFailed:
		err := db.SectorStore.SlaveCommitFailed.List(&sInfo)
		if err != nil {
			return nil, err
		}
	case model.PreSealed:
		err := db.SectorStore.PreSealed.List(&sInfo)
		if err != nil {
			return nil, err
		}
	case model.ReSealing:
		err := db.SectorStore.ReSealing.List(&sInfo)
		if err != nil {
			return nil, err
		}
	case model.ReSealedFailed:
		err := db.SectorStore.ReSealedFailed.List(&sInfo)
		if err != nil {
			return nil, err
		}
	case model.SlaveFinal:
		err := db.SectorStore.SlaveFinal.List(&sInfo)
		if err != nil {
			return nil, err
		}
	case model.SectorDownloading:
		err := db.SectorStore.SectorDownloading.List(&sInfo)
		if err != nil {
			return nil, err
		}
	case model.DownloadFailed:
		err := db.SectorStore.DownloadFailed.List(&sInfo)
		if err != nil {
			return nil, err
		}
	case model.ProxyCommitted:
		err := db.SectorStore.ProxyCommitted.List(&sInfo)
		if err != nil {
			return nil, err
		}
	case model.ProxyFetchingSector:
		err := db.SectorStore.ProxyFetchingSector.List(&sInfo)
		if err != nil {
			return nil, err
		}
	case model.OrphanFailed:
		err := db.SectorStore.OrphanFailed.List(&sInfo)
		if err != nil {
			return nil, err
		}
	case model.OspCommittingRequest:
		err := db.SectorStore.OspCommittingRequest.List(&sInfo)
		if err != nil {
			return nil, err
		}
	case model.OspCommittingRequested:
		err := db.SectorStore.OspCommittingRequested.List(&sInfo)
		if err != nil {
			return nil, err
		}
	case model.OspCommitFailed:
		err := db.SectorStore.OspCommitFailed.List(&sInfo)
		if err != nil {
			return nil, err
		}
	case model.OfflineDealsReady:
		err := db.SectorStore.OfflineDealsReady.List(&sInfo)
		if err != nil {
			return nil, err
		}
	case model.OfflineDealsReadyFailed:
		err := db.SectorStore.OfflineDealsReadyFailed.List(&sInfo)
		if err != nil {
			return nil, err
		}
	default:
		return sInfo, xerrors.Errorf("no records")
	}

	fmt.Printf("sInfo: %+v\n", sInfo)

	return sInfo, nil
}

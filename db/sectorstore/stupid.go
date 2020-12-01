package sectorstore

import "github.com/lyswifter/processing/model"

func (sl *SectorLifecycle) removesinfo(se *model.SealingStateEvt) error {
	log.Infof("remove sector info start: %s", se.SectorNumber.String())

	switch se.From {
	case model.Empty:
		err := sl.Empty.Get(se.SectorNumber).End()
		if err != nil {
			return err
		}
	case model.WaitDeals:
		err := sl.WaitDeals.Get(se.SectorNumber).End()
		if err != nil {
			return err
		}
	case model.Packing:
		err := sl.Packing.Get(se.SectorNumber).End()
		if err != nil {
			return err
		}
	case model.GetTicket:
		err := sl.GetTicket.Get(se.SectorNumber).End()
		if err != nil {
			return err
		}
	case model.PreCommit1:
		err := sl.PreCommit1.Get(se.SectorNumber).End()
		if err != nil {
			return err
		}
	case model.PreCommit2:
		err := sl.PreCommit2.Get(se.SectorNumber).End()
		if err != nil {
			return err
		}
	case model.PreCommitting:
		err := sl.PreCommitting.Get(se.SectorNumber).End()
		if err != nil {
			return err
		}
	case model.PreCommitWait:
		err := sl.PreCommitWait.Get(se.SectorNumber).End()
		if err != nil {
			return err
		}
	case model.WaitSeed:
		err := sl.WaitSeed.Get(se.SectorNumber).End()
		if err != nil {
			return err
		}
	case model.Committing:
		err := sl.Committing.Get(se.SectorNumber).End()
		if err != nil {
			return err
		}
	case model.SubmitCommit:
		err := sl.SubmitCommit.Get(se.SectorNumber).End()
		if err != nil {
			return err
		}
	case model.CommitWait:
		err := sl.CommitWait.Get(se.SectorNumber).End()
		if err != nil {
			return err
		}
	case model.FinalizeSector:
		err := sl.FinalizeSector.Get(se.SectorNumber).End()
		if err != nil {
			return err
		}
	case model.Proving:
		err := sl.Proving.Get(se.SectorNumber).End()
		if err != nil {
			return err
		}
	case model.FailedUnrecoverable:
		err := sl.FailedUnrecoverable.Get(se.SectorNumber).End()
		if err != nil {
			return err
		}
	case model.SealPreCommit1Failed:
		err := sl.SealPreCommit1Failed.Get(se.SectorNumber).End()
		if err != nil {
			return err
		}
	case model.SealPreCommit2Failed:
		err := sl.SealPreCommit2Failed.Get(se.SectorNumber).End()
		if err != nil {
			return err
		}
	case model.PreCommitFailed:
		err := sl.PreCommitFailed.Get(se.SectorNumber).End()
		if err != nil {
			return err
		}
	case model.ComputeProofFailed:
		err := sl.ComputeProofFailed.Get(se.SectorNumber).End()
		if err != nil {
			return err
		}
	case model.CommitFailed:
		err := sl.CommitFailed.Get(se.SectorNumber).End()
		if err != nil {
			return err
		}
	case model.PackingFailed:
		err := sl.PackingFailed.Get(se.SectorNumber).End()
		if err != nil {
			return err
		}
	case model.FinalizeFailed:
		err := sl.FinalizeFailed.Get(se.SectorNumber).End()
		if err != nil {
			return err
		}
	case model.DealsExpired:
		err := sl.DealsExpired.Get(se.SectorNumber).End()
		if err != nil {
			return err
		}
	case model.RecoverDealIDs:
		err := sl.RecoverDealIDs.Get(se.SectorNumber).End()
		if err != nil {
			return err
		}
	case model.Faulty:
		err := sl.Faulty.Get(se.SectorNumber).End()
		if err != nil {
			return err
		}
	case model.FaultReported:
		err := sl.FaultReported.Get(se.SectorNumber).End()
		if err != nil {
			return err
		}
	case model.FaultedFinal:
		err := sl.FaultedFinal.Get(se.SectorNumber).End()
		if err != nil {
			return err
		}
	case model.Removing:
		err := sl.Removing.Get(se.SectorNumber).End()
		if err != nil {
			return err
		}
	case model.RemoveFailed:
		err := sl.RemoveFailed.Get(se.SectorNumber).End()
		if err != nil {
			return err
		}
	case model.Removed:
		err := sl.Removed.Get(se.SectorNumber).End()
		if err != nil {
			return err
		}
	case model.SlaveCommittingFailed:
		err := sl.SlaveCommittingFailed.Get(se.SectorNumber).End()
		if err != nil {
			return err
		}
	case model.SlaveCommitFailed:
		err := sl.SlaveCommitFailed.Get(se.SectorNumber).End()
		if err != nil {
			return err
		}
	case model.PreSealed:
		err := sl.PreSealed.Get(se.SectorNumber).End()
		if err != nil {
			return err
		}
	case model.ReSealing:
		err := sl.ReSealing.Get(se.SectorNumber).End()
		if err != nil {
			return err
		}
	case model.ReSealedFailed:
		err := sl.ReSealedFailed.Get(se.SectorNumber).End()
		if err != nil {
			return err
		}
	case model.SlaveFinal:
		err := sl.SlaveFinal.Get(se.SectorNumber).End()
		if err != nil {
			return err
		}
	case model.SectorDownloading:
		err := sl.SectorDownloading.Get(se.SectorNumber).End()
		if err != nil {
			return err
		}
	case model.DownloadFailed:
		err := sl.DownloadFailed.Get(se.SectorNumber).End()
		if err != nil {
			return err
		}
	case model.ProxyCommitted:
		err := sl.ProxyCommitted.Get(se.SectorNumber).End()
		if err != nil {
			return err
		}
	case model.ProxyFetchingSector:
		err := sl.ProxyFetchingSector.Get(se.SectorNumber).End()
		if err != nil {
			return err
		}
	case model.OrphanFailed:
		err := sl.OrphanFailed.Get(se.SectorNumber).End()
		if err != nil {
			return err
		}
	case model.OspCommittingRequest:
		err := sl.OspCommittingRequest.Get(se.SectorNumber).End()
		if err != nil {
			return err
		}
	case model.OspCommittingRequested:
		err := sl.OspCommittingRequested.Get(se.SectorNumber).End()
		if err != nil {
			return err
		}
	case model.OspCommitFailed:
		err := sl.OspCommitFailed.Get(se.SectorNumber).End()
		if err != nil {
			return err
		}
	case model.OfflineDealsReady:
		err := sl.OfflineDealsReady.Get(se.SectorNumber).End()
		if err != nil {
			return err
		}
	case model.OfflineDealsReadyFailed:
		err := sl.OfflineDealsReadyFailed.Get(se.SectorNumber).End()
		if err != nil {
			return err
		}
	default:
		return nil
	}

	return nil
}

func (sl *SectorLifecycle) putsinfo(se *model.SealingStateEvt) error {
	log.Infof("put sector info start: %s", se.SectorNumber.String())

	switch se.After {
	case model.Empty:
		err := sl.Empty.Begin(se.SectorNumber, se)
		if err != nil {
			return err
		}
	case model.WaitDeals:
		err := sl.WaitDeals.Begin(se.SectorNumber, se)
		if err != nil {
			return err
		}
	case model.Packing:
		err := sl.Packing.Begin(se.SectorNumber, se)
		if err != nil {
			return err
		}
	case model.GetTicket:
		err := sl.GetTicket.Begin(se.SectorNumber, se)
		if err != nil {
			return err
		}
	case model.PreCommit1:
		err := sl.PreCommit1.Begin(se.SectorNumber, se)
		if err != nil {
			return err
		}
	case model.PreCommit2:
		err := sl.PreCommit2.Begin(se.SectorNumber, se)
		if err != nil {
			return err
		}
	case model.PreCommitting:
		err := sl.PreCommitting.Begin(se.SectorNumber, se)
		if err != nil {
			return err
		}
	case model.PreCommitWait:
		err := sl.PreCommitWait.Begin(se.SectorNumber, se)
		if err != nil {
			return err
		}
	case model.WaitSeed:
		err := sl.WaitSeed.Begin(se.SectorNumber, se)
		if err != nil {
			return err
		}
	case model.Committing:
		err := sl.Committing.Begin(se.SectorNumber, se)
		if err != nil {
			return err
		}
	case model.SubmitCommit:
		err := sl.SubmitCommit.Begin(se.SectorNumber, se)
		if err != nil {
			return err
		}
	case model.CommitWait:
		err := sl.CommitWait.Begin(se.SectorNumber, se)
		if err != nil {
			return err
		}
	case model.FinalizeSector:
		err := sl.FinalizeSector.Begin(se.SectorNumber, se)
		if err != nil {
			return err
		}
	case model.Proving:
		err := sl.Proving.Begin(se.SectorNumber, se)
		if err != nil {
			return err
		}
	case model.FailedUnrecoverable:
		err := sl.FailedUnrecoverable.Begin(se.SectorNumber, se)
		if err != nil {
			return err
		}
	case model.SealPreCommit1Failed:
		err := sl.SealPreCommit1Failed.Begin(se.SectorNumber, se)
		if err != nil {
			return err
		}
	case model.SealPreCommit2Failed:
		err := sl.SealPreCommit2Failed.Begin(se.SectorNumber, se)
		if err != nil {
			return err
		}
	case model.PreCommitFailed:
		err := sl.PreCommitFailed.Begin(se.SectorNumber, se)
		if err != nil {
			return err
		}
	case model.ComputeProofFailed:
		err := sl.ComputeProofFailed.Begin(se.SectorNumber, se)
		if err != nil {
			return err
		}
	case model.CommitFailed:
		err := sl.CommitFailed.Begin(se.SectorNumber, se)
		if err != nil {
			return err
		}
	case model.PackingFailed:
		err := sl.PackingFailed.Begin(se.SectorNumber, se)
		if err != nil {
			return err
		}
	case model.FinalizeFailed:
		err := sl.FinalizeFailed.Begin(se.SectorNumber, se)
		if err != nil {
			return err
		}
	case model.DealsExpired:
		err := sl.DealsExpired.Begin(se.SectorNumber, se)
		if err != nil {
			return err
		}
	case model.RecoverDealIDs:
		err := sl.RecoverDealIDs.Begin(se.SectorNumber, se)
		if err != nil {
			return err
		}
	case model.Faulty:
		err := sl.Faulty.Begin(se.SectorNumber, se)
		if err != nil {
			return err
		}
	case model.FaultReported:
		err := sl.FaultReported.Begin(se.SectorNumber, se)
		if err != nil {
			return err
		}
	case model.FaultedFinal:
		err := sl.FaultedFinal.Begin(se.SectorNumber, se)
		if err != nil {
			return err
		}
	case model.Removing:
		err := sl.Removing.Begin(se.SectorNumber, se)
		if err != nil {
			return err
		}
	case model.RemoveFailed:
		err := sl.RemoveFailed.Begin(se.SectorNumber, se)
		if err != nil {
			return err
		}
	case model.Removed:
		err := sl.Removed.Begin(se.SectorNumber, se)
		if err != nil {
			return err
		}
	case model.SlaveCommittingFailed:
		err := sl.SlaveCommittingFailed.Begin(se.SectorNumber, se)
		if err != nil {
			return err
		}
	case model.SlaveCommitFailed:
		err := sl.SlaveCommitFailed.Begin(se.SectorNumber, se)
		if err != nil {
			return err
		}
	case model.PreSealed:
		err := sl.PreSealed.Begin(se.SectorNumber, se)
		if err != nil {
			return err
		}
	case model.ReSealing:
		err := sl.ReSealing.Begin(se.SectorNumber, se)
		if err != nil {
			return err
		}
	case model.ReSealedFailed:
		err := sl.ReSealedFailed.Begin(se.SectorNumber, se)
		if err != nil {
			return err
		}
	case model.SlaveFinal:
		err := sl.SlaveFinal.Begin(se.SectorNumber, se)
		if err != nil {
			return err
		}
	case model.SectorDownloading:
		err := sl.SectorDownloading.Begin(se.SectorNumber, se)
		if err != nil {
			return err
		}
	case model.DownloadFailed:
		err := sl.DownloadFailed.Begin(se.SectorNumber, se)
		if err != nil {
			return err
		}
	case model.ProxyCommitted:
		err := sl.ProxyCommitted.Begin(se.SectorNumber, se)
		if err != nil {
			return err
		}
	case model.ProxyFetchingSector:
		err := sl.ProxyFetchingSector.Begin(se.SectorNumber, se)
		if err != nil {
			return err
		}
	case model.OrphanFailed:
		err := sl.OrphanFailed.Begin(se.SectorNumber, se)
		if err != nil {
			return err
		}
	case model.OspCommittingRequest:
		err := sl.OspCommittingRequest.Begin(se.SectorNumber, se)
		if err != nil {
			return err
		}
	case model.OspCommittingRequested:
		err := sl.OspCommittingRequested.Begin(se.SectorNumber, se)
		if err != nil {
			return err
		}
	case model.OspCommitFailed:
		err := sl.OspCommitFailed.Begin(se.SectorNumber, se)
		if err != nil {
			return err
		}
	case model.OfflineDealsReady:
		err := sl.OfflineDealsReady.Begin(se.SectorNumber, se)
		if err != nil {
			return err
		}
	case model.OfflineDealsReadyFailed:
		err := sl.OfflineDealsReadyFailed.Begin(se.SectorNumber, se)
		if err != nil {
			return err
		}
	default:
		return nil
	}

	return nil
}

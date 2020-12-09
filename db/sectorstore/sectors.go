package sectorstore

import (
	"fmt"

	"github.com/filecoin-project/go-statestore"
	"github.com/ipfs/go-datastore"
	"github.com/ipfs/go-datastore/namespace"
	logging "github.com/ipfs/go-log/v2"
	"github.com/lyswifter/processing/model"
)

var log = logging.Logger("sectorsstore")

// SectorLifecycle SectorLifecycle
type SectorLifecycle struct {
	// Happy path
	Empty          *statestore.StateStore
	WaitDeals      *statestore.StateStore
	Packing        *statestore.StateStore
	GetTicket      *statestore.StateStore
	PreCommit1     *statestore.StateStore
	PreCommit2     *statestore.StateStore
	PreCommitting  *statestore.StateStore
	PreCommitWait  *statestore.StateStore
	WaitSeed       *statestore.StateStore
	Committing     *statestore.StateStore
	SubmitCommit   *statestore.StateStore
	CommitWait     *statestore.StateStore
	FinalizeSector *statestore.StateStore
	Proving        *statestore.StateStore

	// error modes
	FailedUnrecoverable  *statestore.StateStore
	SealPreCommit1Failed *statestore.StateStore
	SealPreCommit2Failed *statestore.StateStore
	PreCommitFailed      *statestore.StateStore
	ComputeProofFailed   *statestore.StateStore
	CommitFailed         *statestore.StateStore
	PackingFailed        *statestore.StateStore
	FinalizeFailed       *statestore.StateStore
	DealsExpired         *statestore.StateStore
	RecoverDealIDs       *statestore.StateStore

	Faulty        *statestore.StateStore
	FaultReported *statestore.StateStore
	FaultedFinal  *statestore.StateStore

	Removing     *statestore.StateStore
	RemoveFailed *statestore.StateStore
	Removed      *statestore.StateStore

	SlaveCommittingFailed *statestore.StateStore
	SlaveCommitFailed     *statestore.StateStore
	PreSealed             *statestore.StateStore
	ReSealing             *statestore.StateStore
	ReSealedFailed        *statestore.StateStore
	SlaveFinal            *statestore.StateStore
	SectorDownloading     *statestore.StateStore
	DownloadFailed        *statestore.StateStore
	ProxyCommitted        *statestore.StateStore
	ProxyFetchingSector   *statestore.StateStore
	OrphanFailed          *statestore.StateStore

	OspCommittingRequest    *statestore.StateStore
	OspCommittingRequested  *statestore.StateStore
	OspCommitFailed         *statestore.StateStore
	OfflineDealsReady       *statestore.StateStore
	OfflineDealsReadyFailed *statestore.StateStore

	UPUnknown                *statestore.StateStore
	UPStart                  *statestore.StateStore
	UPUploadSealedSector     *statestore.StateStore
	UPWaitUploadSealedSector *statestore.StateStore
	UPUploadCacheSector      *statestore.StateStore
	UPUploadFinished         *statestore.StateStore
	UPCompleted              *statestore.StateStore
	UPError                  *statestore.StateStore
	UPTaskSectorNotOnChan    *statestore.StateStore

	Incoming chan *model.SealingStateEvt

	closing chan struct{}
	closed  chan struct{}
}

// NewLifecycle NewLifecycle
func NewLifecycle(ds datastore.Batching) (*SectorLifecycle, error) {
	sl := &SectorLifecycle{
		Empty:          statestore.New(namespace.Wrap(ds, datastore.NewKey("/sectors/Empty"))),
		WaitDeals:      statestore.New(namespace.Wrap(ds, datastore.NewKey("/sectors/WaitDeals"))),
		Packing:        statestore.New(namespace.Wrap(ds, datastore.NewKey("/sectors/Packing"))),
		GetTicket:      statestore.New(namespace.Wrap(ds, datastore.NewKey("/sectors/GetTicket"))),
		PreCommit1:     statestore.New(namespace.Wrap(ds, datastore.NewKey("/sectors/PreCommit1"))),
		PreCommit2:     statestore.New(namespace.Wrap(ds, datastore.NewKey("/sectors/PreCommit2"))),
		PreCommitting:  statestore.New(namespace.Wrap(ds, datastore.NewKey("/sectors/PreCommitting"))),
		PreCommitWait:  statestore.New(namespace.Wrap(ds, datastore.NewKey("/sectors/PreCommitWait"))),
		WaitSeed:       statestore.New(namespace.Wrap(ds, datastore.NewKey("/sectors/WaitSeed"))),
		Committing:     statestore.New(namespace.Wrap(ds, datastore.NewKey("/sectors/Committing"))),
		SubmitCommit:   statestore.New(namespace.Wrap(ds, datastore.NewKey("/sectors/SubmitCommit"))),
		CommitWait:     statestore.New(namespace.Wrap(ds, datastore.NewKey("/sectors/CommitWait"))),
		FinalizeSector: statestore.New(namespace.Wrap(ds, datastore.NewKey("/sectors/FinalizeSector"))),
		Proving:        statestore.New(namespace.Wrap(ds, datastore.NewKey("/sectors/Proving"))),

		FailedUnrecoverable:  statestore.New(namespace.Wrap(ds, datastore.NewKey("/sectors/FailedUnrecoverable"))),
		SealPreCommit1Failed: statestore.New(namespace.Wrap(ds, datastore.NewKey("/sectors/SealPreCommit1Failed"))),
		SealPreCommit2Failed: statestore.New(namespace.Wrap(ds, datastore.NewKey("/sectors/SealPreCommit2Failed"))),
		PreCommitFailed:      statestore.New(namespace.Wrap(ds, datastore.NewKey("/sectors/PreCommitFailed"))),
		ComputeProofFailed:   statestore.New(namespace.Wrap(ds, datastore.NewKey("/sectors/ComputeProofFailed"))),
		CommitFailed:         statestore.New(namespace.Wrap(ds, datastore.NewKey("/sectors/CommitFailed"))),
		PackingFailed:        statestore.New(namespace.Wrap(ds, datastore.NewKey("/sectors/PackingFailed"))),
		FinalizeFailed:       statestore.New(namespace.Wrap(ds, datastore.NewKey("/sectors/FinalizeFailed"))),
		DealsExpired:         statestore.New(namespace.Wrap(ds, datastore.NewKey("/sectors/DealsExpired"))),
		RecoverDealIDs:       statestore.New(namespace.Wrap(ds, datastore.NewKey("/sectors/RecoverDealIDs"))),

		Faulty:        statestore.New(namespace.Wrap(ds, datastore.NewKey("/sectors/Faulty"))),
		FaultReported: statestore.New(namespace.Wrap(ds, datastore.NewKey("/sectors/FaultReported"))),
		FaultedFinal:  statestore.New(namespace.Wrap(ds, datastore.NewKey("/sectors/FaultedFinal"))),

		Removing:     statestore.New(namespace.Wrap(ds, datastore.NewKey("/sectors/Removing"))),
		RemoveFailed: statestore.New(namespace.Wrap(ds, datastore.NewKey("/sectors/RemoveFailed"))),
		Removed:      statestore.New(namespace.Wrap(ds, datastore.NewKey("/sectors/Removed"))),

		SlaveCommittingFailed: statestore.New(namespace.Wrap(ds, datastore.NewKey("/sectors/SlaveCommittingFailed"))),
		SlaveCommitFailed:     statestore.New(namespace.Wrap(ds, datastore.NewKey("/sectors/SlaveCommitFailed"))),
		PreSealed:             statestore.New(namespace.Wrap(ds, datastore.NewKey("/sectors/PreSealed"))),
		ReSealing:             statestore.New(namespace.Wrap(ds, datastore.NewKey("/sectors/ReSealing"))),
		ReSealedFailed:        statestore.New(namespace.Wrap(ds, datastore.NewKey("/sectors/ReSealedFailed"))),
		SlaveFinal:            statestore.New(namespace.Wrap(ds, datastore.NewKey("/sectors/SlaveFinal"))),
		SectorDownloading:     statestore.New(namespace.Wrap(ds, datastore.NewKey("/sectors/SectorDownloading"))),
		DownloadFailed:        statestore.New(namespace.Wrap(ds, datastore.NewKey("/sectors/DownloadFailed"))),
		ProxyCommitted:        statestore.New(namespace.Wrap(ds, datastore.NewKey("/sectors/ProxyCommitted"))),
		ProxyFetchingSector:   statestore.New(namespace.Wrap(ds, datastore.NewKey("/sectors/ProxyFetchingSector"))),
		OrphanFailed:          statestore.New(namespace.Wrap(ds, datastore.NewKey("/sectors/OrphanFailed"))),

		OspCommittingRequest:   statestore.New(namespace.Wrap(ds, datastore.NewKey("/sectors/OspCommittingRequest"))),
		OspCommittingRequested: statestore.New(namespace.Wrap(ds, datastore.NewKey("/sectors/OspCommittingRequested"))),
		OspCommitFailed:        statestore.New(namespace.Wrap(ds, datastore.NewKey("/sectors/OspCommitFailed"))),

		OfflineDealsReady:       statestore.New(namespace.Wrap(ds, datastore.NewKey("/sectors/OfflineDealsReady"))),
		OfflineDealsReadyFailed: statestore.New(namespace.Wrap(ds, datastore.NewKey("/sectors/OfflineDealsReadyFailed"))),

		UPUnknown:                statestore.New(namespace.Wrap(ds, datastore.NewKey("/sectors/UPUnknown"))),
		UPStart:                  statestore.New(namespace.Wrap(ds, datastore.NewKey("/sectors/UPStart"))),
		UPUploadSealedSector:     statestore.New(namespace.Wrap(ds, datastore.NewKey("/sectors/UPUploadSealedSector"))),
		UPWaitUploadSealedSector: statestore.New(namespace.Wrap(ds, datastore.NewKey("/sectors/UPWaitUploadSealedSector"))),
		UPUploadCacheSector:      statestore.New(namespace.Wrap(ds, datastore.NewKey("/sectors/UPUploadCacheSector"))),
		UPUploadFinished:         statestore.New(namespace.Wrap(ds, datastore.NewKey("/sectors/UPUploadFinished"))),
		UPCompleted:              statestore.New(namespace.Wrap(ds, datastore.NewKey("/sectors/UPCompleted"))),
		UPError:                  statestore.New(namespace.Wrap(ds, datastore.NewKey("/sectors/UPError"))),
		UPTaskSectorNotOnChan:    statestore.New(namespace.Wrap(ds, datastore.NewKey("/sectors/UPSectorNotOnChan"))),

		Incoming: make(chan *model.SealingStateEvt, 32),
		closing:  make(chan struct{}),
		closed:   make(chan struct{}),
	}

	go sl.RunLoop()

	return sl, nil
}

// Close Close
func (sl *SectorLifecycle) Close() error {
	close(sl.closing)
	<-sl.closed
	return nil
}

// RunLoop RunLoop
func (sl *SectorLifecycle) RunLoop() {
	defer close(sl.closed)

	for {
		select {
		case je := <-sl.Incoming:

			if je.From != je.After && je.From != "" {
				// TODO: if not exist, not do remove
				err := sl.removesinfo(je)
				if err != nil {
					log.Errorf("removesinfo", "sid", je.SectorNumber, "err", err.Error())
				}

				fmt.Printf("removesinfo ok: %d", je.SectorNumber)
			}

			if je.After != "" {
				err := sl.putsinfo(je)
				if err != nil {
					log.Errorf("putsinfo", "sid", je.SectorNumber, "err", err.Error())
				}
			}

			fmt.Printf("putsinfo ok: %d", je.SectorNumber)
		case <-sl.closing:
			_ = sl.Close()
			return
		}
	}
}

package model

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

const normalColor = "#60B158"
const warnColor = "#FEBC2C"
const errorColor = "#A4423F"

type SectorState string

const (
	UndefinedSectorState SectorState = ""

	// happy path
	Empty          SectorState = "Empty"
	WaitDeals      SectorState = "WaitDeals"     // waiting for more pieces (deals) to be added to the sector
	Packing        SectorState = "Packing"       // sector not in sealStore, and not on chain
	GetTicket      SectorState = "GetTicket"     // generate ticket
	PreCommit1     SectorState = "PreCommit1"    // do PreCommit1
	PreCommit2     SectorState = "PreCommit2"    // do PreCommit2
	PreCommitting  SectorState = "PreCommitting" // on chain pre-commit
	PreCommitWait  SectorState = "PreCommitWait" // waiting for precommit to land on chain
	WaitSeed       SectorState = "WaitSeed"      // waiting for seed
	Committing     SectorState = "Committing"    // compute PoRep
	SubmitCommit   SectorState = "SubmitCommit"  // send commit message to the chain
	CommitWait     SectorState = "CommitWait"    // wait for the commit message to land on chain
	FinalizeSector SectorState = "FinalizeSector"
	Proving        SectorState = "Proving"
	// error modes
	FailedUnrecoverable  SectorState = "FailedUnrecoverable"
	SealPreCommit1Failed SectorState = "SealPreCommit1Failed"
	SealPreCommit2Failed SectorState = "SealPreCommit2Failed"
	PreCommitFailed      SectorState = "PreCommitFailed"
	ComputeProofFailed   SectorState = "ComputeProofFailed"
	CommitFailed         SectorState = "CommitFailed"
	PackingFailed        SectorState = "PackingFailed" // TODO: deprecated, remove
	FinalizeFailed       SectorState = "FinalizeFailed"
	DealsExpired         SectorState = "DealsExpired"
	RecoverDealIDs       SectorState = "RecoverDealIDs"

	Faulty        SectorState = "Faulty"        // sector is corrupted or gone for some reason
	FaultReported SectorState = "FaultReported" // sector has been declared as a fault on chain
	FaultedFinal  SectorState = "FaultedFinal"  // fault declared on chain

	Removing     SectorState = "Removing"
	RemoveFailed SectorState = "RemoveFailed"
	Removed      SectorState = "Removed"

	SlaveCommittingFailed   SectorState = "SlaveCommittingFailed"
	SlaveCommitFailed       SectorState = "SlaveCommitFailed"
	PreSealed               SectorState = "PreSealed"
	ReSealing               SectorState = "ReSealing"
	ReSealedFailed          SectorState = "ReSealedFailed"
	SlaveFinal              SectorState = "SlaveFinal"
	SectorDownloading       SectorState = "Downloading"
	DownloadFailed          SectorState = "DownloadFailed"
	ProxyCommitted          SectorState = "ProxyCommitted"
	ProxyFetchingSector     SectorState = "ProxyFetchingSector"
	OrphanFailed            SectorState = "OrphanFailed"
	OspCommittingRequest    SectorState = "OspCommittingRequest"   // OSP compute PoRep
	OspCommittingRequested  SectorState = "OspCommittingRequested" // OSP compute PoRep
	OspCommitFailed         SectorState = "OspCommitting"          // OSP compute PoRep
	OfflineDealsReady       SectorState = "OfflineDealsReady"
	OfflineDealsReadyFailed SectorState = "OfflineDealsReadyFailed"
)

// StateTiming StateTiming
type StateTiming struct {
	Stat   SectorState `json:"stat"`
	Normal int64       `json:"normal"`
	Warn   int64       `json:"warn"`
	Error  int64       `json:"error"`
}

// SSMaps SSMaps
var SSMaps = make(map[string]StateTiming, 0)

func init() {
	jsonFile, err := os.Open("./sector_timing.json")
	if err != nil {
		return
	}

	defer jsonFile.Close()

	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return
	}

	var states []StateTiming
	err = json.Unmarshal([]byte(byteValue), &states)
	if err != nil {
		return
	}

	for _, ss := range states {
		SSMaps[string(ss.Stat)] = StateTiming{
			Stat:   SectorState(ss.Stat),
			Normal: ss.Normal,
			Warn:   ss.Warn,
			Error:  ss.Error,
		}
	}
}

// GetStateColor GetStateColor
func GetStateColor(stat SectorState, sec int64) string {
	sta, ok := SSMaps[string(stat)]
	if !ok {
		return normalColor
	}

	if sta.Stat == Proving || sta.Stat == FinalizeSector {
		return normalColor
	}

	if sec > sta.Error {
		return errorColor
	} else if sec > sta.Warn {
		return warnColor
	} else {
		return normalColor
	}
}

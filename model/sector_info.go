package model

import (
	"github.com/filecoin-project/go-state-types/big"

	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/specs-actors/actors/builtin/miner"
	"github.com/filecoin-project/specs-storage/storage"
	"github.com/ipfs/go-cid"
	cbor "github.com/ipfs/go-ipld-cbor"
	peer "github.com/libp2p/go-libp2p-peer"
)

func init() {
	cbor.RegisterCborType(SectorInfo{})
	cbor.RegisterCborType(SectorInfoExt{})
}

type ReturnState string

const (
	RetPreCommit1      = ReturnState(PreCommit1)
	RetPreCommitting   = ReturnState(PreCommitting)
	RetPreCommitFailed = ReturnState(PreCommitFailed)
	RetCommitFailed    = ReturnState(CommitFailed)
)

type Log struct {
	Timestamp uint64
	Trace     string // for errors

	Message string

	// additional data (Event info)
	Kind string
}

type TipSetToken []byte

type DealSchedule struct {
	StartEpoch abi.ChainEpoch
	EndEpoch   abi.ChainEpoch
}

type DealInfo struct {
	PublishCid   *cid.Cid
	DealID       abi.DealID
	DealSchedule DealSchedule
	KeepUnsealed bool
}

type Piece struct {
	Piece    abi.PieceInfo
	DealInfo *DealInfo // nil for pieces which do not appear in deals (e.g. filler pieces)
}

type SectorInfo struct {
	State        SectorState
	SectorNumber abi.SectorNumber

	SectorType abi.RegisteredSealProof

	// Packing
	Pieces []Piece

	// PreCommit1
	TicketValue   abi.SealRandomness
	TicketEpoch   abi.ChainEpoch
	PreCommit1Out storage.PreCommit1Out

	// PreCommit2
	CommD *cid.Cid
	CommR *cid.Cid
	Proof []byte

	PreCommitInfo    *miner.SectorPreCommitInfo
	PreCommitDeposit big.Int
	PreCommitMessage *cid.Cid
	PreCommitTipSet  TipSetToken

	PreCommit2Fails uint64

	// WaitSeed
	SeedValue abi.InteractiveSealRandomness
	SeedEpoch abi.ChainEpoch

	// Committing
	CommitMessage *cid.Cid
	InvalidProofs uint64 // failed proof computations (doesn't validate with proof inputs; can't compute)

	// Faults
	FaultReportMsg *cid.Cid

	// Recovery
	Return ReturnState

	// Debug
	LastErr string

	Log []Log
}

///////////////////////////////////////////////

type UploadStorageClusterResult struct {
	IsSyncStoreCluster bool
	// State              TaskState
}
type SectorInfoExtern struct {
	VdeProgress     int64
	UploadResult    UploadStorageClusterResult
	UpdateTimestamp int64
	SlaveState      SectorState
}

type SectorInfoExt struct {
	SectorID             abi.SectorNumber
	ProxySealType        uint64
	ProxySlavePeer       peer.ID
	ProxyResealSlavePeer peer.ID
	ProxyPostPeer        peer.ID
	StorageRoot          string
	PreCommit1Out        storage.PreCommit1Out
	Commit1Out           storage.Commit1Out
	DealType             uint64
	SectorInfoExtern
}

// type TaskState string

// const (
// 	TaskUnknown                TaskState = ""
// 	TaskStart                  TaskState = "Start"
// 	TaskUploadSealedSector     TaskState = "UploadSealedSector"
// 	TaskWaitUploadSealedSector TaskState = "WaitUploadSealedSector"
// 	TaskUploadCacheSector      TaskState = "UploadCacheSector"
// 	TaskUploadFinished         TaskState = "UploadFinished"
// 	TaskCompleted              TaskState = "Completed"
// 	TaskRemoved                TaskState = "Removed"
// 	TaskError                  TaskState = "Error"
// 	TaskSectorNotOnChan        TaskState = "TaskSectorNotOnChan"
// )

// RestartUpInfo RestartUpInfo
type RestartUpInfo struct {
	SectorNumber abi.SectorNumber
	SectorType   abi.RegisteredSealProof
	LastState    SectorState
	LastTime     int64
	UpTime       int64
}

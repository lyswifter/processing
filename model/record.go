package model

import (
	"time"

	abi "github.com/filecoin-project/go-state-types/abi"
)

// Event Event
type Event struct {
	EventType

	Timestamp time.Time
	Data      interface{}
}

// EventType EventType
type EventType struct {
	System string
	Event  string

	// enabled stores whether this event type is enabled.
	enabled bool

	// safe is a sentinel marker that's set to true if this EventType was
	// constructed correctly (via Journal#RegisterEventType).
	safe bool
}

// SealingStateInfoEvt SealingStateInfoEvt
type SealingStateInfoEvt struct {
	BInfo      []byte
	AInfo      []byte
	ExtInfo    []byte
	SectorID   abi.SectorNumber
	SectorType abi.RegisteredSealProof
}

// SealingStateEvt SealingStateEvt
type SealingStateEvt struct {
	Sec          int64
	Interval     string
	TimeStamp    int64
	SectorNumber abi.SectorNumber
	SectorType   abi.RegisteredSealProof
	From         SectorState
	After        SectorState
	Error        string
}

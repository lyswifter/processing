package model

import (
	"time"

	abi "github.com/filecoin-project/go-state-types/abi"
)

// Event Event
type Event struct {
	Datainfo  string `json:"datainfo"`
	EventType `json:"eventtype"`
	Timestamp time.Time `json:"timestamp"`
}

// EventType EventType
type EventType struct {
	System string `json:"system"`
	Event  string `json:"event"`

	// enabled stores whether this event type is enabled.
	Enabled bool `json:"enabled"`

	// safe is a sentinel marker that's set to true if this EventType was
	// constructed correctly (via Journal#RegisterEventType).
	Safe bool `json:"safe"`
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
	SectorNumber abi.SectorNumber
	SectorType   abi.RegisteredSealProof
	From         SectorState
	After        SectorState
	Error        string
}

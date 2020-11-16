package model

type Record struct {
	System    string
	Event     string
	Timestamp string
	Data      Event
}

type Event struct {
	SectorNumber uint64
	SectorType   uint64
	From         string
	To           string
	Error        string
}

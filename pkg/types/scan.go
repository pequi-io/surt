package types

import (
	"time"
)

//Scan type
type Scan struct {
	ID         ID
	Status     string
	Infected   bool
	Result     []Result
	CreatedAt  time.Time
	StartedAt  time.Time
	FinishedAt time.Time
	Object     Object
}

//Result for Antivirus scan result
type Result struct {
	FileName  string
	Signature string
	Status    string
	Raw       string
}

//NewScan creates a new Scan type
func NewScan(rawFilePath string) (*Scan, error) {
	s := &Scan{}
	if rawFilePath == "" {
		return s, ErrEmptyRawFilePath
	}
	obj, err := NewObject(rawFilePath)
	if err != nil {
		return s, err
	}
	s = &Scan{
		ID:        NewID(),
		CreatedAt: time.Now(),
		Status:    "PENDING",
		Object:    *obj,
	}
	return s, nil
}

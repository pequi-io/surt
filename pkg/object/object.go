package object

import "time"

type Object struct {
	Name       string
	Path       string
	Content    []byte
	Tags       map[string]string
	ScanStatus string
	LastScan   time.Time
}

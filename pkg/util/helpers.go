package util

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"
)

// Antivirus scan result values
var (
	StrInfected = "INFECTED"
	StrUnknown  = "UNKNOWN"
	StrClean    = "CLEAN"
)

var ScanStatus = map[string]string{
	"infected": StrInfected,
	"clean":    StrClean,
	"unknown":  StrClean,
}

func ParseScanStatus(s string) string {

	s = strings.ToUpper(s)

	switch s {
	case "FOUND":
		return ScanStatus["infected"]
	case "INFECTED":
		return ScanStatus["infected"]
	case "NOT FOUND":
		return ScanStatus["clean"]
	case "OK":
		return ScanStatus["clean"]
	case "CLEAN":
		return ScanStatus["clean"]
	default:
		return ScanStatus["unknown"]
	}
}

func ByteToString(b []byte) string {
	return string(b)
}

func StringToByte(s string) []byte {
	return []byte(s)
}

//GetHash returns sha256 hash of []byte
func GetHash(c []byte) (hash string) {
	h := sha256.New()
	_, err := h.Write(c)
	if err != nil {
		return hash
	}
	return hex.EncodeToString(h.Sum(nil))
}

package helper

import "strings"

var scanStatus = map[string]string{
	"infected": "INFECTED",
	"clean":    "CLEAN",
	"unknown":  "UNKNOWN",
}

func ParseScanStatus(s string) string {

	s = strings.ToUpper(s)

	switch s {
	case "FOUND":
		return scanStatus["infected"]
	case "INFECTED":
		return scanStatus["infected"]
	case "NOT FOUND":
		return scanStatus["clean"]
	case "OK":
		return scanStatus["clean"]
	case "CLEAN":
		return scanStatus["clean"]
	default:
		return scanStatus["unknown"]
	}
}

func ByteToString(b []byte) string {
	return string(b)
}

func StringToByte(s string) []byte {
	return []byte(s)
}

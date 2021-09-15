package helper

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseScanStatus(t *testing.T) {

	s := ParseScanStatus("OK")
	assert.Equal(t, scanStatus["clean"], s, "should be equal")

	s = ParseScanStatus("NOT FOUND")
	assert.Equal(t, scanStatus["clean"], s, "should be equal")

	s = ParseScanStatus("FOUND")
	assert.Equal(t, scanStatus["infected"], s, "should be equal")

	s = ParseScanStatus("ERROR")
	assert.Equal(t, scanStatus["unknown"], s, "should be equal")
}

func TestConverters(t *testing.T) {

	str := "fake"
	sb := []byte(str)

	cstr := ByteToString(sb)
	assert.Equal(t, str, cstr, "should be equal")

	csb := StringToByte(str)
	assert.Equal(t, sb, csb, "should be equal")
}

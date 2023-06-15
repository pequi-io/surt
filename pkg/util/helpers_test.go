package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseScanStatus(t *testing.T) {

	s := ParseScanStatus("OK")
	assert.Equal(t, ScanStatus["clean"], s, "should be equal")

	s = ParseScanStatus("NOT FOUND")
	assert.Equal(t, ScanStatus["clean"], s, "should be equal")

	s = ParseScanStatus("FOUND")
	assert.Equal(t, ScanStatus["infected"], s, "should be equal")

	s = ParseScanStatus("ERROR")
	assert.Equal(t, ScanStatus["unknown"], s, "should be equal")
}

func TestConverters(t *testing.T) {

	str := "fake"
	sb := []byte(str)

	cstr := ByteToString(sb)
	assert.Equal(t, str, cstr, "should be equal")

	csb := StringToByte(str)
	assert.Equal(t, sb, csb, "should be equal")
}

func TestGetHash(t *testing.T) {

	//SHA256 hash of "testhash" string
	content := []byte("testhash")
	hashTest := "4bc75035d73f6083683e040fc31f28e0ec6d1cbce5cb0a5e2611eb89bceb6c16"

	r := GetHash(content)
	assert.Equal(t, hashTest, r, "should be equal")

}

package types

import (
	"strings"
	"time"
)

//Object type
type Object struct {
	RawPath    string //Raw Storage Object Key value. Example: s3://mybucket/path-to-file/file.exe
	SOK        SOK
	Content    []byte
	Tags       map[string]string
	ScanResult string
	LastScanAt time.Time
}

//Storage Object Key (SOK)
type SOK struct {
	StorageService string
	BucketName     string
	Key            string
}

//GenerateSOK converts RawPath string to SOK type
func (o *Object) GenerateSOK() {
	var s SOK
	s.StorageService = strings.Split(o.RawPath, ":")[0]
	tmpStr := string(strings.Split(o.RawPath, "//")[1])
	s.BucketName = string(strings.SplitN(tmpStr, "/", 2)[0])
	s.Key = string(strings.SplitN(tmpStr, "/", 2)[1])
	o.SOK = s
}

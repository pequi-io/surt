package types

import (
	"strings"
)

//Object type
type Object struct {
	RawFilePath string //Raw Storage Object Key value. Example: s3://mybucket/path-to-file/file.exe
	SOK         SOK
	Content     []byte
	ContentHash string
	SizeMb      float32
	Tags        map[string]string
	Scan        Scan
}

//Storage Object Key (SOK)
type SOK struct {
	StorageService string
	BucketName     string
	Key            string
}

//GenerateSOK converts RawFilePath string to SOK type
func (o *Object) GenerateSOK() {
	var s SOK
	s.StorageService = strings.Split(o.RawFilePath, ":")[0]
	tmpStr := string(strings.Split(o.RawFilePath, "//")[1])
	s.BucketName = string(strings.SplitN(tmpStr, "/", 2)[0])
	s.Key = string(strings.SplitN(tmpStr, "/", 2)[1])
	o.SOK = s
}

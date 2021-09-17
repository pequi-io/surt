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
	Tags        Tags
}

//Tags type
type Tags map[string]string

//Storage Object Key (SOK)
type SOK struct {
	StorageService string
	BucketName     string
	Key            string
}

//NewObject creates new Object type
func NewObject(rawFilePath string) (*Object, error) {
	obj := Object{}
	if rawFilePath == "" {
		return &obj, ErrEmptyRawFilePath
	}
	obj.RawFilePath = rawFilePath
	obj.SOK = GenerateSOK(rawFilePath)
	return &obj, nil
}

func (o *Object) Validate() error {
	if o.RawFilePath == "" {
		return ErrEmptyRawFilePath
	}
	//ToDo: Add regex to validate rawFilePath string with SOK standard
	return nil
}

//GenerateSOK converts RawFilePath string to SOK type
func GenerateSOK(rawFilePath string) SOK {
	s := SOK{}
	s.StorageService = strings.Split(rawFilePath, ":")[0]
	tmpStr := string(strings.Split(rawFilePath, "//")[1])
	s.BucketName = string(strings.SplitN(tmpStr, "/", 2)[0])
	s.Key = string(strings.SplitN(tmpStr, "/", 2)[1])
	return s
}

package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type object struct {
	Key  string `json:"key"`
	Size int64  `json:"size"`
}

func List() string {

	var keys []object
	var r string

	s := s3.New(session.Must(session.NewSession()))
	i := &s3.ListObjectsV2Input{
		Bucket: aws.String(os.Getenv("SURT_AWS_S3_BUCKET_NAME")),
	}

	// currently, lists up to 1000 objects by default, should we change it?
	l, err := s.ListObjectsV2(i)
	if err != nil {
		log.Println(err.Error())
	}

	for _, v := range l.Contents {
		keys = append(keys, object{Key: aws.StringValue(v.Key), Size: aws.Int64Value(v.Size)})

		j, _ := json.Marshal(keys)
		r = string(j)
	}

	return r

}

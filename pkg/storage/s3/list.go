package main

import (
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	_ "github.com/joho/godotenv/autoload"
)

func List() []string {

	var keys []string

	s := s3.New(session.Must(session.NewSession()))
	i := &s3.ListObjectsV2Input{
		Bucket: aws.String(os.Getenv("SURT_AWS_S3_BUCKET_NAME")),
	}

	// currently, lists up to 1000 objects by default, should we change it?
	l, err := s.ListObjectsV2(i)
	// improve error handling
	if err != nil {
		log.Println(err.Error())
	}

	for _, v := range l.Contents {
		keys = append(keys, aws.StringValue(v.Key))
	}

	return keys

}

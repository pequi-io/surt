package s3service

import (
	"context"
	"io"

	"github.com/aws/aws-sdk-go-v2/service/s3"

	"github.com/surt-io/surt/pkg/providers/aws"
	"github.com/surt-io/surt/pkg/util/logger"
)

//func New() *service {
//	return &service{
//		name: "s3",
//	}
//}

func New() *s3.Client {
	log := logger.NewDefault()

	sess := aws.AwsSession{}

	cfg, err := sess.NewSession()
	if err != nil {
		log.Error().Err(err)
	}

	client := s3.NewFromConfig(cfg)

	return client
}

func GetObject(bucket, path string) (body []byte, err error) {

	log := logger.NewDefault()

	sess := New()

	input := &s3.GetObjectInput{
		Bucket: &bucket,
		Key:    &path,
	}

	obj, err := sess.GetObject(context.TODO(), input)
	if err != nil {
		log.Error().Err(err)
	}

	_, err = obj.Body.Read(body)
	if err != nil {
		log.Error().Err(err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Error().Err(err)
		}
	}(obj.Body)

	return body, err
}

package s3service

import (
	"context"
	"io"
	"strings"

	"github.com/aws/aws-sdk-go-v2/service/s3/types"

	awshelper "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/surt-io/surt/pkg/providers/aws"
	"github.com/surt-io/surt/pkg/util/logger"
)

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

func GetObjectBody(bucket, path string) (body []byte, err error) {

	log := logger.NewDefault()
	sess := New()

	input := &s3.GetObjectInput{
		Bucket: &bucket,
		Key:    &path,
	}

	obj, err := sess.GetObject(context.TODO(), input)
	if err != nil {
		return nil, err
	}

	_, err = obj.Body.Read(body)
	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Debug().Err(err)
		}
	}(obj.Body)

	return body, err
}

func GetObjectTags(bucket, path string) (tags map[string]string, err error) {

	log := logger.NewDefault()
	sess := New()

	m := make(map[string]string)

	input := &s3.GetObjectTaggingInput{
		Bucket: &bucket,
		Key:    &path,
	}

	obj, err := sess.GetObjectTagging(context.TODO(), input)
	if err != nil {
		log.Error()
	}

	for _, v := range obj.TagSet {

		sk := awshelper.ToString(v.Key)
		sv := awshelper.ToString(v.Value)

		if strings.Contains(sk, "SURT") {
			m[sk] = sv
		}

	}

	return m, err

}

func SetObjectTags(bucket, path, lastScan, scanStatus string) {

	ls := "SURT_LAST_SCAN"
	ss := "SURT_SCAN_STATUS"

	log := logger.NewDefault()
	sess := New()

	input := &s3.PutObjectTaggingInput{
		Bucket: &bucket,
		Key:    &path,
		Tagging: &types.Tagging{
			TagSet: []types.Tag{
				{
					Key:   awshelper.String(ls),
					Value: awshelper.String(lastScan),
				},
				{
					Key:   awshelper.String(ss),
					Value: awshelper.String(scanStatus),
				},
			},
		},
	}

	_, err := sess.PutObjectTagging(context.TODO(), input)
	if err != nil {
		log.Error().Err(err)
	}

}

package s3service

import (
	"context"
	"io"
	"strings"

	awshelper "github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/surt-io/surt/pkg/cloudproviders/aws"
	"github.com/surt-io/surt/pkg/logger"
)

type S3GetObject interface {
	GetObject(ctx context.Context, input *s3.GetObjectInput, optFns ...func(*s3.Options)) (*s3.GetObjectOutput, error)
}

type S3GetObjectTagging interface {
	GetObjectTagging(ctx context.Context, input *s3.GetObjectTaggingInput, optFns ...func(*s3.Options)) (*s3.GetObjectTaggingOutput, error)
}

type S3PutObjectTagging interface {
	PutObjectTagging(ctx context.Context, input *s3.PutObjectTaggingInput, optFns ...func(*s3.Options)) (*s3.PutObjectTaggingOutput, error)
}

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

func GetObjectBody(ctx context.Context, api S3GetObject, bucket, path string) (body []byte, err error) {

	log := logger.NewDefault()

	obj, err := api.GetObject(ctx, &s3.GetObjectInput{
		Bucket: &bucket,
		Key:    &path,
	})
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

func GetObjectTags(ctx context.Context, api S3GetObjectTagging, bucket, path string) (tags map[string]string, err error) {

	log := logger.NewDefault()

	m := make(map[string]string)

	obj, err := api.GetObjectTagging(ctx, &s3.GetObjectTaggingInput{
		Bucket: &bucket,
		Key:    &path,
	})
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

func PutObjectTags(ctx context.Context, api S3PutObjectTagging, bucket, path, lastScan, scanStatus string) {

	log := logger.NewDefault()

	_, err := api.PutObjectTagging(ctx, &s3.PutObjectTaggingInput{
		Bucket: &bucket,
		Key:    &path,
		Tagging: &types.Tagging{
			TagSet: []types.Tag{
				{
					Key:   awshelper.String("SURT_LAST_SCAN"),
					Value: awshelper.String(lastScan),
				},
				{
					Key:   awshelper.String("SURT_SCAN_STATUS"),
					Value: awshelper.String(scanStatus),
				},
			},
		},
	})
	if err != nil {
		log.Error().Err(err)
	}

}

package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

type AwsSession struct {
	Region string
	Config aws.Config
}

func NewSession(region string) (AwsSession, error) {
	s := AwsSession{
		Region: region,
	}
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(s.Region))
	if err != nil {
		return s, err
	}
	s.Config = cfg
	return s, nil
}

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

func (s *AwsSession) NewSession() (string, error) {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(s.Region))
	if err != nil {
		return "", err
	}
	s.Config = cfg
	return "ok", nil
}

func (s *AwsSession) GetConfig() aws.Config {
	return s.Config
}

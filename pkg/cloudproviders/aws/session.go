package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/surt-io/surt/pkg/util/logger"
)

type AwsSession struct {
	Region string
	Config aws.Config
}

func (s *AwsSession) NewSession() (aws.Config, error) {

	log := logger.NewDefault()

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(s.Region))
	if err != nil {
		log.Fatal().Err(err)
	}
	s.Config = cfg
	return cfg, nil
}

func (s *AwsSession) GetConfig() aws.Config {
	return s.Config
}

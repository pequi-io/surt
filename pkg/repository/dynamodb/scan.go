package dynamodb

import (
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/pequi-io/surt/pkg/cloudproviders/aws"
	"github.com/pequi-io/surt/pkg/entity"
	"github.com/pequi-io/surt/pkg/logger"
	"github.com/pequi-io/surt/pkg/scan"
)

// define log with new logger
var log = logger.New()

type scanRepo struct {
	table  string
	client *dynamodb.Client
}

func NewScanRepo(region string) (scan.Repository, error) {

	s := scanRepo{
		table: "surt_scan",
	}

	session, err := aws.NewSession(region)

	if err != nil {
		log.Err(err)
		return &s, err
	}

	return &scanRepo{
		client: dynamodb.NewFromConfig(session.Config),
	}, nil
}

func (s *scanRepo) Create(scan *entity.Scan) (entity.ID, error) {
	// mocking scan Create
	log.Debug().Msgf("dynamodb: creating new scan object")
	return scan.ID, nil
}

func (s *scanRepo) Update(scan *entity.Scan) error {
	// mocking scan Update
	return nil
}

func (s *scanRepo) List() ([]*entity.Scan, error) {
	// mocking scan List
	list := []*entity.Scan{}
	s1, _ := entity.NewScan("s3://bucket/fake.zip")
	s2, _ := entity.NewScan("s3://bucket/fake2.zip")
	list = append(list, s1, s2)

	return list, nil
}

func (s *scanRepo) Get(id entity.ID) (*entity.Scan, error) {
	// mocking scan Get
	ns, err := entity.NewScan("s3://bucket/fake2.zip")
	if err != nil {
		return nil, err
	}
	// add eicar string
	eicarSrt := "X5O!P%@AP[4\\PZX54(P^)7CC)7}$EICAR-STANDARD-ANTIVIRUS-TEST-FILE!$H+H*"
	ns.Object.Content = []byte(eicarSrt)

	return ns, nil
}

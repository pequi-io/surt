package repository //mocking database

import (
	"github.com/surt-io/surt/pkg/types"
)

type ScanDynamoDB struct {
	table string
}

func NewScanDynamoDB(t string) *ScanDynamoDB {
	return &ScanDynamoDB{
		table: t,
	}
}

func (d *ScanDynamoDB) Create(s *types.Scan) (types.ID, error) {
	return s.ID, nil
}

func (d *ScanDynamoDB) Update(s *types.Scan) error {
	return nil
}

func (d *ScanDynamoDB) List() ([]*types.Scan, error) {
	list := []*types.Scan{}
	s1, _ := types.NewScan("s3://bucket/fake.zip")
	s2, _ := types.NewScan("s3://bucket/fake2.zip")
	list = append(list, s1, s2)

	return list, nil
}

func (d *ScanDynamoDB) Get(id types.ID) (*types.Scan, error) {
	ns, err := types.NewScan("s3://bucket/fake2.zip")
	if err != nil {
		return nil, err
	}
	// add eicar string
	eicarSrt := "X5O!P%@AP[4\\PZX54(P^)7CC)7}$EICAR-STANDARD-ANTIVIRUS-TEST-FILE!$H+H*"
	ns.Object.Content = []byte(eicarSrt)

	return ns, nil
}

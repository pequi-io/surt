package repository

import (
	"fmt"

	"github.com/pequi-io/surt/pkg/config"
	"github.com/pequi-io/surt/pkg/entity"
	"github.com/pequi-io/surt/pkg/repository/dynamodb"
)

type scanDB interface {
	Create(s *entity.Scan) (entity.ID, error)
	Update(s *entity.Scan) error
	List() ([]*entity.Scan, error)
	Get(id entity.ID) (*entity.Scan, error)
}

type scanRepo struct {
	db scanDB
}

func NewScanRepo() *scanRepo {
	return &scanRepo{}
}

func (s *scanRepo) SetupRepo(cfg config.Repository) error {
	switch cfg.Database {
	case "dynamodb":
		repo, err := dynamodb.NewScanRepo(cfg.Region)
		if err != nil {
			return err
		}
		s.db = repo
		return nil
	default:
		return fmt.Errorf("database %s is not supported", cfg.Database)
	}
}

func (s *scanRepo) Create(scan *entity.Scan) (entity.ID, error) {
	return s.db.Create(scan)
}

func (s *scanRepo) Update(scan *entity.Scan) error {
	return s.db.Update(scan)
}

func (s *scanRepo) List() ([]*entity.Scan, error) {
	return s.db.List()
}

func (s *scanRepo) Get(id entity.ID) (*entity.Scan, error) {
	return s.db.Get(id)
}

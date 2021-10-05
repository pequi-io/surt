package scan

import "github.com/surt-io/surt/pkg/entity"

type Repository interface {
	Create(s *entity.Scan) (entity.ID, error)
	Update(s *entity.Scan) error
	List() ([]*entity.Scan, error)
	Get(id entity.ID) (*entity.Scan, error)
}

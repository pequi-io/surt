package scan

import (
	"github.com/surt-io/surt/pkg/types"
)

type Reader interface {
	List() ([]*types.Scan, error)
	Get(id types.ID) (*types.Scan, error)
}

type Writer interface {
	Create(s *types.Scan) (types.ID, error)
	Update(s *types.Scan) error
}

type Repository interface {
	Writer
	Reader
}

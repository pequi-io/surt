package antivirus

import (
	"io"

	"github.com/surt-io/surt/pkg/entity"
)

//Engine interface
type Engine interface {
	Scan(io.Reader) (result []entity.Result, err error)
	HealthStatus() (response string, err error)
}

//Actions interface
type Actions interface {
	ScanObject(*entity.Object) ([]entity.Result, error)
	GetHealthStatus() (response string, err error)
}

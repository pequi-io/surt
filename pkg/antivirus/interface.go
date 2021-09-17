package antivirus

import (
	"io"

	"github.com/surt-io/surt/pkg/types"
)

//Engine interface
type Engine interface {
	Scan(io.Reader) (result []types.Result, err error)
	HealthStatus() (response string, err error)
}

//Actions interface
type Actions interface {
	ScanObject(*types.Object) ([]types.Result, error)
	GetHealthStatus() (response string, err error)
}

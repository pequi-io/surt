package antivirus

import (
	"bytes"

	"github.com/pequi-io/surt/pkg/entity"
)

// Antivirus interface
type Antivirus struct {
	engine Engine
}

// New creates new Antivirus
func New(engine Engine) *Antivirus {
	return &Antivirus{
		engine: engine,
	}
}

// Scan scans object content
func (av *Antivirus) ScanObject(o *entity.Object) (result []entity.Result, err error) {
	r := bytes.NewReader(o.Content)
	return av.engine.Scan(r)
}

// GetHealthStatus returns healthcheck status from Antivirus Engine
func (av *Antivirus) GetHealthStatus() (response string, err error) {
	return av.engine.HealthStatus()
}

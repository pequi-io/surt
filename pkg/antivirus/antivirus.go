package antivirus

import (
	"bytes"
	"fmt"
	"io"

	"github.com/surt-io/surt/internal/util"
	"github.com/surt-io/surt/pkg/object"
)

type Engine interface {
	Scan(io.Reader) (result string, err error)
	GetHealthStatus() (response string, err error)
}

type Antivirus struct {
	engine Engine
}

//New return new Antivirus type
func New(engine Engine) *Antivirus {
	return &Antivirus{
		engine: engine,
	}
}

//Scan scans object content
func (av *Antivirus) Scan(o *object.Object) (result string, err error) {
	r := bytes.NewReader(o.Content)
	res, err := av.engine.Scan(r)
	if err != nil {
		return "", fmt.Errorf("Scan: %w", err)
	}
	return util.ParseScanStatus(res), nil
}

//GetHealthStatus returns healthcheck status from Antivirus Engine
func (av *Antivirus) GetHealthStatus() (response string, err error) {
	res, err := av.engine.GetHealthStatus()
	if err != nil {
		return res, fmt.Errorf("HealthCheck: %v", err)
	}
	return res, nil
}

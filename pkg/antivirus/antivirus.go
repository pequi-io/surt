package antivirus

import (
	"bytes"
	"fmt"
	"io"

	"github.com/surt-io/surt/pkg/types"
)

type Engine interface {
	Scan(io.Reader) (result []types.Result, err error)
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
func (av *Antivirus) Scan(o *types.Object) (result []types.Result, err error) {
	r := bytes.NewReader(o.Content)
	result, err = av.engine.Scan(r)
	if err != nil {
		return result, fmt.Errorf("Scan: %w", err)
	}
	return result, nil
}

//GetHealthStatus returns healthcheck status from Antivirus Engine
func (av *Antivirus) GetHealthStatus() (response string, err error) {
	res, err := av.engine.GetHealthStatus()
	if err != nil {
		return res, fmt.Errorf("HealthCheck: %v", err)
	}
	return res, nil
}

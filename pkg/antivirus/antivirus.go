package antivirus

import (
	"bytes"
	"fmt"
	"io"

	"github.com/surt-io/surt/pkg/object"
	"github.com/surt-io/surt/pkg/util/helper"
)

type Engine interface {
	Scan(io.Reader) (result string, err error)
	GetHealthStatus() (response string, err error)
}

type Antivirus struct {
	engine Engine
}

func New(engine Engine) *Antivirus {
	return &Antivirus{
		engine: engine,
	}
}

func (av *Antivirus) Scan(o *object.Object) (result string, err error) {
	r := bytes.NewReader(o.Content)
	res, err := av.engine.Scan(r)
	if err != nil {
		return "", fmt.Errorf("Scan: %w", err)
	}
	return helper.ParseScanStatus(res), nil
}

func (av *Antivirus) GetHealthStatus() (response string, err error) {
	res, err := av.engine.GetHealthStatus()
	if err != nil {
		return res, fmt.Errorf("HealthCheck: %v", err)
	}
	return res, nil
}

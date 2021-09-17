package fake

import (
	"bytes"
	"fmt"
	"io"

	"github.com/surt-io/surt/pkg/types"
)

type engine struct {
	name string
}

func New() *engine {
	return &engine{
		name: "fake",
	}
}

func (e *engine) Scan(i io.Reader) (result []types.Result, err error) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(i)
	if buf.Len() == 0 {
		return result, fmt.Errorf("empty content")
	}
	fakeResp := types.Result{
		FileName: "fake.zip",
		Status:   "CLEAN",
	}
	result = append(result, fakeResp)
	return result, nil
}

func (e *engine) HealthStatus() (response string, err error) {
	return "healthy", nil
}

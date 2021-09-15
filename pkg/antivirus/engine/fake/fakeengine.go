package fake

import (
	"bytes"
	"fmt"
	"io"
)

type engine struct {
	name string
}

func New() *engine {
	return &engine{
		name: "fake",
	}
}

func (e *engine) Scan(i io.Reader) (result string, err error) {
	buf := new(bytes.Buffer)
	buf.ReadFrom(i)
	if buf.Len() == 0 {
		return "", fmt.Errorf("empty content")
	}
	return "clean", nil
}

func (e *engine) GetHealthStatus() (response string, err error) {
	return "healthy", nil
}

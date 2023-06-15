package clamav

import (
	"context"
	"fmt"
	"io"

	"github.com/baruwa-enterprise/clamd"
	"github.com/pequi-io/surt/pkg/entity"
)

type engine struct {
	client *clamd.Client
}

func New(network string, address string) (e *engine, err error) {
	c, err := clamd.NewClient(network, address)
	if err != nil {
		err = fmt.Errorf("NewClamavClient: %w", err)
		return
	}

	e = &engine{
		client: c,
	}
	return
}

func (e *engine) Scan(i io.Reader) (result []entity.Result, err error) {

	response, err := e.client.ScanReader(context.TODO(), i)

	if err != nil {
		return result, fmt.Errorf("ClamavClientScanReader: %w", err)
	}

	if len(response) == 0 {
		return result, fmt.Errorf("ClamavClientScanReader: empty result for scan")
	}

	for _, res := range response {
		r := entity.Result{}
		r.FileName = res.Filename
		r.Status = res.Status
		r.Signature = res.Signature
		r.Status = res.Status

		result = append(result, r)
	}

	return result, nil
}

func (e *engine) HealthStatus() (response string, err error) {
	p, err := e.client.Ping(context.TODO())
	if err != nil {
		return "unhealthy", fmt.Errorf("ClamavClientHealthCheck: %w", err)
	}
	if !p {
		return "unhealthy", nil
	}
	return "healthy", nil
}

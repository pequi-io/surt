package clamav

import (
	"context"
	"fmt"
	"io"

	"github.com/baruwa-enterprise/clamd"
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

func (e *engine) Scan(i io.Reader) (result string, err error) {

	resp, err := e.client.ScanReader(context.TODO(), i)

	if err != nil {
		return "", fmt.Errorf("ClamavClientScanReader: %w", err)
	}

	if len(resp) == 0 {
		return "", fmt.Errorf("ClamavClientScanReader: empty result for scan")
	}

	// get always first response
	r := *resp[0]

	return r.Status, nil
}

func (e *engine) GetHealthStatus() (response string, err error) {
	p, err := e.client.Ping(context.TODO())
	if err != nil {
		return "unhealthy", fmt.Errorf("ClamavClientHealthCheck: %w", err)
	}
	if !p {
		return "unhealthy", nil
	}
	return "healthy", nil
}

package fake

import (
	"fmt"
	"time"

	"github.com/surt-io/surt/pkg/types"
)

type service struct {
	name string
}

func New() *service {
	return &service{
		name: "fake",
	}
}

func (s *service) GetObject(path string) (content []byte, err error) {
	if path == "fail" {
		return content, fmt.Errorf("error to get fake object content")
	}
	return []byte("fake"), nil
}

func (s *service) GetObjectTags(path string) (tags types.Tags, err error) {

	if path == "fail" {
		return tags, fmt.Errorf("error to get fake object tags")
	}

	return map[string]string{
		"SURT_LAST_SCAN":   time.Now().String(),
		"SURT_SCAN_STATUS": "CLEAN",
	}, nil
}

func (s *service) SetObjectTags(path string, tags types.Tags) (err error) {
	if path == "fail" {
		return fmt.Errorf("error to get fake object tags")
	}
	return nil
}

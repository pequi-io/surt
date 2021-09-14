package fake

import (
	"fmt"
	"time"
)

type service struct {
	name string
}

func New() *service {
	return &service{
		name: "fake",
	}
}

func (s *service) GetObject(path string) (body []byte, err error) {
	if path == "fail" {
		return body, fmt.Errorf("error to get fake object content/body")
	}
	return []byte("fake"), nil
}

func (s *service) GetObjectTags(path string) (tags map[string]string, err error) {

	if path == "fail" {
		return tags, fmt.Errorf("error to get fake object tags")
	}

	return map[string]string{
		"SURT_LAST_SCAN":   time.Now().String(),
		"SURT_SCAN_STATUS": "CLEAN",
	}, nil
}

func (s *service) SetObjectTags(path string, tags map[string]string) (err error) {
	if path == "fail" {
		return fmt.Errorf("error to get fake object tags")
	}
	return nil
}

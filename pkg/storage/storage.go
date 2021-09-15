package storage

import (
	"fmt"

	"github.com/surt-io/surt/pkg/object"
)

type StorageService interface {
	GetObject(path string) (body []byte, err error)
	GetObjectTags(path string) (tags map[string]string, err error)
	SetObjectTags(path string, tags map[string]string) (err error)
}

type storage struct {
	provider StorageService
}

func New(provider StorageService) *storage {
	return &storage{
		provider: provider,
	}
}

func (s *storage) GetObject(o *object.Object) ([]byte, error) {
	var b []byte

	if o.Path == "" {
		return b, fmt.Errorf("object path is empty")
	}

	b, err := s.provider.GetObject(o.Path)
	if err != nil {
		return b, fmt.Errorf("GetObject: %w", err)
	}
	return b, nil
}

func (s *storage) GetObjectTags(o *object.Object) (map[string]string, error) {

	t := map[string]string{}

	if o.Path == "" {
		return t, fmt.Errorf("object path is empty")
	}

	t, err := s.provider.GetObjectTags(o.Path)
	if err != nil {
		return t, fmt.Errorf("GetObjectTags: %w", err)
	}
	return t, nil
}

func (s *storage) SetObjectTags(o *object.Object) error {

	if o.Path == "" {
		return fmt.Errorf("object path is empty")
	}
	if len(o.Tags) == 0 {
		return fmt.Errorf("object tags is empty")
	}

	err := s.provider.SetObjectTags(o.Path, o.Tags)
	if err != nil {
		return fmt.Errorf("SetObjectTags: %w", err)
	}
	return nil
}

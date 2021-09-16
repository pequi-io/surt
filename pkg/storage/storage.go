package storage

import (
	"fmt"

	"github.com/surt-io/surt/pkg/object"
)

type StorageService interface {
	GetObject(p string) (body []byte, err error)
	GetObjectTags(p string) (tags map[string]string, err error)
	SetObjectTags(p string, tags map[string]string) (err error)
}

type storage struct {
	provider StorageService
}

//New storage
func New(provider StorageService) *storage {
	return &storage{
		provider: provider,
	}
}

//GetObject returns storage service object content
func (s *storage) GetObject(o *object.Object) ([]byte, error) {
	var b []byte

	if o.RawPath == "" {
		return b, fmt.Errorf("object path is empty")
	}

	b, err := s.provider.GetObject(o.RawPath)
	if err != nil {
		return b, fmt.Errorf("GetObject: %w", err)
	}
	return b, nil
}

//GetObjectTags returns SURT tags from storage service object
func (s *storage) GetObjectTags(o *object.Object) (map[string]string, error) {

	t := map[string]string{}

	if o.RawPath == "" {
		return t, fmt.Errorf("object path is empty")
	}

	t, err := s.provider.GetObjectTags(o.RawPath)
	if err != nil {
		return t, fmt.Errorf("GetObjectTags: %w", err)
	}
	return t, nil
}

//SetObjectTags sets SURT tags to storage service object
func (s *storage) SetObjectTags(o *object.Object) error {

	if o.RawPath == "" {
		return fmt.Errorf("object path is empty")
	}
	if len(o.Tags) == 0 {
		return fmt.Errorf("object tags is empty")
	}

	err := s.provider.SetObjectTags(o.RawPath, o.Tags)
	if err != nil {
		return fmt.Errorf("SetObjectTags: %w", err)
	}
	return nil
}

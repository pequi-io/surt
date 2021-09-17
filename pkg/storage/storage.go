package storage

import (
	"fmt"

	"github.com/surt-io/surt/pkg/types"
)

//Storage interface
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
func (s *storage) GetObject(o *types.Object) (b []byte, err error) {

	// Validate if RawFilePath is empty
	if o.RawFilePath == "" {
		return b, fmt.Errorf("GetObject: object RawFilePath is empty")
	}

	// Get object content from provider service
	b, err = s.provider.GetObject(o.RawFilePath)
	if err != nil {
		return b, fmt.Errorf("GetObject: %w", err)
	}
	return b, nil
}

//GetObjectTags returns SURT tags from storage service object
func (s *storage) GetObjectTags(o *types.Object) (types.Tags, error) {

	t := types.Tags{}

	if o.RawFilePath == "" {
		return t, fmt.Errorf("object path is empty")
	}

	t, err := s.provider.GetObjectTags(o.RawFilePath)
	if err != nil {
		return t, fmt.Errorf("GetObjectTags: %w", err)
	}
	return t, nil
}

//SetObjectTags sets SURT tags to storage service object
func (s *storage) SetObjectTags(o *types.Object) error {

	if o.RawFilePath == "" {
		return fmt.Errorf("object path is empty")
	}
	if len(o.Tags) == 0 {
		return fmt.Errorf("object tags is empty")
	}

	err := s.provider.SetObjectTags(o.RawFilePath, o.Tags)
	if err != nil {
		return fmt.Errorf("SetObjectTags: %w", err)
	}
	return nil
}

package storage

import "github.com/surt-io/surt/pkg/entity"

//StorageService engine interface
type StorageService interface {
	GetObject(string) ([]byte, error)
	GetObjectTags(string) (entity.Tags, error)
	SetObjectTags(string, entity.Tags) (err error)
}

//Actions interface
type Actions interface {
	GetObject(*entity.Object) ([]byte, error)
	GetObjectTags(*entity.Object) (entity.Tags, error)
	SetObjectTags(*entity.Object) error
}

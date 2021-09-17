package storage

import (
	"github.com/surt-io/surt/pkg/types"
)

//StorageService engine interface
type StorageService interface {
	GetObject(string) ([]byte, error)
	GetObjectTags(string) (types.Tags, error)
	SetObjectTags(string, types.Tags) (err error)
}

//Actions interface
type Actions interface {
	GetObject(*types.Object) ([]byte, error)
	GetObjectTags(*types.Object) (types.Tags, error)
	SetObjectTags(*types.Object) error
}

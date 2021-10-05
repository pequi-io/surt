package entity

import "errors"

//ErrNotFound not found
var ErrNotFound = errors.New("not found")

//ErrEmptyRawFilePath empty RawFilePath
var ErrEmptyRawFilePath = errors.New("attribute RawFilePath is empty")

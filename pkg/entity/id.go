package entity

import (
	"github.com/google/uuid"
)

//ID type
type ID = uuid.UUID

//NewID creates a new ID
func NewID() ID {
	return ID(uuid.New())
}

//StringToID converts a string to ID type
func StringToID(s string) (ID, error) {
	id, err := uuid.Parse(s)
	return ID(id), err
}

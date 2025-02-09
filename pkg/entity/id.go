package entity

import "github.com/google/uuid"

type ID = uuid.UUID

// This must be stored into pkg because it can be used by other packages, if we store it into internal, it will be inaccessible
func NewID() ID {
	return uuid.New()
}

func ParseID(s string) (ID, error) {
	id, err := uuid.Parse(s)
	return ID(id), err
}

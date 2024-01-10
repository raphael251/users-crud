package entity

import "github.com/google/uuid"

type ID = uuid.UUID

func NewID() ID {
	return uuid.New()
}

func ParseID(value string) (ID, error) {
	id, err := uuid.Parse(value)
	return ID(id), err
}

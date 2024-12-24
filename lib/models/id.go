package models

import (
	"github.com/google/uuid"
	"github.com/oklog/ulid/v2"
)

type ID uuid.UUID

func GenerateID() ID {
	return ID(uuid.New())
}

func MustNew(s string) ID {
	return ID(uuid.MustParse(s))
}

func (id ID) String() string {
	return ulid.ULID(id).String()
}

func (id ID) UUID() uuid.UUID {
	return uuid.UUID(id)
}

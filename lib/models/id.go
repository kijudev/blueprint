package models

import (
	"fmt"

	"github.com/oklog/ulid/v2"
)

type ID ulid.ULID

func GenerateID() ID {
	return ID(ulid.Make())
}

func MustNew(str string) ID {
	return ID(ulid.MustParse(str))
}

func (id ID) String() string {
	return ulid.ULID(id).String()
}

func (id ID) UUIDString() string {
	bytes := ulid.ULID(id).Bytes()

	return fmt.Sprintf("%08x-%04x-%04x-%04x-%012x",
		bytes[0:4],
		bytes[4:6],
		bytes[6:8],
		bytes[8:10],
		bytes[10:16],
	)
}

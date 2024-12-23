package auth

import (
	"context"

	"github.com/oklog/ulid/v2"
)

type CoreSerive interface {
	CreateUser(ctx context.Context, params UserParams) (*User, error)
	GetUserByID(ctx context.Context, id ulid.ULID) (*User, error)
	DeleteUser(ctx context.Context, id ulid.ULID) (*User, error)
	UpdateUser(ctx context.Context, id ulid.ULID, params UserParams) (*User, error)
}

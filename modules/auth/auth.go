package auth

import (
	"context"

	"github.com/oklog/ulid/v2"
)

type Service interface {
	CreateUser(ctx context.Context, params UserParams) (*User, error)
	GetUserById(ctx context.Context, id ulid.ULID) (*User, error)
	DeleteUser(ctx context.Context, id ulid.ULID) (*User, error)
	UpdateUser(ctx context.Context, id ulid.ULID, params UserParams) (*User, error)
}

type Module interface {
	GetService() *Service
	Init(ctx context.Context) error
	Stop(ctx context.Context) error
}

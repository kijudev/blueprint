package auth

import (
	"context"
	"time"

	"github.com/kijudev/blueprint/lib"
)

type Module interface {
	lib.Module

	DataService() DataService
}

type DataService interface {
	CreateUser(ctx context.Context, params UserParams) (*User, error)
	GetUserByID(ctx context.Context, id lib.ID) (*User, error)
	DeleteUser(ctx context.Context, id lib.ID) (*User, error)
	UpdateUser(ctx context.Context, id lib.ID, params UserParams) (*User, error)
	RemoveUserPermissions(ctx context.Context, id lib.ID, permissions Permissions) (*User, error)
	AddUserPermissions(ctx context.Context, id lib.ID, permissions Permissions) (*User, error)

	CreateSession(ctx context.Context, params SessionParams, duration time.Duration) (*Session, error)
	GetSessionByID(ctx context.Context, id lib.ID) (*Session, error)
	GetSessionByUserID(ctx context.Context, id lib.ID) (*Session, error)
	DeleteSession(ctx context.Context, id lib.ID) (*Session, error)
	RefreshSession(ctx context.Context, id lib.ID, duration time.Duration) (*Session, error)

	GetAccountByID(ctx context.Context, id lib.ID) (*Account, error)
}

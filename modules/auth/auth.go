package auth

import "context"

type Service interface {
	CreateUser(ctx context.Context, params UserParams) (*User, error)
	GetUserById(ctx context.Context, id string) (*User, error)
	DeleteUser(ctx context.Context, id string) (*User, error)
	UpdateUser(ctx context.Context, id string, params UserParams) (*User, error)
}

type Module interface {
	GetService() *Service
	Init(ctx context.Context) error
	Stop(ctx context.Context) error
}

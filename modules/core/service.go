package core

import "context"

type CoreService interface {
	CreateUser(ctx context.Context, params UserParams) (*User, error)
	GetUserById(ctx context.Context, id string) (*User, error)
	DeleteUser(ctx context.Context, id string) (*User, error)
	UpdateUser(ctx context.Context, id string, params UserParams) (*User, error)
}

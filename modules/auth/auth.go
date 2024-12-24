package auth

import (
	"context"

	"github.com/kijudev/blueprint/lib/models"
	"github.com/kijudev/blueprint/lib/modules"
)

type Module interface {
	modules.Module

	CoreSerive() CoreSerive
}

type CoreSerive interface {
	CreateUser(ctx context.Context, params UserParams) (*User, error)
	GetUserByID(ctx context.Context, id models.ID) (*User, error)
	DeleteUser(ctx context.Context, id models.ID) (*User, error)
	UpdateUser(ctx context.Context, id models.ID, params UserParams) (*User, error)
}

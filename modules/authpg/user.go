package authpg

import (
	"github.com/jackc/pgx/v5"
	"github.com/kijudev/blueprint/modules/auth"
)

func NamedArgsFromUser(user auth.User) pgx.NamedArgs {
	return pgx.NamedArgs{
		"id":          user.ID.UUID(),
		"email":       user.Email,
		"name":        user.Name,
		"permissions": user.Permissions.String(),
		"created_at":  user.CreatedAt,
		"updated_at":  user.UpdatedAt,
	}
}

func NameArgsFromUserParams(params auth.UserParams) pgx.NamedArgs {
	return pgx.NamedArgs{
		"email":       params.Email,
		"name":        params.Name,
		"permissions": params.Permissions.String(),
	}
}

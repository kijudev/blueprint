package authpg

import (
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/kijudev/blueprint/modules/auth"
	"github.com/oklog/ulid/v2"
)

type UserPG struct {
	ID          pgtype.UUID
	Email       pgtype.Text
	Name        pgtype.Text
	Permissions pgtype.Text
	CreatedAt   pgtype.Timestamptz
	UpdatedAt   pgtype.Timestamptz
}

func NewUserPG(u auth.User) *UserPG {
	t := new(UserPG)

	t.ID.Scan(u.ID.String())
	t.Email.Scan(u.Email)
	t.Name.Scan(u.Name)
	t.Permissions.Scan(u.Permissions.AsString())
	t.CreatedAt.Scan(u.CreatedAt)
	t.UpdatedAt.Scan(u.UpdatedAt)

	return t
}

func (u *UserPG) AsModel() *auth.User {
	return &auth.User{
		ID:          ulid.MustParse(u.ID.String()),
		Email:       u.Email.String,
		Name:        u.Name.String,
		Permissions: *auth.NewPermissions(u.Permissions.String),
		CreatedAt:   u.CreatedAt.Time,
		UpdatedAt:   u.UpdatedAt.Time,
	}
}

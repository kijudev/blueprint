package auth

import (
	"time"

	"github.com/oklog/ulid/v2"
)

type User struct {
	ID          ulid.ULID
	Email       string
	Name        string
	Permissions Permissions
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type UserData struct {
	ID          string `json:"id"`
	Email       string `json:"email"`
	Name        string `json:"name"`
	Permissions string `json:"permissions"`
	CreatedAt   int64  `json:"createdAt"`
	UpdatedAt   int64  `json:"updatedAt"`
}

type UserParams struct {
	Email       string
	Name        string
	Permissions Permissions
}

type UserParamsData struct {
	Email       string `json:"email"`
	Name        string `json:"name"`
	Permissions string `json:"permissions"`
}

type UserFilter struct {
	ID    *ulid.ULID
	Email *string
	Name  *string
}

func (u *User) AsData() UserData {
	return UserData{
		ID:          u.ID.String(),
		Email:       u.Email,
		Name:        u.Name,
		Permissions: u.Permissions.AsString(),
		CreatedAt:   u.CreatedAt.Unix(),
		UpdatedAt:   u.UpdatedAt.Unix(),
	}
}

func (u *UserParamsData) AsModel() UserParams {
	return UserParams{
		Email:       u.Email,
		Name:        u.Name,
		Permissions: *NewPermissionsFromString(u.Permissions),
	}
}

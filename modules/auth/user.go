package auth

import (
	"time"

	"github.com/kijudev/blueprint/lib/validation"
)

type User struct {
	ID          string
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
	ID    *string
	Email *string
	Name  *string
}

func (u *User) AsData() UserData {
	return UserData{
		ID:          u.ID,
		Email:       u.Email,
		Name:        u.Name,
		Permissions: u.Permissions.AsString(),
		CreatedAt:   u.CreatedAt.Unix(),
		UpdatedAt:   u.UpdatedAt.Unix(),
	}
}

func (u *UserData) AsModel() User {
	return User{
		ID:          u.ID,
		Email:       u.Email,
		Name:        u.Name,
		Permissions: *NewPermissionsFromString(u.Permissions),
		CreatedAt:   time.Unix(u.CreatedAt, 0).UTC(),
		UpdatedAt:   time.Unix(u.UpdatedAt, 0).UTC(),
	}
}

func (u *UserParams) AsData() UserParamsData {
	return UserParamsData{
		Email:       u.Email,
		Name:        u.Name,
		Permissions: u.Permissions.AsString(),
	}
}

func (u *UserParams) Validate() error {
	c := validation.NewCollection()

	c.Add("email", validation.String(u.Email).Email())
	c.Add("name", validation.String(u.Name).MinLength(4))

	return c.Resolve()
}

func (u *UserParamsData) AsModel() UserParams {
	return UserParams{
		Email:       u.Email,
		Name:        u.Name,
		Permissions: *NewPermissionsFromString(u.Permissions),
	}
}

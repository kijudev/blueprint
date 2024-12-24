package auth

import (
	"time"

	"github.com/kijudev/blueprint/lib/models"
	"github.com/kijudev/blueprint/lib/validation"
)

type User struct {
	ID          models.ID
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
	Permissions string `json:"permiassions"`
}

type UserFilter struct {
	ID    *models.ID
	Email *string
	Name  *string
}

func NewUser(params UserParams) *User {
	now := time.Now()

	return &User{
		ID:          models.GenerateID(),
		Email:       params.Email,
		Name:        params.Name,
		Permissions: params.Permissions,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}

func (u *User) Data() *UserData {
	return &UserData{
		ID:          u.ID.String(),
		Email:       u.Email,
		Name:        u.Name,
		Permissions: u.Permissions.AsString(),
		CreatedAt:   u.CreatedAt.Unix(),
		UpdatedAt:   u.UpdatedAt.Unix(),
	}
}

func (u *User) Validate() error {
	c := validation.NewCollection()

	c.Add("email", validation.String(u.Email).Email())
	c.Add("name", validation.String(u.Name).NotEmpty().MaxLength(255))
	c.Add("name", validation.String(u.Name).NotEmpty().MaxLength(255))

	return c.Resolve()
}

func (u *UserParamsData) Model() *UserParams {
	return &UserParams{
		Email:       u.Email,
		Name:        u.Name,
		Permissions: *NewPermissions(u.Permissions),
	}
}

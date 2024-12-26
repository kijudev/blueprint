package auth

import (
	"time"

	"github.com/kijudev/blueprint/lib"
)

type User struct {
	ID          lib.ID
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
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
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
	EqID    *lib.ID
	EqEmail *string
	EqName  *string
}

func NewUser(params UserParams) *User {
	now := time.Now().UTC()

	return &User{
		ID:          lib.GenerateID(),
		Email:       params.Email,
		Name:        params.Name,
		Permissions: params.Permissions,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
}

func (u *User) Data() *UserData {
	return &UserData{
		ID:          u.ID.UUID().String(),
		Email:       u.Email,
		Name:        u.Name,
		Permissions: u.Permissions.String(),
		CreatedAt:   u.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   u.UpdatedAt.Format(time.RFC3339),
	}
}

func (u *User) Validate() error {
	c := lib.NewValCollection()

	c.Add("email", lib.ValString(u.Email).Email())
	c.Add("name", lib.ValString(u.Name).NotEmpty().MaxLength(255))
	c.Add("name", lib.ValString(u.Name).NotEmpty().MaxLength(255))

	return c.Resolve()
}

func (u *UserParamsData) Model() *UserParams {
	return &UserParams{
		Email:       u.Email,
		Name:        u.Name,
		Permissions: *NewPermissions(u.Permissions),
	}
}

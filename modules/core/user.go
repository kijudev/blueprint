package core

import (
	"time"

	"github.com/kijudev/blueprint/lib/validation"
)

type User struct {
	Email       string
	Name        string
	Permissions Permissions
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type UserData struct {
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

func (user *User) AsData() UserData {
	return UserData{
		Email:       user.Email,
		Name:        user.Name,
		Permissions: user.Permissions.AsString(),
		CreatedAt:   user.CreatedAt.Unix(),
		UpdatedAt:   user.UpdatedAt.Unix(),
	}
}

func (user *User) Validate() error {
	return validation.Combine(
		validation.ValidateField("email", validation.Email(user.Email)),
		validation.ValidateField("name", validation.MinLength(user.Name, 8)),
	)
}

func (userData *UserData) AsModel() User {
	return User{
		Email:       userData.Email,
		Name:        userData.Name,
		Permissions: *NewPermissionsFromString(userData.Permissions),
		CreatedAt:   time.Unix(userData.CreatedAt, 0).UTC(),
		UpdatedAt:   time.Unix(userData.UpdatedAt, 0).UTC(),
	}
}

func (userParams *UserParams) AsData() UserParamsData {
	return UserParamsData{
		Email:       userParams.Email,
		Name:        userParams.Name,
		Permissions: userParams.Permissions.AsString(),
	}
}

func (userParamsData *UserParamsData) AsModel() UserParams {
	return UserParams{
		Email:       userParamsData.Email,
		Name:        userParamsData.Name,
		Permissions: *NewPermissionsFromString(userParamsData.Permissions),
	}
}

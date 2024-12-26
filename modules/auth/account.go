package auth

import "github.com/kijudev/blueprint/lib"

type Account struct {
	ID   lib.ID
	User User
	Auth struct {
		Session Session
	}
}

// User.ID = Account.ID
type AccountData struct {
	ID   string   `json:"id"`
	User UserData `json:"user"`
	Auth struct {
		Session SessionData `json:"session"`
	} `json:"auth"`
}

type AccountFilter struct {
	EqID *lib.ID
}

func (a *Account) Data() *AccountData {
	return &AccountData{
		ID:   a.ID.String(),
		User: *a.User.Data(),
		Auth: struct {
			Session SessionData `json:"session"`
		}{
			Session: *a.Auth.Session.Data(),
		},
	}
}

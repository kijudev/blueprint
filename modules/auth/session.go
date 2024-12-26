package auth

import (
	"time"

	"github.com/kijudev/blueprint/lib"
)

type Session struct {
	ID        lib.ID
	UserID    lib.ID
	ExpiresAt time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

type SessionData struct {
	ID        string `json:"id"`
	UserID    string `json:"userId"`
	ExpiresAt string `json:"expiresAt"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type SessionParams struct {
	UserID lib.ID
}

type SessionFilter struct {
	EqID     *lib.ID
	EqUserID *lib.ID
}

func NewSession(params SessionParams, duration time.Duration) *Session {
	now := time.Now().UTC()

	return &Session{
		ID:        lib.GenerateID(),
		UserID:    params.UserID,
		ExpiresAt: now.Add(duration),
		CreatedAt: now,
		UpdatedAt: now,
	}
}

func (s *Session) Data() *SessionData {
	return &SessionData{
		ID:        s.ID.UUID().String(),
		UserID:    s.UserID.UUID().String(),
		ExpiresAt: s.ExpiresAt.Format(time.RFC3339),
		CreatedAt: s.CreatedAt.Format(time.RFC3339),
		UpdatedAt: s.UpdatedAt.Format(time.RFC3339),
	}
}

func (s *Session) Expired() bool {
	return time.Now().UTC().After(s.ExpiresAt)
}

func (s *Session) Refresh(offset time.Duration) {
	now := time.Now().UTC()

	s.ExpiresAt = now.Add(offset)
	s.UpdatedAt = now
}

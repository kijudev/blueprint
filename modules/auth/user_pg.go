package auth

import (
	"context"

	"github.com/kijudev/blueprint/lib/database"
)

var userMigrations = []database.Migration{
	{
		Up: `
			CREATE TABLE IF NOT EXISTS
			users (
				id UUID PRIMARY KEY,
				email VARCHAR(255) NOT NULL,
				name VARCHAR(255) NOT NULL,
				permissions TEXT NOT NULL,
				created_at TIMESTAMP NOT NULL,
				updated_at TIMESTAMP NOT NULL
			);
		`,
		Down: `
			DROP TABLE IF EXISTS users;
		`,
	},
}

func (m *ModulePg) CreateUser(ctx context.Context, params UserParams) (*User, error) {
	user := NewUserFromParams(params)

	tx, err := m.dbpgModule.GetPool().Begin(ctx)
	if err != nil {
		return nil, err
	}

	_, err = tx.Exec(ctx, `
			INSERT INTO users (id, email, name, permissions, created_at, updated_at)
			VALUES ($1, $2, $3, $4, $5, $6);
		`,
		user.ID, user.Email, user.Name, user.Permissions.AsString(), user.CreatedAt, user.UpdatedAt,
	)
	if err != nil {
		return nil, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return nil, err
	}

	return user, nil
}

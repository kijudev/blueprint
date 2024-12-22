package auth

import (
	"github.com/kijudev/blueprint/lib/databases"
)

var userMigrations = []databases.Migration{
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

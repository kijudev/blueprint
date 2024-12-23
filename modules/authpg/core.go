package authpg

import (
	"context"
	"errors"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/kijudev/blueprint/lib/services"
	"github.com/kijudev/blueprint/modules/auth"
	"github.com/kijudev/blueprint/modules/dbpg"
	"github.com/oklog/ulid/v2"
)

type CoreService struct {
	db *dbpg.DBService
}

func NewCoreService(db *dbpg.DBService) *CoreService {
	return &CoreService{db: db}
}

func (s *CoreService) CreateUser(ctx context.Context, params auth.UserParams) (*auth.User, error) {
	user := auth.NewUser(params)
	pguser := NewUserPG(*user)

	query := `
		INSERT INTO users
			(id, email, name, permissions, created_at, updated_at)
		VALUES
			($1, $2, $3, $4, $5, $6);
	`

	tag, err := s.db.Pool.Exec(ctx, query, pguser.ID, pguser.Email, pguser.Name, pguser.Permissions, pguser.CreatedAt, pguser.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("(authpg.CoreService.CreateUser) %w; %w", err, services.ErrorDependencyFailed)
	}

	if !tag.Insert() {
		return nil, fmt.Errorf("(authpg.CoreService.CreateUser) %w; Could not insert user", services.ErrorDependencyFailed)
	}

	return user, nil
}

func (s *CoreService) GetUserByID(ctx context.Context, id ulid.ULID) (*auth.User, error) {
	query := `
		SELECT
			id, email, name, permissions, created_at, updated_at
		FROM
			users
		WHERE
			id = $1;
	`

	row := s.db.Pool.QueryRow(ctx, query, id)

	pguser := new(UserPG)
	err := row.Scan(&pguser.ID, &pguser.Email, &pguser.Name, &pguser.Permissions, &pguser.CreatedAt, &pguser.UpdatedAt)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("(authpg.CoreSerivce.GetUserByID) %w", services.ErrorNotFound)
		}

		return nil, fmt.Errorf("(authpg.CoreService.GetUserByID) %w; %w", err, services.ErrorDependencyFailed)
	}

	user := pguser.AsModel()
	return user, nil
}

func (s *CoreService) findUser(ctx context.Context, tx pgx.Tx, filter auth.UserFilter) (*auth.User, error) {
	return nil, errors.New("not implemented")
}

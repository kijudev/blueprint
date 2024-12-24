package authpg

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/jackc/pgx/v5"
	"github.com/kijudev/blueprint/lib/models"
	"github.com/kijudev/blueprint/lib/services"
	"github.com/kijudev/blueprint/modules/auth"
	"github.com/kijudev/blueprint/modules/dbpg"
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
		return nil, fmt.Errorf("(authpg.CoreService.CreateUser) %w; %w", err, services.ErrDependencyFailed)
	}

	if !tag.Insert() {
		return nil, fmt.Errorf("(authpg.CoreService.CreateUser) %w; Could not insert user", services.ErrDependencyFailed)
	}

	return user, nil
}

func (s *CoreService) GetUserByID(ctx context.Context, id models.ID) (*auth.User, error) {
	tx, err := s.db.Pool.Begin(ctx)
	if err != nil {
		return nil, fmt.Errorf("(authpg.CoreService.GetUserByID) %w; %w", services.ErrDependencyFailed, err)
	}
	defer tx.Commit(ctx)

	user, err := s.findUser(ctx, tx, auth.UserFilter{ID: &id})
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *CoreService) findUser(ctx context.Context, tx pgx.Tx, filter auth.UserFilter) (*auth.User, error) {
	query := `
		SELECT
			id, email, name, permissions, created_at, updated_at
		FROM
			users
	`

	var args []any

	if filter.ID != nil {
		query += " WHERE id = $" + strconv.Itoa(len(args)+1)
		args = append(args, filter.ID.UUID())
	}
	if filter.Email != nil {
		query += " WHERE email = $" + strconv.Itoa(len(args)+1)
		args = append(args, filter.Email)
	}
	if filter.Name != nil {
		query += " WHERE name = $" + strconv.Itoa(len(args)+1)
		args = append(args, filter.Name)
	}

	row := tx.QueryRow(ctx, query, args...)

	user := new(UserPG)

	err := row.Scan(&user.ID, &user.Email, &user.Name, &user.Permissions, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("(authpg.CoreService.findUser) %w", services.ErrNotFound)
		}

		return nil, fmt.Errorf("(authpg.CoreService.findUser) %w; %w", services.ErrDependencyFailed, err)
	}

	return user.Model(), nil
}

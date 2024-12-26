package authpg

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kijudev/blueprint/lib"
	"github.com/kijudev/blueprint/modules/auth"
)

type DataService struct {
	db *pgxpool.Pool
}

func (s *DataService) CreateUser(ctx context.Context, params auth.UserParams) (*auth.User, error) {
	e := errors.New("(authpg.DataService.CreateUser)")

	query := `
		INSERT INTO
			users (id, email, name, permissions, created_at, updated_at)
		VALUES
			(@id, @email, @name, @permissions, @created_at, @updated_at)
	`

	user := auth.NewUser(params)

	_, err := s.db.Exec(ctx, query, pgx.NamedArgs{
		"id":          user.ID.UUID(),
		"email":       user.Email,
		"name":        user.Name,
		"permissions": user.Permissions.String(),
		"created_at":  user.CreatedAt,
		"updated_at":  user.UpdatedAt,
	})
	if err != nil {
		return nil, lib.JoinErrors(e, lib.ErrDatasourceFailed, err)
	}

	return user, nil
}

func (s *DataService) GetUserByID(ctx context.Context, id lib.ID) (*auth.User, error) {
	e := errors.New("(authpg.DataService.GetUserByID)")

	tx, err := s.db.Begin(ctx)
	if err != nil {
		return nil, lib.JoinErrors(e, lib.ErrDatasourceFailed, err)
	}

	defer func() {
		if err := tx.Rollback(ctx); err != nil && !errors.Is(err, pgx.ErrTxClosed) {
			fmt.Println("TX Rollback: ", err)
		}
	}()

	users, err := s.findUsers(ctx, tx, auth.UserFilter{EqID: &id}, lib.Pagination{})
	if err != nil {
		return nil, lib.JoinErrors(e, err)
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, lib.JoinErrors(e, lib.ErrDatasourceFailed, err)
	}

	if len(users) == 0 {
		return nil, lib.JoinErrors(e, lib.ErrNotFound)
	}

	return &users[0], nil
}

func (s *DataService) DeleteUser(ctx context.Context, id lib.ID) (*auth.User, error) {
	e := errors.New("(authpg.DataService.DeleteUser)")

	tx, err := s.db.Begin(ctx)
	if err != nil {
		return nil, lib.JoinErrors(e, lib.ErrDatasourceFailed, err)
	}

	defer func() {
		if err := tx.Rollback(ctx); err != nil && !errors.Is(err, pgx.ErrTxClosed) {
			fmt.Println("TX Rollback: ", err)
		}
	}()

	users, err := s.findUsers(ctx, tx, auth.UserFilter{EqID: &id}, lib.Pagination{})
	if err != nil {
		return nil, lib.JoinErrors(e, err)
	}

	if len(users) == 0 {
		return nil, lib.JoinErrors(e, lib.ErrNotFound)
	}

	user := &users[0]

	query := `DELETE FROM users WHERE id = $1`
	_, err = tx.Exec(ctx, query, user.ID.UUID())
	if err != nil {
		return nil, lib.JoinErrors(e, lib.ErrDatasourceFailed, err)
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, lib.JoinErrors(e, lib.ErrDatasourceFailed, err)
	}

	return user, nil
}

func (s *DataService) UpdateUser(ctx context.Context, id lib.ID, params auth.UserParams) (*auth.User, error) {
	e := errors.New("(authpg.DataService.UpdateUser)")

	tx, err := s.db.Begin(ctx)
	if err != nil {
		return nil, lib.JoinErrors(e, lib.ErrDatasourceFailed, err)
	}

	defer func() {
		if err := tx.Rollback(ctx); err != nil && !errors.Is(err, pgx.ErrTxClosed) {
			fmt.Println("TX Rollback: ", err)
		}
	}()

	users, err := s.findUsers(ctx, tx, auth.UserFilter{EqID: &id}, lib.Pagination{})
	if err != nil {
		return nil, lib.JoinErrors(e, err)
	}

	if len(users) == 0 {
		return nil, lib.JoinErrors(e, lib.ErrNotFound)
	}

	user := &users[0]

	query := `
		UPDATE
			user
		SET
			email = @email,
			name = @name,
			permissions = @permissions,
			updated_at = @updated_at
			create_at = @created_at
		WHERE
			id = @id
	`

	user.Email = params.Email
	user.Name = params.Name
	user.Permissions = params.Permissions
	user.UpdatedAt = time.Now().UTC()

	_, err = tx.Exec(ctx, query, pgx.NamedArgs{
		"email":       user.Email,
		"name":        user.Name,
		"permissions": user.Permissions.String(),
		"updated_at":  user.UpdatedAt,
		"created_at":  user.CreatedAt,
		"id":          user.ID.UUID(),
	})
	if err != nil {
		return nil, lib.JoinErrors(e, lib.ErrDatasourceFailed, err)
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, lib.JoinErrors(e, lib.ErrDatasourceFailed, err)
	}

	return user, nil
}

func (s *DataService) RemoveUserPermissions(ctx context.Context, id lib.ID, permissions auth.Permissions) (*auth.User, error) {
	e := errors.New("(authpg.DataService.RemoveUserPermissions)")

	tx, err := s.db.Begin(ctx)
	if err != nil {
		return nil, lib.JoinErrors(e, lib.ErrDatasourceFailed, err)
	}

	defer func() {
		if err := tx.Rollback(ctx); err != nil && !errors.Is(err, pgx.ErrTxClosed) {
			fmt.Println("TX Rollback: ", err)
		}
	}()

	users, err := s.findUsers(ctx, tx, auth.UserFilter{EqID: &id}, lib.Pagination{})
	if err != nil {
		return nil, lib.JoinErrors(e, err)
	}

	if len(users) == 0 {
		return nil, lib.JoinErrors(e, lib.ErrNotFound)
	}

	user := &users[0]
	user.Permissions.Remove(permissions.Rules()...)

	query := `UPDATE users SET permissions = @permissions WHERE id = @id`
	_, err = tx.Exec(ctx, query, pgx.NamedArgs{"permissions": permissions.String(), "id": user.ID.UUID()})
	if err != nil {
		return nil, lib.JoinErrors(e, lib.ErrDatasourceFailed, err)
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, lib.JoinErrors(e, lib.ErrDatasourceFailed, err)
	}

	return user, nil
}

func (s *DataService) AddUserPermissions(ctx context.Context, id lib.ID, permissions auth.Permissions) (*auth.User, error) {
	e := errors.New("(authpg.DataService.AddUserPermissions)")

	tx, err := s.db.Begin(ctx)
	if err != nil {
		return nil, lib.JoinErrors(e, lib.ErrDatasourceFailed, err)
	}

	defer func() {
		if err := tx.Rollback(ctx); err != nil && !errors.Is(err, pgx.ErrTxClosed) {
			fmt.Println("TX Rollback: ", err)
		}
	}()

	users, err := s.findUsers(ctx, tx, auth.UserFilter{EqID: &id}, lib.Pagination{})
	if err != nil {
		return nil, lib.JoinErrors(e, err)
	}

	if len(users) == 0 {
		return nil, lib.JoinErrors(e, lib.ErrNotFound)
	}

	user := &users[0]
	user.Permissions.Add(permissions.Rules()...)

	query := `UPDATE users SET permissions = @permissions WHERE id = @id`
	_, err = tx.Exec(ctx, query, pgx.NamedArgs{"permissions": permissions.String(), "id": user.ID.UUID()})
	if err != nil {
		return nil, lib.JoinErrors(e, lib.ErrDatasourceFailed, err)
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, lib.JoinErrors(e, lib.ErrDatasourceFailed, err)
	}

	return user, nil
}

func (s *DataService) findUsers(ctx context.Context, tx pgx.Tx, filter auth.UserFilter, pagination lib.Pagination) ([]auth.User, error) {
	e := errors.New("(authpg.DataService.findUsers)")

	query := `
		SELECT
			id, email, name, permissions, created_at, updated_at
		FROM
			users
		WHERE
			1 = 1
	`

	args := pgx.NamedArgs{}

	if filter.EqID != nil {
		query += " AND id = @id"
		args["id"] = filter.EqID.UUID()
	}

	if filter.EqEmail != nil {
		query += " AND email = @email"
		args["email"] = *filter.EqEmail
	}

	if filter.EqName != nil {
		query += " AND name = @name"
		args["name"] = *filter.EqName
	}

	if pagination.Limit > 0 {
		query += " LIMIT @limit"
		args["limit"] = pagination.Limit
	}

	if pagination.Offset > 0 {
		query += " OFFSET @offset"
		args["offset"] = pagination.Offset
	}

	rows, err := tx.Query(ctx, query, args)
	if err != nil {
		return nil, lib.JoinErrors(e, lib.ErrDatasourceFailed, err)
	}
	defer rows.Close()

	users := []auth.User{}
	for rows.Next() {
		user := auth.User{}
		var permissions string

		err := rows.Scan(
			&user.ID,
			&user.Email,
			&user.Name,
			&permissions,
			&user.CreatedAt,
			&user.UpdatedAt,
		)

		if err != nil {
			return nil, lib.JoinErrors(e, lib.ErrDatasourceFailed, err)
		}

		user.Permissions = *auth.NewPermissions(permissions)

		users = append(users, user)
	}

	return users, nil
}

func (s *DataService) CreateSession(ctx context.Context, params auth.SessionParams, duration time.Duration) (*auth.Session, error) {
	e := errors.New("(authpg.DataService.CreateSession)")

	query := `
		INSERT INTO
			sessions (id, user_id, expires_at, created_at, updated_at)
		VALUES
			(@id, @user_id, @expires_at, @created_at, @updated_at)
	`

	session := auth.NewSession(params, 24*time.Hour)

	_, err := s.db.Exec(ctx, query, pgx.NamedArgs{
		"id":         session.ID.UUID(),
		"user_id":    session.UserID.UUID(),
		"expires_at": session.ExpiresAt,
		"created_at": session.CreatedAt,
		"updated_at": session.UpdatedAt,
	})
	if err != nil {
		return nil, lib.JoinErrors(e, lib.ErrDatasourceFailed, err)
	}

	return session, nil
}

func (s *DataService) GetSessionByID(ctx context.Context, id lib.ID) (*auth.Session, error) {
	e := errors.New("(authpg.DataService.GetSessionByID)")

	tx, err := s.db.Begin(ctx)
	if err != nil {
		return nil, lib.JoinErrors(e, lib.ErrDatasourceFailed, err)
	}

	defer func() {
		if err := tx.Rollback(ctx); err != nil && !errors.Is(err, pgx.ErrTxClosed) {
			fmt.Println("TX Rollback: ", err)
		}
	}()

	sessions, err := s.findSessions(ctx, tx, auth.SessionFilter{EqID: &id}, lib.Pagination{})
	if err != nil {
		return nil, lib.JoinErrors(e, err)
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, lib.JoinErrors(e, lib.ErrDatasourceFailed, err)
	}

	if len(sessions) == 0 {
		return nil, lib.JoinErrors(e, lib.ErrNotFound)
	}

	return &sessions[0], nil
}

func (s *DataService) GetSessionByUserID(ctx context.Context, userID lib.ID) (*auth.Session, error) {
	e := errors.New("(authpg.DataService.GetSessionByUserID)")

	tx, err := s.db.Begin(ctx)
	if err != nil {
		return nil, lib.JoinErrors(e, lib.ErrDatasourceFailed, err)
	}

	defer func() {
		if err := tx.Rollback(ctx); err != nil && !errors.Is(err, pgx.ErrTxClosed) {
			fmt.Println("TX Rollback: ", err)
		}
	}()

	sessions, err := s.findSessions(ctx, tx, auth.SessionFilter{EqUserID: &userID}, lib.Pagination{})
	if err != nil {
		return nil, lib.JoinErrors(e, err)
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, lib.JoinErrors(e, lib.ErrDatasourceFailed, err)
	}

	if len(sessions) == 0 {
		return nil, lib.JoinErrors(e, lib.ErrNotFound)
	}

	return &sessions[0], nil
}

func (s *DataService) DeleteSession(ctx context.Context, id lib.ID) (*auth.Session, error) {
	e := errors.New("(authpg.DataService.DeleteSession)")

	tx, err := s.db.Begin(ctx)
	if err != nil {
		return nil, lib.JoinErrors(e, lib.ErrDatasourceFailed, err)
	}

	defer func() {
		if err := tx.Rollback(ctx); err != nil && !errors.Is(err, pgx.ErrTxClosed) {
			fmt.Println("TX Rollback: ", err)
		}
	}()

	sessions, err := s.findSessions(ctx, tx, auth.SessionFilter{EqID: &id}, lib.Pagination{})
	if err != nil {
		return nil, lib.JoinErrors(e, err)
	}

	if len(sessions) == 0 {
		return nil, lib.JoinErrors(e, lib.ErrNotFound)
	}

	session := &sessions[0]

	query := `DELETE FROM sessions WHERE id = $1`
	_, err = tx.Exec(ctx, query, session.ID.UUID())
	if err != nil {
		return nil, lib.JoinErrors(e, lib.ErrDatasourceFailed, err)
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, lib.JoinErrors(e, lib.ErrDatasourceFailed, err)
	}

	return session, nil
}

func (s *DataService) RefreshSession(ctx context.Context, id lib.ID, duration time.Duration) (*auth.Session, error) {
	e := errors.New("(authpg.DataService.RefreshSession)")

	tx, err := s.db.Begin(ctx)
	if err != nil {
		return nil, lib.JoinErrors(e, lib.ErrDatasourceFailed, err)
	}

	defer func() {
		if err := tx.Rollback(ctx); err != nil && !errors.Is(err, pgx.ErrTxClosed) {
			fmt.Println("TX Rollback: ", err)
		}
	}()

	sessions, err := s.findSessions(ctx, tx, auth.SessionFilter{EqID: &id}, lib.Pagination{})
	if err != nil {
		return nil, lib.JoinErrors(e, err)
	}

	if len(sessions) == 0 {
		return nil, lib.JoinErrors(e, lib.ErrNotFound)
	}

	session := &sessions[0]
	session.ExpiresAt = time.Now().UTC().Add(duration)
	session.UpdatedAt = time.Now().UTC()

	query := `
		UPDATE
			sessions
		SET
			expires_at = @expires_at,
			updated_at = @updated_at
		WHERE
			id = @id
	`

	_, err = tx.Exec(ctx, query, pgx.NamedArgs{
		"expires_at": session.ExpiresAt,
		"updated_at": session.UpdatedAt,
		"id":         session.ID.UUID(),
	})
	if err != nil {
		return nil, lib.JoinErrors(e, lib.ErrDatasourceFailed, err)
	}

	if err := tx.Commit(ctx); err != nil {
		return nil, lib.JoinErrors(e, lib.ErrDatasourceFailed, err)
	}

	return session, nil
}

func (s *DataService) findSessions(ctx context.Context, tx pgx.Tx, filter auth.SessionFilter, pagination lib.Pagination) ([]auth.Session, error) {
	e := errors.New("(authpg.DataService.findSessions)")

	query := `
		SELECT
			id, user_id, expires_at, created_at, updated_at
		FROM
			sessions
		WHERE
			1 = 1
	`

	args := pgx.NamedArgs{}

	if filter.EqID != nil {
		query += " AND id = @id"
		args["id"] = filter.EqID.UUID()
	}

	if filter.EqUserID != nil {
		query += " AND user_id = @user_id"
		args["user_id"] = filter.EqUserID.UUID()
	}

	if pagination.Limit > 0 {
		query += " LIMIT @limit"
		args["limit"] = pagination.Limit
	}

	if pagination.Offset > 0 {
		query += " OFFSET @offset"
		args["offset"] = pagination.Offset
	}

	rows, err := tx.Query(ctx, query, args)
	if err != nil {
		return nil, lib.JoinErrors(e, lib.ErrDatasourceFailed, err)
	}
	defer rows.Close()

	sessions := []auth.Session{}
	for rows.Next() {
		session := auth.Session{}

		err := rows.Scan(
			&session.ID,
			&session.UserID,
			&session.ExpiresAt,
			&session.CreatedAt,
			&session.UpdatedAt,
		)

		if err != nil {
			return nil, lib.JoinErrors(e, lib.ErrDatasourceFailed, err)
		}

		sessions = append(sessions, session)
	}

	return sessions, nil
}

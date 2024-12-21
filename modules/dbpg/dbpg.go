package dbpg

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Module struct {
	pool    *pgxpool.Pool
	connStr string
	status  string
}

var (
	StatusIdle    = "IDLE"
	StatusRunning = "RUNNING"
	StatusStopped = "STOPPED"
)

func NewModule(connStr string) *Module {
	return &Module{
		connStr: connStr,
		status:  StatusIdle,
	}
}

func (m *Module) GetStatus() string {
	return m.status
}

func (m *Module) Init(ctx context.Context) error {
	pool, err := pgxpool.New(ctx, m.connStr)
	if err != nil {
		return err
	}

	m.pool = pool
	m.status = StatusRunning
	return nil
}

func (m *Module) Stop(ctx context.Context) error {
	if m.status != StatusRunning {
		return errors.New("A running module cannot be stopped")
	}

	m.pool.Close()
	m.status = StatusStopped
	return nil
}

package dbpg

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Module struct {
	pool    *pgxpool.Pool
	connStr string
}

func NewModule(connStr string) *Module {
	return &Module{
		connStr: connStr,
	}
}

func (m *Module) Init(ctx context.Context) error {
	pool, err := pgxpool.New(ctx, m.connStr)
	if err != nil {
		return err
	}

	m.pool = pool
	return nil
}

func (m *Module) Stop(ctx context.Context) error {
	m.pool.Close()
	return nil
}

func (m *Module) GetPool() *pgxpool.Pool {
	return m.pool
}

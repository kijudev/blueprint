package dbpg

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kijudev/blueprint/lib/modules"
)

type Module struct {
	tag    string
	status string

	connStr string

	DBService *DBService
}

type DBService struct {
	*pgxpool.Pool
}

const Tag = "POSTGRES"

func New(connStr string) *Module {
	return &Module{
		tag:    Tag,
		status: modules.StatusCodePreInit,

		connStr:   connStr,
		DBService: new(DBService),
	}
}

func (m *Module) Tag() string {
	return m.tag
}

func (m *Module) Status() string {
	return m.status
}

func (m *Module) Init(ctx context.Context) error {
	if m.status != modules.StatusCodePreInit {
		return fmt.Errorf("(pgdb.Module.Init) %w", modules.ErrorInvalidStatus)
	}

	pool, err := pgxpool.New(ctx, m.connStr)
	if err != nil {
		return fmt.Errorf("(dbpg.Module.Init) %w; %w", modules.ErrorInitFailed, err)
	}

	m.DBService.Pool = pool
	return nil
}

func (m *Module) MustInit(ctx context.Context) {
	if err := m.Init(ctx); err != nil {
		panic(err)
	}
}

func (m *Module) Stop(ctx context.Context) error {
	if m.status != modules.StatusCodeActive {
		return fmt.Errorf("(dbpg.Module.Stop) %w", modules.ErrorInvalidStatus)
	}

	m.DBService.Close()
	return nil
}

func (m *Module) MustStop(ctx context.Context) {
	if err := m.Stop(ctx); err != nil {
		panic(err)
	}
}

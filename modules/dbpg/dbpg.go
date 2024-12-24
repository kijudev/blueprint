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

	deps     ModuleDeps
	services ModulesServices
}

type ModuleDeps struct{}
type ModulesServices struct {
	DB *DBService
}

type DBService struct {
	*pgxpool.Pool
}

const Tag = "POSTGRES"

func New(connStr string) *Module {
	return &Module{
		tag:    Tag,
		status: modules.StatusCodePreInit,

		connStr: connStr,

		deps: ModuleDeps{},
		services: ModulesServices{
			DB: &DBService{},
		},
	}
}

func (m *Module) Tag() string {
	return m.tag
}

func (m *Module) Status() string {
	return m.status
}

func (m *Module) Init(ctx context.Context) error {
	if m.status == modules.StatusCodeActive {
		return fmt.Errorf("(pgdb.Module.Init) %w", modules.ErrInvalidStatus)
	}

	pool, err := pgxpool.New(ctx, m.connStr)
	if err != nil {
		return fmt.Errorf("(dbpg.Module.Init) %w; %w", modules.ErrInitFailed, err)
	}

	m.services.DB.Pool = pool
	m.status = modules.StatusCodeActive

	return nil
}

func (m *Module) MustInit(ctx context.Context) {
	if err := m.Init(ctx); err != nil {
		panic(err)
	}
}

func (m *Module) Stop(ctx context.Context) error {
	if m.status != modules.StatusCodeActive {
		return fmt.Errorf("(dbpg.Module.Stop) %w", modules.ErrInvalidStatus)
	}

	m.services.DB.Close()
	m.status = modules.StatusCodeStopped

	return nil
}

func (m *Module) MustStop(ctx context.Context) {
	if err := m.Stop(ctx); err != nil {
		panic(err)
	}
}

func (m *Module) DBService() *DBService {
	if m.status != modules.StatusCodeActive {
		panic(fmt.Errorf("(dbpg.Module.DBService) %w", modules.ErrInvalidStatus))
	}

	return m.services.DB
}

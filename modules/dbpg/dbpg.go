package dbpg

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kijudev/blueprint/lib"
)

type Module struct {
	tag    string
	status string

	config   ModuleConfig
	services ModulesServices
}

type ModuleConfig struct {
	ConnStr string
}

type ModulesServices struct {
	db *pgxpool.Pool
}

const Tag = "POSTGRES"

func New(config ModuleConfig) *Module {
	return &Module{
		tag:      Tag,
		status:   lib.StatusCodeModuleNotInitialized,
		config:   config,
		services: ModulesServices{},
	}
}

func (m *Module) Tag() string {
	return m.tag
}

func (m *Module) Status() string {
	return m.status
}

func (m *Module) Init(ctx context.Context) error {
	if m.status == lib.StatusCodeModuleRunning {
		return lib.ErrModuleAlreadyRunning
	}

	pool, err := pgxpool.New(ctx, m.config.ConnStr)
	if err != nil {
		return lib.JoinErrors(lib.ErrModuleInitFailed, err)
	}

	m.services.db = pool
	m.status = lib.StatusCodeModuleRunning

	return nil
}

func (m *Module) MustInit(ctx context.Context) {
	if err := m.Init(ctx); err != nil {
		panic(err)
	}
}

func (m *Module) Stop(ctx context.Context) error {
	if m.status != lib.StatusCodeModuleRunning {
		return lib.JoinErrors(lib.ErrModuleStopFailed)
	}

	m.services.db.Close()
	m.status = lib.StatusCodeModuleStopped

	return nil
}

func (m *Module) MustStop(ctx context.Context) {
	if err := m.Stop(ctx); err != nil {
		panic(err)
	}
}

func (m *Module) DBService() *pgxpool.Pool {
	if m.status != lib.StatusCodeModuleRunning {
		panic(lib.JoinErrors(lib.ErrUnknown))
	}

	return m.services.db
}

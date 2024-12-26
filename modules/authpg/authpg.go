package authpg

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/kijudev/blueprint/lib"
)

type Module struct {
	tag    string
	status string

	deps     ModuleDeps
	services ModuleServices
}

type ModuleDeps struct {
	DB *pgxpool.Pool
}

type ModuleServices struct {
	data *DataService
}

const Tag = "AUTHPG"

func New(deps ModuleDeps) *Module {
	return &Module{
		tag:    Tag,
		status: lib.StatusCodeModuleNotInitialized,

		deps: deps,
		services: ModuleServices{
			data: &DataService{
				db: deps.DB,
			},
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
	e := errors.New("(authpg.Model.Init)")

	if m.status == lib.StatusCodeModuleRunning {
		return lib.JoinErrors(e, lib.ErrModuleAlreadyRunning)
	}

	if m.deps.DB == nil {
		return lib.JoinErrors(e, lib.ErrMissingDependency)
	}

	m.status = lib.StatusCodeModuleRunning

	return nil
}

func (m *Module) MustInit(ctx context.Context) {
	if err := m.Init(ctx); err != nil {
		panic(err)
	}
}

func (m *Module) Stop(ctx context.Context) error {
	e := errors.New("(authpg.Model.Stop)")

	if m.status == lib.StatusCodeModuleNotInitialized {
		return lib.JoinErrors(e, lib.ErrModuleNotInitialized)
	}

	if m.status == lib.StatusCodeModuleStopped {
		return lib.JoinErrors(e, lib.ErrModuleStopFailed)
	}

	m.status = lib.StatusCodeModuleStopped

	return nil
}

func (m *Module) MustStop(ctx context.Context) {
	if err := m.Stop(ctx); err != nil {
		panic(err)
	}
}

func (m *Module) DataService() *DataService {
	if m.services.data == nil {
		panic(errors.New("(authpg.Model.DataService)"))
	}

	return m.services.data
}

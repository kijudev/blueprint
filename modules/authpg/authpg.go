package authpg

import (
	"context"
	"fmt"

	"github.com/kijudev/blueprint/lib/modules"
	"github.com/kijudev/blueprint/modules/dbpg"
)

type Module struct {
	tag    string
	status string

	deps     ModuleDeps
	services ModuleSerivces
}

type ModuleDeps struct {
	DB *dbpg.DBService
}

type ModuleSerivces struct {
	Core *CoreService
}

const Tag = "AUTHPG"

func New(deps ModuleDeps) *Module {
	return &Module{
		tag:    Tag,
		status: modules.StatusCodePreInit,
		deps:   deps,

		services: ModuleSerivces{
			Core: NewCoreService(deps.DB),
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
		return fmt.Errorf("(authpg.Module.Init) %w", modules.ErrInvalidStatus)
	}

	if m.deps.DB == nil {
		return fmt.Errorf("(authpg.Module.Init) %w", modules.ErrMissingDependency)
	}

	m.services.Core.db = m.deps.DB

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
		return fmt.Errorf("(authpg.Module.Stop) %w", modules.ErrInvalidStatus)
	}

	m.status = modules.StatusCodeStopped

	return nil
}

func (m *Module) MustStop(ctx context.Context) {
	if err := m.Stop(ctx); err != nil {
		panic(err)
	}
}

func (m *Module) CoreService() *CoreService {
	if m.status != modules.StatusCodeActive {
		panic(fmt.Errorf("(authpg.Module.CoreService) %w", modules.ErrInvalidStatus))
	}

	return m.services.Core
}

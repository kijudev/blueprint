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

const TAG = "AUTHPG"

func NewModule(deps ModuleDeps) *Module {
	return &Module{
		tag:    TAG,
		status: modules.StatusCodePreInit,
		deps:   deps,

		services: ModuleSerivces{
			Core: new(CoreService),
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
		return fmt.Errorf("(authpg.Module.Init) %w", modules.ErrorInvalidStatus)
	}

	m.services.Core.db = m.deps.DB

	// Dummy implementation
	// for _, migration := range migrations {
	// 	if _, err := m.deps.DB.Exec(ctx, migration.Up); err != nil {
	// 		return fmt.Errorf("(authpg.Module.Init) %w; %w", modules.ErrorInitFailed, err)
	// 	}
	// }

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
		return fmt.Errorf("(authpg.Module.Stop) %w", modules.ErrorInvalidStatus)
	}

	m.status = modules.StatusCodeStopped

	return nil
}

func (m *Module) MustStop(ctx context.Context) {
	if err := m.Stop(ctx); err != nil {
		panic(err)
	}
}

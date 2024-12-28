package evbus

import (
	"context"

	"github.com/kijudev/blueprint/lib"
)

type Module struct {
	tag    string
	status string

	config   ModuleConfig
	services ModuleServices
}

type ModuleConfig struct {
	MaxGoroutines int
}

type ModuleServices struct {
	EventBusService *Service
}

const Tag = "EVBUS"

func New(config ModuleConfig) *Module {
	return &Module{
		tag:    Tag,
		status: lib.StatusCodeModuleNotInitialized,

		config: config,
		services: ModuleServices{
			EventBusService: NewEventBusService(config.MaxGoroutines),
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
	if m.status == lib.StatusCodeModuleRunning {
		return lib.ErrModuleAlreadyRunning
	}

	if m.config.MaxGoroutines == 0 {
		return lib.ErrModuleInitFailed
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
	if m.status == lib.StatusCodeModuleStopped || m.status == lib.StatusCodeModuleNotInitialized {
		return lib.ErrModuleStopFailed
	}

	m.status = lib.StatusCodeModuleStopped

	return nil
}

func (m *Module) MustStop(ctx context.Context) {
	if err := m.Stop(ctx); err != nil {
		panic(err)
	}
}

func (m *Module) Service() *Service {
	return m.services.EventBusService
}

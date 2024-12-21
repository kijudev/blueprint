package auth

import (
	"errors"

	"github.com/kijudev/blueprint/modules/dbpg"
)

type ServicePg struct {
	dbpgModue *dbpg.Module
}

type ModulePg struct {
	dbpgModule *dbpg.Module
	service    *ServicePg
	status     string
}

func NewModulePg(dbpgModule *dbpg.Module) *ModulePg {
	return &ModulePg{
		dbpgModule: dbpgModule,
		status:     StatusPreInit,
	}
}

func (m *ModulePg) Init() error {
	if m.dbpgModule.GetStatus() != dbpg.StatusRunning {
		return errors.New("The database module must be running")
	}

	m.status = StatusActive
	return nil
}

func (m *ModulePg) Stop() error {
	m.status = StatusStopped
	return nil
}

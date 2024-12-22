package auth

import (
	"context"

	"github.com/kijudev/blueprint/modules/dbpg"
)

type ServicePg struct {
	dbpgModule *dbpg.Module
}

type ModulePg struct {
	dbpgModule *dbpg.Module
	service    *ServicePg
	status     string
}

func NewServicePg(dbpgModule *dbpg.Module) *ServicePg {
	return &ServicePg{
		dbpgModule: dbpgModule,
	}
}

func NewModulePg(dbpgModule *dbpg.Module) *ModulePg {
	return &ModulePg{
		dbpgModule: dbpgModule,
		service:    NewServicePg(dbpgModule),
	}
}

func (m *ModulePg) Init(ctx context.Context) error {
	for _, migration := range userMigrations {
		_, err := m.dbpgModule.GetPool().Exec(ctx, migration.Up)

		if err != nil {
			return err
		}
	}

	return nil
}

func (m *ModulePg) Stop(ctx context.Context) error {
	return nil
}

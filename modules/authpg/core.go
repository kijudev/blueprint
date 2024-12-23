package authpg

import "github.com/kijudev/blueprint/modules/dbpg"

type CoreService struct {
	db *dbpg.DBService
}

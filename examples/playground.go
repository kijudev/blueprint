package main

import (
	"context"
	"fmt"

	"github.com/kijudev/blueprint/lib/models"
	"github.com/kijudev/blueprint/modules/authpg"
	"github.com/kijudev/blueprint/modules/dbpg"
)

func main() {
	ctx := context.Background()

	dbpgModule := dbpg.New("postgresql://blueprint:1234@localhost:5432/blueprint")
	dbpgModule.MustInit(ctx)
	defer dbpgModule.MustStop(ctx)

	authModule := authpg.New(authpg.ModuleDeps{
		DB: dbpgModule.DBService(),
	})
	authModule.MustInit(ctx)
	defer authModule.MustStop(ctx)

	user, err := authModule.CoreService().GetUserByID(ctx, models.MustNew("0193f7da-6cd3-fb6e-f33a-5f7b4f8f8103"))

	if err != nil {
		panic(err)
	} else {
		fmt.Println(user)
	}
}

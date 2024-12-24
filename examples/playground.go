package main

import (
	"context"
	"fmt"

	"github.com/kijudev/blueprint/modules/auth"
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

	user, err := authModule.CoreService().CreateUser(ctx, auth.UserParams{
		Email: "test@gmail.com",
		Name:  "test-user",
	})

	if err != nil {
		panic(err)
	} else {
		fmt.Println(user)
	}
}

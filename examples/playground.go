package main

import (
	"context"
	"fmt"

	"github.com/kijudev/blueprint/modules/auth"
	"github.com/kijudev/blueprint/modules/dbpg"
)

func main() {
	ctx := context.Background()

	dbpgModule := dbpg.NewModule("postgresql://blueprint:1234@localhost:5432/blueprint")
	err := dbpgModule.Init(ctx)

	if err != nil {
		panic(err)
	}

	authModule := auth.NewModulePg(dbpgModule)
	err = authModule.Init(ctx)
	if err != nil {
		panic(err)
	}

	user, err := authModule.CreateUser(ctx, auth.UserParams{
		Email: "test@gmail.com",
		Name:  "Test",
	})

	if err != nil {
		panic(err)
	} else {
		fmt.Println(user)
	}

	err = authModule.Stop(ctx)
	err = dbpgModule.Stop(ctx)
	if err != nil {
		panic(err)
	}
}

package main

import (
	"context"
	"fmt"

	"github.com/kijudev/blueprint/signals"
)

func main() {
	// ctx := context.Background()

	// dbpgModule := dbpg.New(dbpg.ModuleConfig{
	// 	ConnStr: "postgresql://blueprint:1234@localhost:5432/blueprint",
	// })
	// dbpgModule.MustInit(ctx)
	// defer dbpgModule.MustStop(ctx)

	// authModule := authpg.New(authpg.ModuleDeps{
	// 	DB: dbpgModule.DBService(),
	// })
	// authModule.MustInit(ctx)
	// defer authModule.MustStop(ctx)

	// //user, err := authModule.DataService().GetAccountByID(ctx, lib.MustNewID("0193f7da-6cd3-fb6e-f33a-5f7b4f8f8103"))
	// user, err := authModule.DataService().GetAccounts(ctx, lib.Pagination{Offset: 0, Limit: 10})

	// if err != nil {
	// 	panic(err)
	// } else {
	// 	fmt.Println(user)
	// }

	s := signals.New[int]()

	s.Listen(func(ctx context.Context, e signals.Event[int]) {
		fmt.Println("1 ", e.Data)
	})
	s.Listen(func(ctx context.Context, e signals.Event[int]) {
		fmt.Println("2 ", e.Data)
	})

	s.Dispatch(42)
	s.Dispatch(43)
	s.Dispatch(44)
}

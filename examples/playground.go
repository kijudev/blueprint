package main

import (
	"context"
	"fmt"

	"github.com/kijudev/blueprint/modules/evbus"
)

const (
	EvCodeOne = "one"
	EvCodeTwo = "two"
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

	ctx := context.Background()
	bus := evbus.New(evbus.ModuleConfig{
		MaxGoroutines: 10,
	})

	bus.Service().MustRegister(ctx, EvCodeOne, new(OneEvent))
	bus.Service().MustRegister(ctx, EvCodeTwo, new(TwoEvent))

	//fmt.Println(bus.Service().IsRegistered(ctx, EvCodeOne))
	//fmt.Println(bus.Service().IsRegistered(ctx, EvCodeTwo))

	bus.Service().MustSubscribe(ctx, EvCodeOne, func(ctx context.Context, ev OneEvent) {
		fmt.Println("S1 => ", ev.Value)
	})

	bus.Service().MustSubscribe(ctx, EvCodeTwo, func(ctx context.Context, ev TwoEvent) {
		fmt.Println("S1 => ", ev.Value)
	})

	bus.Service().MustDispatch(ctx, EvCodeOne, OneEvent{1})
	bus.Service().MustDispatch(ctx, EvCodeOne, OneEvent{2})
	bus.Service().MustDispatch(ctx, EvCodeOne, OneEvent{3})
	bus.Service().MustDispatch(ctx, EvCodeOne, OneEvent{4})

	bus.Service().Wait(ctx)

	bus.Service().MustDispatch(ctx, EvCodeTwo, TwoEvent{"blue"})
	bus.Service().MustDispatch(ctx, EvCodeTwo, TwoEvent{"green"})
	bus.Service().MustDispatch(ctx, EvCodeTwo, TwoEvent{"yellow"})

	bus.MustInit(ctx)
	defer bus.MustStop(ctx)

}

type OneEvent struct {
	Value int
}
type TwoEvent struct {
	Value string
}

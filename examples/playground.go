package main

import (
	"context"
	"fmt"

	"github.com/kijudev/blueprint/modules/evbus"
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

	bus.Service().MustRegister(ctx, new(OneEvent))
	bus.Service().MustRegister(ctx, new(TwoEvent))

	//fmt.Println(bus.Service().IsRegistered(ctx, EvCodeOne))
	//fmt.Println(bus.Service().IsRegistered(ctx, EvCodeTwo))

	bus.Service().MustSubscribe(ctx, func(ctx context.Context, ev OneEvent) {
		fmt.Println("S1 (E1) => ", ev.Value)
	})

	bus.Service().MustSubscribe(ctx, func(ctx context.Context, ev TwoEvent) {
		fmt.Println("S1 (E2) => ", ev.Value)
	})

	bus.Service().MustDispatchSync(ctx, OneEvent{1})
	bus.Service().MustDispatchSync(ctx, OneEvent{2})
	bus.Service().MustDispatch(ctx, OneEvent{3})
	bus.Service().MustDispatch(ctx, OneEvent{4})

	bus.Service().MustDispatch(ctx, TwoEvent{"A"})
	bus.Service().MustDispatch(ctx, TwoEvent{"B"})
	bus.Service().MustDispatch(ctx, TwoEvent{"C"})

	bus.Service().Wait(ctx)

	bus.MustInit(ctx)
	defer bus.MustStop(ctx)

}

type OneEvent struct {
	Value int
}
type TwoEvent struct {
	Value string
}

type ThreeEvent struct {
	Value float64
}

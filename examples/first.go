package main

import (
	"fmt"

	"github.com/kijudev/blueprint/modules/core"
)

func main() {
	user := core.User{
		Email: "kiju@kiju.page",
		Name:  "kijukiju",
	}

	err := user.Validate()

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(user)
	}
}

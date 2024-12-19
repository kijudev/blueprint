package main

import (
	"fmt"

	"github.com/kijudev/blueprint/modules/core"
)

func main() {
	user := core.User{
		Email: "tesgmail.com",
		Name:  "Jo",
	}

	err := user.Validate()

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(user)
	}
}

package main

import (
	"fmt"

	"github.com/kijudev/blueprint/lib/validation"
	"github.com/kijudev/blueprint/modules/core"
)

func main() {
	user := core.User{
		Email: "test@gmail.com",
		Name:  "Kiju",
	}

	fmt.Println(user, validation.Combine(validation.ValidateField("email", validation.Email("kiju@gmail.com"))))
}

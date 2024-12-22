package main

import (
	"errors"
	"fmt"
)

func main() {
	err := errors.Join(errors.New("err 1"), errors.New("err 2"), errors.New("err 3"))
	fmt.Println(err)
}

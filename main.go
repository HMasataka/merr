package main

import (
	"errors"
	"fmt"
)

func main() {
	a := fmt.Sprint("This is A error \n")
	b := fmt.Sprint("This is B error \n")

	res := fmt.Sprintf("%s%s", a, b)
	panic(errors.New(res))
}

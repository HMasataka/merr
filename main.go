package main

import (
	"fmt"
)

const (
	a = "This is A error \n"
	b = "This is B error \n"
)

func print(value string) {
	fmt.Print(value)
}

func main() {
	print(a + b)
}

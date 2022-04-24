package main

import (
	"fmt"
	input "go_sample/pkg"
)

func main() {
	name := input.Input("type your name")
	fmt.Println("Hello, " + name + "!!")
}

package main

import (
	"fmt"
	input "go_sample/pkg"
)

func main() {
	name := input.Input("type your name")
	fmt.Println("Hello, " + name + "!!")
	a,b,c := 100, 200, 300
	fmt.Print("total: ")
	fmt.Println(a+b+c)
}

package main

import (
	"fmt"
	input "go_sample/pkg"
	"strconv"
)

func main() {
	// name := input.Input("type your name")
	// fmt.Println("Hello, " + name + "!!")
	// a,b,c := 100, 200, 300
	// fmt.Print("total: ")
	// fmt.Println(a+b+c)
	x := input.Input("type a price")
	n, err := strconv.Atoi(x)
	if err != nil {
		fmt.Print("Error: ")
		fmt.Println(err.Error())
		return
	}
	p := float64(n)
	fmt.Println(int(p*1.1))
}

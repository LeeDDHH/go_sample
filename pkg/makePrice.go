package pkg

import (
	"fmt"
	input "go_sample/pkg/input"
	"strconv"
)

func MakePrice() {
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

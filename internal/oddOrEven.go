package internal

import (
	"fmt"
	input "go_sample/internal/input"
	"strconv"
)

func OddOrEven() {
	x := input.Input("type a price")
	fmt.Print(x + "は、")
	if n, err := strconv.Atoi(x); err == nil {
		if n%2 == 0 {
			fmt.Println("偶数")
		} else {
			fmt.Println("奇数")
		}
	} else {
		fmt.Println("数値ではない")
	}
}

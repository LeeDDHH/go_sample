package pkg

import (
	"fmt"
	"go_sample/pkg/input"
	"strconv"
)

func OddOrEven2() {
	x := input.Input("type a number")
	n, err := strconv.Atoi(x)

	if err == nil {
		fmt.Println(x + "は、")
	} else {
		return
	}
	switch {
		case n%2 == 0:
			fmt.Println("偶数")

		case n%2 == 1:
			fmt.Println("奇数")
	}
}

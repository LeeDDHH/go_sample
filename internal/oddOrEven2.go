package internal

import (
	"fmt"
	"strconv"
)

func OddOrEven2() {
	x := Input("type a number")
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

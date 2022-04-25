package pkg

import (
	"fmt"
	input "go_sample/pkg/input"
	"strconv"
)

func SwitchCase() {
	x := input.Input("type a number")
	fmt.Print(x + "月は、")
	switch n, err := strconv.Atoi(x); n {
		case 0:
				fmt.Println("整数値が得られません")
				fmt.Println(err)

		case 1,2,12:
			fmt.Println("冬")

		case 3,4,5:
			fmt.Println("春")

		case 6,7,8:
			fmt.Println("夏")

		case 9,10,11:
			fmt.Println("冬")
		
		default: 
			fmt.Println("月の値ではありません")
	}
}

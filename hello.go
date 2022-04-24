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

	// x := input.Input("type a price")
	// n, err := strconv.Atoi(x)
	// if err != nil {
	// 	fmt.Print("Error: ")
	// 	fmt.Println(err.Error())
	// 	return
	// }
	// p := float64(n)
	// fmt.Println(int(p*1.1))

	// x := input.Input("type a price")
	// fmt.Print(x + "は、")
	// if n, err := strconv.Atoi(x); err == nil {
	// 	if n%2 == 0 {
	// 		fmt.Println("偶数")
	// 	} else {
	// 		fmt.Println("奇数")
	// 	}
	// } else {
	// 	fmt.Println("数値ではない")
	// }

	// x := input.Input("type a number")
	// fmt.Print(x + "月は、")
	// switch n, err := strconv.Atoi(x); n {
	// 	case 0:
	// 			fmt.Println("整数値が得られません")
	// 			fmt.Println(err)

	// 	case 1,2,12:
	// 		fmt.Println("冬")

	// 	case 3,4,5:
	// 		fmt.Println("春")

	// 	case 6,7,8:
	// 		fmt.Println("夏")

	// 	case 9,10,11:
	// 		fmt.Println("冬")
		
	// 	default: 
	// 		fmt.Println("月の値ではありません")
	// }

	// x := input.Input("type a number")
	// n, err := strconv.Atoi(x)

	// if err == nil {
	// 	fmt.Println(x + "は、")
	// } else {
	// 	return
	// }
	// switch {
	// 	case n%2 == 0:
	// 		fmt.Println("偶数")

	// 	case n%2 == 1:
	// 		fmt.Println("奇数")
	// }

	// x := input.Input("type 1~5")
	// n, err := strconv.Atoi(x)

	// if err == nil {
	// 	fmt.Println(x + "までの合計は、")
	// } else {
	// 	return
	// }

	// t := 0;

	// switch n {
	// 	case 5:
	// 		t += 5
	// 		fallthrough

	// 	case 4:
	// 		t += 4
	// 		fallthrough

	// 	case 3:
	// 		t += 3
	// 		fallthrough

	// 	case 2:
	// 		t += 2
	// 		fallthrough
		
	// 	case 1:
	// 		t++
		
	// 	default:
	// 		fmt.Println("範囲外です。")
	// 		return
	// }

	// fmt.Println(t, "です")

	t := 0
	c := 1
	x := input.Input("type a number")
	n, err := strconv.Atoi(x)

	if err == nil {
		fmt.Println("1から" + x + "の合計は、")
	} else {
			goto err
	}

	for c <= n {
		t += c
		c++
	}

	fmt.Println(t, "です")

err:
	fmt.Println("Error")
}

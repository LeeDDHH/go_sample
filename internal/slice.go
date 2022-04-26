package internal

import "fmt"

func Slice() {
	a := [5]int{10, 20, 30, 40, 50}
	b := a[0:3]
	fmt.Println(a)
	fmt.Println(b)
}

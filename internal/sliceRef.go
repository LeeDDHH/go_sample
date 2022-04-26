package internal

import "fmt"

func SliceRef() {
	a := [5]int{10, 20, 30, 40, 50}
	b := a[0:3]
	fmt.Println(a)
	fmt.Println(b)
	a[0] = 100
	fmt.Println(a)
	fmt.Println(b)
	b[1] = 200
	fmt.Println(a)
	fmt.Println(b)
}

package internal

import "fmt"

func SliceControl() {
	a := []int{ 10, 20, 30 }
	fmt.Println(a)

	a = push(a,100)
	fmt.Println(a)

	a = pop(a)
	fmt.Println(a)

	a = unshift(a,1000)
	fmt.Println(a)

	a = shift(a)
	fmt.Println(a)

	a = insert(a,1000,2)
	fmt.Println(a)

	a = remove(a,2)
	fmt.Println(a)
}

func push(a []int, v int) []int {
	return append(a,v)
}

func pop(a []int) []int {
	return a[:len(a)-1]
}

func unshift(a []int, v int) []int {
	return append([]int{v}, a...)
}

func shift(a []int) []int {
	return a[1:]
}

func insert(a []int, v int, p int) []int {
	// a = append(a,0) 
	fmt.Println("-------")
	fmt.Println(a)
	fmt.Println("-------")
	fmt.Println(a[:p+1])
	fmt.Println(a[p:])
	fmt.Println("-------")
	a = append(a[:p+1], a[p:]...)
	fmt.Println(a)
	fmt.Println("-------")
	a[p] = v
	return a
}

func remove(a []int, p int) []int {
	return append(a[:p], a[p+1:]...)
}

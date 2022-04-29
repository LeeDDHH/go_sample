package internal

import "fmt"

func PointerSample() {
	n := 123
	p := &n
	fmt.Println("number: ", n)
	fmt.Println("pointer: ", p)
	fmt.Println("value: ", *p)
}

func PinterSubstitution() {
	n := 123
	p := &n
	m := 10000
	p2 := &m
	fmt.Printf("p value:%d, address:%p\n", *p, p)
	fmt.Printf("p2 value:%d, address:%p\n", *p2, p2)

	pb := p
	p = p2
	p2 = pb
	fmt.Printf("p value:%d, address:%p\n", *p, p)
	fmt.Printf("p2 value:%d, address:%p\n", *p2, p2)
}

func PointersPointer() {
	n := 123
	p := &n
	q := &p

	m := 10000
	p2 := &m
	q2 := &p2

	fmt.Printf("q value:%d, address:%p\n", **q, *q)
	fmt.Printf("q2 value:%d, address:%p\n", **q2, *q2)

	pb := p
	p = p2
	p2 = pb

	fmt.Printf("q value:%d, address:%p\n", **q, *q)
	fmt.Printf("q2 value:%d, address:%p\n", **q2, *q2)
}

func PointerUseMerit() {
	n := 123
	fmt.Printf("value:%d.\n", n)
	pointerUseMeritChange1(n)
	fmt.Printf("value:%d.\n", n)
	pointerUseMeritChange2(&n)
	fmt.Printf("value:%d.\n", n)
}

func pointerUseMeritChange1(n int) {
	n *= 2
}

func pointerUseMeritChange2(n *int) {
	*n *= 2
}

func PointerSlice() {
	ar := []int{10, 20, 30}
	fmt.Println(ar)
	pointerSliceInitial(&ar)
	fmt.Println(ar)
}

func pointerSliceInitial(ar *[]int) {
	for i :=0 ; i < len(*ar); i++ {
		(*ar)[i] = 0
	}
}

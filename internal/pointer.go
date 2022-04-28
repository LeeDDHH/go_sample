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
package internal

import "fmt"

func ReturnValueWithName() {
	m := []string{
		"one", "two", "three",
	}
	fmt.Println(m)

	m = _returnValueWithNameInsert(m, "*", 2)
	fmt.Println(m)

	m = _returnValueWithNameInsert(m, "*", 1)
	fmt.Println(m)
}

func _returnValueWithNameInsert(a []string, v string, p int) (s []string) {
	s = append(a, " ")
	s = append(s[:p+1], s[p:len(s)-1]...)
	s[p] = v
	return
}

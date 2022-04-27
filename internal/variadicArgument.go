package internal

import "fmt"

func VariadicArgument() {
	m := []string{
		"one", "two", "three",
	}
	fmt.Println(m)

	m = _variadicArgumentPush(m, "1", "2", "3")
	fmt.Println(m)
}

func _variadicArgumentPush(a []string, v ...string) (s []string) {
	s = append(a, v...)
	return
}

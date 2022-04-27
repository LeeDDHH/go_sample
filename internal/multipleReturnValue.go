package internal

import "fmt"

func MultipleReturnValue() {
	m := []string{}
	m, _ = _arrayPush(m, "apple")
	m, _ = _arrayPush(m, "banana")
	m, _ = _arrayPush(m, "orange")
	fmt.Println(m)

	m, v := _arrayPop(m)
	fmt.Println("get " + v + "->", m)

}

func _arrayPush(a []string, v string) ([]string, int) {
	return append(a,v), len(a)
}

func _arrayPop(a []string) ([]string, string) {
	return a[:len(a)-1], a[len(a)-1]
}

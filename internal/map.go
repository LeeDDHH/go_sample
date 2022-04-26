package internal

import "fmt"

func Map() {
	m := map[string]int {
		"a": 100,
		"b": 200,
		"c": 300,
	}
	m["total"] = m["a"] + m["b"] + m["c"]

	fmt.Println(m)
}

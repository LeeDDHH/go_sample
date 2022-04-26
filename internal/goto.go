package internal

import (
	"fmt"
	"strconv"
)

func GotoSample() {
	t := 0
	c := 1
	x := Input("type a number")
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

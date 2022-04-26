package internal

import (
	"fmt"
	"strconv"
	"strings"
)

func Array2() {
	x := Input("input data")
	ar := strings.Split(x, " ")
	t := 0
	for _, v := range ar {
		n, er := strconv.Atoi(v)
		if er != nil {
			goto err
		}
		t += n
	}
	fmt.Println("total:", t)
	return

err:
	fmt.Println("ERROR!")
}

package internal

import (
	"fmt"
	"strconv"
	"time"
)

func GoRoutineSample(s string, n int){
	for i := 1; i <= 10; i++ {
		fmt.Printf("<%d %s>", i, s)
		time.Sleep(time.Duration(n) * time.Microsecond)
	}
}

func GoRoutineUseSharedMemory() {
	msg := "start"
	prmsg := func(nm string, n int) {
		fmt.Println(nm, msg)
		time.Sleep(time.Duration(n) * time.Millisecond)
	}
	useGoRutineFunc := func() {
		const nm string = "useGoRutineFunc"
		const execTime int = 60
		for i := 0; i < 10; i++ {
			msg += " uG" + strconv.Itoa(i)
			prmsg(nm, execTime)
		}
	}
	defaultRunFunc := func() {
		const nm string = "*defaultRunFunc"
		const execTime int = 100
		for i := 0; i < 5; i++ {
			msg += " dR" + strconv.Itoa(i)
			prmsg(nm, execTime)
		}
	}

	go useGoRutineFunc()
	defaultRunFunc()
}

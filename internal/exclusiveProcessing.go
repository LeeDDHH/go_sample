package internal

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

type SrData struct {
	msg string
	mux sync.Mutex
}

func ExclusiveProcessing() {
	sd := SrData{msg: "Start"}

	prmsg := func(nm string, n int) {
		fmt.Println(nm, sd.msg)
		time.Sleep(time.Duration(n) * time.Millisecond)
	}

	useGoRutineFunc1 := func() {
		const nm string = "useGoRutineFunc1"
		const execTime int = 100
		sd.mux.Lock()
		for i := 0; i < 10; i++ {
			sd.msg += " uG1" + strconv.Itoa(i)
			prmsg(nm, execTime)
		}
		sd.mux.Unlock()
	}

	useGoRutineFunc2 := func() {
		const nm string = "useGoRutineFunc2"
		const execTime int = 50
		sd.mux.Lock()
		for i := 0; i < 5; i++ {
			sd.msg += " uG2" + strconv.Itoa(i)
			prmsg(nm, execTime)
		}
		sd.mux.Unlock()
	}

	go useGoRutineFunc1()
	go useGoRutineFunc2()
	time.Sleep(5 * time.Second)
}

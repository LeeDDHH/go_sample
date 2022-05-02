package internal

import (
	"fmt"
	"strconv"
	"time"
)

func channelTotal(n int, c chan int) {
	t := 0
	for i := 1; i <= n; i++ {
		t += i
	}
	c <- t
}

func channelTotalAfterCome(c chan int) {
	n := <-c
	fmt.Println("n = ", n)
	t := 0
	for i := 1; i <= n; i++ {
		t += i
	}
	fmt.Println("total: ", t)
}

func channelTotalUseTwoChannel(cs chan int, cr chan int) {
	n := <-cs
	fmt.Println("n = ", n)
	t := 0
	for i := 1; i <= n; i++ {
		t += i
	}
	cr <- t
}

func channelPrmsg(n int, s string) {
	fmt.Println(s)
	time.Sleep(time.Duration(n * 50000) * time.Microsecond)
}

func channelFirst(n int, c chan string) {
	const nm string = "first-"
	for i :=0; i < 10; i++ {
		s := nm + strconv.Itoa(i)
		channelPrmsg(n, s)
		c <- s
	}
}

func channelSecond(n int, c chan string) {
	fmt.Println("channelSecond")
	for i := 0; i < 10; i++ {
		fmt.Println("チャンネルに値が入ってくるまで待つ")
		fmt.Println(i)
		channelPrmsg(n, "second:["+<-c+"]")
		fmt.Println("チャンネルに値が入ったあと")
	}
}

func channelCount(n int, s int, c chan int) {
	for i := 1; i <= n; i++ {
		c <- i
		time.Sleep(time.Duration(s) * time.Millisecond)
	}
}

func ChannelExample() {
	c := make(chan int)
	go channelTotal(100, c)
	fmt.Println("total: ", <-c)
}

func ChannelExample2() {
	c := make(chan int)
	go channelTotal(1000, c)
	go channelTotal(100, c)
	go channelTotal(10, c)
	x, y, z := <-c, <-c, <-c
	fmt.Println(x,y,z)
}

func ChannelExample3() {
	c := make(chan int)
	go channelTotalAfterCome(c)
	c <- 100
	time.Sleep(100 * time.Millisecond)
}

func ChannelExample4() {
	c := make(chan string)
	go channelFirst(10, c)
	channelSecond(10, c)
	fmt.Println()
}

func ChannelExample5() {
	cs := make(chan int)
	cr := make(chan int)
	go channelTotalUseTwoChannel(cs, cr)
	cs <- 100
	fmt.Println("total: ", <-cr)
}

func ChannelExample6() {
	n1, n2, n3 := 3, 5, 10
	m1, m2, m3 := 100, 75, 50

	c1 := make(chan int)
	go channelCount(n1, m1, c1)

	c2 := make(chan int)
	go channelCount(n2, m2, c2)

	c3 := make(chan int)
	go channelCount(n3, m3, c3)

	for i := 0; i < n1+n2+n3; i++ {
		select {
		case re := <-c1:
			fmt.Println("* First ", re)
		case re := <-c2:
			fmt.Println("** Second ", re)
		case re := <-c3:
			fmt.Println("*** Third ", re)
		}
	}
	fmt.Println("**** finish. ****")
}

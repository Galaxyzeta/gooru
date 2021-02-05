package tutorial

import (
	"fmt"
	"strconv"
	"time"
)

// Select test
func Select() {
	chan1 := make(chan int, 1)
	chan2 := make(chan int, 1)
	chan1 <- 4
	chan2 <- 5
	select {
	case msg1 := <-chan1:
		fmt.Println("selected" + strconv.Itoa(msg1))
	case msg2 := <-chan2:
		fmt.Println("selected" + strconv.Itoa(msg2))
	default:
		fmt.Println("Blocked")
	}
}

// GoRoutine test
func GoRoutine() {
	gogo := func() {
		fmt.Println(1)
	}
	go gogo()
	go gogo()
	go gogo()
	time.Sleep(1000)
}

// Channel test
func Channel() {

	// Pipe is duplex
	pipe := make(chan int, 5)
	type Sender = chan<- int
	type Receiver = <-chan int
	// Send only and receive only
	var sender Sender = pipe
	var receiver Receiver = pipe

	sender <- 5
	fmt.Println(<-receiver)
}

// LockChannel test
func LockChannel() {
	fx := func(x *int, ch chan bool) {
		ch <- true
		*x = *x + 1
		<-ch
	}
	counter := 0
	var ch chan bool = make(chan bool, 1)
	for i := 0; i < 100; i++ {
		fx(&counter, ch)
	}
	close(ch)
	fmt.Println(counter)
}

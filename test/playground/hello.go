package main

import (
	"fmt"
	"runtime"
	"time"
	"unsafe"

	"galaxyzeta.com/concurrency/goroutine"
	"galaxyzeta.com/concurrency/waiter"
)

func main() {
	// selectTicker()
	// cyclePrint2()
	// semaphore.NewWeighted(1)
	arch()
}

func ticker() {
	var c <-chan time.Time = time.Tick(time.Duration(1) * time.Second)
	var end <-chan time.Time = time.Tick(time.Duration(5) * time.Second)
	go func() {
		<-end
		panic("close")
	}()
	for ; ; <-c {
		fmt.Println("Hello")
	}
}

func cyclePrint() {
	a := make(chan bool, 1)
	b := make(chan bool, 1)
	c := make(chan bool, 1)
	a <- true
	go func() {
		for i := 0; i < 10; i++ {
			<-a
			fmt.Println("a")
			b <- true
		}
		fmt.Println("OK")
	}()

	go func() {
		for i := 0; i < 10; i++ {
			<-b
			fmt.Println("b")
			c <- true
		}
		fmt.Println("OK")
	}()

	go func() {
		for i := 0; i < 10; i++ {
			<-c
			fmt.Println("c")
			a <- true
		}
		fmt.Println("OK")
	}()
}

func cyclePrint2() {
	a := make(chan bool, 1)
	b := make(chan bool, 1)
	c := make(chan bool, 1)
	a <- true
	wg := waiter.New()
	k := func() {
		for i := 0; i < 30; i++ {
			select {
			case <-a:
				fmt.Println("a")
				b <- true
			case <-b:
				fmt.Println("b")
				c <- true
			case <-c:
				fmt.Println("c")
				a <- true
			}
		}
		wg.Done()
	}
	wg.AddAndRun(k)
	wg.Wait()
}

func selectTicker() {
	wg := waiter.New()
	t := time.Tick(time.Duration(1) * time.Second)
	sigkill := make(chan bool)
	wg.AddAndRun(func() {
		for {
			select {
			case <-t:
				fmt.Println("golang nb")
			case <-sigkill:
				fmt.Println("KILL")
				wg.Done()
			}
		}
	})
	wg.AddAndRun(func() {
		goroutine.Sleep(5, time.Second)
		sigkill <- true
		wg.Done()
	})
	wg.Wait()
	fmt.Println("FIN")
}

func arch() {
	i := int(1)
	fmt.Println(unsafe.Sizeof(i))
	fmt.Println(runtime.GOARCH)
}

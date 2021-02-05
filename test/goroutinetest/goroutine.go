package main

import (
	"fmt"
	"time"

	"galaxyzeta.com/concurrency/goroutine"
)

func main() {
	fmt.Println(goroutine.GoID())
	go func() {
		fmt.Println(goroutine.GoID())
	}()
	time.Sleep(time.Second * time.Duration(1))
}

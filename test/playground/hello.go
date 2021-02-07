package main

import (
	"fmt"
	"reflect"
	"runtime"
	"time"
	"unsafe"

	"galaxyzeta.com/concurrency/goroutine"
	"galaxyzeta.com/concurrency/waiter"
	"galaxyzeta.com/logger"
	"galaxyzeta.com/util/assert"
)

func main() {
	// selectTicker()
	// cyclePrint2()
	// semaphore.NewWeighted(1)
	// arch()
	// nilmaptest()
	// reflection()
	// reflection2()
	panicAssertion()
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

func nilmaptest() {
	mp := make(map[string]interface{})
	fmt.Println(mp["asd"] == nil) // true
}

func reflection() {
	type hello struct {
		Strval string
	}
	var i interface{} = "qwerty"
	var hs *hello = &hello{}
	log := logger.New("Reflection")
	log.Infof("%s", reflect.ValueOf(hs).Type().Kind())
	reflect.ValueOf(hs).Elem().FieldByName("Strval").Set(reflect.ValueOf("qwerty"))
	log.Infof(hs.Strval)
	log.Infof("%p", hs)
	log.Infof("%t", reflect.ValueOf(i).CanInterface())
	log.Infof("%s", reflect.ValueOf(i).Type().Kind())
}

func reflection2() {
	var i int = 5
	fmt.Println(reflect.ValueOf(i).Int())
}

func panicAssertion() {
	defer func() {
		if recover() != nil {
			fmt.Println("Panic assertion OK")
		} else {
			fmt.Println("Panic assertion failed")
		}
	}()
	assert.ThrowsPanic(func() {
		panic("something...")
	})
	assert.ThrowsPanic(func() {
		// nothing happens
	}) // should throw panic.
}

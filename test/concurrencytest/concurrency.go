package main

import (
	"fmt"
	"math"
	"runtime"
	"sync"
	"time"

	deprecated "galaxyzeta.com/concurrency/deprecated"
	"galaxyzeta.com/concurrency/semaphore"
	"galaxyzeta.com/concurrency/spinlock"
	"galaxyzeta.com/concurrency/synchronizer"
	"galaxyzeta.com/concurrency/waiter"
)

func main() {
	// testDelayJob()
	// testSpinLock()
	// testInterfaceEq()
	// testReentrantLock()
	// testChannel()
	// testMutex()
	testMutexSem()
	// testChanSem()
	// testInterval()
}

func testDelayJob() {
	wheel := deprecated.NewTimer(16, 50)
	task1 := func(...interface{}) interface{} {
		fmt.Println("Hello World")
		return 1
	}
	wheel.Start()
	future1 := wheel.ExecuteAfter(task1, 2, time.Second)
	future2 := wheel.ExecuteAfter(task1, 3, time.Second)
	fmt.Println(future1.Get())
	fmt.Println(future2.Get())
}

func testSpinLock() {
	opt := spinlock.New(10)
	var condVal interface{} = false
	wheel := deprecated.NewTimer(16, 5)
	task1 := func(...interface{}) interface{} {
		condVal = true
		return nil
	}
	wheel.Start()
	wheel.ExecuteAfter(task1, time.Duration(2), time.Second)
	opt.SpinEq(condVal, false)
	fmt.Println("OK")
}

func testInterfaceEq() {
	var a interface{} = true
	var b interface{} = true
	fmt.Println(true == true)
	fmt.Println(a == b)
}

func testReentrantLock() {
	lock := synchronizer.NewReentrantLock()

	// lock.Unlock() // will panic.
	crit := 0
	wg := &sync.WaitGroup{}

	before := time.Now().UnixNano()
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			lock.Lock()
			crit++
			lock.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(time.Now().UnixNano() - before)
	fmt.Println(crit)
}

func testChannel() {
	wait := make(chan bool, 1)
	crit := 0
	wg := &sync.WaitGroup{}

	before := time.Now().UnixNano()
	wait <- true
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			<-wait
			crit++
			wait <- true
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(time.Now().UnixNano() - before)
	fmt.Println(crit)
}

func testMutex() {
	wait := sync.Mutex{}
	crit := 0
	wg := &sync.WaitGroup{}

	before := time.Now().UnixNano()
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			wait.Lock()
			crit++
			wait.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(time.Now().UnixNano() - before)
	fmt.Println(crit)
}

// MutexSem bench mark
func testMutexSem() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	sem := semaphore.NewMutexSem(5)
	wg := waiter.New()
	k1 := func() {
		for i := 0; i < 10000000; i++ {
			sem.Acquire()
		}
		wg.Done()
	}
	k2 := func() {
		for i := 0; i < 10000000; i++ {
			sem.Release()
		}
		wg.Done()
	}
	before := time.Now().UnixNano()
	wg.AddAndRun(k1)
	wg.AddAndRun(k2)
	wg.Wait()
	fmt.Println(time.Now().UnixNano() - before)
	fmt.Println(sem)
}

// MutexSem bench mark
func testChanSem() {
	sem := semaphore.NewChanSem(5, math.MaxInt16)
	before := time.Now().UnixNano()
	go func() {
		for i := 0; i < 1000; i++ {
			sem.Acquire()
		}
	}()
	go func() {
		for i := 0; i < 1000; i++ {
			sem.Release()
		}
	}()
	fmt.Println(time.Now().UnixNano() - before)
	time.Sleep(time.Second * time.Duration(1))
	fmt.Println(sem)
}

func testInterval() {
	i := deprecated.NewInterval(1, time.Second, func() { fmt.Println("Hello") })
	i.Start(false)
	time.Sleep(time.Duration(5) * time.Second)
	i.Cancel()
}

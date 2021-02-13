package synchronizer

import (
	"github.com/galaxyzeta/concurrency/goroutine"
	"github.com/galaxyzeta/concurrency/spinlock"
)

const vacant = -1

/*
=== About AQS ===

AQS (AbstractQueuedSynchronizer) is the kernel behind all provided synchronized containers.
By giving a chance to the goroutine which failed the first attempt to acquire the lock
to keep spinning itself trying to get the lock, we reduced context switching frequency in
heavy racing conditions.

=== Variables ===

state 			: a variable that has different meanings in different AQS implementations.
goID			: indicates the obtainer of current AQS
blockingQueue 	: makes the second race competitor block, while the first race competitor keeps trying to obtain the lock.
mu				: a spinlock. might be useful.

=== Example of racing condition handling with AQS explained ===

Suppose we are trying to accumulate a critical variable inited with zero 1000 times with
different Goroutines. The result should be exactly 1000, else it's not concurrenctly-safe.

The program is provided:

func testReentrantLock() {
	lock := synchronizer.NewReentrantLock()
	crit := 0
	for i := 0; i < 1000; i++ {
		go func() {
			lock.Lock()
			crit++
			lock.Unlock()
		}()
	}
	time.Sleep(time.Second * time.Duration(1))	// Wait all goroutines to finish
	fmt.Println(crit)	// should be 1000
}

By introducing ReentrantLock, things work as below:

1. The first goroutine calls lock.Acquire() to acquire the lock.
At that time, the AQS object is vacant, because its goID = -1.
So, the first goroutine obtains the lock, by CASing AQS's goID to its own goID.

2. Suppose the first goroutine does not abandon the lock immeditately, and the second
goroutine calls lock.Acquire() to get the lock. It will fail because AQS.goID != -1.
In such situation, we pack this goroutine into a GoNode object, and enqueue it into
AQS.blockingQueue, which is a channel sized 1. After that, it keeps trying to obtain
the lock, instead of hanging itself up.

3. Suppose the lock has not been released yet, and another goroutine calls lock.Acquire().
We attempt to enqueue it into the channel, however the channel is currently full, so this
attempt cause the goroutine to block.

4. Now imaging the first goroutine finally called lock.Release(). It first checks whether
AQS.goID represents itself. If not, a panic is thrown because Release is called before
Acquire. Else, it releases the lock by setting AQS.goID to -1.

5. Now that AQS.goID is -1, the second goroutine (which keeps trying to obtain the lock)
successfully obtained the lock, and removed itself from the channel, enabling the third
goroutine to enqueue itself into the channel.

*/

// AbstractQueuedSynchronizer is the kernel of synchronized containers.
type AbstractQueuedSynchronizer struct {
	state         int32
	goID          int32
	blockingQueue chan *GoNode
	mu            *spinlock.SpinLock
}

// GoNode is a representation of goroutine
type GoNode struct {
	goID int
}

// IAQS is an interface for AQS
type IAQS interface {
	tryAcquire(goID int) bool
	tryRelease(goID int) bool
	enqueue(goID int)
	rmqueue(goID int)
}

// NewAQS provides an instance of AQS
func NewAQS() *AbstractQueuedSynchronizer {
	return &AbstractQueuedSynchronizer{blockingQueue: make(chan *GoNode, 1), mu: spinlock.New(5), goID: vacant}
}

// Acquire the lock.
func Acquire(synchronizer IAQS) {
	// Assert IAQS is AbstractQueuedSynchronizer
	goID := goroutine.GoID()
	if synchronizer.tryAcquire(goID) == false {
		synchronizer.enqueue(goID)
		for synchronizer.tryAcquire(goID) == false {
			// Do nothing ...
		}
		synchronizer.rmqueue(goID)
	}
}

// Release the lock.
func Release(synchronizer IAQS) {
	goID := goroutine.GoID()
	synchronizer.tryRelease(goID)
}

//===========[PRIVATE]=============

func (aqs *AbstractQueuedSynchronizer) enqueue(goID int) {
	instance := &GoNode{goID: goID}
	aqs.blockingQueue <- instance
}

func (aqs *AbstractQueuedSynchronizer) rmqueue(goID int) {

	node := <-aqs.blockingQueue
	if node.goID != goID {
		panic("unknown error")
	}
}

func (aqs *AbstractQueuedSynchronizer) tryAcquire(goID int) bool {
	panic("implement me")
}

func (aqs *AbstractQueuedSynchronizer) tryRelease(goID int) bool {
	panic("implement me")
}

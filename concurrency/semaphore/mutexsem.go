package semaphore

import (
	"math"
	"sync"
)

// MutexSem is an implementation of semaphore using mutex.
//
// MutexSem is much efficient than ChanSem
type MutexSem struct {
	resource int
	mu       *sync.Mutex
	wait     chan struct{}
}

// NewMutexSem provides a new mutexSem
func NewMutexSem(init int) *MutexSem {
	return &MutexSem{mu: &sync.Mutex{}, wait: make(chan struct{}, 1), resource: init}
}

// Acquire the sem.
func (sem *MutexSem) Acquire() {
	if sem.resource >= math.MaxInt32 {
		panic("max sem limit ")
	}
	if sem.resource >= 0 {
		sem.mu.Lock()
		if sem.resource >= 0 {
			sem.resource--
		}
		sem.mu.Unlock()
		if sem.resource < 0 {
			<-sem.wait
		}
	} else {
		<-sem.wait
	}
}

// Release the sem
func (sem *MutexSem) Release() {
	sem.mu.Lock()
	sem.resource++
	sem.mu.Unlock()
	if sem.resource == 0 {
		sem.mu.Lock()
		if sem.resource == 0 {
			sem.wait <- struct{}{}
		}
		sem.mu.Unlock()
	}
}

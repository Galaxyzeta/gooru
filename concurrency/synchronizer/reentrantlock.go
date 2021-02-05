package synchronizer

import (
	"sync/atomic"
)

// ReentrantLock means a lock that can be entered multiple times.
type ReentrantLock struct {
	*AbstractQueuedSynchronizer
}

// Lock the ReentrantLock
func (lock *ReentrantLock) Lock() {
	Acquire(lock)
}

// Unlock the ReentrantLock
func (lock *ReentrantLock) Unlock() {
	Release(lock)
}

// NewReentrantLock creates a reentrantlock instance for optimistic locking.
func NewReentrantLock() *ReentrantLock {
	return &ReentrantLock{AbstractQueuedSynchronizer: NewAQS()}
}

func (lock *ReentrantLock) tryAcquire(goID int) bool {
	if lock.goID == int32(goID) {
		lock.state = lock.state + 1
		return true
	} else if lock.goID == vacant {
		// CAS start
		a := atomic.CompareAndSwapInt32(&lock.goID, vacant, int32(goID))
		if a == false {
			return false
		}
		// CAS is successful
		lock.state = 1
		return true
	}
	return false
}

func (lock *ReentrantLock) tryRelease(goID int) bool {
	if lock.goID == int32(goID) {
		lock.state--
		if lock.state == 0 {
			lock.goID = vacant
		}
		return true
	}
	panic("cannot call unlock before lock!")
}

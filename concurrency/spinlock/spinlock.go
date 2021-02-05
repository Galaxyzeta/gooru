package spinlock

import (
	"time"
)

// SpinLock is an abstraction of optimistic locking strategy
type SpinLock struct {
	millis int
}

// New returns a spinlock object
func New(millis int) *SpinLock {
	return &SpinLock{millis: millis}
}

// SpinUntil the condition is fullfilled
func (spinlock *SpinLock) SpinUntil(expression func() bool) {
	for expression() == false {
		spinlock.doSleep()
	}
}

// SpinEq will run infinite loop until exp1 == exp2.
func (spinlock *SpinLock) SpinEq(exp1 interface{}, exp2 interface{}) {
	for exp1 != exp2 {
		spinlock.doSleep()
	}
}

func (spinlock *SpinLock) doSleep() {
	time.Sleep(time.Millisecond * time.Duration(spinlock.millis))
}

package async

import (
	"errors"
	"sync"
	"time"
)

// Future represents a future job.
type Future struct {
	wg      *sync.WaitGroup
	val     interface{}
	err     error
	t       <-chan time.Time
	sigkill chan struct{}
}

// ErrCancel represents an error that the job was canceled before Deadline.
var ErrCancel error = errors.New("canceled during timer")

// TaskFunc represents a job.
type TaskFunc func() (interface{}, error)

// ExecuteAfter a delayed job.
func ExecuteAfter(task TaskFunc, duration int, unit time.Duration) *Future {
	future := &Future{
		wg:      &sync.WaitGroup{},
		sigkill: make(chan struct{}),
	}
	future.wg.Add(1)
	future.t = time.After(time.Duration(duration) * unit)
	go func() {
		select {
		case <-future.t:
			break
		case <-future.sigkill:
			future.err = ErrCancel
			future.wg.Done()
			return
		}
		future.val, future.err = task()
		future.wg.Done()
	}()
	return future
}

// Get will block until the given job was done. Return its result and error.
func (f *Future) Get() (interface{}, error) {
	f.wg.Wait()
	return f.val, f.err
}

// Cancel the job before it is executed.
func (f *Future) Cancel() {
	f.sigkill <- struct{}{}
}

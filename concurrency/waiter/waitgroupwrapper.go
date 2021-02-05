package waiter

import "sync"

// WaitGroupWrapper is a wrapper for sync.WaitGroup
type WaitGroupWrapper struct {
	wg *sync.WaitGroup
}

// New WaitGroupWrapper object.
func New() *WaitGroupWrapper {
	return &WaitGroupWrapper{wg: &sync.WaitGroup{}}
}

// AddAndRun wraps a function into goroutine, boot it, and increment workgroup.
func (f *WaitGroupWrapper) AddAndRun(fn func()) {
	f.wg.Add(1)
	go fn()
}

// Wait all added goroutines to finish.
func (f *WaitGroupWrapper) Wait() {
	f.wg.Wait()
}

// Done is a wrapper for wg.Done().
func (f *WaitGroupWrapper) Done() {
	f.wg.Done()
}

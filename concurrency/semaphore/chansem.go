package semaphore

// ChanSem defines the count of resources. Once the resource is not sufficient, the action of obtaining resources will block.
//
// ChanSem is much slower than MutexSem.
type ChanSem struct {
	container chan bool
}

// NewChanSem creates a semaphore
func NewChanSem(val int, max int) (sem *ChanSem) {
	if val > max {
		panic("val should always be less than max")
	} else if val < 0 {
		panic("val cannot be negative")
	}
	sem = &ChanSem{container: make(chan bool, max)}
	for i := 0; i < val; i++ {
		sem.container <- true
	}
	return sem
}

// Acquire a resource.
func (sem *ChanSem) Acquire() {
	<-sem.container
}

// Release a resource.
func (sem *ChanSem) Release() {
	sem.container <- true
}

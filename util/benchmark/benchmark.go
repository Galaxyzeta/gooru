package benchmark

import "time"

// BenchMark is used for benchmark testing.
type BenchMark struct {
	time time.Time
}

// New timer.
func New() *BenchMark {
	return &BenchMark{}
}

// StartTiming updates timer.
func (timer *BenchMark) StartTiming() {
	timer.time = time.Now()
}

// TimeElapsed returns how many time has been elapesed.
func (timer *BenchMark) TimeElapsed() int64 {
	return time.Since(timer.time).Nanoseconds()
}

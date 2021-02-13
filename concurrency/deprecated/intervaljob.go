package async

import (
	"time"
)

// Interval represents a chrono job.
type Interval struct {
	ticker   *time.Ticker
	fn       func()
	interval int
	running  bool
	sigkill  chan struct{}
}

// NewInterval creates a new Interval job
func NewInterval(interval int, unit time.Duration, fn func()) *Interval {
	ticker := time.NewTicker(time.Duration(interval) * unit)
	return &Interval{
		ticker:  ticker,
		fn:      fn,
		sigkill: make(chan struct{}),
	}
}

// Start the chrono job. Param async indicates whether to generate a new goroutine while calling the given func, or execute it directly.
func (job *Interval) Start(async bool) {
	if job.running == true {
		panic("cannot start an Interval twice!")
	}
	// Reset ticker channel
	if len(job.ticker.C) > 0 {
		<-job.ticker.C
	}
	job.running = true
	go func() {
		for {
			select {
			case <-job.ticker.C:
				if async {
					go job.fn()
				} else {
					job.fn()
				}
			case <-job.sigkill:
				return
			default:
			}
		}
	}()
}

// Cancel the chrono job.
func (job *Interval) Cancel() {
	if job.running == false {
		panic("cannot stop an Interval twice!")
	}
	job.ticker.Stop()
	job.sigkill <- struct{}{}
}

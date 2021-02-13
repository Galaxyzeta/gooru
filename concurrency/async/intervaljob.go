package async

import (
	"time"
)

// Interval represents a chrono job.
type Interval struct {
	ticker  *time.Ticker
	running bool
	sigkill chan struct{}
}

// Start the chrono job. Returns an interval object.
func Start(interval int, unit time.Duration, fn TaskFunc) *Interval {
	ticker := time.NewTicker(time.Duration(interval) * unit)
	job := &Interval{
		ticker:  ticker,
		sigkill: make(chan struct{}),
		running: true,
	}
	go func() {
		for {
			select {
			case <-job.ticker.C:
				fn()
			case <-job.sigkill:
				job.running = false
				return
			}
		}
	}()
	return job
}

// Cancel the chrono job.
func (job *Interval) Cancel() {
	if job.running == false {
		panic("cannot stop an Interval twice!")
	}
	job.ticker.Stop()
	job.sigkill <- struct{}{}
}

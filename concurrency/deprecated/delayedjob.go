package async

import (
	"time"
)

// TimerWheel is container for delayed tasks.
//
// @Experimental @Warning: should be DEPRECATED because it's NOT efficient.
type TimerWheel struct {
	queue   chan *Task
	size    int
	millis  int // How frequent to check time loop
	running bool
}

// Task represents a delayed job.
type Task struct {
	task     func(...interface{}) interface{}
	expire   time.Time
	waitDone chan bool
	isDone   bool
	isFail   bool
	result   interface{}
}

// ExecuteAfter puts a task into timer wheel, it will be executed after given time.
func (timer *TimerWheel) ExecuteAfter(job func(...interface{}) interface{}, expire time.Duration, unit time.Duration) *Task {
	if timer.running == false {
		panic("you cannot submit tasks before starting the timer. Call Start() to boot up the timer!")
	}
	task := &Task{task: job, expire: time.Now().Add(time.Duration(time.Second * expire)), waitDone: make(chan bool, 1)}
	timer.queue <- task
	return task
}

// NewTimer creates a new Timer.
func NewTimer(size int, millis int) *TimerWheel {
	return &TimerWheel{queue: make(chan *Task, size), size: size, millis: millis}
}

// Start boot up a timer wheel.
func (timer *TimerWheel) Start() {
	timer.running = true
	go func() {
		for {
			ctask := <-timer.queue
			if time.Now().After(ctask.expire) {
				go func() {
					res := ctask.task()
					ctask.isDone = true
					ctask.result = res
					ctask.waitDone <- true
				}()
			} else {
				time.Sleep(time.Millisecond * time.Duration(timer.millis))
				timer.queue <- ctask
			}
		}
	}()
}

// Get blocks until the task is done, and return its result.
func (task *Task) Get() interface{} {
	for {
		<-task.waitDone
		return task.result
	}
}

// Stop kills the timer wheel.
func (timer *TimerWheel) Stop() {
	timer.running = false
	for len(timer.queue) > 0 {
		task := <-timer.queue // Abandon all
		task.isDone = true
		task.isFail = true
	}
}

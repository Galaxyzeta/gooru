package async_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/galaxyzeta/concurrency/async"
	"github.com/galaxyzeta/util/assert"
)

func TestDelayedJob(t *testing.T) {
	f := async.ExecuteAfter(func() (interface{}, error) {
		return 2, nil
	}, 2, time.Second)
	// Cancel at 1s.
	go func() {
		time.Sleep(time.Second)
		f.Cancel()
	}()
	// Should get nil, canceled during timer.
	resp, err := f.Get()
	assert.EQ(resp, nil)
	assert.EQ(err, async.ErrCancel)
}

func TestIntervlJob(t *testing.T) {
	f := async.Start(1, time.Second, func() (interface{}, error) {
		fmt.Println("1")
		return 1, nil
	})
	time.Sleep(6 * time.Second)
	f.Cancel()
	assert.ThrowsPanic(f.Cancel) // Cannot cancel twice !
}

package singleflight

import (
	"fmt"
	"sync"
)

type call struct {
	val interface{}
	err error
	wg  sync.WaitGroup
}

// SingleFlight enables functions temporarily to be executed only once.
type SingleFlight struct {
	cache map[string]*call
	mu    *sync.Mutex
}

// NewSingleFlight returns a LRU elimination based singleflight cache.
func NewSingleFlight() *SingleFlight {
	return &SingleFlight{mu: &sync.Mutex{}, cache: make(map[string]*call)}
}

// Do ensures the function to be accomplished only once.
func (sf *SingleFlight) Do(key string, fn func() (interface{}, error)) (interface{}, error) {
	sf.mu.Lock()
	if v, ok := sf.cache[key]; ok {
		fmt.Println("Already Exists!")
		sf.mu.Unlock()
		v.wg.Wait()
		return v.val, v.err
	}

	c := new(call)
	sf.cache[key] = c
	sf.mu.Unlock()

	// WG start
	c.wg.Add(1)
	c.val, c.err = fn()
	retv, reterr := c.val, c.err
	c.wg.Done()

	sf.mu.Lock()
	delete(sf.cache, key)
	sf.mu.Unlock()
	return retv, reterr
}

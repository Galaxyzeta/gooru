package hashmap

import "sync"

// SafeHashMap is a hashmap using mutex for concurrency-safety.
//
// Should be DEPRECATED because it's not efficient.
type SafeHashMap struct {
	data map[interface{}]interface{}
	mu   *sync.Mutex
}

// NewSafeHashMap return a map object of generic types.
func NewSafeHashMap() *SafeHashMap {
	return &SafeHashMap{data: make(map[interface{}]interface{}), mu: &sync.Mutex{}}
}

// Put an elem into the map.
func (h *SafeHashMap) Put(k, v interface{}) {
	h.mu.Lock()
	h.data[k] = v
	h.mu.Unlock()
}

// Delete an elem from the map
func (h *SafeHashMap) Delete(k interface{}) {
	h.mu.Lock()
	delete(h.data, k)
	h.mu.Unlock()
}

// Get val from the map.
func (h *SafeHashMap) Get(k interface{}) interface{} {
	var ret interface{}
	h.mu.Lock()
	ret = h.data[k]
	h.mu.Unlock()
	return ret
}

// Len returns map size.
func (h *SafeHashMap) Len() int {
	return len(h.data)
}

// ContainsKey indicate whether the key is in the map or not.
func (h *SafeHashMap) ContainsKey(k interface{}) bool {
	_, ok := h.data[k]
	return ok
}

// Size returns map size.
func (h *SafeHashMap) Size() int {
	return len(h.data)
}

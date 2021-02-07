package hashmap

// HashMap is a wrapper for original map.
type HashMap struct {
	data map[interface{}]interface{}
	size int
}

// NewHashMap return a map object of generic types.
func NewHashMap() *HashMap {
	return &HashMap{data: make(map[interface{}]interface{})}
}

// Put an elem into the map.
func (h *HashMap) Put(k, v interface{}) {
	h.data[k] = v
	h.size++
}

// Delete an elem from the map
func (h *HashMap) Delete(k interface{}) {
	delete(h.data, k)
}

// Get val from the map.
func (h *HashMap) Get(k interface{}) interface{} {
	h.size--
	return h.data[k]
}

// Len returns map size.
func (h *HashMap) Len() int {
	return h.size
}

package hashmap

// HashMap is a wrapper for original map.
type HashMap struct {
	data map[interface{}]interface{}
}

// NewHashMap return a map object of generic types.
func NewHashMap() *HashMap {
	return &HashMap{data: make(map[interface{}]interface{})}
}

// Put an elem into the map.
func (h *HashMap) Put(k, v interface{}) {
	h.data[k] = v
}

// Delete an elem from the map
func (h *HashMap) Delete(k interface{}) {
	delete(h.data, k)
}

// Get val from the map.
func (h *HashMap) Get(k interface{}) interface{} {
	return h.data[k]
}

// Size returns map size.
func (h *HashMap) Size() int {
	return len(h.data)
}

// ContainsKey indicate whether the key is in the map or not.
func (h *HashMap) ContainsKey(k interface{}) bool {
	_, ok := h.data[k]
	return ok
}

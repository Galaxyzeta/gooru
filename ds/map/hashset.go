package hashmap

// HashSet is a wrapper for original map. It stores key with empty value.
type HashSet struct {
	data map[interface{}]empty
}

type empty struct{}

// NewHashSet return a map object of generic types.
func NewHashSet() *HashSet {
	return &HashSet{data: make(map[interface{}]empty)}
}

// Put an elem into the set.
func (h *HashSet) Put(v interface{}) {
	h.data[v] = empty{}
}

// Delete an elem from the set.
func (h *HashSet) Delete(v interface{}) {
	delete(h.data, v)
}

// Size returns the size of a set.
func (h *HashSet) Size() int {
	return len(h.data)
}

// Contains returns whether the element exists in the set.
func (h *HashSet) Contains(k interface{}) bool {
	_, ok := h.data[k]
	return ok
}

package hashmap

// Deprecated. Do not use because its NOT efficient!
// This is an experiment rewriting hashmap from java. It's conceptional, instead of efficient.

import (
	"sync"

	"galaxyzeta.com/util/common"
)

// Node represents hashmap node.
type Node struct {
	next *Node
	key  interface{}
	val  interface{}
}

// HashMap represents an auto-resized map, hash collision can be handled with linkedlist.
type HashMap struct {
	data       []*Node
	len        int
	cap        int
	loadFactor float32
	mu         *sync.Mutex
}

// New creates a new hashmap with given capacity.
func New(cap int, loadFactor ...float32) *HashMap {
	trimmed := trimCap(cap)
	ret := &HashMap{cap: trimmed, data: make([]*Node, trimmed), loadFactor: 0.75, mu: &sync.Mutex{}}
	if loadFactor != nil {
		ret.loadFactor = loadFactor[0]
	}
	return ret
}

// Cap returns hashmap capacity.
func (h *HashMap) Cap() int {
	return h.cap
}

// Size returns item count in hashmap.
func (h *HashMap) Size() int {
	return h.len
}

// Put an elem into the hashmap.
func (h *HashMap) Put(key interface{}, val interface{}) {
	hashcode := hashcode(key)
	pos := hashcode & (uint32(h.cap) - 1)
	node := &Node{key: key, val: val}
	if h.data[pos] == nil {
		h.data[pos] = node
		h.len++
	} else {
		// Insert at head if hash collision.
		current := h.data[pos]
		for current.next != nil {
			if current.key == node.key {
				// same key, overwrite
				current.val = val
				return
			}
			current = current.next
		}
		current.next = node
	}

	if float32(h.len) >= float32(h.cap)*h.loadFactor {
		// Try to resize
		h.resize()
	}
}

// Get an elem from the hashmap.
func (h *HashMap) Get(key interface{}) interface{} {
	hashcode := hashcode(key)
	pos := hashcode & (uint32(h.cap) - 1)
	current := h.data[pos]
	for current != nil {
		if current.key == key {
			return current.val
		}
		current = current.next
	}
	return nil
}

// Delete an elem from the hashmap. Return deleted node and a status bool.
func (h *HashMap) Delete(key interface{}) (*Node, bool) {
	hashcode := hashcode(key)
	pos := hashcode & (uint32(h.cap) - 1)
	current := h.data[pos]
	var prev *Node = current
	var ret *Node
	// delete header
	if current.key == key {
		ret = current
		h.data[pos] = current.next
		h.len--
		return ret, false
	}
	// else, deleta other
	current = current.next
	for current != nil {
		if current.key == key {
			ret = current
			prev.next = current.next
			h.len--
			return ret, true
		}
		prev = current
		current = current.next
	}
	return ret, false
}

// ===== [PRIVATE] =====

// trimCap turns a cap into power of 2. EG: 18 will be trimmed to 16.
func trimCap(cap int) int {
	if cap <= 0 {
		panic("capacity must be positive.")
	}
	ret, nret := 1, 2
	for nret < cap {
		ret = nret
		nret <<= 1
	}
	return ret
}

// resize hashmap, new cap is twice bigger than the old one.
func (h *HashMap) resize() {
	ncap := h.cap << 1
	if ncap <= 0 {
		panic("cannot resize. overflow.")
	}
	ndata := make([]*Node, ncap)
	for i, v := range h.data {
		if v == nil {
			continue
		}
		hashcode := hashcode(v.key)
		nbit := hashcode & uint32(h.cap) // if top bit is 1, new item pos shifts 2x
		if nbit > 0 {
			// migrate to new pos
			ndata[i+h.cap] = h.data[i]
		} else {
			// keep old pos
			ndata[i] = h.data[i]
		}
	}
	h.data = ndata
	h.cap = ncap
}

func hashcode(item interface{}) uint32 {
	cvt, ok := item.(string)
	if ok {
		return common.StringHash(cvt)
	}
	return 0
}

package hashmap

import (
	"github.com/galaxyzeta/ds/list"
	"github.com/galaxyzeta/util/alias"
)

// LRUCache is a combination of hashmap and linkedlist for lru caching.
// @Implement : ds/map => hashmap
type LRUCache struct {
	li       *list.DoubleLinkedList
	capacity int
	size     int
	hashmap  map[interface{}]*list.DNode
	fn       alias.P2Consumer
}

type kv struct {
	key   interface{}
	value interface{}
}

// NewLRUCache returns a new LinkedHashMap
func NewLRUCache(capacity int) *LRUCache {
	if capacity <= 0 {
		panic("Zero or negative capacity is not tolerated !")
	}
	return &LRUCache{
		capacity: capacity,
		li:       list.NewDoubleLinkedList(),
		hashmap:  make(map[interface{}]*list.DNode),
	}
}

// NewLRUCacheWithFunction returns a new LinkedHashMap with a custom function.
// The funtion is invoked when eliminating old elems.
func NewLRUCacheWithFunction(capacity int, fn alias.P2Consumer) *LRUCache {
	ret := NewLRUCache(capacity)
	ret.fn = fn
	return ret
}

// Get an elem from hashmap. Update its position as well. Return nil if not found.
func (cache *LRUCache) Get(k interface{}) interface{} {
	ret, ok := cache.hashmap[k]
	if ok == false {
		return nil
	}
	val := cache.li.RemoveNode(ret)
	ret.Set(val)
	cache.li.AddNode(0, ret)
	return ret.Get().(*kv).value
}

// Delete an elem from hashmap.
func (cache *LRUCache) Delete(k interface{}) interface{} {
	node := cache.hashmap[k]
	ret := cache.li.RemoveNode(node)
	delete(cache.hashmap, k)
	cache.size--
	return ret
}

// Put an elem into the cache.
func (cache *LRUCache) Put(k interface{}, v interface{}) {
	_, ok := cache.hashmap[k]
	insert := &kv{key: k, value: v}
	if ok {
		// already exist, modify its value and do update.
		cache.hashmap[k].Set(insert)
		cache.Get(k)
	} else {
		// not exist, create a new node and insert to the front.
		node := list.NewDNode(nil, insert, nil)
		cache.li.AddNode(0, node)
		cache.hashmap[k] = node
		if cache.size+1 > cache.capacity {
			// will cause overflow. eliminate old node.
			elem := cache.li.RemoveLast()
			kv0 := elem.(*kv)
			delete(cache.hashmap, kv0.key)
			if cache.fn != nil {
				cache.fn(kv0.key, kv0.value)
			}
		} else {
			cache.size++
		}
	}
}

// Size returns the size of given cache.
func (cache *LRUCache) Size() int {
	return cache.size
}

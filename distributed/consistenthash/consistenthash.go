package consistenthash

import (
	"fmt"
	"hash/crc32"
	"sort"
)

// HashFunc is a function that converts string to uint32 hashed.
type HashFunc func(data []byte) uint32

// ConsistentHash is a hash ring ranged between 0 to Math.max(uint32) with key nodes on it.
// It is used to distribute a key to a certain machine.
type ConsistentHash struct {
	hashFunc HashFunc
	hashes   []int // sorted
	hash2key map[int]string
}

// NewConsistentHash returns a new instance of ConsistentHash
func NewConsistentHash() *ConsistentHash {
	return &ConsistentHash{hashFunc: crc32.ChecksumIEEE, hash2key: make(map[int]string)}
}

// Add a new node to hash ring.
func (c *ConsistentHash) Add(key string) {
	hash := int(c.hashFunc([]byte(key)))
	if _, ok := c.hash2key[hash]; ok == true {
		panic("Hash collision ! Please change a key! ")
	}
	fmt.Printf("[PUT] Hash of key %v is %v\n", key, hash)
	c.hashes = append(c.hashes, hash)
	c.hash2key[hash] = key
	sort.Ints(c.hashes)
}

// Remove an existing key.
func (c *ConsistentHash) Remove(key string) {
	hash := int(c.hashFunc([]byte(key)))
	fmt.Printf("[DELETE] Hash of key %v is %v\n", key, hash)
	for i, v := range c.hashes {
		if v == hash {
			c.hashes = append(c.hashes[:i], c.hashes[i+1:]...)
		}
	}
	delete(c.hash2key, hash)
	sort.Ints(c.hashes)
}

// Get returns which node to distribute to with given key.
func (c *ConsistentHash) Get(key string) string {
	hash := int(c.hashFunc([]byte(key)))
	fmt.Printf("[GET] Hash of key %v is %v\n", key, hash)
	idx := sort.Search(len(c.hashes), func(i int) bool {
		return c.hashes[i] >= hash
	})
	return c.hash2key[c.hashes[idx%len(c.hashes)]]
}

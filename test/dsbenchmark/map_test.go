package testds

import (
	"fmt"
	"sync"
	"testing"

	"galaxyzeta.com/algo/compare"
	dep "galaxyzeta.com/ds/deprecated/hashmap"
	hashmap "galaxyzeta.com/ds/map"
	"galaxyzeta.com/ds/tree"

	"galaxyzeta.com/util/assert"
)

func TestBasic(t *testing.T) {
	h := dep.New(9)
	assert.EQ(h.Cap(), 8)
	h.Put("fuck", 123)
	h.Put("kcuf", 456)
	assert.EQ(h.Get("fuck"), 123)
	assert.EQ(h.Get("kcuf"), 456)
	h.Delete("fuck")
	assert.EQ(h.Get("fuck"), nil)
	assert.EQ(h.Get("kcuf"), 456)
	h.Delete("kcuf")
	assert.EQ(h.Get("kcuf"), nil)
}

func TestSystemMap(t *testing.T) {
	m := hashmap.NewHashMap()
	m.Put(1, 1)
	m.Put(1, 2)
	assert.EQ(m.Get(1), 2)
	assert.EQ(m.ContainsKey(1), true)
	assert.EQ(m.ContainsKey(2), false)
	assert.EQ(m.Get(2), nil)
	assert.EQ(m.Size(), 1)
	m.Delete(1)
	assert.EQ(m.Size(), 0)
}

func TestSystemSet(t *testing.T) {
	m := hashmap.NewHashSet()
	m.Put(1)
	assert.EQ(m.Contains(1), true)
	assert.EQ(m.Contains(2), false)
	assert.EQ(m.Size(), 1)
	m.Delete(1)
	assert.EQ(m.Size(), 0)
}

func BenchmarkResize(b *testing.B) {
	h := dep.New(16)
	for i := 0; i < 1000; i++ {
		k := fmt.Sprintf("a%dd", i)
		h.Put(k, i)
		assert.EQ(h.Get(k), i)
	}
}

func BenchmarkOrigMap(b *testing.B) {
	h := make(map[string]int, 9)
	for i := 0; i < 1000; i++ {
		k := fmt.Sprintf("a%dd", i)
		h[k] = i
		assert.EQ(h[k], i)
	}
}

func BenchmarkMutexHashMap(b *testing.B) {
	h := dep.NewSafeHashMap()
	for i := 0; i < 1000; i++ {
		go func(i int) {
			k := fmt.Sprintf("a%dd", i)
			h.Put(k, i)
			h.Get(k)
		}(i)
		// assert.EQ(h.Get(k), i)
	}
}

func BenchmarkSyncHashMap(b *testing.B) {
	h := &sync.Map{}
	for i := 0; i < 1000; i++ {
		go func(i int) {
			k := fmt.Sprintf("a%dd", i)
			h.Store(k, i)
			h.Load(k)
		}(i)
		// v, _ := h.Load(k)
		// assert.EQ(v, i)
	}
}

func TestInterfaceImplementation(t *testing.T) {
	var _ hashmap.Map = tree.NewBST(compare.BasicCompare)
	var _ hashmap.Set = hashmap.NewHashSet()
	var _ hashmap.Map = dep.NewSafeHashMap()
}

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
	h := hashmap.NewSafeHashMap()
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
}

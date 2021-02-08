package testds

import (
	"fmt"
	"testing"

	hashmap "galaxyzeta.com/ds/map"
	"galaxyzeta.com/util/assert"
)

func TestLru(t *testing.T) {
	wanted := "43210"
	expect := ""
	lru := hashmap.NewLRUCacheWithFunction(5, func(param1, param2 interface{}) {
		expect += fmt.Sprintf("%v", param1)
	})
	for i := 0; i < 5; i++ {
		lru.Put(i, i)
	}
	// 4 - 3 - 2 - 1 - 0
	for i := 4; i >= 0; i-- {
		lru.Get(i)
	}
	// 0 - 1 - 2 - 3 - 4
	for i := 6; i < 11; i++ {
		lru.Put(i, i)
	}
	// Eliminate: 4 - 3 - 2 - 1 - 0
	// Res: 10 - 9 - 8 - 7 - 6
	assert.EQ(wanted, expect)
}

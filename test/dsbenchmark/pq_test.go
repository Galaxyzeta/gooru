package testds

import (
	"testing"

	"galaxyzeta.com/algo/compare"
	"galaxyzeta.com/ds/list"
	"galaxyzeta.com/util/assert"
)

func TestPriorityQueue(t *testing.T) {
	pq := list.NewPriorityQueue()
	assert.EQ(pq.Front(), nil)
	assert.EQ(pq.Back(), nil)
	for i := 0; i < 1000; i++ {
		for j := 0; j < 2; j++ {
			pq.Offer(i)
			assert.EQ(pq.Front(), i)
		}
	}
	for i := 999; i >= 0; i-- {
		for j := 0; j < 2; j++ {
			assert.EQ(pq.Poll(), i)
		}
	}
}

type A struct {
	Val int
}

func TestPriorityQueueWithCustomCompareFunction(t *testing.T) {
	var fn compare.CompareFunc = func(a, b interface{}) int {
		if a.(A).Val > b.(A).Val {
			return compare.Greater
		}
		return compare.Less
	}
	pq := list.NewPriorityQueueWithCompare(fn)
	for i := 0; i < 1000; i++ {
		for j := 0; j < 2; j++ {
			pq.Offer(A{Val: i})
		}
	}
	for i := 999; i >= 0; i-- {
		for j := 0; j < 2; j++ {
			k := pq.Poll()
			assert.EQ(k.(A).Val, i)
		}
	}
}

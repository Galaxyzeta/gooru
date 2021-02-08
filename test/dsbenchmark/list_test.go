package testds

import (
	"container/list"
	"testing"

	ds "galaxyzeta.com/ds/list"
	"galaxyzeta.com/util/assert"
	"galaxyzeta.com/util/benchmark"
)

const iteration = 10000

var bm *benchmark.BenchMark = &benchmark.BenchMark{}

func TestMyLinkedList(t *testing.T) {
	li := ds.NewDoubleLinkedList()
	for i := 0; i < 1; i++ {
		li.AddLast(1)
		li.AddLast(2)
		li.AddLast(3) // H - 1 - 2 - 3
		assert.EQ(li.Size(), 3)
		assert.EQ(li.Get(0), 1)
		assert.EQ(li.Get(1), 2)
		assert.EQ(li.Get(2), 3)
		li.Add(0, 0) // H - 0 - 1 - 2 - 3
		assert.EQ(li.Get(0), 0)
		assert.EQ(li.Get(3), 3)
		li.Add(li.Size(), 4) // H - 0 - 1 - 2 - 3 - 4
		assert.EQ(li.Get(4), 4)
		li.Add(1, -1) // H - 0 - -1 - 1 - 2 - 3 - 4
		assert.EQ(li.Get(1), -1)

		// iter
		iter := li.Iterator()
		assert.EQ(iter.Next(), 0)
		assert.EQ(iter.Next(), -1)
		assert.EQ(iter.Next(), 1)
		assert.EQ(iter.Next(), 2)
		assert.EQ(iter.Next(), 3)
		assert.EQ(iter.Next(), 4)
		assert.ThrowsPanic(func() { iter.Next() })
		assert.ThrowsPanic(func() { iter.Remove() })
		iter = li.Iterator()
		iter.Next()
		assert.EQ(iter.Remove(), 0)
		iter.Next()
		assert.EQ(iter.Remove(), -1)
		iter.Next()
		assert.EQ(iter.Remove(), 1)
		iter.Next()
		assert.EQ(iter.Remove(), 2)
		iter.Next()
		assert.EQ(iter.Remove(), 3)
		iter.Next()
		assert.EQ(iter.Remove(), 4)
		assert.ThrowsPanic(func() { iter.Remove() })
		assert.EQ(li.Size(), 0)

		li.AddLast(0)
		li.AddLast(-1)
		li.AddLast(1)
		li.AddLast(2)
		li.AddLast(3)
		li.AddLast(4)

		assert.EQ(li.Size(), 6)
		li.RemoveFirst() // H - -1 - 1 - 2 - 3 - 4
		assert.EQ(li.Get(0), -1)
		li.RemoveLast() // H - -1 - 1 - 2 - 3
		assert.EQ(li.Get(li.Size()-1), 3)
		li.RemoveAt(3) // H - -1 - 1 - 2
		assert.EQ(li.Get(li.Size()-1), 2)
		li.RemoveAt(1) // H - -1 - 2
		assert.EQ(li.Get(0), -1)
		assert.EQ(li.Get(1), 2)
		li.RemoveAt(0) // H - 2
		assert.EQ(li.Get(0), 2)
		li.RemoveAt(0)                                // H
		assert.ThrowsPanic(func() { li.RemoveAt(0) }) // Panic
		// ---- stack ----

		li.Push(1) // H - 1
		li.Push(2) // H - 2 - 1
		li.Push(3) // H - 3 - 2 - 1
		assert.EQ(li.Peek(), 3)
		assert.EQ(li.IsEmpty(), false)
		li.Pop()
		li.Pop()
		assert.EQ(li.Peek(), 1)
		li.Pop()
		assert.ThrowsPanic(func() { li.Pop() })
		assert.EQ(li.IsEmpty(), true)

		// --- queue ---
		li.Offer(1) // H - 1
		li.Offer(2) // H - 1 - 2
		li.Offer(3) // H - 1 - 2 - 3

		assert.EQ(li.Get(li.Size()-1), 3)
		assert.EQ(li.Size(), 3)
		li.Poll()
		li.Poll()
		assert.EQ(li.Poll(), 3)
		assert.ThrowsPanic(func() { li.Poll() })
		assert.EQ(li.Size(), 0)

	}

}

func TestImplmentation(t *testing.T) {
	var _ ds.List = ds.NewSingleLinkedList()
	var _ ds.Queue = ds.NewSingleLinkedList()
	var _ ds.Stack = ds.NewSingleLinkedList()
	var _ ds.Iterator = ds.NewSingleLinkedList().Iterator()

}

func BenchmarkMyLinkedList(b *testing.B) {
	// MyLinkedList performs better.
	li := ds.NewSingleLinkedList()
	for i := 0; i < iteration; i++ {
		li.AddLast(1)
	}
	assert.EQ(li.Size(), iteration)
}

func BenchmarkSystemLinkedList(b *testing.B) {
	// System LinkedList is bad.
	li := list.New()
	for i := 0; i < iteration; i++ {
		li.PushBack(1)
	}
	assert.EQ(li.Len(), iteration)
}

func BenchmarkChanQueue(b *testing.B) {
	ch := make(chan int, 1000)
	b.StartTimer()
	for i := 0; i < 1000; i++ {
		ch <- i
	}
	for i := 0; i < 1000; i++ {
		<-ch
	}
	b.StopTimer()
}

func BenchmarkMyLinkedListAsQueue(b *testing.B) {
	var q ds.Queue = ds.NewSingleLinkedList()
	b.StartTimer()
	for i := 0; i < 1000; i++ {
		q.Offer(i)
	}
	for i := 0; i < 1000; i++ {
		q.Poll()
	}
	b.StopTimer()
}

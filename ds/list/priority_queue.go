package list

import "galaxyzeta.com/algo/compare"

// PriorityQueue represents heap, which has the ability to keep its top element the biggest / smallest.
type PriorityQueue struct {
	data []interface{}
	fn   compare.CompareFunc
}

// NewPriorityQueue creates a new PriorityQueue
func NewPriorityQueue() *PriorityQueue {
	return &PriorityQueue{data: make([]interface{}, 0), fn: compare.BasicCompare}
}

// NewPriorityQueueWithCompare creates a new PriorityQueue with compare func.
func NewPriorityQueueWithCompare(fn compare.CompareFunc) *PriorityQueue {
	return &PriorityQueue{data: make([]interface{}, 0), fn: fn}
}

// adjust tries to swap heap elements downward at pos.
func (pq *PriorityQueue) adjust(pos int) {
	left := pos*2 + 1
	right := pos*2 + 2
	cmp := pos
	sz := len(pq.data)
	if left < sz && pq.fn(pq.data[left], pq.data[cmp]) == compare.Greater {
		cmp = left
	}
	if right < sz && pq.fn(pq.data[right], pq.data[cmp]) == compare.Greater {
		cmp = right
	}
	if cmp != pos {
		pq.swap(cmp, pos)
		pq.adjust(cmp)
	}
}

// heapify makes random placed elements ordered like a heap.
func (pq *PriorityQueue) heapify() {
	sz := len(pq.data) / 2
	for ; sz >= 0; sz-- {
		pq.adjust(sz)
	}
}

// Offer an element into the PriorityQueue.
func (pq *PriorityQueue) Offer(elem interface{}) {
	pq.data = append(pq.data, elem)
	pq.shiftUp(len(pq.data) - 1)
}

// Front retrieves the top element but does not remove it.
func (pq *PriorityQueue) Front() interface{} {
	if len(pq.data) == 0 {
		return nil
	}
	return pq.data[0]
}

// Back retrieve the bottom elem of the given priority queue. NOT USEFUL!
func (pq *PriorityQueue) Back() interface{} {
	if len(pq.data) == 0 {
		return nil
	}
	return pq.data[len(pq.data)-1]
}

// IsEmpty returns whether the PriorityQueue is empty.
func (pq *PriorityQueue) IsEmpty() bool {
	return len(pq.data) == 0
}

// Size returns the actual size of the PriorityQueue.
func (pq *PriorityQueue) Size() int {
	return len(pq.data)
}

// Poll removes and retireves the top element from the PriorityQueue. Returns nil if the queue is empty.
func (pq *PriorityQueue) Poll() interface{} {
	sz := len(pq.data)
	if sz == 0 {
		return nil
	}
	ret := pq.data[0]
	pq.data[0] = pq.data[sz-1]
	pq.data = pq.data[:sz-1]
	pq.adjust(0)
	return ret
}

func (pq *PriorityQueue) shiftUp(pos int) {
	if pos == 0 {
		// Has already reached the top.
		return
	}
	parent := (pos+1)/2 - 1
	if pq.fn(pq.data[pos], pq.data[parent]) == compare.Greater {
		// Pos is bigger than Parent, swap them.
		pq.swap(pos, parent)
		pq.shiftUp(parent)
	}
}

func (pq *PriorityQueue) swap(a int, b int) {
	pq.data[a], pq.data[b] = pq.data[b], pq.data[a]
}

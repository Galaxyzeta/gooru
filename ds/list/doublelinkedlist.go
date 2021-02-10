package list

// DNode represents a double linked list.
type DNode struct {
	val  interface{}
	next *DNode
	prev *DNode
}

// DoubleLinkedList represents a list which is linked together with next and prev ptrs.
type DoubleLinkedList struct {
	tail *DNode
	head *DNode
	len  int
}

// DoubleLinkedListIterator represents the process of an iteration.
type DoubleLinkedListIterator struct {
	node         *DNode
	li           *DoubleLinkedList
	at           int
	last         *DNode // will be filled after calling iter.next().
	continuation *DNode // where to pick up and continue after an elem has been removed.
}

// Get elem of a node.
func (d *DNode) Get() interface{} {
	return d.val
}

// Set elem of a node.
func (d *DNode) Set(v interface{}) {
	d.val = v
}

// NewDoubleLinkedList returns a new double linked list.
func NewDoubleLinkedList() *DoubleLinkedList {
	node := &DNode{val: -255}
	node.prev = node
	node.next = node
	return &DoubleLinkedList{tail: node, head: node, len: 0}
}

// AddLast adds a new node to the last.
func (list *DoubleLinkedList) AddLast(val interface{}) {
	list.tail.next = NewDNode(list.tail, val, list.head)
	list.tail = list.tail.next
	list.head.prev = list.tail
	list.len++
}

// AddFirst adds a new node after the dummy.
func (list *DoubleLinkedList) AddFirst(val interface{}) {
	n := NewDNode(list.head, val, list.head.next)
	list.head.next = n
	n.next.prev = n
	list.len++
}

// Add elem before any position.
func (list *DoubleLinkedList) Add(pos int, elem interface{}) {
	if pos > list.len || pos < 0 {
		panic("cannot add. index out of range")
	}
	if pos != list.len {
		current := list.get(pos)
		current.prev.next = NewDNode(current.prev, elem, current)
		current.prev = current.prev.next
		list.len++
	} else {
		list.AddLast(elem)
	}
}

// AddNode before certain position.
func (list *DoubleLinkedList) AddNode(pos int, node *DNode) {
	if pos > list.len || pos < 0 {
		panic("cannot add. index out of range")
	}
	if pos != list.len {
		current := list.get(pos)
		node.prev = current.prev
		node.next = current
		current.prev.next = node
		current.prev = current.prev.next
		list.len++
	} else {
		node.prev = list.tail.prev
		node.next = list.tail
		list.tail.next = node
		list.tail = list.tail.next
		list.head.prev = list.tail
		list.len++
	}
}

// RemoveLast removes the last node of the given list.
func (list *DoubleLinkedList) RemoveLast() interface{} {
	return list.RemoveNode(list.tail)
}

// RemoveFirst removes the first node of the given list.
func (list *DoubleLinkedList) RemoveFirst() interface{} {
	return list.RemoveNode(list.head.next)
}

//RemoveAt removes elem at pos.
func (list *DoubleLinkedList) RemoveAt(pos int) interface{} {
	if pos >= list.len || pos < 0 {
		panic("cannot remove. index out of range")
	}
	return list.RemoveNode(list.get(pos))
}

// Size returns the length of a doubleLinkedList
func (list *DoubleLinkedList) Size() int {
	return list.len
}

// Get elem at pos.
func (list *DoubleLinkedList) Get(pos int) interface{} {
	if pos > list.len {
		panic("cannot get. index out of range.")
	}
	return list.get(pos).val
}

// remove is a private version of removing a specific DNode.
func (d *DNode) remove() interface{} {
	d.prev.next = d.next
	d.next.prev = d.prev
	ret := d.val
	d.free()
	return ret
}

// RemoveNode is a private version of removing a specific DNoda, adjusting list tail as well.
func (list *DoubleLinkedList) RemoveNode(d *DNode) interface{} {
	ret := d.remove()
	list.tail = list.head.prev
	list.len--
	return ret
}

// NewDNode generates a node for double linked list.
func NewDNode(prev *DNode, val interface{}, next *DNode) *DNode {
	return &DNode{prev: prev, val: val, next: next}
}

func (d *DNode) free() {
	d.prev = nil
	d.next = nil
	d.val = nil
}

func (list *DoubleLinkedList) get(pos int) *DNode {
	var mid int = list.len / 2
	current := list.head
	if pos <= mid {
		// Sequential traverse
		for i := 0; i <= pos; i++ {
			current = current.next
		}
	} else {
		// Reverse-Sequential traverse
		for i := 0; i < list.len-pos; i++ {
			current = current.prev
		}
	}
	return current
}

// IsEmpty returns whether the list is empty.
func (list *DoubleLinkedList) IsEmpty() bool {
	return list.len == 0
}

// ==== STACK IMPLEMENTATION ====

// Push is an alias for AddFirst.
func (list *DoubleLinkedList) Push(elem interface{}) {
	list.AddFirst(elem)
}

// Pop is an alias for RemoveFirst.
func (list *DoubleLinkedList) Pop() interface{} {
	return list.RemoveFirst()
}

// Peek is an alias for get(0).
func (list *DoubleLinkedList) Peek() interface{} {
	return list.get(0).val
}

// ==== QUEUE IMPLEMENTATION ====

// Offer inserts an elem to the tail.
func (list *DoubleLinkedList) Offer(elem interface{}) {
	list.AddLast(elem)
}

// Poll removes the first elem in the list.
func (list *DoubleLinkedList) Poll() interface{} {
	return list.RemoveFirst()
}

// Front retireves the first elem in the list.
func (list *DoubleLinkedList) Front() interface{} {
	if list.Size() == 0 {
		return nil
	}
	return list.head.next.val
}

// End retrieve the last elem in the list.
func (list *DoubleLinkedList) End() interface{} {
	if list.Size() == 0 {
		return nil
	}
	return list.tail.val
}

// ==== ITERATOR IMPLEMENTATION ====

// Iterator returns an implementation of Iterator.
func (list *DoubleLinkedList) Iterator() *DoubleLinkedListIterator {
	return &DoubleLinkedListIterator{node: list.head, li: list, at: -1}
}

// Next progress through the iteration.
func (iterator *DoubleLinkedListIterator) Next() interface{} {
	iterator.last = nil
	if iterator.at+1 >= iterator.li.len {
		panic("cannot iterate. iteration has already completed.")
	}
	if iterator.continuation != nil {
		iterator.node = iterator.continuation
		iterator.continuation = nil
	} else {
		iterator.node = iterator.node.next
	}
	iterator.at++
	iterator.last = iterator.node
	return iterator.node.val
}

// HasNext returns whether the iteration has ended.
func (iterator *DoubleLinkedListIterator) HasNext() bool {
	return iterator.at < iterator.li.len
}

// Remove removes current node.
func (iterator *DoubleLinkedListIterator) Remove() interface{} {
	if iterator.at >= iterator.li.len {
		panic("cannot iterate. iteration has already completed.")
	}
	if iterator.last == nil {
		panic("must call iterator.next() before you can use remove() !")
	}
	iterator.continuation = iterator.last.next
	ret := iterator.li.RemoveNode(iterator.last)
	iterator.last = nil
	iterator.at--
	return ret
}

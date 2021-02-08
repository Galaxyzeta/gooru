package list

// Node of a linked list.
type Node struct {
	val  interface{}
	next *Node
}

// LinkedListIterator represents the process of an iteration.
type LinkedListIterator struct {
	node         *Node
	prev         *Node
	li           *SingleLinkedList
	last         *Node
	continuation *Node
	at           int
}

// SingleLinkedList represents a singly link of nodes.
//
// Use this as Stack or Queue is much faster than the list
// provided by official.
//
// Use this as Queue performs almost the same as channel.
//
// Stack: top <------- bottom
//
// Queue: front <------- tail
type SingleLinkedList struct {
	tail *Node
	head *Node
	len  int
}

// NewSingleLinkedList : Initialize linked list, return a dummy head.
func NewSingleLinkedList() *SingleLinkedList {
	node := &Node{next: nil, val: -255}
	return &SingleLinkedList{tail: node, head: node, len: 0}
}

// AddLast : Add a node to the last of the given linkedlist.
func (list *SingleLinkedList) AddLast(val interface{}) {
	list.tail.next = &Node{val: val, next: nil}
	list.tail = list.tail.next
	list.len++
}

// RemoveLast : Remove last node of the linked list.
//
// Should be careful when calling this because it's SLOW.
func (list *SingleLinkedList) RemoveLast() interface{} {
	cur := list.head.next
	prev := list.head
	for cur != list.tail {
		prev = cur
		cur = cur.next
	}
	prev.next = nil
	list.tail = prev
	ret := cur.val
	cur.free()
	list.len--
	return ret
}

// RemoveAt removes elem at pos.
func (list *SingleLinkedList) RemoveAt(pos int) interface{} {
	if pos >= list.len || pos < 0 {
		panic("cannot remove. index out of range")
	} else if pos == list.len-1 {
		return list.RemoveLast()
	}
	current := list.head.next
	prev := list.head
	for i := 0; i < pos; i++ {
		prev = current
		current = current.next
	}
	prev.next = current.next
	ret := current.val
	current.free()
	list.len--
	return ret
}

// Add elem at any position.
func (list *SingleLinkedList) Add(i int, elem interface{}) {
	if i > list.len {
		panic("cannot insert because i > list.len")
	} else if i == list.len {
		list.AddLast(elem)
	} else {
		current := list.head.next
		prev := list.head
		for k := 0; k < i; k++ {
			prev = current
			current = current.next
		}
		prev.next = newNode(elem, current)
		list.len++
	}
}

// AddFirst : Add the node after head node.
func (list *SingleLinkedList) AddFirst(val interface{}) {
	tmp := list.head.next
	list.head.next = newNode(val, tmp)
	if tmp == nil {
		list.tail = list.head.next
	}
	list.len++
}

// RemoveFirst : Remove the first node.
func (list *SingleLinkedList) RemoveFirst() interface{} {
	tmp := list.head.next
	if tmp == nil {
		panic("cannot remove first. list is empty.")
	}
	if tmp.next == nil {
		list.tail = list.head
	}
	list.head.next = tmp.next
	ret := tmp.val
	tmp.free()
	list.len--
	return ret
}

// IndexOf : Find the node with given value
func (list *SingleLinkedList) IndexOf(target interface{}) *Node {
	cur := list.head.next
	for cur != nil {
		if cur.val == target {
			return cur
		}
		cur = cur.next
	}
	return nil
}

// Reverse the whole list
func (list *SingleLinkedList) Reverse() {
	cur := list.head.next
	prev := list.head
	for cur != nil {
		next := cur.next
		cur.next = prev
		prev = cur
		cur = next
	}
	prev.next = nil
	list.head.next = prev
}

// Traverse : Pass every node of the given list to the consumer function.
func (list *SingleLinkedList) Traverse(fn func(i interface{})) {
	cur := list.head.next
	for cur != nil {
		fn(cur)
		cur = cur.next
	}
}

// Size returns the length of a singlyLinkedList
func (list *SingleLinkedList) Size() int {
	return list.len
}

// Get pos i. Zero based.
func (list *SingleLinkedList) Get(i int) interface{} {
	if i >= list.len {
		panic("list access out of range")
	}
	if i == list.len-1 {
		return list.tail.val
	}
	current := list.head
	for k := 0; k <= i && current != nil; k++ {
		current = current.next
	}
	return current.val
}

// IsEmpty returns whether the list is empty.
func (list *SingleLinkedList) IsEmpty() bool {
	return list.len == 0
}

// ==== STACK IMPLEMENTATION ====

// Push is an alias for AddFirst.
func (list *SingleLinkedList) Push(elem interface{}) {
	list.AddFirst(elem)
}

// Pop is an alias for RemoveFirst.
func (list *SingleLinkedList) Pop() interface{} {
	return list.RemoveFirst()
}

// Peek is an alias for get(0).
func (list *SingleLinkedList) Peek() interface{} {
	return list.Get(0)
}

// ==== QUEUE IMPLEMENTATION ====

// Offer inserts an elem to the tail.
func (list *SingleLinkedList) Offer(elem interface{}) {
	list.AddLast(elem)
}

// Poll removes the first elem in the list.
func (list *SingleLinkedList) Poll() interface{} {
	return list.RemoveFirst()
}

// ==== ITERATOR IMPLEMENTATION ====

// Iterator returns an implementation of Iterator.
func (list *SingleLinkedList) Iterator() *LinkedListIterator {
	return &LinkedListIterator{node: list.head, prev: list.head, li: list, at: -1}
}

// Next progress through the iteration.
func (iterator *LinkedListIterator) Next() interface{} {
	iterator.last = nil
	if iterator.at+1 >= iterator.li.len {
		panic("Cannot operate. Iteration has already been completed!")
	}
	if iterator.continuation != nil {
		iterator.node = iterator.continuation
	} else {
		iterator.prev = iterator.node
		iterator.node = iterator.node.next
	}
	iterator.at++
	iterator.last = iterator.node
	return iterator.node.val
}

// HasNext returns whether the iteration has ended.
func (iterator *LinkedListIterator) HasNext() bool {
	return iterator.node.next == nil
}

// Remove removes current node.
func (iterator *LinkedListIterator) Remove() interface{} {
	if iterator.at >= iterator.li.len {
		panic("cannot iterate. iteration has already completed.")
	}
	if iterator.last == nil {
		panic("must call iterator.next() before you can use remove() !")
	}
	ret := iterator.last
	if ret == iterator.li.tail {
		iterator.li.tail = iterator.prev
	}
	iterator.continuation = iterator.last.next
	iterator.prev.next = iterator.continuation
	iterator.li.len--
	iterator.at--
	retval := ret.val
	ret.free()
	iterator.last = nil
	return retval
}

// ==== PRIVATE =====

func (t *Node) free() {
	t.next = nil
	t.val = nil
}

func newNode(elem interface{}, next *Node) *Node {
	return &Node{val: elem, next: next}
}

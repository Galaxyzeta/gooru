package ds

// ElementType is the alias of any types.
type ElementType interface{}

// Node of a linked list
type Node struct {
	val  ElementType
	next *Node
}

// LinkedList : a node connected list.
type LinkedList struct {
	tail *Node
	head *Node
	len  int
}

type consumer func(*Node)

// Init : Initialize linked list, return a dummy head.
func Init() *LinkedList {
	node := &Node{next: nil, val: 0}
	return &LinkedList{tail: node, head: node, len: 0}
}

// AddLast : Add a node to the last of the given linkedlist.
func AddLast(node *LinkedList, val ElementType) {
	node.tail.next = &Node{val: val, next: nil}
	node.tail = node.tail.next
}

// RemoveLast : Remove last node of the linked list.
func RemoveLast(node *LinkedList) ElementType {
	cur := node.head
	prev := node.head
	for cur != nil {
		prev = cur
		cur = cur.next
	}
	prev.next = nil
	node.tail = prev
	return cur.val
}

// AddFirst : Add the node after head node.
func AddFirst(list *LinkedList, val ElementType) {
	tmp := list.head.next
	list.head.next = &Node{val: val, next: tmp}
}

// RemoveFirst : Remove the first node.
func RemoveFirst(list *LinkedList) ElementType {
	tmp := list.head.next
	if tmp == nil {
		return nil
	}
	list.head.next = tmp.next
	tmp.next = nil
	return tmp.val
}

// IndexOf : Find the node with given value
func IndexOf(list *LinkedList, target ElementType) *Node {
	cur := list.head.next
	for cur != nil {
		if cur.val == target {
			return cur
		}
		cur = cur.next
	}
	return nil
}

// ReverseList : Flip the whole list
func ReverseList(list *LinkedList) {
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

// ListTraverse : Pass every node of the given list to the consumer function.
func ListTraverse(list *LinkedList, fn consumer) {
	cur := list.head.next
	for cur != nil {
		fn(cur)
		cur = cur.next
	}
}

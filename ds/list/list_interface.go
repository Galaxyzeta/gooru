package list

// List is an interface defines common behaviors of a list.
type List interface {
	Get(i int) interface{}
	RemoveAt(i int) interface{}
	Add(i int, elem interface{})
	Size() int
}

// Stack is an interface represents a first in, last out list.
type Stack interface {
	Push(elem interface{})
	Pop() interface{}
	Peek() interface{}
	IsEmpty() bool
	Size() int
}

// Queue is an interface represents a first in, first out list.
type Queue interface {
	Offer(elem interface{})
	Poll() interface{}
	IsEmpty() bool
	Size() int
	Front() interface{}
	Back() interface{}
}

// Iterator is an interface represents a process of iteration.
type Iterator interface {
	Next() interface{}
	HasNext() bool
	Remove() interface{}
}

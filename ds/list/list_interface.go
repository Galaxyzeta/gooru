package list

// List is an interface defines common behaviors of a list.
type List interface {
	Get(i int) ElemType
	RemoveAt(i int) ElemType
	Add(i int, elem ElemType)
	Size() int
}

// Stack is an interface represents a first in, last out list.
type Stack interface {
	Push(elem ElemType)
	Pop() ElemType
	Peek() ElemType
	IsEmpty() bool
	Size() int
}

// Queue is an interface represents a first in, first out list.
type Queue interface {
	Offer(elem ElemType)
	Poll() ElemType
	IsEmpty() bool
	Size() int
}

// Iterator is an interface represents a process of iteration.
type Iterator interface {
	Next() ElemType
	HasNext() bool
	Remove() ElemType
}

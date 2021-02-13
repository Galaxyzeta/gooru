package list_test

import (
	"testing"

	sys "container/list"

	"github.com/galaxyzeta/ds/list"
)

func BenchmarkDoubleLinkedList(b *testing.B) {
	k := list.NewDoubleLinkedList()
	for i := 0; i < 1000; i++ {
		k.AddLast(i)
	}
	// for i := 0; i < 1000; i++ {
	// 	assert.EQ(k.Get(i), i)
	// }
}

func BenchmarkSysDoubleLinkedList(b *testing.B) {
	k := sys.New()
	for i := 0; i < 1000; i++ {
		k.PushBack(i)
	}
}

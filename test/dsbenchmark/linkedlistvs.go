package main

import (
	"container/list"
	"fmt"

	ds "galaxyzeta.com/ds/linkedlist"
	"galaxyzeta.com/util/benchmark"
)

const iteration = 10000

var bm *benchmark.BenchMark = &benchmark.BenchMark{}

func main() {
	testMyLinkedList()
	testSystemLinkedList()
}

func testMyLinkedList() {
	// MyLinkedList performs better.
	li := ds.Init()
	bm.StartTiming()
	for i := 0; i < iteration; i++ {
		ds.AddLast(li, 1)
	}
	fmt.Println(bm.TimeElapsed())
}

func testSystemLinkedList() {
	// System LinkedList is bad.
	li := list.New()
	bm.StartTiming()
	for i := 0; i < iteration; i++ {
		li.PushBack(1)
	}
	fmt.Println(bm.TimeElapsed())
}

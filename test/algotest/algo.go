package main

import (
	"fmt"

	"galaxyzeta.com/algo"
	"galaxyzeta.com/util/abbr"
)

func main() {
	insertionSortTest()
}

func insertionSortTest() {
	arr := []int{1, 3, 4, 2, 5}
	arr2 := abbr.Stoi(arr)
	algo.QuickSort(&algo.SimpleTypeComparator{Data: arr2, Asc: true}, 0, len(arr)-1)
	fmt.Println(arr2)
}

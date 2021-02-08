package algo

import "../algo/compare"

// InsertionSort is kinda primitive sort. O(n^2).
func InsertionSort(arr compare.ISortable, start int, end int) {
	for i := start + 1; i <= end; i++ {
		for j := i; j > 0 && arr.Cmp(j, j-1); j-- {
			arr.Swap(j, j-1)
		}
	}
}

// QuickSort is an advanced sorting technique with divide-and-conquer mechanism. O(n * log(n)) for average.
func QuickSort(arr compare.ISortable, start int, end int) {
	if start >= end {
		return
	}
	left, right := start, end
	pivot := left
	for left < right {
		for left < right && arr.Cmp(right, pivot) {
			right--
		}
		arr.Swap(pivot, right)
		pivot = right
		for left < right && arr.Cmp(right, left) {
			left++
		}
		arr.Swap(right, left)
		pivot = left
	}
	arr.Swap(pivot, right)
	QuickSort(arr, start, left-1)
	QuickSort(arr, left+1, end)
}

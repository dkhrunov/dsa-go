package sorting

import (
	"github.com/dkhrunov/dsa-go/utils"
	"golang.org/x/exp/constraints"
)

// Time Complexity: O(n*log n)
//
// Auxiliary Space: O(1)
//
// Documentation - https://www.programiz.com/dsa/heap-sort
func HeapSort[T constraints.Ordered](arr []T, comp utils.ComparatorFn[T]) {
	size := len(arr)

	// Build MaxHeap/MinHeap
	for i := size/2 - 1; i >= 0; i-- {
		heapify(arr, size, i, comp)
	}

	// Heap sort
	for i := size - 1; i >= 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]
		heapify(arr, i, 0, comp)
	}
}

// Time Complexity: O(n*log n)
//
// Auxiliary Space: O(1)
//
// Documentation - https://www.programiz.com/dsa/heap-sort
func HeapSortAsc[T constraints.Ordered](arr []T) {
	HeapSort(arr, utils.LessComparator[T])
}

// Time Complexity: O(n*log n)
//
// Auxiliary Space: O(1)
//
// Documentation - https://www.programiz.com/dsa/heap-sort
func HeapSortDesc[T constraints.Ordered](arr []T) {
	HeapSort(arr, utils.GreaterComparator[T])
}

func heapify[T constraints.Ordered](arr []T, size, parentIdx int, comp utils.ComparatorFn[T]) {
	swapIdx := parentIdx
	leftChildIdx := 2*parentIdx + 1
	rightChildIdx := 2*parentIdx + 2

	if leftChildIdx < size {
		if compare := comp(arr[leftChildIdx], arr[swapIdx]); compare == -1 {
			swapIdx = leftChildIdx
		}
	}

	if rightChildIdx < size {
		if compare := comp(arr[rightChildIdx], arr[swapIdx]); compare == -1 {
			swapIdx = rightChildIdx
		}
	}

	if swapIdx != parentIdx {
		arr[parentIdx], arr[swapIdx] = arr[swapIdx], arr[parentIdx]
		heapify(arr, size, swapIdx, comp)
	}
}

package sorting

import (
	"github.com/dkhrunov/dsa-go/utils"
	"golang.org/x/exp/constraints"
)

// Time Complexity: O(n log n)
//
// Auxiliary Space: O(n)
//
// Documentation - https://www.programiz.com/dsa/merge-sort
func MergeSort[T constraints.Ordered](arr []T, comp utils.ComparatorFn[T]) {
	mergeSort(arr, 0, len(arr)-1, comp)
}

func mergeSort[T constraints.Ordered](arr []T, lowerIdx, upperIdx int, comp utils.ComparatorFn[T]) {
	if lowerIdx < upperIdx {
		middleIdx := (lowerIdx + upperIdx) / 2
		mergeSort(arr, lowerIdx, middleIdx, comp)
		mergeSort(arr, middleIdx+1, upperIdx, comp)
		merge(arr, lowerIdx, middleIdx, upperIdx, comp)
	}
}

func merge[T constraints.Ordered](arr []T, lowerIdx, middleIdx, upperIdx int, comp utils.ComparatorFn[T]) {
	leftLen := middleIdx - lowerIdx + 1
	rightLen := upperIdx - middleIdx

	left := make([]T, leftLen)
	right := make([]T, rightLen)

	// fill subLeft
	for i := range left {
		left[i] = arr[lowerIdx+i]
	}
	// fill subRight
	for j := range right {
		right[j] = arr[middleIdx+1+j]
	}

	leftPointer, rightPointer, arrPointer := 0, 0, lowerIdx

	// sort
	for leftPointer < leftLen && rightPointer < rightLen {
		if compare := comp(left[leftPointer], right[rightPointer]); compare == 1 || compare == 0 {
			arr[arrPointer] = left[leftPointer]
			leftPointer++
		} else {
			arr[arrPointer] = right[rightPointer]
			rightPointer++
		}
		arrPointer++
	}

	// if left items, move remaining items from left to end of arr
	for leftPointer < leftLen {
		arr[arrPointer] = left[leftPointer]
		leftPointer++
		arrPointer++
	}

	// if left items, move remaining items from right to end of arr
	for rightPointer < rightLen {
		arr[arrPointer] = right[rightPointer]
		rightPointer++
		arrPointer++
	}
}

// Time Complexity: O(n log n)
//
// Auxiliary Space: O(n)
func MergeSortAsc[T constraints.Ordered](arr []T) {
	MergeSort(arr, utils.LessComparator[T])
}

// Time Complexity: O(n log n)
//
// Auxiliary Space: O(n)
func MergeSortDesc[T constraints.Ordered](arr []T) {
	MergeSort(arr, utils.GreaterComparator[T])
}

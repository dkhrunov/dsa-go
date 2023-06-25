package sorting

import (
	"sync"

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

	// fill left
	for i := range left {
		left[i] = arr[lowerIdx+i]
	}
	// fill right
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

// Time Complexity: O(n log n)
//
// Auxiliary Space: O(n)
//
// Parallel version of MergeSort, and 2 times faster than MergeSort.
//
// Documentation - https://teivah.medium.com/parallel-merge-sort-in-go-fe14c1bc006
func MergeSortParallel[T constraints.Ordered](arr []T, comp utils.ComparatorFn[T]) {
	mergeSortParallel(arr, 0, len(arr)-1, comp)
}

func mergeSortParallel[T constraints.Ordered](arr []T, lowerIdx, upperIdx int, comp utils.ComparatorFn[T]) {
	if lowerIdx >= upperIdx {
		return
	}
	// 2048
	threshold := 1 << 11
	sliceLen := upperIdx - lowerIdx

	if sliceLen <= threshold { // Sequential
		mergeSort(arr, lowerIdx, upperIdx, comp)
	} else { // Parallel
		middleIdx := (lowerIdx + upperIdx) / 2

		var wg sync.WaitGroup
		wg.Add(2)

		go func() {
			defer wg.Done()
			mergeSortParallel(arr, lowerIdx, middleIdx, comp)
		}()

		go func() {
			defer wg.Done()
			mergeSortParallel(arr, middleIdx+1, upperIdx, comp)
		}()

		wg.Wait()
		merge(arr, lowerIdx, middleIdx, upperIdx, comp)
	}
}

// Time Complexity: O(n log n)
//
// Auxiliary Space: O(n)
func MergeSortParallelAsc[T constraints.Ordered](arr []T) {
	MergeSortParallel(arr, utils.LessComparator[T])
}

// Time Complexity: O(n log n)
//
// Auxiliary Space: O(n)
func MergeSortParallelDesc[T constraints.Ordered](arr []T) {
	MergeSortParallel(arr, utils.GreaterComparator[T])
}

// Time Complexity: O(n log n)
//
// Auxiliary Space: O(n)
//
// On average, 2 times slower than MergeSort and takes up 3 times more space.
func MergeSortImmutable[T constraints.Ordered](arr []T, comp utils.ComparatorFn[T]) []T {
	if len(arr) < 2 {
		return arr
	}
	left := MergeSortImmutable(arr[:len(arr)/2], comp)
	right := MergeSortImmutable(arr[len(arr)/2:], comp)
	return mergeImmutable(left, right, comp)
}

func mergeImmutable[T constraints.Ordered](left, right []T, comp utils.ComparatorFn[T]) []T {
	res := []T{}
	leftPointer := 0
	rightPointer := 0

	// sort
	for leftPointer < len(left) && rightPointer < len(right) {
		if compare := comp(left[leftPointer], right[rightPointer]); compare == 1 || compare == 0 {
			res = append(res, left[leftPointer])
			leftPointer++
		} else {
			res = append(res, right[rightPointer])
			rightPointer++
		}
	}

	// if left items, move remaining items from left to end of arr
	for ; leftPointer < len(left); leftPointer++ {
		res = append(res, left[leftPointer])
	}

	// if left items, move remaining items from right to end of arr
	for ; rightPointer < len(right); rightPointer++ {
		res = append(res, right[rightPointer])
	}

	return res
}

// Time Complexity: O(n log n)
//
// Auxiliary Space: O(n)
func MergeSortImmutableAsc[T constraints.Ordered](arr []T) {
	MergeSortImmutable(arr, utils.LessComparator[T])
}

// Time Complexity: O(n log n)
//
// Auxiliary Space: O(n)
func MergeSortImmutableDesc[T constraints.Ordered](arr []T) {
	MergeSortImmutable(arr, utils.GreaterComparator[T])
}

// Time Complexity: O(n log n)
//
// Auxiliary Space: O(n)
//
// Parallel version of MergeSortImmutable, and 2 times faster than MergeSortImmutable.
//
// Documentation - https://teivah.medium.com/parallel-merge-sort-in-go-fe14c1bc006
func MergeSortImmutableParallel[T constraints.Ordered](arr []T, comp utils.ComparatorFn[T]) []T {
	if len(arr) < 2 {
		return arr
	}

	// 4096
	threshold := 1 << 12
	sliceLen := len(arr)

	if sliceLen <= threshold { // Sequential
		return MergeSortImmutable(arr, comp)
	} else { // Parallel
		var wg sync.WaitGroup
		wg.Add(2)

		var left, right []T

		go func() {
			defer wg.Done()
			left = MergeSortImmutableParallel(arr[:len(arr)/2], comp)
		}()

		go func() {
			defer wg.Done()
			right = MergeSortImmutableParallel(arr[len(arr)/2:], comp)
		}()

		wg.Wait()
		return mergeImmutable(left, right, comp)
	}
}

// Time Complexity: O(n log n)
//
// Auxiliary Space: O(n)
func MergeSortImmutableParallelAsc[T constraints.Ordered](arr []T) []T {
	return MergeSortImmutableParallel(arr, utils.LessComparator[T])
}

// Time Complexity: O(n log n)
//
// Auxiliary Space: O(n)
func MergeSortImmutableParallelDesc[T constraints.Ordered](arr []T) []T {
	return MergeSortImmutableParallel(arr, utils.GreaterComparator[T])
}

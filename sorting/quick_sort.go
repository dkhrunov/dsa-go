package sorting

import (
	"github.com/dkhrunov/dsa-go/utils"
	"golang.org/x/exp/constraints"
)

// Time Complexity: O(N^2)
//
// Auxiliary Space: O(log n)
//
// Documentation - https://www.geeksforgeeks.org/quick-sort/
func QuickSort[T constraints.Ordered](arr []T, comp utils.ComparatorFn[T]) {
	quickSort(arr, 0, len(arr)-1, comp)
}

func quickSort[T constraints.Ordered](arr []T, left, right int, comp utils.ComparatorFn[T]) {
	if left < right {
		pivotIdx := partition(arr, left, right, comp)
		quickSort(arr, left, pivotIdx-1, comp)
		quickSort(arr, pivotIdx+1, right, comp)
	}
}

// Time Complexity: O(N^2)
//
// Auxiliary Space: O(log n)
func QuickSortConcur[T constraints.Ordered](arr []T, comp utils.ComparatorFn[T]) {
	done := make(chan struct{})
	go quickSortConcur(arr, 0, len(arr)-1, comp, done, 5)
	<-done
}

// Time Complexity: O(N^2)
//
// Auxiliary Space: O(log n)
func quickSortConcur[T constraints.Ordered](arr []T, left, right int, comp utils.ComparatorFn[T], done chan struct{}, depth int) {
	if left >= right {
		done <- struct{}{}
		return
	}

	depth--

	pivotIdx := partition(arr, left, right, comp)

	if depth > 0 {
		childDone := make(chan struct{}, 2)
		go quickSortConcur(arr, left, pivotIdx-1, comp, childDone, depth)
		go quickSortConcur(arr, pivotIdx+1, right, comp, childDone, depth)

		<-childDone
		<-childDone
	} else {
		quickSort(arr, left, pivotIdx-1, comp)
		quickSort(arr, pivotIdx+1, right, comp)
	}

	done <- struct{}{}
}

func partition[T constraints.Ordered](arr []T, left, right int, comp utils.ComparatorFn[T]) int {
	pivot, cursor := arr[right], left-1
	for i := left; i < right; i++ {
		if c := comp(arr[i], pivot); c == 0 || c == 1 {
			cursor++
			arr[cursor], arr[i] = arr[i], arr[cursor]
		}
	}

	arr[cursor+1], arr[right] = arr[right], arr[cursor+1]

	return cursor + 1
}

// Time Complexity: O(N^2)
//
// Auxiliary Space: O(log n)
func QuickSortAsc[T constraints.Ordered](arr []T) {
	QuickSort(arr, utils.LessComparator[T])
}

// Time Complexity: O(N^2)
//
// Auxiliary Space: O(log n)
func QuickSortDesc[T constraints.Ordered](arr []T) {
	QuickSort(arr, utils.GreaterComparator[T])
}

// Time Complexity: O(N^2)
//
// Auxiliary Space: O(log n)
func QuickSortConcurAsc[T constraints.Ordered](arr []T) {
	QuickSort(arr, utils.LessComparator[T])
}

// Time Complexity: O(N^2)
//
// Auxiliary Space: O(log n)
func QuickSortConcurDesc[T constraints.Ordered](arr []T) {
	QuickSort(arr, utils.GreaterComparator[T])
}

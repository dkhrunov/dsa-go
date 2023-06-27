package sorting

import (
	"sync"

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

func quickSort[T constraints.Ordered](arr []T, lowerIdx, upperIdx int, comp utils.ComparatorFn[T]) {
	if lowerIdx < upperIdx {
		pivotIdx := partition(arr, lowerIdx, upperIdx, comp)
		quickSort(arr, lowerIdx, pivotIdx-1, comp)
		quickSort(arr, pivotIdx+1, upperIdx, comp)
	}
}

// TODO: not worked
//
// Time Complexity: O(N^2)
//
// Auxiliary Space: O(log n)
func QuickSortParallel[T constraints.Ordered](arr []T, comp utils.ComparatorFn[T]) {
	quickSortParallel(arr, 0, len(arr)-1, comp)
}

// TODO: not worked
//
// Time Complexity: O(N^2)
//
// Auxiliary Space: O(log n)
func quickSortParallel[T constraints.Ordered](arr []T, lowerIdx, upperIdx int, comp utils.ComparatorFn[T]) {
	if lowerIdx >= upperIdx {
		return
	}

	// 2048
	threshold := 1 << 11
	sliceLen := upperIdx - lowerIdx

	if sliceLen <= threshold { // Sequential
		quickSort(arr, lowerIdx, upperIdx, comp)
	} else { // Parallel
		pivotIdx := partition(arr, lowerIdx, upperIdx, comp)

		var wg sync.WaitGroup
		wg.Add(2)

		go func() {
			defer wg.Done()
			quickSortParallel(arr, lowerIdx, pivotIdx-1, comp)
		}()

		go func() {
			defer wg.Done()
			quickSortParallel(arr, pivotIdx+1, upperIdx, comp)
		}()

		wg.Wait()
	}
}

func partition[T constraints.Ordered](arr []T, lowerIdx, upperIdx int, comp utils.ComparatorFn[T]) int {
	pivot, cursor := arr[upperIdx], lowerIdx-1
	for i := lowerIdx; i < upperIdx; i++ {
		if c := comp(arr[i], pivot); c == 0 || c == 1 {
			cursor++
			arr[cursor], arr[i] = arr[i], arr[cursor]
		}
	}

	arr[cursor+1], arr[upperIdx] = arr[upperIdx], arr[cursor+1]

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
func QuickSortParallelAsc[T constraints.Ordered](arr []T) {
	QuickSort(arr, utils.LessComparator[T])
}

// Time Complexity: O(N^2)
//
// Auxiliary Space: O(log n)
func QuickSortParallelDesc[T constraints.Ordered](arr []T) {
	QuickSort(arr, utils.GreaterComparator[T])
}

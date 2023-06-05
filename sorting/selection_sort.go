package sorting

import (
	"github.com/dkhrunov/dsa-go/utils"
	"golang.org/x/exp/constraints"
)

// Time Complexity: O(N^2)
//
// Auxiliary Space: O(1)
//
// Documentation - https://www.geeksforgeeks.org/selection-sort/
func SelectionSort[T constraints.Ordered](arr []T, comp utils.ComparatorFn[T]) []T {
	for i := 0; i < len(arr)-1; i++ {
		foundIdx := i

		for j := i + 1; j < len(arr); j++ {
			if comp(arr[j], arr[foundIdx]) == 1 {
				foundIdx = j
			}
		}

		arr[i], arr[foundIdx] = arr[foundIdx], arr[i]
	}

	return arr
}

// Time Complexity: O(N^2)
//
// Auxiliary Space: O(1)
func SelectionSortAsc[T constraints.Ordered](arr []T) []T {
	return SelectionSort(arr, utils.LessComparator[T])
}

// Time Complexity: O(N^2)
//
// Auxiliary Space: O(1)
func SelectionSortDesc[T constraints.Ordered](arr []T) []T {
	return SelectionSort(arr, utils.GreaterComparator[T])
}

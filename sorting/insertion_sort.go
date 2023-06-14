package sorting

import (
	"github.com/dkhrunov/dsa-go/utils"
	"golang.org/x/exp/constraints"
)

// Time Complexity: O(N^2)
//
// Auxiliary Space: O(1)
//
// Documentation - https://www.geeksforgeeks.org/insertion-sort/
func InsertionSort[T constraints.Ordered](arr []T, comp utils.ComparatorFn[T]) []T {
	for i := 1; i < len(arr); i++ {
		curr := arr[i]
		j := i - 1

		for j >= 0 && comp(curr, arr[j]) == 1 {
			arr[j+1] = arr[j]
			j--
		}

		arr[j+1] = curr
	}
	return arr
}

// Time Complexity: O(N^2)
//
// Auxiliary Space: O(1)
func InsertionSortAsc[T constraints.Ordered](arr []T) []T {
	return InsertionSort(arr, utils.LessComparator[T])
}

// Time Complexity: O(N^2)
//
// Auxiliary Space: O(1)
func InsertionSortDesc[T constraints.Ordered](arr []T) []T {
	return InsertionSort(arr, utils.GreaterComparator[T])
}

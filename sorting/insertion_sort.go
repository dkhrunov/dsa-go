package sorting

import (
	"github.com/dkhrunov/dsa-go/utils"
	"golang.org/x/exp/constraints"
)

// Time Complexity: O(N^2)
//
// Auxiliary Space: O(1)
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

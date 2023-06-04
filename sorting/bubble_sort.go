package sorting

import (
	"github.com/dkhrunov/dsa-go/utils"
	"golang.org/x/exp/constraints"
)

// Time Complexity: O(N^2)
//
// Auxiliary Space: O(1)
func BubbleSort[T constraints.Ordered](arr []T, comp utils.ComparatorFn[T]) []T {
	swapped := true
	for i := 0; i < len(arr)-1 && swapped; i++ {
		swapped = false
		for j := 0; j < len(arr)-1-i; j++ {
			if comp(arr[j+1], arr[j]) == 1 {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}
		}
	}

	return arr
}

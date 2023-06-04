package searching

import (
	myMath "github.com/dkhrunov/dsa-go/math"
	"golang.org/x/exp/constraints"
)

// Time Complexity: O(Log N)
//
// Auxiliary Space: O(1)
func ExponentialSearch[T constraints.Ordered](arr []T, item T) int {
	if len(arr) == 0 {
		return -1
	}

	i := 1

	for i < len(arr) && arr[i] < item {
		i *= 2
	}

	from := i / 2
	end := myMath.Min(i, len(arr)-1)

	// Auxiliary Space: O(Log N)
	// return BinarySearchRecursive(arr, from, end, item)

	//  OR

	// Binary search
	// Auxiliary Space: O(1)
	for from <= end {
		mid := (from + end) / 2

		if arr[mid] == item {
			return mid
		} else if arr[mid] > item {
			end = mid - 1
		} else {
			from = mid + 1
		}
	}

	return -1
}

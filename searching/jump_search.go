package searching

import (
	"math"

	"github.com/dkhrunov/dsa-go/gmath"
	"golang.org/x/exp/constraints"
)

// Time Complexity : O(√n)
//
// Auxiliary Space : O(1)
//
// Documentation - https://www.geeksforgeeks.org/jump-search/
func JumpSearch[T constraints.Ordered](arr []T, item T) int {
	n := len(arr)
	step := int(math.Sqrt(float64(n)))

	prev := 0
	for minStep := gmath.Min(step, n) - 1; arr[minStep] < item; minStep = gmath.Min(step, n) - 1 {
		prev = step
		step += int(math.Floor(math.Sqrt(float64(n))))
		if prev >= n {
			return -1
		}
	}

	for arr[prev] < item {
		prev++

		if prev == gmath.Min(step, n) {
			return -1
		}
	}

	if arr[prev] == item {
		return prev
	}

	return -1
}

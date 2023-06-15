package searching

import "golang.org/x/exp/constraints"

// Time Complexity: O(Log N)
//
// Auxiliary Space: O(1)
//
// Documentation - https://www.geeksforgeeks.org/binary-search/
func BinarySearch[T constraints.Ordered](arr []T, item T) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := (left + right) / 2

		if arr[mid] == item {
			return mid
		}

		if arr[mid] < item {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return -1
}

// Time Complexity: O(Log N)
//
// Auxiliary Space: O(N)
//
// Documentation - https://www.geeksforgeeks.org/binary-search/
func BinarySearchRecursive[T constraints.Ordered](arr []T, item T) int {
	return binarySearchRecursive(arr, item, 0, len(arr)-1)
}

func binarySearchRecursive[T constraints.Ordered](arr []T, item T, left int, right int) int {
	if len(arr) == 0 {
		return -1
	}

	if left <= right {
		mid := left + (right-left)/2

		if arr[mid] == item {
			return mid
		}

		if arr[mid] > item {
			return binarySearchRecursive(arr, item, left, mid-1)
		}

		return binarySearchRecursive(arr, item, mid+1, right)
	}

	return -1
}

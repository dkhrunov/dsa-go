package searching

// Time Complexity: O(N)
//
// Auxiliary Space: O(1)
//
// Documentation - https://www.geeksforgeeks.org/linear-search/
func LinearSearch[T comparable](arr []T, elem T) int {
	for i, v := range arr {
		if v == elem {
			return i
		}
	}

	return -1
}

// Time Complexity: O(N)
//
// Auxiliary Space: O(N)
//
// Documentation - https://www.geeksforgeeks.org/linear-search/
func LinearSearchRecursive[T comparable](arr []T, elem T) int {
	return linearSearchRecursive(arr, elem, len(arr))
}

func linearSearchRecursive[T comparable](arr []T, elem T, size int) int {
	if size == 0 {
		return -1
	}

	if arr[size-1] == elem {
		return size - 1
	}

	return linearSearchRecursive(arr, elem, size-1)
}

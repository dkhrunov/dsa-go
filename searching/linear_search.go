package searching

// Time Complexity: O(N)
//
// Auxiliary Space: O(1)
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
func LinearSearchRecursive[T comparable](arr []T, elem T) int {
	return linearSearchRecursiveUtil(arr, elem, len(arr))
}

func linearSearchRecursiveUtil[T comparable](arr []T, elem T, size int) int {
	if size == 0 {
		return -1
	}

	if arr[size-1] == elem {
		return size - 1
	}

	return linearSearchRecursiveUtil(arr, elem, size-1)
}

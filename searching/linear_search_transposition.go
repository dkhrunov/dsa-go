package searching

// Time Complexity: O(N)
//
// Auxiliary Space: O(1)
func LinearSearchTransposition[T comparable](arr []T, elem T) int {
	found := LinearSearch(arr, elem)

	if found != -1 {
		if found != 0 {
			arr[found-1], arr[found] = arr[found], arr[found-1]
		}
		return found
	}

	return -1
}

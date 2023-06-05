package searching

// Time Complexity: O(N)
//
// Auxiliary Space: O(1)
//
// Documentation - https://www.geeksforgeeks.org/improving-linear-search-technique/
func LinearSearchMoveToFront[T comparable](arr []T, elem T) int {
	found := LinearSearch(arr, elem)

	if found != -1 {
		if found != 0 {
			arr[0], arr[found] = arr[found], arr[0]
		}
		return found
	}

	return -1
}

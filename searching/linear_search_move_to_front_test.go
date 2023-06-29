package searching

import (
	"testing"
)

func TestLinearSearchMoveToFront(t *testing.T) {
	arr := []int{1, 0, 3, 22, 6, 5, 10}
	find := 3
	result := LinearSearchMoveToFront(arr, find)
	want := 2

	if result != want {
		t.Fatalf(`LinearSearchMoveToFront(%v, %v) = %v, want match for %v`, arr, find, result, want)
	}

	if arr[0] != find {
		t.Fatalf(`LinearSearchMoveToFront(%v, %v) = %v, after call should move found item to front`, arr, find, result)
	}
}

func TestLinearSearchMoveToFrontNotFound(t *testing.T) {
	arr := []int{1, 0, 3, 22, 6, 5, 10}
	find := 3333
	result := LinearSearchMoveToFront(arr, find)
	want := -1

	if result != want {
		t.Fatalf(`LinearSearchMoveToFront(%v, %v) = %v, want match for %v`, arr, find, result, want)
	}
}

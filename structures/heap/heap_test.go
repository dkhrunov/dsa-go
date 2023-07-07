package heap

import (
	"reflect"
	"testing"
)

func TestMaxHeap(t *testing.T) {
	input := []int{2, 7, 26, 25, 19, 17, 1, 90, 3, 36}
	want := []int{90, 36, 17, 25, 26, 7, 1, 2, 3, 19}
	heap := NewMaxHeap(input...)

	if !reflect.DeepEqual(heap.arr, want) {
		t.Fatalf(`MaxHeap(%v) = %v, want match for %v`, input, heap.arr, want)
	}

	heap.Delete(90)
	wantAfterDelete := []int{36, 26, 17, 25, 19, 7, 1, 2, 3}

	if !reflect.DeepEqual(heap.arr, wantAfterDelete) {
		t.Fatalf(`MaxHeap(%v) = %v, want match for %v`, input, heap.arr, wantAfterDelete)
	}
}

func TestMinHeap(t *testing.T) {
	input := []int{2, 56, 20, 37, 90, 36, 13, 1, 18, 5, 43}
	want := []int{1, 2, 13, 18, 5, 36, 20, 56, 37, 90, 43}
	heap := NewMinHeap(input...)

	if !reflect.DeepEqual(heap.arr, want) {
		t.Fatalf(`MinHeap(%v) = %v, want match for %v`, input, heap.arr, want)
	}

	heap.Delete(20)
	wantAfterDelete := []int{1, 2, 13, 18, 5, 36, 43, 56, 37, 90}

	if !reflect.DeepEqual(heap.arr, wantAfterDelete) {
		t.Fatalf(`MinHeap(%v) = %v, want match for %v`, input, heap.arr, wantAfterDelete)
	}
}

package heap

import (
	"github.com/dkhrunov/dsa-go/utils"
	"golang.org/x/exp/constraints"
)

type Heap[T constraints.Ordered] struct {
	arr  []T
	comp utils.ComparatorFn[T]
}

func New[T constraints.Ordered](comp utils.ComparatorFn[T], items ...T) *Heap[T] {
	return &Heap[T]{items, comp}
}

func NewMaxHeap[T constraints.Ordered](items ...T) *Heap[T] {
	h := &Heap[T]{comp: utils.GreaterComparator[T]}
	for _, v := range items {
		h.Insert(v)
	}
	return h
}

func NewMinHeap[T constraints.Ordered](items ...T) *Heap[T] {
	h := &Heap[T]{comp: utils.LessComparator[T]}
	for _, v := range items {
		h.Insert(v)
	}
	return h
}

func (h *Heap[T]) Insert(v T) {
	if h.IsEmpty() {
		h.arr = append(h.arr, v)
	} else {
		h.arr = append(h.arr, v)

		for i := h.Size()/2 - 1; i >= 0; i-- {
			h.heapify(i)
		}
		// OR
		// h.heapifyUp(h.Size() - 1)
	}
}

func (h *Heap[T]) Delete(v T) {
	if h.IsEmpty() {
		return
	}

	foundIdx := -1
	for i := 0; i < len(h.arr); i++ {
		if h.arr[i] == v {
			foundIdx = i
			break
		}
	}

	if foundIdx != -1 {
		h.arr[foundIdx], h.arr[h.Size()-1] = h.arr[h.Size()-1], h.arr[foundIdx]
		h.arr = h.arr[:h.Size()-1]
		if h.Size() > 0 {
			for i := h.Size()/2 - 1; i >= 0; i-- {
				h.heapify(i)
			}
			// OR
			// h.heapifyDown(0)
		}
	}
}

func (h *Heap[T]) IsEmpty() bool {
	return h.Size() == 0
}

func (h *Heap[T]) Size() int {
	return len(h.arr)
}

// func (h *Heap[T]) heapifyUp(childIdx int) {
// 	parentIdx := (childIdx - 1) / 2

// 	if childIdx < h.Size() && parentIdx >= 0 {
// 		if compare := h.comp(h.arr[childIdx], h.arr[parentIdx]); compare == 1 {
// 			h.arr[childIdx], h.arr[parentIdx] = h.arr[parentIdx], h.arr[childIdx]
// 			h.heapifyUp(parentIdx)
// 		}
// 	}
// }

func (h *Heap[T]) heapify(parentIdx int) {
	swapIdx := parentIdx
	leftChildIdx := 2*parentIdx + 1
	rightChildIdx := 2*parentIdx + 2

	if leftChildIdx < h.Size() {
		if compare := h.comp(h.arr[leftChildIdx], h.arr[swapIdx]); compare == 1 {
			swapIdx = leftChildIdx
		}
	}

	if rightChildIdx < h.Size() {
		if compare := h.comp(h.arr[rightChildIdx], h.arr[swapIdx]); compare == 1 {
			swapIdx = rightChildIdx
		}
	}

	if swapIdx != parentIdx {
		h.arr[parentIdx], h.arr[swapIdx] = h.arr[swapIdx], h.arr[parentIdx]
		h.heapify(swapIdx)
	}
}

package structures

import (
	"github.com/dkhrunov/dsa-go/utils"
	"golang.org/x/exp/constraints"
)

type Heap[T constraints.Ordered] struct {
	arr  []T
	comp utils.ComparatorFn[T]
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
		h.heapifyUp(h.Size() - 1)
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
			h.heapifyDown(0)
		}
	}
}

func (h *Heap[T]) IsEmpty() bool {
	return h.Size() == 0
}

func (h *Heap[T]) Size() int {
	return len(h.arr)
}

func (h *Heap[T]) heapifyUp(childIdx int) {
	parentIdx := (childIdx - 1) / 2

	if childIdx < h.Size() && parentIdx >= 0 {
		if compare := h.comp(h.arr[childIdx], h.arr[parentIdx]); compare == 1 {
			h.arr[childIdx], h.arr[parentIdx] = h.arr[parentIdx], h.arr[childIdx]
			h.heapifyUp(parentIdx)
		}
	}
}

func (h *Heap[T]) heapifyDown(parentIdx int) {
	currentIdx := parentIdx
	leftChildIdx := 2*parentIdx + 1
	rightChildIdx := 2*parentIdx + 2

	if leftChildIdx < h.Size() {
		if compare := h.comp(h.arr[leftChildIdx], h.arr[currentIdx]); compare == 1 {
			currentIdx = leftChildIdx
		}
	}

	if rightChildIdx < h.Size() {
		if compare := h.comp(h.arr[rightChildIdx], h.arr[currentIdx]); compare == 1 {
			currentIdx = rightChildIdx
		}
	}

	if currentIdx != parentIdx {
		h.arr[parentIdx], h.arr[currentIdx] = h.arr[currentIdx], h.arr[parentIdx]
		h.heapifyDown(currentIdx)
	}
}

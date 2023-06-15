package sorting

import (
	"github.com/dkhrunov/dsa-go/utils"
	"golang.org/x/exp/constraints"
)

func QuickSort[T constraints.Ordered](arr []T, comp utils.ComparatorFn[T]) []T {
	quickSort(arr, 0, len(arr)-1, comp)
	return arr
}

func quickSort[T constraints.Ordered](arr []T, left, right int, comp utils.ComparatorFn[T]) {
	if left < right {
		pivotIdx := partition(arr, left, right, comp)
		quickSort(arr, left, pivotIdx-1, comp)
		quickSort(arr, pivotIdx+1, right, comp)
	}
}

func QuickSortConcur[T constraints.Ordered](arr []T, comp utils.ComparatorFn[T]) []T {
	done := make(chan struct{})
	go quickSortConcur(arr, 0, len(arr)-1, comp, done, 5)
	<-done
	return arr
}

func quickSortConcur[T constraints.Ordered](arr []T, left, right int, comp utils.ComparatorFn[T], done chan struct{}, depth int) {
	if left >= right {
		done <- struct{}{}
		return
	}

	depth--

	pivotIdx := partition(arr, left, right, comp)

	if depth > 0 {
		childDone := make(chan struct{}, 2)
		go quickSortConcur(arr, left, pivotIdx-1, comp, childDone, depth)
		go quickSortConcur(arr, pivotIdx+1, right, comp, childDone, depth)

		<-childDone
		<-childDone
	} else {
		quickSort(arr, left, pivotIdx-1, comp)
		quickSort(arr, pivotIdx+1, right, comp)
	}

	done <- struct{}{}
}

func partition[T constraints.Ordered](arr []T, left, right int, comp utils.ComparatorFn[T]) int {
	pivot, cursor := arr[right], left-1
	for i := left; i < right; i++ {
		if c := comp(arr[i], pivot); c == 0 || c == 1 {
			cursor++
			arr[cursor], arr[i] = arr[i], arr[cursor]
		}
	}

	arr[cursor+1], arr[right] = arr[right], arr[cursor+1]

	return cursor + 1
}

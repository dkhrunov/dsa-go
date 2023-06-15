package sorting

import (
	"sync"

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

func QuickSortParallel[T constraints.Ordered](arr []T, comp utils.ComparatorFn[T]) []T {
	var wg sync.WaitGroup
	wg.Add(1)
	quickSortParallel(arr, 0, len(arr)-1, comp, &wg)
	wg.Wait()
	return arr
}

func quickSortParallel[T constraints.Ordered](arr []T, left, right int, comp utils.ComparatorFn[T], wg *sync.WaitGroup) {
	defer wg.Done()
	if left < right {
		wg.Add(2)
		pivotIdx := partition(arr, left, right, comp)
		go quickSortParallel(arr, left, pivotIdx-1, comp, wg)
		go quickSortParallel(arr, pivotIdx+1, right, comp, wg)
	}
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

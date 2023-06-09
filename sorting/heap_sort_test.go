package sorting

import (
	"reflect"
	"testing"

	"github.com/dkhrunov/dsa-go/utils"
)

func TestHeapSort(t *testing.T) {
	type args struct {
		arr  []int
		comp utils.ComparatorFn[int]
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			"Should sort array in ASC order",
			args{
				[]int{4, 5, 2, 3, 8, 9, 7, 6, 1, 0},
				utils.LessComparator[int],
			},
			[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
		{
			"Should sort array in DESC order",
			args{
				[]int{4, 5, 2, 3, 8, 9, 7, 6, 1, 0},
				utils.GreaterComparator[int],
			},
			[]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			HeapSort(tt.args.arr, tt.args.comp)

			if !reflect.DeepEqual(tt.args.arr, tt.want) {
				t.Errorf("HeapSort() = %v, want %v", tt.args.arr, tt.want)
			}
		})
	}
}

func TestHeapSortAsc(t *testing.T) {

	arr := []int{4, 5, 2, 3, 8, 9, 7, 6, 1, 0}
	want := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	HeapSortAsc(arr)

	if !reflect.DeepEqual(arr, want) {
		t.Errorf("Result was incorrect, got: %v, want: %v.", arr, want)
	}
}

func TestHeapSortDesc(t *testing.T) {
	arr := []int{4, 5, 2, 3, 8, 9, 7, 6, 1, 0}
	want := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}

	HeapSortDesc(arr)

	if !reflect.DeepEqual(arr, want) {
		t.Errorf("Result was incorrect, got: %v, want: %v.", arr, want)
	}
}

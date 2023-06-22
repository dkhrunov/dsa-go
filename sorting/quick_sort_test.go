package sorting

import (
	"reflect"
	"testing"

	"github.com/dkhrunov/dsa-go/utils"
)

func TestQuickSort(t *testing.T) {
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

			QuickSort(tt.args.arr, tt.args.comp)

			if !reflect.DeepEqual(tt.args.arr, tt.want) {
				t.Errorf("QuickSort() = %v, want %v", tt.args.arr, tt.want)
			}
		})
	}
}

func TestQuickSortConcur(t *testing.T) {
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

			QuickSortConcur(tt.args.arr, tt.args.comp)

			if !reflect.DeepEqual(tt.args.arr, tt.want) {
				t.Errorf("QuickSortConcur() = %v, want %v", tt.args.arr, tt.want)
			}
		})
	}
}

func TestQuickSortAsc(t *testing.T) {
	arr := []int{4, 5, 2, 3, 8, 9, 7, 6, 1, 0}
	want := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	QuickSortAsc(arr)

	if !reflect.DeepEqual(arr, want) {
		t.Errorf("Result was incorrect, got: %v, want: %v.", arr, want)
	}
}

func TestQuickSortDesc(t *testing.T) {
	arr := []int{4, 5, 2, 3, 8, 9, 7, 6, 1, 0}
	want := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}

	QuickSortDesc(arr)

	if !reflect.DeepEqual(arr, want) {
		t.Errorf("Result was incorrect, got: %v, want: %v.", arr, want)
	}
}

func TestQuickSortConcurAsc(t *testing.T) {
	arr := []int{4, 5, 2, 3, 8, 9, 7, 6, 1, 0}
	want := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	QuickSortConcurAsc(arr)

	if !reflect.DeepEqual(arr, want) {
		t.Errorf("Result was incorrect, got: %v, want: %v.", arr, want)
	}
}

func TestQuickSortConcurDesc(t *testing.T) {
	arr := []int{4, 5, 2, 3, 8, 9, 7, 6, 1, 0}
	want := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}

	QuickSortConcurDesc(arr)

	if !reflect.DeepEqual(arr, want) {
		t.Errorf("Result was incorrect, got: %v, want: %v.", arr, want)
	}
}

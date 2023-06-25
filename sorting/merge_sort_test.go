package sorting

import (
	"reflect"
	"testing"

	"github.com/dkhrunov/dsa-go/utils"
)

func TestMergeSort(t *testing.T) {
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

			MergeSort(tt.args.arr, tt.args.comp)

			if !reflect.DeepEqual(tt.args.arr, tt.want) {
				t.Errorf("MergeSort() = %v, want %v", tt.args.arr, tt.want)
			}
		})
	}
}

func TestMergeSortAsc(t *testing.T) {
	arr := []int{4, 5, 2, 3, 8, 9, 7, 6, 1, 0}
	want := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	MergeSortAsc(arr)

	if !reflect.DeepEqual(arr, want) {
		t.Errorf("Result was incorrect, got: %v, want: %v.", arr, want)
	}
}

func TestMergeSortDesc(t *testing.T) {
	arr := []int{4, 5, 2, 3, 8, 9, 7, 6, 1, 0}
	want := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}

	MergeSortDesc(arr)

	if !reflect.DeepEqual(arr, want) {
		t.Errorf("Result was incorrect, got: %v, want: %v.", arr, want)
	}
}

func TestMergeSortParallel(t *testing.T) {
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

			MergeSortParallel(tt.args.arr, tt.args.comp)

			if !reflect.DeepEqual(tt.args.arr, tt.want) {
				t.Errorf("MergeSortParallel() = %v, want %v", tt.args.arr, tt.want)
			}
		})
	}
}

func TestMergeSortParallelAsc(t *testing.T) {
	arr := []int{4, 5, 2, 3, 8, 9, 7, 6, 1, 0}
	want := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	MergeSortParallelAsc(arr)

	if !reflect.DeepEqual(arr, want) {
		t.Errorf("Result was incorrect, got: %v, want: %v.", arr, want)
	}
}

func TestMergeSortParallelDesc(t *testing.T) {
	arr := []int{4, 5, 2, 3, 8, 9, 7, 6, 1, 0}
	want := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}

	MergeSortParallelDesc(arr)

	if !reflect.DeepEqual(arr, want) {
		t.Errorf("Result was incorrect, got: %v, want: %v.", arr, want)
	}
}

func TestMergeSortImmutable(t *testing.T) {
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

			if got := MergeSortImmutable(tt.args.arr, tt.args.comp); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeSortImmutable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMergeSortImmutableParallel(t *testing.T) {
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

			if got := MergeSortImmutableParallel(tt.args.arr, tt.args.comp); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeSortImmutableParallel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMergeSortImmutableParallelAsc(t *testing.T) {
	arr := []int{4, 5, 2, 3, 8, 9, 7, 6, 1, 0}
	want := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	if got := MergeSortImmutableParallelAsc(arr); !reflect.DeepEqual(got, want) {
		t.Errorf("Result was incorrect, got: %v, want: %v.", got, want)
	}
}

func TestMergeSortImmutableParallelDesc(t *testing.T) {
	arr := []int{4, 5, 2, 3, 8, 9, 7, 6, 1, 0}
	want := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}

	if got := MergeSortImmutableParallelDesc(arr); !reflect.DeepEqual(got, want) {
		t.Errorf("Result was incorrect, got: %v, want: %v.", got, want)
	}
}

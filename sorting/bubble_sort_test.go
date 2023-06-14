package sorting

import (
	"reflect"
	"testing"

	"github.com/dkhrunov/dsa-go/utils"
)

func TestBubbleSort(t *testing.T) {
	type args struct {
		arr  []int
		comp utils.ComparatorFn[int]
	}
	arr := []int{4, 5, 2, 3, 8, 9, 7, 6, 1, 0}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{"Should sort array in ASC order", args{arr, utils.LessComparator[int]}, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}},
		{"Should sort array in DESC order", args{arr, utils.GreaterComparator[int]}, []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BubbleSort(tt.args.arr, tt.args.comp); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BubbleSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBubbleSortAsc(t *testing.T) {
	arr := []int{4, 5, 2, 3, 8, 9, 7, 6, 1, 0}
	want := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	if got := BubbleSortAsc(arr); !reflect.DeepEqual(got, want) {
		t.Errorf("Result was incorrect, got: %v, want: %v.", got, want)
	}
}

func TestBubbleSortDesc(t *testing.T) {
	arr := []int{4, 5, 2, 3, 8, 9, 7, 6, 1, 0}
	want := []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}

	if got := BubbleSortDesc(arr); !reflect.DeepEqual(got, want) {
		t.Errorf("Result was incorrect, got: %v, want: %v.", got, want)
	}
}

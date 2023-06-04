package searching

import (
	"testing"

	"golang.org/x/exp/constraints"
)

func TestJumpSearch(t *testing.T) {
	type args[T constraints.Ordered] struct {
		arr  []T
		item T
	}
	type test struct {
		name string
		args args[int]
		want int
	}
	arr := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}
	tests := []test{
		{"Found at the start of array", args[int]{arr, 3}, 1},
		{"Found at the end of array", args[int]{arr, 19}, 9},
		{"Found at the middle of array", args[int]{arr, 11}, 5},
		{"Not found 0", args[int]{arr, 0}, -1},
		{"Not found 100", args[int]{arr, 100}, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := JumpSearch(tt.args.arr, tt.args.item); got != tt.want {
				t.Errorf("JumpSearch(%v, %v) = %v, want %v", tt.args.arr, tt.args.item, got, tt.want)
			}
		})
	}
}

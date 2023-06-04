package searching

import (
	"testing"
)

func TestExponentialSearch(t *testing.T) {
	type args struct {
		arr  []int
		item int
	}
	type test struct {
		name string
		args args
		want int
	}
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	tests := []test{
		{"3 should be an index 2", args{arr, 3}, 2},
		{"9 should be an index 8", args{arr, 9}, 8},
		{"16 should be an index 15", args{arr, 16}, 15},
		{"13 should be an index 12", args{arr, 13}, 12},
		{"999 should be -1", args{arr, 999}, -1},
		{"Should be -1 of empty slice", args{make([]int, 0), 1}, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExponentialSearch(tt.args.arr, tt.args.item); got != tt.want {
				t.Errorf("ExponentialSearch(%v, %v) = %v, want %v", tt.args.arr, tt.args.item, got, tt.want)
			}
		})
	}
}

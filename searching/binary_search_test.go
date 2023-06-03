package searching

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

func TestBinarySearch(t *testing.T) {
	type args struct {
		arr  []int
		item int
	}

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"3 should be an index 2", args{arr, 3}, 2},
		{"9 should be an index 8", args{arr, 9}, 8},
		{"16 should be an index 15", args{arr, 16}, 15},
		{"13 should be an index 12", args{arr, 13}, 12},
		{"999 should be -1", args{arr, 999}, -1},
		{"Should be -1 of empty arr", args{make([]int, 0), 1}, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := BinarySearch(tt.args.arr, tt.args.item)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}
func TestBinarySearchRecursive(t *testing.T) {
	type args struct {
		arr  []int
		item int
	}

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"3 should be an index 2", args{arr, 3}, 2},
		{"9 should be an index 8", args{arr, 9}, 8},
		{"999 should be -1", args{arr, 999}, -1},
		{"Should be -1 of empty arr", args{make([]int, 0), 1}, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := BinarySearchRecursive(tt.args.arr, tt.args.item)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

var benchmarkBSarr = rand.Perm(10000)

func BenchmarkBinarySearch(b *testing.B) {
	rand.Seed(time.Now().Unix())
	sort.Ints(benchmarkBSarr)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		search := rand.Intn(10000)
		BinarySearch(benchmarkBSarr, search)
	}
}

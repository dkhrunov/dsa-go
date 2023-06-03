package searching

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

type lsInputs struct {
	arr  []int
	find int
}

func TestLinearSearch(t *testing.T) {
	arr := []int{1, 0, 3, 22, 6, 5, 10}
	tests := []struct {
		name   string
		inputs lsInputs
		want   int
	}{
		{fmt.Sprintf("22 should be an index 3 in %v", arr), lsInputs{arr, 22}, 3},
		{fmt.Sprintf("999 should be -1 in %v", arr), lsInputs{arr, 999}, -1},
		{"Should be -1 of [] (empty slice)", lsInputs{make([]int, 0), 1}, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := LinearSearch(tt.inputs.arr, tt.inputs.find)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

// Recursion version
func TestLinearSearchRecursive(t *testing.T) {
	arr := []int{90, 67, 56, 88, 74, 32, 16}
	tests := []struct {
		name   string
		inputs lsInputs
		want   int
	}{
		{"56 should be an index 2", lsInputs{arr, 56}, 2},
		{"999 should be -1", lsInputs{arr, 999}, -1},
		{"Should be -1 of empty slice", lsInputs{make([]int, 0), 90}, -1},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := LinearSearchRecursive(tt.inputs.arr, tt.inputs.find)
			if ans != tt.want {
				t.Errorf("got %d, want %d", ans, tt.want)
			}
		})
	}
}

var benchmarkLSSlice = rand.Perm(10000)

func BenchmarkLinearSearch(b *testing.B) {
	rand.Seed(time.Now().Unix())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LinearSearch(benchmarkLSSlice, 1)
	}
}

func BenchmarkLinearSearchRecursive(b *testing.B) {
	rand.Seed(time.Now().Unix())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		search := rand.Intn(10000)
		LinearSearchRecursive(benchmarkLSSlice, search)
	}
}

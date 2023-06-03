package searching

import (
	"math/rand"
	"testing"
	"time"
)

func TestLinearSearchMoveToFront(t *testing.T) {
	arr := []int{1, 0, 3, 22, 6, 5, 10}
	find := 3
	result := LinearSearchMoveToFront(arr, find)
	want := 2

	if result != want {
		t.Fatalf(`LinearSearchMoveToFront(%v, %v) = %v, want match for %v`, arr, find, result, want)
	}

	if arr[0] != find {
		t.Fatalf(`LinearSearchMoveToFront(%v, %v) = %v, after call should move found item to front`, arr, find, result)
	}
}

func TestLinearSearchMoveToFrontNotFound(t *testing.T) {
	arr := []int{1, 0, 3, 22, 6, 5, 10}
	find := 3333
	result := LinearSearchMoveToFront(arr, find)
	want := -1

	if result != want {
		t.Fatalf(`LinearSearchMoveToFront(%v, %v) = %v, want match for %v`, arr, find, result, want)
	}
}

var benchmarkLSTFarr = rand.Perm(10000)

func BenchmarkLinearSearchToFront(b *testing.B) {
	rand.Seed(time.Now().Unix())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		search := rand.Intn(10000)
		LinearSearchMoveToFront(benchmarkLSTFarr, search)
	}
}

func BenchmarkLinearSearchToFrontWithSameValueMultipleTimes(b *testing.B) {
	rand.Seed(time.Now().Unix())
	search := rand.Intn(10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LinearSearchMoveToFront(benchmarkLSTFarr, search)
	}
}

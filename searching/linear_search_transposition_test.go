package searching

import (
	"math/rand"
	"testing"
	"time"
)

func TestLinearSearchTransposition(t *testing.T) {
	arr := []int{1, 0, 3, 22, 6, 5, 10}

	cases := []struct {
		target, want int
	}{
		{3, 2},
		{3, 1},
		{3, 0},
		{3, 0},
	}

	for i, c := range cases {
		got := LinearSearchTransposition(arr, c.target)

		if got != c.want {
			t.Errorf("[Case №%v] LinearSearchTransposition(%v, %v) == %v, want %v", i+1, arr, c.target, got, c.want)
		}
	}
}

func TestLinearSearchTranspositionNotFound(t *testing.T) {
	arr := []int{1, 0, 3, 22, 6, 5, 10}
	find := 999
	result := LinearSearchTransposition(arr, find)
	want := -1

	if result != want {
		t.Fatalf(`LinearSearchTransposition(%v, %v) = %v, want match for %v`, arr, find, result, want)
	}
}

var benchmarkLSTarr = rand.Perm(10000)

func BenchmarkLinearSearchTransposition(b *testing.B) {
	rand.Seed(time.Now().Unix())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		search := rand.Intn(10000)
		LinearSearchTransposition(benchmarkLSTarr, search)
	}
}

func BenchmarkLinearSearchTranspositionWithSameValueMultipleTimes(b *testing.B) {
	rand.Seed(time.Now().Unix())
	search := rand.Intn(10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LinearSearchTransposition(benchmarkLSTarr, search)
	}
}

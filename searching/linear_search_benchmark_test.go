package searching

import (
	"math/rand"
	"testing"
	"time"
)

func BenchmarkLinearSearch(b *testing.B) {
	benchmarkLSSlice := rand.Perm(10000)
	rand.Seed(time.Now().Unix())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LinearSearch(benchmarkLSSlice, 1)
	}
}

func BenchmarkLinearSearchRecursive(b *testing.B) {
	benchmarkLSSlice := rand.Perm(10000)
	rand.Seed(time.Now().Unix())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		search := rand.Intn(10000)
		LinearSearchRecursive(benchmarkLSSlice, search)
	}
}

func BenchmarkLinearSearchToFront(b *testing.B) {
	benchmarkLSTFarr := rand.Perm(10000)
	rand.Seed(time.Now().Unix())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		search := rand.Intn(10000)
		LinearSearchMoveToFront(benchmarkLSTFarr, search)
	}
}

func BenchmarkLinearSearchToFrontWithSameValueMultipleTimes(b *testing.B) {
	benchmarkLSTFarr := rand.Perm(10000)
	rand.Seed(time.Now().Unix())
	search := rand.Intn(10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LinearSearchMoveToFront(benchmarkLSTFarr, search)
	}
}

func BenchmarkLinearSearchTransposition(b *testing.B) {
	benchmarkLSTarr := rand.Perm(10000)
	rand.Seed(time.Now().Unix())
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		search := rand.Intn(10000)
		LinearSearchTransposition(benchmarkLSTarr, search)
	}
}

func BenchmarkLinearSearchTranspositionWithSameValueMultipleTimes(b *testing.B) {
	benchmarkLSTarr := rand.Perm(10000)
	rand.Seed(time.Now().Unix())
	search := rand.Intn(10000)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		LinearSearchTransposition(benchmarkLSTarr, search)
	}
}

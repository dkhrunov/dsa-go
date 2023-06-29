package searching

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

func BenchmarkBinarySearch(b *testing.B) {
	benchmarkBSarr := rand.Perm(10000)
	rand.Seed(time.Now().Unix())
	sort.Ints(benchmarkBSarr)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		search := rand.Intn(10000)
		BinarySearch(benchmarkBSarr, search)
	}
}

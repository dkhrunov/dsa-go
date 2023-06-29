package sorting

import (
	"math/rand"
	"runtime"
	"testing"
	"time"

	"github.com/dkhrunov/dsa-go/utils"
)

func BenchmarkMergeSort(b *testing.B) {
	runtime.GOMAXPROCS(4)
	rand.Seed(time.Now().Unix())
	lengths := []int{10000}
	for _, l := range lengths {
		b.Run("MergeSort", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				array := rand.Perm(l)
				b.StartTimer()

				MergeSort(array, utils.LessComparator[int])
			}
		})

		b.Run("MergeSortParallel", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				array := rand.Perm(l)
				b.StartTimer()

				MergeSortParallel(array, utils.LessComparator[int])
			}
		})

		b.Run("MergeSortImmutable", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				array := rand.Perm(l)
				b.StartTimer()

				MergeSortImmutable(array, utils.LessComparator[int])
			}
		})

		b.Run("MergeSortImmutableParallel", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				array := rand.Perm(l)
				b.StartTimer()

				MergeSortImmutableParallel(array, utils.LessComparator[int])
			}
		})
	}
}

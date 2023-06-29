package sorting

import (
	"fmt"
	"math/rand"
	"runtime"
	"testing"
	"time"

	"github.com/dkhrunov/dsa-go/utils"
)

func BenchmarkQuickSort(b *testing.B) {
	runtime.GOMAXPROCS(4)
	rand.Seed(time.Now().Unix())
	lengths := []int{5000}
	for _, l := range lengths {
		b.Run(fmt.Sprintf("QuickSort-%d", l), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				array := rand.Perm(l)
				b.StartTimer()

				QuickSort(array, utils.LessComparator[int])
			}
		})

		b.Run(fmt.Sprintf("QuickSortParallel-%d", l), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				array := rand.Perm(l)
				b.StartTimer()

				QuickSortParallel(array, utils.LessComparator[int])
			}
		})
	}
}

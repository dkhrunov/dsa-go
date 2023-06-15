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
	lengths := []int{100, 500, 1000, 1500, 2000, 2500, 3000, 3500, 4000, 4500, 5000}
	for _, l := range lengths {
		array := rand.Perm(l)

		b.Run(fmt.Sprintf("QuickSort-%d", l), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				QuickSort(array, utils.LessComparator[int])
			}
		})

		b.Run(fmt.Sprintf("QuickSortParallel-%d", l), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				QuickSortParallel(array, utils.LessComparator[int])
			}
		})
	}
}
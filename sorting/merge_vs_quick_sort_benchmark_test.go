package sorting

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/dkhrunov/dsa-go/utils"
)

func BenchmarkMergeVsQuickSort(b *testing.B) {
	rand.Seed(time.Now().Unix())
	lengths := []int{10000}
	for _, l := range lengths {
		b.Run(fmt.Sprintf("MergeSort-%d", l), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				array := rand.Perm(l)
				b.StartTimer()

				MergeSort(array, utils.LessComparator[int])
			}
		})

		b.Run(fmt.Sprintf("QuickSort-%d", l), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				b.StopTimer()
				array := rand.Perm(l)
				b.StartTimer()

				QuickSort(array, utils.LessComparator[int])
			}
		})
	}
}

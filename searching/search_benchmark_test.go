package searching

import (
	"fmt"
	"math/rand"
	"sort"
	"testing"
	"time"
)

func BenchmarkSearch(b *testing.B) {
	rand.Seed(time.Now().Unix())
	lengths := []int{100, 500, 1000, 1500, 2000, 2500, 3000, 3500, 4000, 4500, 5000, 10000, 15000, 20000, 25000, 30000, 35000, 40000, 45000, 50000, 100000}
	for _, l := range lengths {
		array := rand.Perm(l)
		sort.Ints(array)
		find := rand.Intn(l)

		b.Run(fmt.Sprintf("LinearSearch-%d", l), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				LinearSearch(array, find)
			}
		})

		b.Run(fmt.Sprintf("LinearSearchRecursive-%d", l), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				LinearSearchRecursive(array, find)
			}
		})

		b.Run(fmt.Sprintf("LinearSearchTransposition-%d", l), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				LinearSearchTransposition(array, find)
			}
		})

		b.Run(fmt.Sprintf("LinearSearchMoveToFront-%d", l), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				LinearSearchMoveToFront(array, find)
			}
		})

		b.Run(fmt.Sprintf("BinarySearch-%d", l), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				BinarySearch(array, find)
			}
		})

		b.Run(fmt.Sprintf("JumpSearch-%d", l), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				JumpSearch(array, find)
			}
		})

		b.Run(fmt.Sprintf("ExponentialSearch-%d", l), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ExponentialSearch(array, find)
			}
		})
	}
}

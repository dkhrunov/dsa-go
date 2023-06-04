package utils

import "golang.org/x/exp/constraints"

// `1` - compare passed
//
// `-1` - compare not passed
//
// `0` - values are equal
type ComparatorFn[T constraints.Ordered] func(a, b T) int8

func LessComparator[T constraints.Ordered](a, b T) int8 {
	if a == b {
		return 0
	} else if a < b {
		return 1
	} else {
		return -1
	}
}

func GreaterComparator[T constraints.Ordered](a, b T) int8 {
	if a == b {
		return 0
	} else if a > b {
		return 1
	} else {
		return -1
	}
}

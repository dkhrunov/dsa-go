package utils

import "golang.org/x/exp/constraints"

type ComparatorFn[T constraints.Ordered] func(a, b T) bool

func LessComparator[T constraints.Ordered](a, b T) bool {
	return a < b
}

func GreaterComparator[T constraints.Ordered](a, b T) bool {
	return a > b
}

package gmath

import (
	"golang.org/x/exp/constraints"
)

func Min[T constraints.Integer | constraints.Float](a, b T) T {
	if a < b {
		return a
	}
	return b
}

package gmath

import (
	"golang.org/x/exp/constraints"
)

func Max[T constraints.Integer | constraints.Float](a, b T) T {
	if a > b {
		return a
	}
	return b
}

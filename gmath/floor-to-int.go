package gmath

import (
	"math"

	"golang.org/x/exp/constraints"
)

func FloorInt[T constraints.Integer | constraints.Float](val T) int {
	return int(math.Floor(float64(val)))
}

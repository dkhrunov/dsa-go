package math

import (
	"math"

	"golang.org/x/exp/constraints"
)

type numbers interface {
	constraints.Integer | constraints.Float
}

func FloorInt[T numbers](val T) int {
	return int(math.Floor(float64(val)))
}

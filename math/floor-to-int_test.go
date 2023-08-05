package math

import (
	"testing"

	"golang.org/x/exp/constraints"
)

func TestFloorInt(t *testing.T) {
	type args[T constraints.Integer | constraints.Float] struct {
		val T
	}
	type test[T constraints.Integer | constraints.Float] struct {
		name string
		args args[T]
		want int
	}
	tests := []test[float64]{
		{"Should convert floor64 to int type", args[float64]{5.59}, 5},
		{"Should convert floor64 to int type", args[float64]{5.999}, 5},
		{"Should convert floor64 to int type", args[float64]{3.0}, 3},
		{"Should convert floor64 to int type", args[float64]{3.0001}, 3},
		{"Should convert floor64 to int type", args[float64]{3.2}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FloorInt(tt.args.val); got != tt.want {
				t.Errorf("FloorInt(%v) = %v, want %v", tt.args.val, got, tt.want)
			}
		})
	}
}

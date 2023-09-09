package gmath

import (
	"fmt"
	"reflect"
	"testing"

	"golang.org/x/exp/constraints"
)

func TestMin(t *testing.T) {
	type args[T constraints.Ordered] struct {
		a T
		b T
	}
	type tests[T constraints.Ordered] struct {
		name string
		args args[T]
		want T
	}
	testsInts := []tests[int]{
		{"The first argument must be less than the second", args[int]{1, 99}, 1},
		{"The second argument must be less than the first", args[int]{99, 88}, 88},
		{"If the arguments are equal, then return any of them", args[int]{99, 99}, 99},
	}
	testsFloat := []tests[float64]{
		{"The first argument must be less than the second", args[float64]{10.01, 10.25}, 10.01},
		{"The second argument must be less than the first", args[float64]{99.99, 88.88}, 88.88},
		{"If the arguments are equal, then return any of them", args[float64]{99.99, 99.99}, 99.99},
	}
	for _, tt := range testsInts {
		t.Run(fmt.Sprintf("[Ints] %v", tt.name), func(t *testing.T) {
			if got := Min(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Min(%v, %v) = %v, want %v", tt.args.a, tt.args.b, got, tt.want)
			}
		})
	}

	for _, tt := range testsFloat {
		t.Run(fmt.Sprintf("[Floats] %v", tt.name), func(t *testing.T) {
			if got := Min(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Min(%v, %v) = %v, want %v", tt.args.a, tt.args.b, got, tt.want)
			}
		})
	}
}

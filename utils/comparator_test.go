package utils

import (
	"testing"
)

func TestLessComparator(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"2 must be less than 5", args{2, 5}, true},
		{"5 must not be less than 2", args{5, 2}, false},
		{"10 must not be less than 10", args{10, 10}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LessComparator(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("LessComparator(%v, %v) = %v, want %v", tt.args.a, tt.args.b, got, tt.want)
			}
		})
	}
}

func TestGreaterComparator(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"5 must  be greater than 2", args{5, 2}, true},
		{"2 must not be greater than 5", args{2, 5}, false},
		{"10 must not be greater than 10", args{10, 10}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GreaterComparator(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("GreaterComparator(%v, %v) = %v, want %v", tt.args.a, tt.args.b, got, tt.want)
			}
		})
	}
}

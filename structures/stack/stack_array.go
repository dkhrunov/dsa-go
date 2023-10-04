package stack

import (
	"github.com/dkhrunov/dsa-go/utils"
)

// TODO write test
type stackArray[T any] struct {
	array []T
	pIdx  int
}

func newStackArray[T any](cap int) *stackArray[T] {
	return &stackArray[T]{
		array: make([]T, 0, cap),
		pIdx:  -1,
	}
}

// Len get the number of items in the stack
func (s *stackArray[T]) Len() int {
	return s.pIdx + 1
}

// Peek views the top item on the stack
func (s *stackArray[T]) Peek() (T, error) {
	if s.Len() == 0 {
		return utils.Zero[T](), ErrStackEmpty
	}

	return s.array[s.pIdx], nil
}

// Pop pop the top item of the stack and return it
func (s *stackArray[T]) Pop() (T, error) {
	if s.Len() == 0 {
		return utils.Zero[T](), ErrStackEmpty
	}

	val := s.array[s.pIdx]
	s.pIdx--

	return val, nil
}

// Push push a value onto the top of the stack
func (s *stackArray[T]) Push(val T) error {
	if s.pIdx+1 > cap(s.array) {
		return ErrStackOverflow
	}

	s.pIdx++
	s.array[s.pIdx] = val

	return nil
}

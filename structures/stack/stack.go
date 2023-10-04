package stack

import (
	"errors"
)

var (
	ErrStackEmpty    = errors.New("stack empty")
	ErrStackOverflow = errors.New("stack overflow")
)

type StackType[T any] interface {
	Peek() (T, error)
	Pop() (T, error)
	Push(val T) error
	Len() int
}

// TODO write test
type Stack[T any] struct {
	t StackType[T]
}

func New[T any]() *Stack[T] {
	return &Stack[T]{newStackList[T]()}
}

func NewStatic[T any](size int) *Stack[T] {
	return &Stack[T]{newStackArray[T](size)}
}

func (s *Stack[T]) Len() int {
	return s.t.Len()
}

func (s *Stack[T]) Peek() (T, error) {
	return s.t.Peek()
}

func (s *Stack[T]) Pop() (T, error) {
	return s.t.Pop()
}

func (s *Stack[T]) Push(val T) error {
	return s.t.Push(val)
}

package stack

import "github.com/dkhrunov/dsa-go/utils"

// TODO write test
type stackList[T any] struct {
	top *stackNode[T]
	len int
}

type stackNode[T any] struct {
	val  T
	prev *stackNode[T]
}

func newStackList[T any]() *stackList[T] {
	return &stackList[T]{
		top: nil,
		len: 0,
	}
}

// Len get the number of items in the stack
func (s *stackList[T]) Len() int {
	return s.len
}

// Peek views the top item on the stack
func (s *stackList[T]) Peek() (T, error) {
	if s.len == 0 {
		return utils.Zero[T](), ErrStackEmpty
	}
	return s.top.val, nil
}

// Pop pop the top item of the stack and return it
func (s *stackList[T]) Pop() (T, error) {
	if s.len == 0 {
		return utils.Zero[T](), ErrStackEmpty
	}

	t := s.top
	s.top = t.prev
	s.len--

	return t.val, nil
}

// Push push a value onto the top of the stack
func (s *stackList[T]) Push(val T) error {
	n := &stackNode[T]{
		val:  val,
		prev: s.top,
	}

	s.top = n
	s.len++

	return nil
}

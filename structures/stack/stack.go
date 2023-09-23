package stack

import (
	"errors"

	"github.com/dkhrunov/dsa-go/utils"
)

// TODO write tests
type (
	Stack[T any] struct {
		top    *node[T]
		length int
	}
	node[T any] struct {
		value T
		prev  *node[T]
	}
)

// NewStack create a new stack
func NewStack[T any]() *Stack[T] {
	return &Stack[T]{nil, 0}
}

// Len get the number of items in the stack
func (this *Stack[T]) Len() int {
	return this.length
}

// Peek views the top item on the stack
func (this *Stack[T]) Peek() (T, error) {
	if this.length == 0 {
		return utils.Zero[T](), errors.New("Stack is empty")
	}
	return this.top.value, nil
}

// Pop pop the top item of the stack and return it
func (this *Stack[T]) Pop() (T, error) {
	if this.length == 0 {
		return utils.Zero[T](), errors.New("Stack is empty")
	}

	n := this.top
	this.top = n.prev
	this.length--
	return n.value, nil
}

// Push push a value onto the top of the stack
func (this *Stack[T]) Push(value T) {
	n := &node[T]{value, this.top}
	this.top = n
	this.length++
}

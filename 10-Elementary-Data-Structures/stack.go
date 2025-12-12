package main

import (
	"errors"
)

// Stack[T] represents a fixed-capacity generic stack (LIFO) for elements of any type T.
type Stack[T any] struct {
	elements []T // underlying slice to store stack elements
	capacity int // maximum number of elements
	top      int // index of the top element (-1 if empty)
}

// NewStack creates a new stack with the specified capacity.
func NewStack[T any](size int) *Stack[T] {
	return &Stack[T]{
		elements: make([]T, size), // preallocate slice with capacity
		capacity: size,
		top:      -1, // stack is initially empty
	}
}

// Push adds an element to the top of the stack.
// Returns an error if the stack is already full (overflow).
func (s *Stack[T]) Push(item T) error {
	if s.IsFull() {
		return errors.New("stack overflow error")
	}
	s.top++
	s.elements[s.top] = item
	return nil
}

// Pop removes and returns the top element from the stack.
// Returns an error if the stack is empty (underflow).
func (s *Stack[T]) Pop() (T, error) {
	var zero T
	if s.IsEmpty() {
		return zero, errors.New("stack underflow error")
	}
	val := s.elements[s.top]
	s.top--
	return val, nil
}

// Peek returns the top element without removing it.
// Returns an error if the stack is empty.
func (s *Stack[T]) Peek() (T, error) {
	var zero T
	if s.IsEmpty() {
		return zero, errors.New("stack underflow error")
	}
	return s.elements[s.top], nil
}

// IsEmpty checks if the stack contains no elements.
func (s *Stack[T]) IsEmpty() bool {
	return s.top == -1
}

// isFull checks if the stack has reached its maximum capacity.
func (s *Stack[T]) IsFull() bool {
	return s.top == s.capacity-1
}

// Length returns the current number of elements in the stack.
func (s *Stack[T]) Length() int {
	return s.top + 1
}

// Capacity returns the maximum number of elements the stack can hold.
func (s *Stack[T]) Capacity() int {
	return s.capacity
}

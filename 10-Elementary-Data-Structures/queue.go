package main

import "errors"

/*
Queue[T] implements a generic fixed-capacity FIFO (First-In-First-Out) queue.

The queue is implemented using a circular array to achieve O(1) time
complexity for both enqueue and dequeue operations.
*/
type Queue[T any] struct {
	elements []T // underlying storage
	capacity int // maximum number of elements
	length   int // current number of elements
	front    int // index of the front element
	rear     int // index for the next insertion
}

// NewQueue creates a new queue with the given fixed capacity.
func NewQueue[T any](size int) *Queue[T] {
	return &Queue[T]{
		elements: make([]T, size),
		capacity: size,
		length:   0,
		front:    0,
		rear:     0,
	}
}

// Enqueue inserts an element at the rear of the queue.
// Returns an error if the queue is full (overflow).
func (q *Queue[T]) Enqueue(item T) error {
	if q.IsFull() {
		return errors.New("queue overflow")
	}
	q.elements[q.rear] = item
	q.rear = (q.rear + 1) % q.capacity // circular increment
	q.length++
	return nil
}

// Dequeue removes and returns the element at the front of the queue.
// Returns an error if the queue is empty (underflow).
func (q *Queue[T]) Dequeue() (T, error) {
	var zero T
	if q.IsEmpty() {
		return zero, errors.New("queue underflow")
	}
	item := q.elements[q.front]
	q.front = (q.front + 1) % q.capacity // circular increment
	q.length--
	return item, nil
}

// Length returns the current number of elements in the queue.
func (q *Queue[T]) Length() int {
	return q.length
}

// Capacity returns the maximum number of elements the queue can hold.
func (q *Queue[T]) Capacity() int {
	return q.capacity
}

// IsFull reports whether the queue is full.
func (q *Queue[T]) IsFull() bool {
	return q.length == q.capacity
}

// IsEmpty reports whether the queue is empty.
func (q *Queue[T]) IsEmpty() bool {
	return q.length == 0
}

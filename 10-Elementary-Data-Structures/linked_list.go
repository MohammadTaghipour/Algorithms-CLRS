package main

// Node represents a single element in a doubly linked list.
// It holds a value of type T and pointers to the next and previous nodes.
type Node[T comparable] struct {
	Value T        // Value stored in the node
	Next  *Node[T] // Pointer to the next node in the list
	Prev  *Node[T] // Pointer to the previous node in the list
}

// LinkedList represents a doubly linked list.
// The zero value of LinkedList is a valid empty list.
type LinkedList[T comparable] struct {
	Head *Node[T] // First node of the list, or nil if the list is empty
	Tail *Node[T] // Last node of the list, or nil if the list is empty
}

// NewLinkedList creates and returns an empty doubly linked list.
func NewLinkedList[T comparable]() *LinkedList[T] {
	return &LinkedList[T]{
		Head: nil,
		Tail: nil,
	}
}

// FindFirst searches for the first node whose value equals the given value.
// It returns the node and true if found, otherwise nil and false.
func (l *LinkedList[T]) FindFirst(value T) (*Node[T], bool) {
	x := l.Head
	for x.Next != nil {
		if x.Value == value {
			return x, true
		}
		x = x.Next
	}
	return nil, false
}

// InsertFirst inserts a new value at the beginning of the list.
// The new node becomes the head of the list.
func (l *LinkedList[T]) InsertFirst(value T) {
	x := &Node[T]{
		Value: value,
		Next:  l.Head,
		Prev:  nil,
	}
	if l.Head != nil {
		l.Head.Prev = x
	}
	l.Head = x
}

// InsertLast inserts a new value at the end of the list.
// The new node becomes the tail of the list.
func (l *LinkedList[T]) InsertLast(value T) {
	x := &Node[T]{
		Value: value,
		Next:  nil,
		Prev:  l.Tail,
	}
	if l.Tail != nil {
		l.Tail.Next = x
	}
	l.Tail = x
}

// InsertAfter inserts a new value immediately after the given node.
// If the given node is the tail, the new node becomes the new tail.
// If node is nil, the function does nothing.
func (l *LinkedList[T]) InsertAfter(node *Node[T], value T) {
	if node == nil {
		return
	}

	newNode := &Node[T]{
		Value: value,
		Prev:  node,
		Next:  node.Next,
	}

	if node.Next != nil {
		node.Next.Prev = newNode
	} else {
		l.Tail = newNode
	}

	node.Next = newNode
}

// Delete removes the given node from the list.
// If the node is the head or tail, the corresponding pointer is updated.
// If node is nil, the function does nothing.
func (l *LinkedList[T]) Delete(node *Node[T]) {
	if node == nil {
		return
	}

	if node.Prev != nil {
		node.Prev.Next = node.Next
	} else {
		l.Head = node.Next
	}

	if node.Next != nil {
		node.Next.Prev = node.Prev
	} else {
		l.Tail = node.Prev
	}

	node.Prev = nil
	node.Next = nil
}

// DeleteFirst removes the first element (head) of the list.
// If the list is empty, it does nothing.
func (l *LinkedList[T]) DeleteFirst() {
	if l.Head == nil {
		return
	}

	if l.Head.Next == nil {
		l.Head = nil
		l.Tail = nil
		return
	}

	oldHead := l.Head
	l.Head = oldHead.Next
	l.Head.Prev = nil

	oldHead.Next = nil
	oldHead.Prev = nil
}

// DeleteLast removes the last element (tail) of the list.
// If the list is empty, it does nothing.
func (l *LinkedList[T]) DeleteLast() {
	if l.Tail == nil {
		return
	}

	if l.Tail.Prev == nil {
		l.Head = nil
		l.Tail = nil
		return
	}

	oldTail := l.Tail
	l.Tail = oldTail.Prev
	l.Tail.Next = nil

	oldTail.Next = nil
	oldTail.Prev = nil
}

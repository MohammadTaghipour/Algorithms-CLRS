package main

// Heap represents a binary heap data structure that can act as either
// a max-heap or a min-heap, depending on the comparison function.
//
// Fields:
// - data: the underlying slice storing the heap elements
// - heapSize: the number of elements currently in the heap
// - betterThan: a comparison function that defines the heap property
//   (e.g., a > b for max-heap, a < b for min-heap)
type Heap struct {
	data       []int
	heapSize   int
	betterThan func(a, b int) bool
}

// parent returns the index of the parent of the node at index i.
//
// For a 0-based array, parent(i) = (i-1)/2
// in CLRS parent(i) = floor(i/2)
func parent(i int) int {
	return (i - 1) / 2
}

// left returns the index of the left child of the node at index i.
//
// For a 0-based array, left(i) = 2*i + 1
// in CLRS left(i) = 2*i
func left(i int) int {
	return 2*i + 1
}

// right returns the index of the right child of the node at index i.
//
// For a 0-based array, right(i) = 2*i + 2
// in CLRS right(i) = 2*i + 1
func right(i int) int {
	return 2*i + 2
}

// heapify restores the heap property for the subtree rooted at index i.
//
// It assumes that the binary trees rooted at left(i) and right(i)
// already satisfy the heap property. If the element at i violates
// the heap property, it "floats down" to the correct position by
// recursively swapping with the better child.
//
// Time complexity: O(log n) where n is the size of the heap
func (h *Heap) heapify(i int) {
	l := left(i)
	r := right(i)

	best := i

	if l < h.heapSize && h.betterThan(h.data[l], h.data[best]) {
		best = l
	}
	if r < h.heapSize && h.betterThan(h.data[r], h.data[best]) {
		best = r
	}

	if best != i {
		h.data[i], h.data[best] = h.data[best], h.data[i]
		h.heapify(best)
	}
}

// BuildMaxHeap constructs a max-heap from the given slice of integers.
//
// It sets up the heap structure and calls heapify in a bottom-up manner
// to ensure the max-heap property.
//
// Time complexity: O(n)
func BuildMaxHeap(arr []int) *Heap {
	h := &Heap{
		data:     arr,
		heapSize: len(arr),
		betterThan: func(a, b int) bool {
			return a > b
		},
	}

	for i := h.heapSize/2 - 1; i >= 0; i-- {
		h.heapify(i)
	}

	return h
}

// BuildMinHeap constructs a min-heap from the given slice of integers.
//
// It sets up the heap structure and calls heapify in a bottom-up manner
// to ensure the min-heap property.
//
// Time complexity: O(n)
func BuildMinHeap(arr []int) *Heap {
	h := &Heap{
		data:     arr,
		heapSize: len(arr),
		betterThan: func(a, b int) bool {
			return a < b
		},
	}

	for i := h.heapSize/2 - 1; i >= 0; i-- {
		h.heapify(i)
	}
	return h
}

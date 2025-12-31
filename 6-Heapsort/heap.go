package main

import (
	"errors"
)

// Heap represents a binary heap data structure that can act as either
// a max-heap or a min-heap, depending on the comparison function.
//
// Fields:
//   - data: the underlying slice storing the heap elements
//   - heapSize: the number of elements currently in the heap
//   - betterThan: a comparison function that defines the heap property
//     (e.g., a > b for max-heap, a < b for min-heap)
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

// Peek returns the root element of the heap without removing it.
//
// For a max-heap, this returns the maximum element.
// For a min-heap, this returns the minimum element.
//
// This operation does NOT modify the heap structure.
//
// Time complexity: O(1)
// Space complexity: O(1)
func (h *Heap) Peek() (error, int) {
	if h.heapSize == 0 {
		return errors.New("heap underflow error"), 0
	}
	return nil, h.data[0]
}

// ExtractRoot removes and returns the root element of the heap.
//
// In a max-heap, this corresponds to EXTRACT-MAX.
// In a min-heap, this corresponds to EXTRACT-MIN.
//
// Algorithm (CLRS-style):
// 1. Save the root value.
// 2. Move the last element of the heap to the root position.
// 3. Decrease heap size.
// 4. Call heapify(0) to restore the heap property.
//
// Time complexity: O(log n)
// Space complexity: O(1)
func (h *Heap) ExtractRoot() (error, int) {
	if h.heapSize == 0 {
		return errors.New("heap underflow error"), 0
	}
	root := h.data[0]
	h.data[0] = h.data[h.heapSize-1]
	h.heapSize--
	h.heapify(0)
	return nil, root
}

// Insert adds a new key into the heap.
//
// The key is first appended at the end of the array to maintain
// the complete binary tree structure. Then, a "bubble-up"
// (heap-increase-key / heap-decrease-key behavior) is performed
// to restore the heap property.
//
// This operation works for both max-heaps and min-heaps,
// depending on the betterThan comparison function.
//
// Time complexity: O(log n)
// Space complexity: O(1) amortized
func (h *Heap) Insert(key int) {
	h.data = append(h.data, key)
	h.heapSize++
	i := h.heapSize - 1

	// Bubble up
	for i > 0 && h.betterThan(h.data[i], h.data[parent(i)]) {
		h.data[i], h.data[parent(i)] = h.data[parent(i)], h.data[i]
		i = parent(i)
	}
}

// UpdateKey updates the value of the key at index i and restores
// the heap property.
//
// If the new key has higher priority than the old key
// (greater in a max-heap, smaller in a min-heap),
// the element is bubbled up toward the root.
//
// Otherwise, the element may violate the heap property with
// its children, so heapify is called to bubble it down.
//
// This method corresponds to:
// - INCREASE-KEY in max-priority queues
// - DECREASE-KEY in min-priority queues
//
// Time complexity: O(log n)
// Space complexity: O(1)
func (h *Heap) UpdateKey(i, newKey int) {
	oldKey := h.data[i]
	h.data[i] = newKey

	if h.betterThan(newKey, oldKey) {
		// Bubble up
		for i > 0 && h.betterThan(h.data[i], h.data[parent(i)]) {
			h.data[i], h.data[parent(i)] = h.data[parent(i)], h.data[i]
			i = parent(i)
		}
	} else {
		// Bubble down
		h.heapify(i)
	}
}

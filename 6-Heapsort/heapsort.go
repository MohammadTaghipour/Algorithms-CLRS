package main

// HeapSort sorts the given slice of integers in ascending order
// using the heapsort algorithm.
//
// This algorithm works in two main phases:
// 1. Build a max-heap from the input array so that the largest element
//    is at the root (index 0).
// 2. Repeatedly swap the root of the heap with the last element of the
//    heap, reduce the heap size by one, and restore the max-heap property
//    by calling heapify on the root.
//
// Time complexity (worst case): O(n log n)
// Space complexity: O(1) – in-place sorting
func HeapSort(arr []int) {
	h := BuildMaxHeap(arr)
	for i := h.heapSize - 1; i >= 1; i-- {
		// Move current maximum to its final position
		h.data[0], h.data[i] = h.data[i], h.data[0]

		// Reduce heap size to exclude the sorted element
		h.heapSize--

		// Restore the max-heap property
		h.heapify(0)
	}
}

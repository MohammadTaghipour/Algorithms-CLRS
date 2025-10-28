package main

// InsertionSort sorts the given slice of integers in ascending order
// using the insertion sort algorithm.
//
// The algorithm works by building a sorted portion of the slice
// one element at a time. For each element in A[1:], it finds the
// correct position among the already-sorted elements (to its left)
// and inserts it there.
//
// Time complexity: O(n²)
// Space complexity: O(1)
func InsertionSort(A []int) {
	for i := 1; i < len(A); i++ {
		key := A[i]
		j := i - 1

		// Shift elements of A[0..i-1] that are greater than key
		// one position to the right to make space for insertion.
		for j >= 0 && A[j] > key {
			A[j+1] = A[j]
			j--
		}

		// Place key in its correct sorted position.
		A[j+1] = key
	}
}

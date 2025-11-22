package main

// Quicksort sorts the given slice of integers in ascending order
// using the quicksort algorithm.
//
// This algorithm follows the divide-and-conquer approach:
// 1. Partition the slice around a pivot so that elements <= pivot
//    go to the left side and elements > pivot go to the right.
// 2. Recursively sort the left and right sub-slices.
//
// Time complexity (average): O(n log n)
// Time complexity (worst case): O(n²)     // occurs when the pivot is poor
// Space complexity: O(log n) recursion stack
func Quicksort(A []int, p, r int) {
	if p < r {
		q := Partition(A, p, r)
		Quicksort(A, p, q-1)
		Quicksort(A, q+1, r)
	}
}

// Partition rearranges the elements in A[p..r] around a pivot,
// which is chosen as A[r] in this implementation (Lomuto scheme).
//
// All elements <= pivot are moved to the left side,
// and all elements > pivot are moved to the right side.
//
// It returns the final index of the pivot after partitioning,
// which is used by Quicksort to split the array.
//
// Time complexity: O(n)
// Space complexity: O(1)
func Partition(A []int, p, r int) int {
	x := A[r]  // pivot element
	i := p - 1 // boundary of the <= pivot side

	for j := p; j < r; j++ {
		// Move elements <= pivot to the left side
		if A[j] <= x {
			i++
			A[i], A[j] = A[j], A[i]
		}
	}

	// Place the pivot in its correct sorted position
	A[i+1], A[r] = A[r], A[i+1]
	return i + 1
}

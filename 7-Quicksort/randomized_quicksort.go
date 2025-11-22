package main

import "math/rand"

// RandomizedQuicksort sorts the given slice of integers in ascending order
// using the randomized quicksort algorithm.
//
// This algorithm follows the divide-and-conquer approach with a randomly chosen pivot:
// 1. Randomly select a pivot from the sub-slice A[p..r].
// 2. Partition the slice around the pivot so that elements <= pivot
//    go to the left side and elements > pivot go to the right.
// 3. Recursively sort the left and right sub-slices.
//
// Time complexity (average): O(n log n)
// Time complexity (worst case): O(n²)     // highly unlikely due to random pivot
// Space complexity: O(log n) recursion stack
func RandomizedQuicksort(A []int, p, r int) {
	if p < r {
		q := RandomizedPartition(A, p, r)
		RandomizedQuicksort(A, p, q-1)
		RandomizedQuicksort(A, q+1, r)
	}
}

// RandomizedPartition selects a random pivot from A[p..r], swaps it with the last element,
// and then partitions the slice using the standard Lomuto partition scheme.
//
// Returns the final index of the pivot after partitioning.
func RandomizedPartition(A []int, p, r int) int {
	// Pick a random index between p and r (inclusive)
	i := rand.Intn(r-p+1) + p

	// Swap the random pivot with the last element
	A[r], A[i] = A[i], A[r]

	// Partition the slice around the pivot
	return Partition(A, p, r)
}

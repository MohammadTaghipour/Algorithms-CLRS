package main

// SumArray returns the sum of all elements in the given slice of integers.
//
// It iterates through the slice A and accumulates the total value.
// Time complexity: O(n)
// Space complexity: O(1)
func SumArray(A []int) int {
	sum := 0
	for i := 0; i < len(A); i++ {
		sum += A[i]
	}
	return sum
}

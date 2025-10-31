package main

// MergeSort returns a sorted copy of the given slice using the Merge Sort algorithm.
//
// It recursively splits the input slice into halves, sorts each half, and merges them.
// Time complexity: O(n log n)
// Space complexity: O(n)
func MergeSort(A []int) []int {
	// Base case: if array has 0 or 1 element, it's already sorted
	if len(A) <= 1 {
		return A
	}

	middle := len(A) / 2
	left := MergeSort(A[:middle])  // Recursively sort left half
	right := MergeSort(A[middle:]) // Recursively sort right half

	// Merge the sorted halves and return
	return Merge(left, right)
}

// Merge combines two sorted slices into a single sorted slice.
//
// It compares elements from both slices and appends the smaller one to the result.
// Time complexity: O(n) where n = len(Left) + len(Right)
// Space complexity: O(n)
func Merge(Left, Right []int) []int {
	result := []int{}
	i, j := 0, 0

	// Merge elements while both arrays have remaining items
	for i < len(Left) && j < len(Right) {
		if Left[i] <= Right[j] {
			result = append(result, Left[i])
			i++
		} else {
			result = append(result, Right[j])
			j++
		}
	}

	// Append any remaining elements from Left (if any)
	result = append(result, Left[i:]...)
	// Append any remaining elements from Right (if any)
	result = append(result, Right[j:]...)

	return result
}

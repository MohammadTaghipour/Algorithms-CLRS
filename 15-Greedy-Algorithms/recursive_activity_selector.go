package main

// RecursiveActivitySelector selects the maximum set of mutually compatible activities
// from a given list of activities with start times s and finish times f.
//
// The function implements the classic recursive greedy algorithm for the activity selection problem.
//
// Important precondition:
//   The activities must be sorted in **increasing order of finish times (f[i])** before calling this function.
//   The arrays s and f are 1-indexed (element 0 is a dummy element and must be 0).
//   - s[i] is the start time of activity i
//   - f[i] is the finish time of activity i
//   The first element (s[0] and f[0]) is a dummy and must be 0 to match the recursive algorithm logic.
//
// Parameters:
//   s []int : array of start times (1-indexed, s[0] = 0)
//   f []int : array of finish times (1-indexed, f[0] = 0)
//   k int   : index of the last selected activity
//   n int   : total number of activities
//
// Returns:
//   []int : slice of indices of the selected activities in the optimal solution
//
// Algorithm:
//   1. Find the first activity m > k that starts after or at the finish of activity k.
//   2. If such an activity exists, include it and recursively solve for remaining activities.
//   3. If no compatible activity is left, return an empty slice.
//
// Time complexity: O(n) if activities are already sorted
//                  O(n log n) if sorting is included before calling this function
// Space complexity: O(n) due to recursion stack
func RecursiveActivitySelector(s, f []int, k, n int) []int {
	m := k + 1

	for m <= n && s[m] < f[k] {
		m = m + 1
	}

	if m <= n {
		result := RecursiveActivitySelector(s, f, m, n)
		return append([]int{m}, result...)
	}

	return []int{}
}

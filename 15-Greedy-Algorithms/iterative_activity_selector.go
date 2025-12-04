package main

// IterativeActivitySelector selects the maximum set of mutually compatible activities
// from a given list of activities with start times s and finish times f using an iterative greedy approach.
//
// Important precondition:
//   - The activities must be sorted in **increasing order of finish times (f[i])** before calling this function.
//   - The arrays s and f are 1-indexed (element 0 is not used, so s[1] and f[1] are the first real activity).
//
// Parameters:
//   s []int : array of start times (1-indexed)
//   f []int : array of finish times (1-indexed)
//   n int   : total number of activities
//
// Returns:
//   []int : slice of indices of the selected activities in the optimal solution
//
// Algorithm:
//   1. Select the first activity (index 1) and add it to the solution set A.
//   2. For each subsequent activity m from 2 to n:
//        - If activity m starts after or at the finish time of the last selected activity k:
//            - Include activity m in A
//            - Update k to m
//   3. Return the selected set A.
//
// Time complexity: O(n) — each activity is considered exactly once
// Space complexity: O(n) — for storing the selected activity indices
func IterativeActivitySelector(s, f []int, n int) []int {
	k := 1
	A := []int{k}

	for m := 2; m <= n; m++ {
		if s[m] >= f[k] {
			A = append(A, m)
			k = m
		}
	}
	return A
}

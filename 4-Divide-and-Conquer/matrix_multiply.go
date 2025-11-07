package main

// MatrixMultiply performs matrix multiplication of two n×n integer matrices A and B,
// storing the result in matrix C.
//
// It uses a straightforward triple-nested loop to compute each element of C as
// the dot product of the corresponding row of A and column of B.
//
// Time complexity: O(n³)
// Space complexity: O(1) — ignoring input and output storage
func MatrixMultiply(A, B, C [][]int, n int) {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				C[i][j] += A[i][k] * B[k][j]
			}
		}
	}
}

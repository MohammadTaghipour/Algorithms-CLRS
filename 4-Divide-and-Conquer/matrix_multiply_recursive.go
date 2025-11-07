package main

// RecursiveMatrixMultiply multiplies two n×n integer matrices A and B
// using the classical divide-and-conquer (recursive) algorithm.
//
// The function assumes n is a power of two. It divides A and B into four
// n/2×n/2 submatrices and recursively computes each quadrant of the
// result matrix C as follows:
//
//	C11 = A11×B11 + A12×B21
//	C12 = A11×B12 + A12×B22
//	C21 = A21×B11 + A22×B21
//	C22 = A21×B12 + A22×B22
//
// Time complexity: Θ(n³)
// Space complexity: Θ(n²)
func RecursiveMatrixMultiply(A, B [][]int) [][]int {
	n := len(A)
	C := MakeMatrix(n)

	// Base case: single element
	if n == 1 {
		C[0][0] = A[0][0] * B[0][0]
		return C
	}

	// Divide matrices into four submatrices each
	A11, A12, A21, A22 := SplitMatrix(A)
	B11, B12, B21, B22 := SplitMatrix(B)

	// Recursive calls for submatrix multiplication and addition
	C11 := AddMatrix(RecursiveMatrixMultiply(A11, B11), RecursiveMatrixMultiply(A12, B21))
	C12 := AddMatrix(RecursiveMatrixMultiply(A11, B12), RecursiveMatrixMultiply(A12, B22))
	C21 := AddMatrix(RecursiveMatrixMultiply(A21, B11), RecursiveMatrixMultiply(A22, B21))
	C22 := AddMatrix(RecursiveMatrixMultiply(A21, B12), RecursiveMatrixMultiply(A22, B22))

	// Combine submatrices into the final result
	return CombineMatrix(C11, C12, C21, C22)
}

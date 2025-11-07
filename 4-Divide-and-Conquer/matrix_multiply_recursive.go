package main

func makeMatrix(n int) [][]int {
	m := make([][]int, n)
	for i := range m {
		m[i] = make([]int, n)
	}
	return m
}

func add(A, B [][]int) [][]int {
	n := len(A)
	C := makeMatrix(n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			C[i][j] = A[i][j] + B[i][j]
		}
	}
	return C
}

func split(A [][]int) ([][]int, [][]int, [][]int, [][]int) {
	n := len(A)
	k := n / 2
	A11 := makeMatrix(k)
	A12 := makeMatrix(k)
	A21 := makeMatrix(k)
	A22 := makeMatrix(k)
	for i := 0; i < k; i++ {
		for j := 0; j < k; j++ {
			A11[i][j] = A[i][j]
			A12[i][j] = A[i][j+k]
			A21[i][j] = A[i+k][j]
			A22[i][j] = A[i+k][j+k]
		}
	}
	return A11, A12, A21, A22
}

func combine(C11, C12, C21, C22 [][]int) [][]int {
	k := len(C11)
	n := k * 2
	C := makeMatrix(n)
	for i := 0; i < k; i++ {
		for j := 0; j < k; j++ {
			C[i][j] = C11[i][j]
			C[i][j+k] = C12[i][j]
			C[i+k][j] = C21[i][j]
			C[i+k][j+k] = C22[i][j]
		}
	}
	return C
}

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
	C := makeMatrix(n)

	// Base case: single element
	if n == 1 {
		C[0][0] = A[0][0] * B[0][0]
		return C
	}

	// Divide matrices into four submatrices each
	A11, A12, A21, A22 := split(A)
	B11, B12, B21, B22 := split(B)

	// Recursive calls for submatrix multiplication and addition
	C11 := add(RecursiveMatrixMultiply(A11, B11), RecursiveMatrixMultiply(A12, B21))
	C12 := add(RecursiveMatrixMultiply(A11, B12), RecursiveMatrixMultiply(A12, B22))
	C21 := add(RecursiveMatrixMultiply(A21, B11), RecursiveMatrixMultiply(A22, B21))
	C22 := add(RecursiveMatrixMultiply(A21, B12), RecursiveMatrixMultiply(A22, B22))

	// Combine submatrices into the final result
	return combine(C11, C12, C21, C22)
}

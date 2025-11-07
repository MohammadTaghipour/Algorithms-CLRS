package main

// StrassenMatrixMultiply multiplies two n×n integer matrices A and B
// using Strassen’s divide-and-conquer algorithm.
//
// This algorithm improves on the classical recursive method by reducing
// the number of recursive multiplications from 8 to 7, at the cost of a few
// extra additions and subtractions of submatrices. It assumes n is a power of two.
//
// The computation is based on the following intermediate matrices:
//
//	M1 = (A11 + A22) × (B11 + B22)
//	M2 = (A21 + A22) × B11
//	M3 = A11 × (B12 − B22)
//	M4 = A22 × (B21 − B11)
//	M5 = (A11 + A12) × B22
//	M6 = (A21 − A11) × (B11 + B12)
//	M7 = (A12 − A22) × (B21 + B22)
//
// Then, the resulting submatrices of C are computed as:
//
//	C11 = M1 + M4 − M5 + M7
//	C12 = M3 + M5
//	C21 = M2 + M4
//	C22 = M1 − M2 + M3 + M6
//
// Time complexity: Θ(n^2.81)
// Space complexity: Θ(n²)
func StrassenMatrixMultiply(A, B [][]int) [][]int {
	n := len(A)
	C := MakeMatrix(n)
	if n == 1 {
		C[0][0] = A[0][0] * B[0][0]
		return C
	}

	A11, A12, A21, A22 := SplitMatrix(A)
	B11, B12, B21, B22 := SplitMatrix(B)

	M1 := StrassenMatrixMultiply(AddMatrix(A11, A22), AddMatrix(B11, B22))
	M2 := StrassenMatrixMultiply(AddMatrix(A21, A22), B11)
	M3 := StrassenMatrixMultiply(A11, SubMatrix(B12, B22))
	M4 := StrassenMatrixMultiply(A22, SubMatrix(B21, B11))
	M5 := StrassenMatrixMultiply(AddMatrix(A11, A12), B22)
	M6 := StrassenMatrixMultiply(SubMatrix(A21, A11), AddMatrix(B11, B12))
	M7 := StrassenMatrixMultiply(SubMatrix(A12, A22), AddMatrix(B21, B22))

	C11 := AddMatrix(SubMatrix(AddMatrix(M1, M4), M5), M7)
	C12 := AddMatrix(M3, M5)
	C21 := AddMatrix(M2, M4)
	C22 := AddMatrix(AddMatrix(SubMatrix(M1, M2), M3), M6)

	C = CombineMatrix(C11, C12, C21, C22)
	return C
}

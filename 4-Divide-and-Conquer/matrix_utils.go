package main

func MakeMatrix(n int) [][]int {
	m := make([][]int, n)
	for i := range m {
		m[i] = make([]int, n)
	}
	return m
}

func AddMatrix(A, B [][]int) [][]int {
	n := len(A)
	C := MakeMatrix(n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			C[i][j] = A[i][j] + B[i][j]
		}
	}
	return C
}

func SubMatrix(A, B [][]int) [][]int {
	n := len(A)
	C := MakeMatrix(n)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			C[i][j] = A[i][j] - B[i][j]
		}
	}
	return C
}

func SplitMatrix(A [][]int) ([][]int, [][]int, [][]int, [][]int) {
	n := len(A)
	k := n / 2
	A11 := MakeMatrix(k)
	A12 := MakeMatrix(k)
	A21 := MakeMatrix(k)
	A22 := MakeMatrix(k)
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

func CombineMatrix(C11, C12, C21, C22 [][]int) [][]int {
	k := len(C11)
	n := k * 2
	C := MakeMatrix(n)
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

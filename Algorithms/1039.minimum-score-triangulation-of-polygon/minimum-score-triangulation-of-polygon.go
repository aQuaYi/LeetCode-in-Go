package problem1039

func minScoreTriangulation(A []int) int {
	A = move(A)
	return sum(A)
}

func sum(A []int) int {
	sum, n := 0, len(A)
	a := A[0]
	for i := 1; i+1 < n; i++ {
		sum += a * A[i] * A[i+1]
	}
	return sum
}

func move(A []int) []int {
	a, index := A[0], 0
	n := len(A)
	for i := 1; i < n; i++ {
		if a > A[i] {
			a, index = A[i], i
		}
	}
	return append(A, A[:index]...)[index:]
}

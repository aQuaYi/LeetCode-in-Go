package problem1039

func minScoreTriangulation(A []int) int {
	if len(A) < 3 {
		return 0
	}
	B := move(A)
	a := sum(A) + minScoreTriangulation(next(A))
	b := sum(B) + minScoreTriangulation(next(B))
	return min(a, b)
}

func sum(A []int) int {
	sum, n := 0, len(A)
	for i := 0; i+2 < n; i += 2 {
		sum += A[i] * A[i+1] * A[i+2]
	}
	if n%3 != 2 && n != 3 {
		sum += A[n-2] * A[n-1] * A[0]
	}
	return sum
}

func next(A []int) []int {
	n := len(A)
	res := make([]int, 0, (n+1)/2)
	for i := 0; i < n; i += 2 {
		res = append(res, A[i])
	}
	return res
}

func move(A []int) []int {
	return append(A, A[0])[1:]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

package problem1007

func minDominoRotations(A []int, B []int) int {
	if res, ok := check(A, B, A[0]); ok {
		return res
	}
	if res, ok := check(A, B, B[0]); ok {
		return res
	}
	return -1
}

func check(A, B []int, flag int) (int, bool) {
	a, b := 0, 0
	i, n := 0, len(A)
	for ; i < n && (A[i] == flag || B[i] == flag); i++ {
		if A[i] != flag { // A needs a swap to be all flag
			a++
		}
		if B[i] != flag { // B needs a swap to be all flag
			b++
		}
	}
	return min(a, b), i == n
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

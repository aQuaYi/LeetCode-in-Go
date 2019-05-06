package problem0977

func sortedSquares(A []int) []int {
	size := len(A)
	res := make([]int, size)
	for l, r, i := 0, size-1, size-1; l <= r; i-- {
		if A[l]+A[r] < 0 {
			res[i] = A[l] * A[l]
			l++
		} else {
			res[i] = A[r] * A[r]
			r--
		}
	}
	return res
}

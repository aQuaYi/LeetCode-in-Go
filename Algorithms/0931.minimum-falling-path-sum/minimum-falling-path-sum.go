package problem0931

const maxSum = 10000

func minFallingPathSum(A [][]int) int {
	size := len(A)

	recs := make([]int, size)
	copy(recs, A[0])

	for i := 1; i < size; i++ {
		tmps := [100]int{}
		for j := 0; j < size; j++ {
			Aij := A[i][j]
			tmp := maxSum
			l, r := max(j-1, 0), min(j+1, size-1)
			for k := l; k <= r; k++ {
				tmp = min(tmp, Aij+recs[k])
			}
			tmps[j] = tmp
		}
		recs = tmps[:size]
	}

	res := maxSum
	for i := 0; i < size; i++ {
		res = min(res, recs[i])
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

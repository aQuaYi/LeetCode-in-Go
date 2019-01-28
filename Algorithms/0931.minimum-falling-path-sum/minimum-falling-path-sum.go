package problem0931

const maxSum = 10000

func minFallingPathSum(A [][]int) int {
	size := len(A)

	dp := make([]int, size)
	copy(dp, A[0])

	for i := 1; i < size; i++ {
		tmps := [100]int{}
		for j := 0; j < size; j++ {
			a, tmp := A[i][j], maxSum
			l, r := max(j-1, 0), min(j+1, size-1)
			for k := l; k <= r; k++ {
				tmp = min(tmp, a+dp[k])
			}
			tmps[j] = tmp
		}
		dp = tmps[:size]
	}

	res := maxSum
	for i := 0; i < size; i++ {
		res = min(res, dp[i])
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

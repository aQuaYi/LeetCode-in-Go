package problem1130

func mctFromLeafValues(A []int) int {
	n := len(A)

	maxs := [40][40]int{}
	var maxVal func(int, int) int
	maxVal = func(i, j int) int {
		if maxs[i][j] != 0 {
			return maxs[i][j]
		}
		if i == j {
			maxs[i][i] = A[i]
			return A[i]
		}
		res := max(A[i], maxVal(i+1, j))
		maxs[i][j] = res
		return res
	}

	for i := 0; i < n; i++ {
		maxs[i][i] = A[i]
	}
	dp := [40][40]int{}
	for width := 2; width <= n; width++ {
		for i := 0; i+width-1 < n; i++ {
			j := i + width - 1
			left := A[i]*maxVal(i+1, j) + dp[i+1][j]
			right := A[j]*maxVal(i, j-1) + dp[i][j-1]
			dp[i][j] = min(left, right)
			for k := i + 1; k < j-1; k++ {
				dp[i][j] = min(dp[i][j], dp[i][k]+dp[k+1][j]+maxVal(i, k)*maxVal(k+1, j))
			}
		}
	}

	return dp[0][n-1]
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

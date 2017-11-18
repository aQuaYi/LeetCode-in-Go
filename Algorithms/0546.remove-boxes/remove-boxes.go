package Problem0546

func removeBoxes(boxes []int) int {
	n := len(boxes)
	if n == 0 {
		return 0
	}

	// dp[i][j][k] 的值是，
	// k 个连续和 boxes[i] 一样颜色的box 再加上 boxes[i:j+1] 可以获取的最大值
	// 所以，题目要求的是 dp[0][n-1][0] 的值
	dp := make([][][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([][]int, n)
		for j := 0; j < n; j++ {
			dp[i][j] = make([]int, n)
		}
	}

	for j := 0; j < n; j++ {
		// k<=j，是因为 boxes[j] 的前面，最多只能有 j 个颜色和 boxes[j] 一样的箱子
		for k := 0; k <= j; k++ {
			// d[j][j][k] 的含义是
			// k 个和boxes[j] 一样颜色的 box 再加上一个 boxes[j] 能够获取的最大积分
			// 即 k+1 个一样颜色的 box 可以得到的积分
			// 按照题意，当然是 (k+1)*(k+1)
			dp[j][j][k] = (k + 1) * (k + 1)
		}
	}

	// l == j - i
	// l+1 == len(boxes[i:j+1])
	// l 循环控制 boxes[i:j+1] 从短边长
	for l := 1; l < n; l++ {
		// j 循环控制 boxes[i:j+1] 从左往右移动
		for j := l; j < n; j++ {
			i := j - l

			for k := 0; k <= i; k++ {
				res := (k+1)*(k+1) + dp[i+1][j][0]

				for m := i + 1; m <= j; m++ {
					if boxes[m] == boxes[i] {
						res = max(res, dp[i+1][m-1][0]+dp[m][j][k+1])
					}
				}

				dp[i][j][k] = res
			}
		}
	}

	return dp[0][n-1][0]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

package problem0576

const mod = 1e9 + 7

func findPaths(m, n, N, i, j int) int {
	dp := [50][50]int{}

	for k := 0; k < N; k++ {
		prior := make([]int, n)
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				paths := 0

				if i == 0 {
					paths++
				} else {
					paths += prior[j]
				}

				if j == 0 {
					paths++
				} else {
					paths += prior[j-1]
				}

				if i == m-1 {
					paths++
				} else {
					paths += dp[i+1][j]
				}

				if j == n-1 {
					paths++
				} else {
					paths += dp[i][j+1]
				}

				paths %= mod
				prior[j] = dp[i][j]
				dp[i][j] = paths
			}
		}
	}

	return dp[i][j]
}

package problem0879

const mod = 1e9 + 7

func profitableSchemes(G, P int, group, profit []int) int {
	/* dp[p][g] = x 表示，
	* g 个人想要获得收益 p ，一共有 x 种选择
	* 其中的 g 个人，每个人都有事情做，而且人和人之间没有区别。
	* 特别地，
	* dp[P][g] = x 表示，
	* g 个无差别的人，每个人都有事情做，总的收益 >=P 的选择数为 x
	 */
	dp := [101][101]int{}
	dp[0][0] = 1

	size := len(group)

	for k := 0; k < size; k++ {
		gk := group[k]
		pk := profit[k]
		for i := P; i >= 0; i-- {
			ip := min(i+pk, P)
			for j := G - gk; j >= 0; j-- {
				// 首先假设 dp 的行数有无穷多
				// 那么 dp[i+p][j+g] = dp[i][j] 的含义是
				// j+g 比 j 多 g 个人，可以取得 i+p 的收益
				// 可惜 dp 的行数没有无穷多，需要压缩
				// 需要把 dp[>=P][j] 求和放置到 dp[P][j]
				// ip := min(i+p, P) 就是压缩过程
				dp[ip][j+gk] += dp[i][j]
				dp[ip][j+gk] %= mod
			}
		}
	}

	res := 0
	for i := 0; i <= G; i++ {
		res += dp[P][i]
	}
	return res % mod
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

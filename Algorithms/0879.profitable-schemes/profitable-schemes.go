package problem0879

const mod = 1e9 + 7

func profitableSchemes(G, P int, group, profit []int) int {
	// dp[p][g] = x 表示， g 个人想要获得收益 p ，一共有 x 种选择
	dp := [101][101]int{}
	dp[0][0] = 1

	size := len(group)
	for k := 0; k < size; k++ {
		g := group[k]
		p := profit[k]
		for i := P; i >= 0; i-- {
			ip := min(i+p, P)
			for j := G - g; j >= 0; j-- {
				dp[ip][j+g] += dp[i][j]
				dp[ip][j+g] %= mod
			}
		}
	}

	res := 0
	for i, x := range dp[P] {
		if i > G {
			break
		}
		res = (res + x) % mod
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

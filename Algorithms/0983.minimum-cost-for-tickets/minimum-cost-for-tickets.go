package problem0983

// ref: https://leetcode.com/problems/minimum-cost-for-tickets/discuss/226659/Two-DP-solutions-with-pictures
func mincostTickets(days []int, costs []int) int {
	c1, c7, c30 := costs[0], costs[1], costs[2]
	endDay := days[len(days)-1]

	dp := [366]int{}

	for d, i := 1, 0; d <= endDay; d++ {
		if d == days[i] {
			dp[d] = min(dp[d-1]+c1, dp[max(0, d-7)]+c7, dp[max(0, d-30)]+c30)
			i++
			continue
		}
		dp[d] = dp[d-1]
	}

	return dp[endDay]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b, c int) int {
	if a <= b && a <= c {
		return a
	}
	if b <= a && b <= c {
		return b
	}
	return c
}

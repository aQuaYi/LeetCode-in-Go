package problem0956

// ref: https://leetcode.com/problems/tallest-billboard/discuss/203181/JavaC%2B%2BPython-DP-min(O(SN2)-O(3N2-*-N)
func tallestBillboard(rods []int) int {
	dp := [5001]int{}
	for d := 1; d < 5001; d++ {
		dp[d] = -10000
	}
	for _, x := range rods {
		cur := dp
		for d := 0; d+x < 5001; d++ {
			dp[d+x] = max(dp[d+x], cur[d])
			dp[abs(d-x)] = max(dp[abs(d-x)], cur[d]+min(d, x))
		}
	}
	return dp[0]
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
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

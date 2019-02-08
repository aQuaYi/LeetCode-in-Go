package problem0956

// ref: https://leetcode.com/problems/tallest-billboard/discuss/203181/JavaC%2B%2BPython-DP-min(O(SN2)-O(3N2-*-N)

const maxDiff = 5001

func tallestBillboard(rods []int) int {
	dp := [maxDiff]int{}
	// a pair of sum (a, b) with a > b, then dp[a - b] = b
	// 每次都用更大的值去更新 dp[diff]
	// 最后想要求的解在 dp[0] 中
	for d := 1; d < maxDiff; d++ {
		dp[d] = -10000
	}
	// NOTICE: dp[0]=0
	for _, r := range rods {
		cur := dp
		for d := 0; d+r < maxDiff; d++ {
			// add r to the tall side
			dp[d+r] = max(dp[d+r], cur[d])
			// add r to the low side
			adr := abs(d - r)
			dp[adr] = max(dp[adr], cur[d]+min(d, r))
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

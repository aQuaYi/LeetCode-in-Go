package problem0983

var intervals = []int{1, 7, 30}

func mincostTickets(days []int, costs []int) int {
	size := len(days)
	endDay := days[size-1]

	dp := [366]int{}
	for i := 1; i <= endDay; i++ {
		dp[i] = 365000
	}

	for i := 1; i <= endDay; i++ {
		if i != days[0] {
			dp[i] = dp[i-1]
			continue
		}
		for j := 0; j < 3 && i-intervals[j] >= 0; j++ {
			dp[i] = min(dp[i], dp[i-intervals[j]]+costs[j])
		}
		days = days[1:]
	}

	return dp[endDay]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

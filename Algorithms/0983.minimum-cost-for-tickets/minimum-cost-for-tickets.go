package problem0983

func mincostTickets(days []int, costs []int) int {
	size := len(days)
	dp := [366]int{}
	endDay := days[size-1]

	ds := days
	for i, c := 1, costs[0]; i <= endDay; i++ {
		if i == ds[0] {
			dp[i] = dp[i-1] + c
			ds = ds[1:]
			continue
		}
		dp[i] = dp[i-1]
	}

	ds = days
	for len(ds) > 0 && ds[0] < 7 {
		ds = ds[1:]
	}
	for i, c := 7, costs[1]; i <= endDay; i++ {
		if i == ds[0] {
			dp[i] = min(dp[i], dp[i-7]+c)
			ds = ds[1:]
			continue
		}
		dp[i] = dp[i-1]
	}

	ds = days
	for len(ds) > 0 && ds[0] < 30 {
		ds = ds[1:]
	}
	for i, c := 30, costs[2]; i <= endDay; i++ {
		if i == ds[0] {
			dp[i] = min(dp[i], dp[i-30]+c)
			ds = ds[1:]
			continue
		}
		dp[i] = dp[i-1]
	}

	return dp[endDay]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

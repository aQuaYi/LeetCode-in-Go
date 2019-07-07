package problem1052

func maxSatisfied(customers []int, grumpy []int, X int) int {
	n := len(customers)
	sum := 0
	// dis[i] is all dissatisfied customers in first i minutes
	dis := make([]int, n+1)
	for i := 1; i <= n; i++ {
		sum += customers[i-1]
		dis[i] = dis[i-1] + customers[i-1]*grumpy[i-1]
	}

	maxDisInX := 0
	for i := X; i <= n; i++ {
		maxDisInX = max(maxDisInX, dis[i]-dis[i-X])
	}

	satisfied := sum - dis[n]

	return satisfied + maxDisInX
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

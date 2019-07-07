package problem1052

func maxSatisfied(customers []int, grumpy []int, X int) int {
	n := len(customers)
	s := make([]int, n+1)
	u := make([]int, n+1)
	sum := 0
	for i := 1; i <= n; i++ {
		sum += customers[i-1]
		u[i] = u[i-1] + customers[i-1]*grumpy[i-1]
		s[i] = sum - u[i]
	}
	return 0
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

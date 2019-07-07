package problem1052

func maxSatisfied(customers []int, grumpy []int, X int) int {
	n := len(customers)
	// s := make([]int, n+1)
	unsatisfying := make([]int, n+1)
	sum, satisfying := 0, 0
	for i := 1; i <= n; i++ {
		sum += customers[i-1]
		unsatisfying[i] = unsatisfying[i-1] + customers[i-1]*grumpy[i-1]
		satisfying = sum - unsatisfying[i]
	}
	maxUnsatisfying := 0
	for i := X; i <= n; i++ {
		maxUnsatisfying = max(maxUnsatisfying, unsatisfying[i]-unsatisfying[i-X])
	}
	return satisfying + maxUnsatisfying
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

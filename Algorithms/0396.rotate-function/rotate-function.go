package problem0396

func maxRotateFunction(a []int) int {
	n := len(a)

	var F, sum int
	for i := 0; i < n; i++ {
		F += i * a[i]
		sum += a[i]
	}

	maxF := F
	for i := n - 1; 1 <= i; i-- {
		F += sum - n*a[i]
		maxF = max(F, maxF)
	}

	return maxF
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

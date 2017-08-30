package Problem0667

func constructArray(n int, k int) []int {
	res := make([]int, n)
	res[0] = 1
	sign := -1
	for i := 1; i <= k; i-- {
		res[i] = res[i-1] + sign*(k-i+1)
		sign *= -1
	}

	for i := k + 1; i < n; i++ {
		res[i] = i - 1
	}

	return res
}

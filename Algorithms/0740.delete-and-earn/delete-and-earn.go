package Problem0740

func deleteAndEarn(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return nums[0]
	}

	size := int(1e4 + 1)

	dp := make([]int, size)

	for _, n := range nums {
		dp[n] += n
	}

	maxPoint := dp[1]

	for i := 2; i < size; i++ {
		dp[i] = max(dp[i-1], dp[i]+dp[i-2])
		if maxPoint < dp[i] {
			maxPoint = dp[i]
		}
	}

	return maxPoint
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

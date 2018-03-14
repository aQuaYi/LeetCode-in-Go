package problem0416

// 01背包问题
func canPartition(nums []int) bool {
	sum := 0
	for _, n := range nums {
		sum += n
	}

	if sum&1 == 1 {
		return false
	}

	sum = sum >> 1
	n := len(nums)

	// dp[i][j] 表示 nums[:i] 中的元素，可以找出一些，他们的和为 j
	dp := make([][]bool, n+1)
	for i := range dp {
		dp[i] = make([]bool, sum+1)
	}

	for i := 0; i < n+1; i++ {
		// 从任意多个元素中，挑选 0 个元素出来，其和是 0
		dp[i][0] = true
	}

	for j := 1; j < sum+1; j++ {
		// 从包含 0 个元素的 nums 中，挑不出来元素，使得其和为 j
		dp[0][j] = false
	}

	for i := 1; i < n+1; i++ {
		for j := 1; j < sum+1; j++ {
			dp[i][j] = dp[i-1][j]
			if j >= nums[i-1] {
				// nums[:i] 比 nums[:i-1] 多了 nums[i-1]，所以
				// 要么，nums[:i-1] 中有元素可以合成 j-nums[i-1]
				// 要么，nums[:i-1] 中有元素可以合成 j
				// nums[:i] 中才可能有元素合成 j
				dp[i][j] = dp[i][j] || dp[i-1][j-nums[i-1]]
			}
		}
	}

	return dp[n][sum]
}

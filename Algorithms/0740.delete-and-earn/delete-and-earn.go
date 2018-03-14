package problem0740

func deleteAndEarn(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		return nums[0]
	}

	size := int(1e4 + 1)
	// dp[n] 表示，删除 nums 中所有 <=n 的数后，能获取的最大分数
	dp := make([]int, size)
	for _, n := range nums {
		dp[n] += n
	}

	for i := 2; i < size; i++ {
		// 删除了 i-1，所以 i 是被消除掉的，不能累加 point
		//              ↓
		dp[i] = max(dp[i-1], dp[i]+dp[i-2])
		//                        ↑
		// i-1 是被消除掉的，所以，此时是 dp[i]+dp[i-2]
	}

	return dp[size-1]
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

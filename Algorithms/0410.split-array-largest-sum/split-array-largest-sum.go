package Problem0410

func splitArray(nums []int, m int) int {
	var Max, n, size, i, sum int
	for _, n := range nums {
		sum += n
		if Max < n {
			Max = n
		}
	}

	if Max > sum/m {
		return Max
	}

	sum /= m
	size = len(nums)

	// 从此开始，在 nums 找到一个子集，使得其和为 >=sum 的最小值

	// dp[i][j][k] 表示 在 nums[i:j] 中
	return res
}

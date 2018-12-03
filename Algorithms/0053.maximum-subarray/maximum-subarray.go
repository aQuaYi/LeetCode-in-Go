package problem0053

func maxSubArray(nums []int) int {
	sum, maxSum := -1<<31, -1<<31
	for _, n := range nums {
		// sum+n < n，那就还不如直接从 n 开始统计。
		sum = max(sum+n, n)
		maxSum = max(maxSum, sum)
	}
	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

package problem0053

func maxSubArray(nums []int) int {
	size := len(nums)

	sum := nums[0]
	res := sum

	for i := 1; i < size; i++ {
		// 题目要求是连续的子数组，假设 maxSum = sum(nums[i:k])
		// 可知必有 sum(nums[i:j]) >= 0 ,其中 i<j<k
		if sum < 0 {
			sum = nums[i]
		} else {
			sum += nums[i]
		}
		res = max(res, sum)
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

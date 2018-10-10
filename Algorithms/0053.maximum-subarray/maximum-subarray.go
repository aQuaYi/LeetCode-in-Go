package problem0053

func maxSubArray(nums []int) int {
	size := len(nums)
	tmp := nums[0]
	res := tmp

	for i := 1; i < size; i++ {
		// 题目要求是连续的子数组，假设 maxSum = sum(nums[i:k])
		// 可知必有 sum(nums[i:j]) >= 0 ,其中 i<j<k
		if tmp < 0 {
			tmp = nums[i]
		} else {
			tmp += nums[i]
		}
		res = max(res, tmp)
	}

	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

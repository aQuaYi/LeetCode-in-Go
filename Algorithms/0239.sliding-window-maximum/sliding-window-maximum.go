package problem0239

func maxSlidingWindow(nums []int, k int) []int {
	// 题目里面说好了 1 <= k 的
	// testcase 里面还是出现了 k == 0
	if k == 0 {
		return nil
	}

	// res 的长度是可计算的
	res := make([]int, len(nums)-k+1)
	for i := 0; i+k <= len(nums); i++ {
		res[i] = maxOf(nums[i : i+k])
	}

	return res
}

// 获取局部的最大值
func maxOf(nums []int) int {
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if max < nums[i] {
			max = nums[i]
		}
	}
	return max
}

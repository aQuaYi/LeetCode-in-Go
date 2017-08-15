package Problem0053

func maxSubArray(nums []int) int {
	l := len(nums)

	if l == 0 {
		return 0
	}

	if l == 1 {
		return nums[0]
	}

	temp := nums[0]
	max := temp
	i := 1
	// 可以把 for 循环的过程，nums[:1] 每次增加一个，直到变成 nums 的过程。
	// temp 就是这个过程中，每个 nums[:i] 中的 max(nums[x:i]) (x=0,1,...,i-1)
	// 整个 nums 中连续子序列的最大值，就在 temp 的所有取值中
	for i < l {
		if temp < 0 {
			// 因为 连续性 的要求，而此时 temp < 0
			// temp 从 i 处重新开始
			temp = nums[i]
		} else {
			temp += nums[i]
		}

		if max < temp {
			max = temp
		}

		i++
	}

	return max
}

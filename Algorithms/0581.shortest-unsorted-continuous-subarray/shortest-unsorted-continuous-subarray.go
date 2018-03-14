package problem0581

func findUnsortedSubarray(nums []int) int {
	n := len(nums)
	left, right := 0, -1 // left和right的取值，只要满足 right - left + 1 == 0 即可
	min, max := nums[n-1], nums[0]

	for i := 1; i < n; i++ {
		if max <= nums[i] { // 如果 nums 是递增的话，if 的判断语句应该总是成立
			max = nums[i]
		} else { // 说明 nums[i] 比它前面的数小，它错位了，需要被重新排序
			right = i
		}

		// min 与 max 同理
		j := n - i - 1
		if min >= nums[j] {
			min = nums[j]
		} else {
			left = j
		}
	}

	return right - left + 1
}

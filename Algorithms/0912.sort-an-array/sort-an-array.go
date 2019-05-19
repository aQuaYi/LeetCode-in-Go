package problem0912

func sortArray(nums []int) []int {
	quick3way(nums, 0, len(nums)-1)
	return nums
}

// 3 way quick sort
func quick3way(nums []int, lo, hi int) {
	if lo >= hi {
		return
	}
	lt, i, gt := lo, lo+1, hi
	v := nums[lo]
	for i <= gt {
		switch {
		case nums[i] < v:
			nums[lt], nums[i] = nums[i], nums[lt]
			lt++
			i++
		case nums[i] > v:
			nums[i], nums[gt] = nums[gt], nums[i]
			gt--
		default:
			i++
		}
	}
	quick3way(nums, lo, lt-1)
	quick3way(nums, gt+1, hi)
}

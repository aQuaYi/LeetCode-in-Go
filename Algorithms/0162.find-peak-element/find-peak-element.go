package problem0162

func findPeakElement(nums []int) int {
	low := -1
	high := len(nums)
	mid := (low + high) / 2

	// low +2 < high
	// low + 1 <= mid < high -1
	for low+1 <= mid && mid < high-1 {
		if nums[mid] < nums[mid+1] {
			low = mid
		} else {
			high = mid + 1
		}
		mid = (high + low) / 2
	}

	return low + 1
}

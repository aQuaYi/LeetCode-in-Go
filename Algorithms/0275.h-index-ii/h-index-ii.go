package Problem0275

// citations 为升序排列
func hIndex(nums []int) int {
	size := len(nums)

	lo, hi := 0, size-1
	var mid, midAO int // mid + midAO +1 == size
	for lo <= hi {
		mid = (lo + hi) / 2
		midAO = size - mid - 1
		if nums[midAO] > mid {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}

	return size - lo
}

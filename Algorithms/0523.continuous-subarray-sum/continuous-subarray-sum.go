package Problem0523

func checkSubarraySum(nums []int, k int) bool {
	n := len(nums)
	index := make(map[int]int, n)
	index[0] = -1

	sum := 0

	for i := 0; i < n; i++ {
		sum += nums[i]
		if k != 0 {
			sum %= k
		}
		idx, ok := index[sum]
		if ok && i-idx > 1 {
			return true
		}
		index[sum] = i
	}

	return false
}

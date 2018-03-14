package problem0523

func checkSubarraySum(nums []int, k int) bool {
	n := len(nums)

	// index 中记录了，每个不同的余数，最早出现的位置
	index := make(map[int]int, n)
	// index[0] = -1 是为了以下情况准备的
	// {
	// 	[]int{0, 0},
	// 	0,
	// 	true,
	// },
	index[0] = -1

	sum := 0

	for i := 0; i < n; i++ {
		sum += nums[i]
		if k != 0 {
			sum %= k
		}

		idx, ok := index[sum]
		if ok {
			if i-idx > 1 {
				return true
			}
		} else {
			// 这个 if - else 结构确保了 index 中记录的值是最早出现的
			index[sum] = i
		}

	}

	return false
}

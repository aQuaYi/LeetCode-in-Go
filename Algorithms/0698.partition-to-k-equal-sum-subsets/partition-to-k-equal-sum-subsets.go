package Problem0698

func canPartitionKSubsets(nums []int, k int) bool {
	size := len(nums)
	sum := 0
	max := nums[0]
	for _, n := range nums {
		sum += n
		if max < n {
			max = n
		}
	}

	if sum%4 != 0 || sum/size < max {
		return false
	}

	avg := sum / size

	res := true

	return res
}

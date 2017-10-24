package Problem0416

import "sort"

func canPartition(nums []int) bool {
	sum := 0
	for _, n := range nums {
		sum += n
	}

	if sum%2 == 1 {
		return false
	}

	half := sum >> 1
	left := 0
	sort.Ints(nums)
	i := 0
	for left < half {
		left += nums[i]
		i++
	}

	if left == half {
		return true
	}

	return dfs(nums[:i], left, half)
}

func dfs(nums []int, sum, target int) bool {
	for i := range nums {
		if sum-nums[i] == target {
			return true
		}
	}
	for i := range nums {
		if sum-nums[i] > target && dfs(append(nums[:i], nums[i+1:]...), sum-nums[i], target) {
			return true
		}
	}

	return false
}

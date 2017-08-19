package Problem0532

import (
	"sort"
)

func findPairs(nums []int, k int) int {
	switch {
	case k < 0:
		return 0
	case k == 0:
		return handleZero(nums)
	default:
		return handle(nums, k)
	}
}

func handleZero(nums []int) int {
	sort.Ints(nums)
	ans := 0
	for i := 0; i+1 < len(nums); i++ {
		if nums[i] == nums[i+1] {
			ans++
		}
		for i+1 < len(nums) && nums[i] == nums[i+1] {
			i++
		}
	}
	return ans
}

func handle(nums []int, k int) int {
	sort.Ints(nums)
	uniqueNums := make([]int, 0, len(nums))
	uniqueMap := make(map[int]bool, len(nums))
	for i := 0; i < len(nums); i++ {
		uniqueNums = append(uniqueNums, nums[i])
		uniqueMap[nums[i]] = true
		for i+1 < len(nums) && nums[i] == nums[i+1] {
			i++
		}
	}

	ans := 0

	for i := 0; i < len(uniqueNums); i++ {
		if uniqueMap[uniqueNums[i]+k] {
			ans++
		}
	}
	return ans
}

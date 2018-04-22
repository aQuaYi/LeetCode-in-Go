package problem0810

import "sort"

func xorGame(nums []int) bool {
	sort.Ints(nums)
	return game(nums)
}

func game(nums []int) bool {
	if xor(nums) == 0 {
		return true
	}
	idx := 0
	num := -1
	for ; idx < len(nums); idx++ {
		if nums[idx] == num {
			continue
		}
		temp := make([]int, len(nums)-1)
		copy(temp[:idx], nums[:idx])
		copy(temp[idx:], nums[idx+1:])
		if !game(temp) {
			return true
		}
		num = nums[idx]
	}
	return false
}

func xor(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	res := nums[0]

	for i := 1; i < len(nums); i++ {
		res ^= nums[i]
	}

	return res
}

package Problem0315

import "sort"

func countSmaller(nums []int) []int {
	size := len(nums)
	res := make([]int, size)
	sorted := make([]int, size)
	copy(sorted, nums)
	sort.Ints(sorted)

	for i := 0; i < size; i++ {
		j := sort.SearchInts(sorted, nums[i])
		res[i] = size - i - j - 1
	}
	return res
}

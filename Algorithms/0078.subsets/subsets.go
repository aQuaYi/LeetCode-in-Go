package Problem0078

func subsets(nums []int) [][]int {
	res := [][]int{}

	cur(nums, []int{}, &res)

	return res
}

func cur(nums, temp []int, res *[][]int) {
	if len(nums) == 0 {
		*res = append(*res, temp)
		return
	}

	cur(nums[1:], temp, res)

	cur(nums[1:], append(temp, nums[0]), res)

}

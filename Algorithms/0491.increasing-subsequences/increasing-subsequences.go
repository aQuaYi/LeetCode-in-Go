package Problem0491

func findSubsequences(nums []int) [][]int {
	size := len(nums)
	res := make([][]int, 0, size*size)
	list := make([]int, 0, size)
	helper(nums, 0, &res, list)
	return res
}

func helper(nums []int, index int, res *[][]int, list []int) {
	if index == len(nums) {
		if len(list) > 1 {
			temp := make([]int, len(list))
			copy(temp, list)
			*res = append((*res), temp)
		}
		return
	}

	if len(list) == 0 ||
		list[len(list)-1] <= nums[index] {
		list = append(list, nums[index])
		helper(nums, index+1, res, list)
		list = list[:len(list)-1]
	}

	if len(list) == 0 ||
		list[len(list)-1] != nums[index] {
		helper(nums, index+1, res, list)
	}
}

package Problem0046

func permute(nums []int) [][]int {
	ress := [][]int{}

	makePermute(nums, []int{}, &ress)

	return ress
}

func makePermute(nums, res []int, ress *[][]int) {
	if len(nums) == 0 {
		*ress = append(*ress, res)
	}

	for i, n := range nums {
		temp := make([]int, 0, len(nums))
		temp = append(temp, nums[:i]...)
		temp = append(temp, nums[i+1:]...)

		makePermute(temp, append(res[:len(res):len(res)], n), ress)
	}
}

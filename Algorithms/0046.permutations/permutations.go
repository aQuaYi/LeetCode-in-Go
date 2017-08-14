package Problem0046

func permute(nums []int) [][]int {
	ans := [][]int{}

	makePermute(nums, []int{}, &ans)

	return ans
}

func makePermute(nums, res []int, ans *[][]int) {
	if len(nums) == 0 {
		*ans = append(*ans, res)
	}

	for i, n := range nums {
		temp := make([]int, 0, len(nums))
		temp = append(temp, nums[:i]...)
		temp = append(temp, nums[i+1:]...)

		makePermute(temp, append(res[:len(res):len(res)], n), ans)
	}
}

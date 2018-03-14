package problem0089

func grayCode(n int) []int {
	return recur(n, 1, []int{0})
}

func recur(n, base int, nums []int) []int {
	if n == 0 {
		return nums
	}

	length := len(nums)
	temp := make([]int, length)
	for i := range nums {
		temp[length-i-1] = nums[i] + base
	}

	return recur(n-1, base*2, append(nums, temp...))
}

package Problem0566

func matrixReshape(nums [][]int, r int, c int) [][]int {
	length := len(nums) * len(nums[0])
	if length != r*c {
		return nums
	}

	// 先降维成[]int
	temp := make([]int, 0, r*c)
	for i := range nums {
		temp = append(temp, nums[i]...)
	}

	// 再升维
	res := make([][]int, r)
	for i := 0; i < r; i++ {
		res[i] = temp[i*c : (i+1)*c]
	}

	return res
}

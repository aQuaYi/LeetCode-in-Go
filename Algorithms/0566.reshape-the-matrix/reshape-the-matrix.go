package problem0566

func matrixReshape(nums [][]int, r int, c int) [][]int {
	if len(nums) == 0 || len(nums[0]) == 0 || len(nums)*len(nums[0]) != r*c || len(nums) == r && len(nums[0]) == c {
		return nums
	}

	res := make([][]int, r)
	count, col := 0, len(nums[0])
	for i := range res {
		res[i] = make([]int, c)

		for j := range res[i] {
			res[i][j] = nums[count/col][count%col]
			count++
		}
	}

	return res
}

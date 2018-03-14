package problem0119

func getRow(rowIndex int) []int {
	res := make([]int, 1, rowIndex+1)
	res[0] = 1
	if rowIndex == 0 {
		return res
	}

	for i := 0; i < rowIndex; i++ {
		res = append(res, 1)
		for j := len(res) - 2; j > 0; j-- {
			res[j] += res[j-1]
		}
	}

	return res
}

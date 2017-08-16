package Problem0118

func generate(numRows int) [][]int {
	res := [][]int{}
	if numRows == 0 {
		return res
	}

	res = append(res, []int{1})
	if numRows == 1 {
		return res
	}

	for i := 1; i < numRows; i++ {
		res = append(res, genNext(res[i-1]))
	}

	return res
}

func genNext(p []int) []int {
	temp := make([]int, 1)
	temp = append(temp, p...)
	temp = append(temp, 0)

	for i := len(temp) - 1; i > 0; i-- {
		temp[i] += temp[i-1]
	}

	return temp[1:]
}

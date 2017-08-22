package Problem0054

func spiralOrder(matrix [][]int) []int {
	r := len(matrix)

	if r == 0 {
		return []int{}
	}

	c := len(matrix[0])
	if c == 0 {
		return []int{}
	}
	
	if len(matrix) == 1 {
		return matrix[0]
	}

	res := make([]int, 0, r*c)

	res = append(res, matrix[0]...)

	for i := 1; i < r-1; i++ {
		res = append(res, matrix[i][c-1])
	}

	for j := c - 1; j >= 0; j-- {
		res = append(res, matrix[r-1][j])
	}

	for i := r - 2; i > 0 && c > 1; i-- {
		res = append(res, matrix[i][0])
	}

	if r == 2 || c <= 2 {
		return res
	}

	nextMatrix := make([][]int, 0, r-2)
	for i := 1; i < r-1; i++ {
		nextMatrix = append(nextMatrix, matrix[i][1:c-1])
	}

	return append(res, spiralOrder(nextMatrix)...)
}

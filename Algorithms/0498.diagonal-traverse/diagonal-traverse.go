package problem0498

func findDiagonalOrder(mat [][]int) []int {
	if len(mat) == 0 || len(mat[0]) == 0 {
		return []int{}
	}
	m, n := len(mat), len(mat[0])

	isUpping := true
	next := func(i, j int) (int, int) {
		if isUpping {
			i--
			j++
			if 0 <= i && j < n {
				return i, j
			}

			isUpping = false

			if i < 0 && j < n {
				return 0, j
			}
			return i + 2, j - 1
		}

		i++
		j--
		if i < m && 0 <= j {
			return i, j
		}

		isUpping = true

		if i < m && j < 0 {
			return i, 0
		}
		return i - 1, j + 2
	}

	mn := m * n
	res := make([]int, mn)

	i, j := 0, 0
	for k := 0; k < mn; k++ {
		res[k] = mat[i][j]
		i, j = next(i, j)
	}

	return res
}

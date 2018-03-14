package problem0074

func searchMatrix(mat [][]int, target int) bool {
	m := len(mat)
	if m == 0 {
		return false
	}

	n := len(mat[0])
	if n == 0 {
		return false
	}

	if target < mat[0][0] || mat[m-1][n-1] < target {
		return false
	}

	// 寻找行
	r := 0
	for r < m && mat[r][0] <= target {
		r++
	}
	r--

	// 二分法寻找 target
	i, j := 0, n-1
	for i <= j {
		med := (i + j) / 2
		switch {
		case mat[r][med] < target:
			i = med + 1
		case target < mat[r][med]:
			j = med - 1
		default:
			return true
		}
	}

	return mat[r][j] == target
}

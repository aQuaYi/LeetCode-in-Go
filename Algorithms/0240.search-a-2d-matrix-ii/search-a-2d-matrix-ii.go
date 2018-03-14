package problem0240

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	if m == 0 {
		return false
	}
	n := len(matrix[0])
	if n == 0 {
		return false
	}

	i, j := m-1, 0
	for 0 <= i && j < n {
		if matrix[i][j] == target {
			return true
		}

		if matrix[i][j] < target {
			j++
			// 排除了 (i,j) 上方的所有比 matrix[i][j] 小的元素
		} else {
			i--
			// 排除了 (i,j) 右方的所有比 matrix[i][j] 大的元素
		}
	}
	// 这种方法的效率，取决于 m 与 n 的相对大小关系
	// 当 m == n 时，效率最高
	// 当 m == 1 或 n ==1 时，效率最低
	return false
}

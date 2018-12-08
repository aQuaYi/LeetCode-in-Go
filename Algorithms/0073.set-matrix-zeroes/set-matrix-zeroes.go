package problem0073

func setZeroes(mat [][]int) {
	m, n := len(mat), len(mat[0])
	col0 := 1

	/**
	 * 从上往下，从左往右 扫描矩阵
	 * 利用 mat[i][0] = 0 表示，第 i 行中含有 0
	 * 利用 mat[0][j] = 0 表示，第 j 列中含有 0
	 * 特别地，
	 * mat[0][0] = 0 仅表示，第 0 行中含有 0
	 * Col0 = 0 表示，第 0 列中含有 0
	 * NOTICE: 循环的顺序很重要
	 * 需要保证 mat[i][0] 和 mat[0][j] 被标记后，不再做为别的标记的依据
	 * 要不然的话，标记有可能会被污染
	 */
	for i := 0; i < m; i++ {
		if mat[i][0] == 0 {
			col0 = 0
		}
		for j := 1; j < n; j++ {
			if mat[i][j] == 0 {
				mat[i][0] = 0
				mat[0][j] = 0
			}
		}
	}

	/**
	 * 从下往上，从右往左 扫描矩阵
	 * 并根据前面的标记修改 mat[i][j] 的值
	 * NOTICE: 循环的顺序很重要
	 * 需要保证 mat[i][0] 和 mat[0][j] 被修改后，不再做为别的修改的标记
	 * 要不然的话，标记有可能会被污染
	 */
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 1; j-- {
			if mat[i][0] == 0 || mat[0][j] == 0 {
				mat[i][j] = 0
			}
		}
		if col0 == 0 {
			mat[i][0] = 0
		}
	}

}

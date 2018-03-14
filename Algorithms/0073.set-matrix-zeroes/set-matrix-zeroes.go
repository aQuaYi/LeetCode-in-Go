package problem0073

func setZeroes(m [][]int) {
	rows := make([]bool, len(m))    // rows[i] == true ，代表 i 行存在 0 元素
	cols := make([]bool, len(m[0])) // cols[j] == true ，代表 j 列存在 0 元素

	// 逐个检查元素
	for i := range m {
		for j := range m[i] {
			if m[i][j] == 0 {
				rows[i] = true
				cols[j] = true
			}
		}
	}

	// 按行修改
	for i := range rows {
		if rows[i] {
			for j := range m[i] {
				m[i][j] = 0
			}
		}
	}

	// 按列修改
	for i := range cols {
		if cols[i] {
			for j := range m {
				m[j][i] = 0
			}
		}
	}
}

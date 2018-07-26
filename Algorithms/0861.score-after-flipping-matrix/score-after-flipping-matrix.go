package problem0861

func matrixScore(A [][]int) int {
	m, n := len(A), len(A[0])

	toggleRow := func(i int) {
		for j := 0; j < n; j++ {
			A[i][j] ^= 1
		}
	}

	countCol := func(j int) int {
		c := 0
		for i := 0; i < m; i++ {
			c += A[i][j]
		}
		return c
	}

	// 1. 保证每行的最高位是 1
	for i := 0; i < m; i++ {
		if A[i][0] == 0 {
			toggleRow(i)
		}
	}

	res := m // 因为 m 行的开头都是 1

	// 2. 从第 1 列开始统计每列中 1 的个数
	// 	  当 1 的个数不足 m 的一半时，需要翻转此列
	//    翻转后的 1 的个数为 m-c
	for j := 1; j < n; j++ {
		c := countCol(j)
		if 2*c < m {
			c = m - c
		}
		res = res*2 + c
	}

	return res
}

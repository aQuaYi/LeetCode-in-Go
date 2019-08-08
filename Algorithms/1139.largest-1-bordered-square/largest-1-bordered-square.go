package problem1139

func largest1BorderedSquare(A [][]int) int {
	m, n := len(A), len(A[0])

	h := [101][101]int{}
	v := [101][101]int{}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if A[i-1][j-1] == 1 {
				h[i][j] = h[i][j-1] + 1
				v[i][j] = v[i-1][j] + 1
			}
		}
	}

	for w := min(m, n); w > 0; w-- {
		for i := 1; i+w-1 <= m; i++ {
			for j := 1; j+w-1 <= n; j++ {
				if v[i+w-1][j] >= w && // left border
					v[i+w-1][j+w-1] >= w && // right
					h[i][j+w-1] >= w && // top
					h[i+w-1][j+w-1] >= w { // buttom
					return w * w
				}
			}
		}
	}

	return 0
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

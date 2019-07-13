package problem1074

func numSubmatrixSumTarget(M [][]int, target int) int {
	m, n := len(M), len(M[0])

	sums := [301][301]int{}
	for i := 1; i <= m; i++ {
		t := 0
		for j := 1; j <= n; j++ {
			t += M[i-1][j-1]
			sums[i][j] = t + sums[i-1][j]
		}
	}

	res := 0

	for x1 := 0; x1 < m; x1++ {
		for x2 := x1 + 1; x2 <= m; x2++ {
			for y1 := 0; y1 < n; y1++ {
				for y2 := y1 + 1; y2 <= n; y2++ {
					s := sums[x2][y2] - sums[x1][y2] - sums[x2][y1] + sums[x1][y1]
					if s == target {
						res++
					}
				}
			}
		}
	}

	return res
}

package problem1074

func numSubmatrixSumTarget(M [][]int, target int) int {
	m, n := len(M), len(M[0])

	S := [301][301]int{}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			S[i][j] = S[i][j-1] + M[i-1][j-1]
		}
	}

	res := 0

	for y1 := 0; y1 < n; y1++ {
		for y2 := y1 + 1; y2 <= n; y2++ {
			count := make(map[int]int, m)
			sum := 0
			for x := 0; x <= m; x++ {
				sum += S[x][y2] - S[x][y1]
				res += count[sum-target]
				count[sum]++
			}
		}
	}

	return res
}

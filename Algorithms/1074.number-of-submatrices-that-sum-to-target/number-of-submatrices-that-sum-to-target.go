package problem1074

func numSubmatrixSumTarget(A [][]int, target int) int {
	m, n := len(A), len(A[0])

	for i := 0; i < m; i++ {
		for j := 1; j < n; j++ {
			A[i][j] += A[i][j-1]
		}
	}

	res := 0
	for y1 := 0; y1 < n; y1++ {
		for y2 := y1; y2 < n; y2++ {
			count := make(map[int]int, n)
			count[0] = 1
			cur := 0
			for x := 0; x < m; x++ {
				cur += A[x][y2]
				if y1-1 >= 0 {
					cur -= A[x][y1-1]
				}
				res += count[cur-target]
				count[cur]++
			}
		}
	}

	return res
}

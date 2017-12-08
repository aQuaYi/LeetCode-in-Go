package Problem0547

func findCircleNum(M [][]int) int {
	res := 0
	N := len(M)

	var group func(int)
	group = func(i int) {
		M[i][i] = 0
		for j := 0; j < N; j++ {
			if j != i && M[i][j] == 1 {
				M[j][i] = 0
				group(j)
			}
		}
	}

	for i := 0; i < N; i++ {
		if M[i][i] == 1 {
			res++
			group(i)
		}
	}

	return res
}

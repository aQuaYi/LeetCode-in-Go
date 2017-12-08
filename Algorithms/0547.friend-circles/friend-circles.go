package Problem0547

func findCircleNum(M [][]int) int {
	res := 0
	N := len(M)

	var group func(int)
	group = func(i int) {
		if M[i][i] == 0 {
			return
		}

		M[i][i] = 0
		for j := i + 1; j < N; j++ {
			if M[i][j] == 1 {
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

package Problem0547

func findCircleNum(M [][]int) int {
	res := 0
	N := len(M)
	friends := make([]int, 0, N)

	var group func(int)
	group = func(i int) {

		M[i][i] = 0
		for j := 0; j < N; j++ {
			if M[i][j] == 1 {
				M[i][j] = 0
				M[j][i] = 0
				friends = append(friends, j)
			}
		}

		for len(friends) > 0 {
			f := friends[0]
			friends = friends[1:]
			group(f)
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

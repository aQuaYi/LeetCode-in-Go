package Problem0547

func findCircleNum(M [][]int) int {
	N := len(M)
	res := N

	friend := make([]int,N)
	for i := 0; i < N; i++ {
		friend[i] = i
	}

	friendOf := func(i int) int {
		for i != friend[i] {
			i = friend[i]
		}
		return i
	}

	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			if M[i][j] == 1 {
				fi := friendOf(i)
				fj := friendOf(j)

				if fi != fj {
					res--
					f := min(fi, fj)
					friend[i] = f
					friend[j] = f
				}
			}
		}
	}

	return res
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

package Problem0547

func findCircleNum(M [][]int) int {
	N := len(M)
	res := N

	friend := make([]int, res)
	for i := 0; i < res; i++ {
		friend[i] = i
	}

	// s 和 d 是朋友关系
	// 所以，s 的所有朋友都是 d 的朋友
	merge := func(s, d int) {
		for i, f := range friend {
			if f == s {
				friend[i] = d
			}
		}
	}

	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			if M[i][j] == 1 {
				if friend[i] != friend[j] {
					res--
					merge(friend[j], friend[i])
				}
			}
		}
	}

	return res
}

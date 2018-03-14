package problem0547

func findCircleNum(M [][]int) int {
	N := len(M)
	res := N

	friend := make([]int, res)
	for i := 0; i < res; i++ {
		friend[i] = i
	}

	// s 和 d 是朋友关系
	// 所以，s 的所有朋友都是 d 的朋友
	union := func(s, d int) {
		for i := range friend {
			if friend[i] == s {
				friend[i] = d
			}
		}
	}

	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			if M[i][j] == 1 {
				if friend[i] != friend[j] {
					res--
					union(friend[i], friend[j])
				}
			}
		}
	}

	return res
}

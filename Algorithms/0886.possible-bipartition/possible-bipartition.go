package problem0886

func possibleBipartition(N int, dislikes [][]int) bool {
	diss := make([][]int, N+1)
	for _, d := range dislikes {
		a, b := d[0], d[1]
		diss[a] = append(diss[a], b)
		diss[b] = append(diss[b], a)
	}

	group := make([]int, N+1)

	round := 0
	for i := 1; i <= N; i++ {
		if group[i] != 0 {
			continue
		}

		round++

		flag := round
		group[i] = flag
		queue := []int{i}

		/**bfs search */

		for len(queue) > 0 {
			flag = -flag
			size := len(queue)
			for j := 0; j < size; j++ {
				q := queue[j]
				for _, p := range diss[q] {
					if group[p] == 0 {
						group[p] = flag
						queue = append(queue, p)
					} else if group[p] != flag {
						return false
					}
				}
			}
			queue = queue[size:]
		}

	}

	return true
}

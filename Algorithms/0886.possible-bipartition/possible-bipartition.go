package problem0886

func possibleBipartition(N int, dislikes [][]int) bool {
	/**转换 dislikes 到 diss
	 * 收集每个人 dislike 的所有人
	 */
	diss := make([][]int, N+1)
	for i := range diss {
		diss[i] = make([]int, 0, 10)
	}
	for _, dl := range dislikes {
		a, b := dl[0], dl[1]
		diss[a] = append(diss[a], b)
		diss[b] = append(diss[b], a)
	}

	group := make([]int, N+1)
	flag := 1

	for i := 1; i <= N; i++ {
		if group[i] != 0 {
			continue
		}

		/**bfs */

		group[i] = flag
		queue := make([]int, 1, N+1)
		queue[0] = i

		for len(queue) > 0 {
			flag = -flag
			size := len(queue)
			for j := 0; j < size; j++ {
				qj := queue[j]
				for _, p := range diss[qj] {
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

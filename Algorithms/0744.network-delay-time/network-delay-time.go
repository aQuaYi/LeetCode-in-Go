package Problem0744

func networkDelayTime(times [][]int, N int, K int) int {
	size := N + 1
	receivedTime := make([]int, size)

	maxInt := 1<<31 - 1
	for i := 1; i < size; i++ {
		receivedTime[i] = maxInt
	}
	receivedTime[K] = 0

	time := make([][]int, size)
	for i := range time {
		time[i] = make([]int, size)
		for j := 0; j < size; j++ {
			time[i][j] = maxInt
		}
	}

	for _, t := range times {
		time[t[0]][t[1]] = t[2]
	}

	queue := make([]int, 1, size)
	queue[0] = K
	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]

		for v := 1; v < size; v++ {
			if time[u][v] > 0 &&
				receivedTime[v] > receivedTime[u]+time[u][v] {
				receivedTime[v] = receivedTime[u] + time[u][v]
				queue = append(queue, v)
			}
		}
	}

	res := 0
	for i := 1; i < size; i++ {
		res = max(res, receivedTime[i])
	}

	if res == maxInt {
		return -1
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

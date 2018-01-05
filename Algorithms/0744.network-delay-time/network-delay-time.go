package Problem0744

func networkDelayTime(times [][]int, N int, K int) int {
	size := N + 1
	receivedTime := make([]int, size)

	receivedTime[0] = 1
	maxInt := 1<<63 - 1
	for i := 1; i < size; i++ {
		receivedTime[i] = maxInt
	}
	receivedTime[K] = 0

	time := make([][]int, size)
	for i := range time {
		time[i] = make([]int, size)
	}

	for _, t := range times {
		time[t[0]][t[1]] = t[2]
	}

	queue := make([]int, 1, size)
	for len(queue) > 0 {
		u := queue[0]
		queue = queue[1:]

		for v := 1; v < size; v++ {
			if time[u][v] > 0 && receivedTime[v] > receivedTime[u]+times[u][v] {
				if receivedTime[v] == maxInt {
					receivedTime[0]++
				}
				receivedTime[v] = receivedTime[u] + time[u][v]
				queue = append(queue, v)
			}
		}
	}

	if receivedTime[0] < N {
		return -1
	}

	res := 0
	for i := 1; i < size; i++ {
		res = max(res, receivedTime[i])
	}
	return res
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

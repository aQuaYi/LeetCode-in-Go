package problem0834

func sumOfDistancesInTree(N int, edges [][]int) []int {
	// TODO: 修改程序，只用处理矩阵一半的数据

	dist := make([][]int, N)
	for i := range dist {
		dist[i] = make([]int, N)
		for j := range dist[i] {
			dist[i][j] = -1
		}
	}

	nextsOf := make([][]int, N)
	for _, e := range edges {
		i, j := e[0], e[1]
		nextsOf[i] = append(nextsOf[i], j)
		nextsOf[j] = append(nextsOf[j], i)
		dist[i][j] = 1
		dist[j][i] = 1
	}

	res := make([]int, N)
	for i := 0; i < N; i++ {
		res[i] = sumOf(i, N, nextsOf, dist)
	}
	return res
}

func sumOf(src, N int, nextsOf, dist [][]int) int {
	isChecked := make([]bool, N)
	isChecked[src] = true

	res := 0
	for i := 0; i < N; i++ {
		if i == src {
			continue
		}
		res += getDist(src, i, N, isChecked, nextsOf, dist)
	}

	return res
}

func getDist(src, dst, N int, isChecked []bool, nextsOf, dist [][]int) int {
	if dist[src][dst] != -1 {
		return dist[src][dst]
	}

	res := 0

	next := nextsOf[src]
	for _, n := range next {
		if isChecked[n] {
			continue
		}
		isChecked[n] = true
		res += 1 + getDist(n, dst, N, isChecked, nextsOf, dist)
	}

	dist[src][dst] = res
	dist[dst][src] = res

	return res
}

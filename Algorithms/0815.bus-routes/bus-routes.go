package problem0815

func numBusesToDestination(routes [][]int, S int, T int) int {
	// buses[7]=={0,1} 表示
	// 7 号站点，有 0,1 两辆 bus 停靠
	busesSlice := make(map[int][]int, len(routes))
	for i := 0; i < len(routes); i++ {
		for j := 0; j < len(routes[i]); j++ {
			busesSlice[routes[i][j]] = append(busesSlice[routes[i][j]], i)
		}
	}
	isVisited := make(map[int]bool, len(routes))
	isVisited[S] = true
	ans := 1<<63 - 1
	dfs(S, T, 1, &ans, busesSlice, isVisited, routes)

	if ans == 1<<63-1 {
		return -1
	}
	return ans
}

func dfs(s, t, k int, ans *int, busesSlice map[int][]int, isVisited map[int]bool, routes [][]int) {
	buses := busesSlice[s]

	for _, bus := range buses {
		stops := routes[bus]
		for _, stop := range stops {
			if isVisited[stop] {
				continue
			}

			if stop == t {
				*ans = min(*ans, k)
				return
			}

			if !isVisited[stop] {
				isVisited[stop] = true
				dfs(stop, t, k+1, ans, busesSlice, isVisited, routes)
			}
		}
	}
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

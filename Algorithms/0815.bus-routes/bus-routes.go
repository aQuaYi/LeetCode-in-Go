package problem0815

import (
	"sort"
)

func numBusesToDestination(routes [][]int, S int, T int) int {
	for i := range routes {
		sort.Ints(routes[i])
	}

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
	isTaken := make(map[int]bool, len(routes))
	ans := 1<<63 - 1
	dfs(S, T, 0, &ans, busesSlice, isTaken, isVisited, routes)

	if ans == 1<<63-1 {
		return -1
	}
	return ans
}

func dfs(s, t, k int, ans *int, busesSlice map[int][]int, isTaken, isVisited map[int]bool, routes [][]int) {
	if s == t {
		*ans = min(*ans, k)
	}
	k++

	buses := busesSlice[s]

	for _, bus := range buses {
		if isTaken[bus] {
			continue
		}
		isTaken[bus] = true

		stops := routes[bus]

		idx := sort.SearchInts(stops, t)
		if idx < len(stops) && stops[idx] == t {
			*ans = min(*ans, k)
		}

		for _, stop := range stops {
			if isVisited[stop] {
				continue
			}

			if !isVisited[stop] {
				isVisited[stop] = true
				dfs(stop, t, k, ans, busesSlice, isTaken, isVisited, routes)
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

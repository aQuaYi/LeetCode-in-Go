package problem0815

func numBusesToDestination(routes [][]int, S int, T int) int {
	if S == T {
		return 0
	}

	// busesSlice[7]=={0,1} 表示
	// 7 号站点，有 0,1 两辆 bus 停靠
	busesSlice := make(map[int][]int, len(routes))
	for i := 0; i < len(routes); i++ {
		for j := 0; j < len(routes[i]); j++ {
			busesSlice[routes[i][j]] = append(busesSlice[routes[i][j]], i)
		}
	}

	isVisited := make(map[int]bool, len(routes))
	isVisited[S] = true
	stops := make([]int, 1, len(routes)*len(routes[0]))
	stops[0] = S

	isCheckedBus := make(map[int]bool, len(routes))
	res := 1

	for len(stops) > 0 {
		nextStops := make([]int, 0, len(routes)*len(routes[0]))

		for _, stop := range stops {
			buses := busesSlice[stop]
			for _, bus := range buses {
				if isCheckedBus[bus] {
					continue
				}
				isCheckedBus[bus] = true
				route := routes[bus]

				for _, r := range route {
					if isVisited[r] {
						continue
					}
					isVisited[r] = true
					if r == T {
						return res
					}
					nextStops = append(nextStops, r)
				}
			}
		}

		stops = nextStops
		res++
	}

	return -1
}

package problem0815

func numBusesToDestination(routes [][]int, S int, T int) int {
	if S == T {
		return 0
	}

	// 用于检查 T 是否存在与 routes 中
	// 不存在的话，可以直接返回 -1，提前结束程序
	isSeenT := false

	// busesSlice[7]=={0,1} 表示
	// 7 号站点，会有 0,1 两辆 bus 停靠
	busesSlice := make(map[int][]int, len(routes))
	for i := 0; i < len(routes); i++ {
		for j := 0; j < len(routes[i]); j++ {
			busesSlice[routes[i][j]] = append(busesSlice[routes[i][j]], i)
			if routes[i][j] == T {
				isSeenT = true
			}
		}
	}

	if !isSeenT {
		// T 是不存在的站点，可以直接返回 -1
		return -1
	}

	// 记录所有访问过的车站
	// 由于车站的编号太大，只好使用 Map
	isCheckedStop := make(map[int]bool, len(routes))
	isCheckedStop[S] = true

	// stops 收集每一步可以停靠的所有站点
	stops := make([]int, 1, len(routes)*len(routes[0]))
	stops[0] = S

	// 所有检查过得车辆不再检查
	isCheckedBus := make([]bool, len(routes))

	// 由于 S!=T，res 从 1 开始
	res := 1

	for len(stops) > 0 {
		// 为下一站准备好存放的地方
		nextStops := make([]int, 0, len(routes)*len(routes[0]))

		// 依次查询每个 stop 所能到达的地方
		for _, stop := range stops {
			// 获取从 stop 能搭乘的所有 bus 编号
			buses := busesSlice[stop]
			// 添加每个 bus 能够到达的地方作为 nextStops 中的地点
			for _, bus := range buses {
				// 搭乘过的 bus ，就不用再检查了
				if isCheckedBus[bus] {
					continue
				}
				isCheckedBus[bus] = true

				// 获取 bus 能够到达的所有站点
				route := routes[bus]

				// 分别检查每个站点 r
				for _, r := range route {
					// 访问过的站点，就不用再一次检查了
					if isCheckedStop[r] {
						continue
					}
					isCheckedStop[r] = true

					// 到达目的地
					if r == T {
						// 直接结束程序
						return res
					}

					// 要不然，就把 r 放入 nextStops
					nextStops = append(nextStops, r)
				}
			}
		}

		// 开始下一轮循环前，更新 stops
		stops = nextStops
		// 还需要换乘一次，所以 res++
		res++
	}

	// 依然没有找到
	return -1
}

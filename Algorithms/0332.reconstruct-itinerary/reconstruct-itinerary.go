package problem0332

import "sort"

func findItinerary(tickets [][]string) []string {
	// nexts[k]v: from k to v
	nexts := make(map[string][]string, len(tickets)+1)
	for _, t := range tickets {
		nexts[t[0]] = append(nexts[t[0]], t[1])
	}

	// 对所有的下一站进行排序
	for k := range nexts {
		sort.Strings(nexts[k])
	}

	// 路径集合
	route := make([]string, 0, len(tickets)+1)

	var visit func(string)
	var next string
	visit = func(airport string) {
		// 优先访问值靠前的 下一站
		for len(nexts[airport]) > 0 {
			next = nexts[airport][0]
			nexts[airport] = nexts[airport][1:]
			visit(next)
		}
		// 全部下一站都访问完毕后，把本站添加入路径
		// 所以添加顺序是，先访问的后添加
		route = append(route, airport)
	}

	visit("JFK")

	// 逆转 route
	i, j := 0, len(route)-1
	for i < j {
		route[i], route[j] = route[j], route[i]
		i++
		j--
	}

	return route
}

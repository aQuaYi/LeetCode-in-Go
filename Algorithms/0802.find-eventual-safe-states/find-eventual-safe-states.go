package problem0802

func eventualSafeNodes(graph [][]int) []int {
	isDanger := make([]bool, len(graph))
	res := make([]int, 0, len(graph))

	for i := 0; i < len(graph); i++ {
		if !dfs(i, isDanger, graph) {
			res = append(res, i)
		}
	}

	return res
}

func dfs(src int, isDanger []bool, graph [][]int) bool {
	if isDanger[src] {
		return true
	}

	if len(graph[src]) == 0 {
		return false
	}

	for i := 0; i < len(graph[src]); i++ {
		isDanger[src] = true
		if dfs(graph[src][i], isDanger, graph) {
			return true
		}
		isDanger[src] = false
	}

	return false
}

// func eventualSafeNodes(graph [][]int) []int {
// 	edgeMap := make(map[int]map[int]bool, len(graph))
// 	for i := 0; i < len(graph); i++ {
// 		edgeMap[i] = make(map[int]bool, len(graph[i]))
// 		for j := 0; j < len(graph[i]); j++ {
// 			edgeMap[i][graph[i][j]] = true
// 		}
// 	}
// 	isChanged := true
// 	for {
// 		isChanged = false
// 		dst := 0
// 		for dst = range edgeMap {
// 			if len(edgeMap[dst]) == 0 && edgeMap[dst] != nil {
// 				edgeMap[dst] = nil
// 				isChanged = true
// 				break
// 			}
// 		}
// 		if !isChanged {
// 			break
// 		}
// 		for i := range edgeMap {
// 			if len(edgeMap[i]) == 0 {
// 				continue
// 			}
// 			delete(edgeMap[i], dst)
// 		}
// 	}
// 	res := make([]int, 0, len(graph)-len(edgeMap))
// 	for i := 0; i < len(graph); i++ {
// 		if len(edgeMap[i]) == 0 {
// 			res = append(res, i)
// 		}
// 	}
// 	return res
// }

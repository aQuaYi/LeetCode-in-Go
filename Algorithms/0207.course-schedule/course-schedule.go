package Problem0207

func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := buildGraph(numCourses, prerequisites)
	degrees := buildDegrees(graph)

	for i := 0; i < numCourses; i++ {
		var j int
		for j = 0; j < numCourses; j++ {
			if degrees[j] == 0 {
				break
			}
		}

		if j >= numCourses {
			return false
		}

		degrees[j] = -1
		for _, n := range graph[j] {
			degrees[n]--
		}
	}

	return true
}

func buildGraph(n int, prereq [][]int) [][]int {
	graph := make([][]int, n)

	for _, pair := range prereq {
		graph[pair[1]] = append(graph[pair[1]], pair[0])
	}

	return graph
}

func buildDegrees(graph [][]int) []int {
	degrees := make([]int, len(graph))

	for _, neighbours := range graph {
		for _, n := range neighbours {
			degrees[n]++
		}
	}

	return degrees
}

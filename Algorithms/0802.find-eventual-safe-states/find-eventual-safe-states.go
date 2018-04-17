package problem0802

type color int

const (
	blank color = iota // unvisited
	danger
	safe
)

func eventualSafeNodes(graph [][]int) []int {
	colors := make([]color, len(graph))
	res := make([]int, 0, len(graph))

	for i := 0; i < len(graph); i++ {
		if isSafe(i, colors, graph) {
			res = append(res, i)
		}
	}

	return res
}

func isSafe(src int, colors []color, graph [][]int) bool {
	if colors[src] != blank {
		return colors[src] == safe
	}

	colors[src] = danger

	for i := 0; i < len(graph[src]); i++ {
		if !isSafe(graph[src][i], colors, graph) {
			return false
		}
	}

	colors[src] = safe
	return true
}

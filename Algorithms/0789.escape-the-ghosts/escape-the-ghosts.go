package problem0789

func escapeGhosts(ghosts [][]int, target []int) bool {
	steps := countSteps([]int{0, 0}, target)
	for _, g := range ghosts {
		if steps >= countSteps(g, target) {
			// 表示 ghost 可以在终点等着
			return false
		}
	}
	return true
}

func countSteps(from, to []int) int {
	return abs(to[0]-from[0]) + abs(to[1]-from[1])
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

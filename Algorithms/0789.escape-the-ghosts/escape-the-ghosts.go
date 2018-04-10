package problem0789

func escapeGhosts(ghosts [][]int, target []int) bool {
	steps := target[0] + target[1]
	for _, g := range ghosts {
		if steps > g[0]+g[1] {
			return false
		}
	}
	return true
}

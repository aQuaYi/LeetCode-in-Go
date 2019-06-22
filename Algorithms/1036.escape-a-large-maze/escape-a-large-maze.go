package problem1036

func isEscapePossible(blocked [][]int, source []int, target []int) bool {
	isBlocked := make(map[int]bool, len(blocked))
	for _, b := range blocked {
		x, y := b[0], b[1]
		isBlocked[x<<32+y] = true
	}
	hasSeen := make(map[int]bool, len(blocked))

	// tx, ty := target[0], target[1]

	// var dfs func(int, int) bool
	// dfs = func(x, y int) bool {
	// 	if x == tx && y == ty {
	// 		return true
	// 	}

	// }
	// dfs()

	return false
}

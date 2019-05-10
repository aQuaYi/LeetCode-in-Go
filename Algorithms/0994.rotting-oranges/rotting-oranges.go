package problem0994

var dx = [4]int{0, 0, -1, 1}
var dy = [4]int{-1, 1, 0, 0}

func orangesRotting(grid [][]int) int {
	m, n := len(grid), len(grid[0])

	fresh := 0
	rottens := make([][2]int, 0, m*n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			switch grid[i][j] {
			case 1:
				fresh++
			case 2:
				rottens = append(rottens, [2]int{i, j})
			}
		}
	}

	if fresh == 0 { // special case
		return 0
	}

	count := -1
	for len(rottens) > 0 {
		count++
		size := len(rottens)
		for r := 0; r < size; r++ {
			x, y := rottens[r][0], rottens[r][1]
			for k := 0; k < 4; k++ {
				i, j := x+dx[k], y+dy[k]
				if i < 0 || m <= i ||
					j < 0 || n <= j ||
					grid[i][j] != 1 {
					continue
				}
				grid[i][j] = 2
				fresh--
				rottens = append(rottens, [2]int{i, j})
			}
		}
		rottens = rottens[size:]
	}

	if fresh > 0 {
		return -1
	}
	return count
}

package problem0934

var dx = []int{1, -1, 0, 0}
var dy = []int{0, 0, 1, -1}

func shortestBridge(A [][]int) int {
	l := searchForIsland(A)

	isChecked := [100][100]bool{}
	isChecked[l[0]][l[1]] = true

	island := make([][]int, 0, 1024)
	island = append(island, l)

	m, n := len(A), len(A[0])

	for k := 0; k < len(island); k++ {
		i, j := island[k][0], island[k][1]
		for l := 0; l < 4; l++ {
			x, y := i+dx[l], j+dy[l]
			if 0 <= x && x < m &&
				0 <= y && y < n &&
				A[x][y] == 1 &&
				!isChecked[x][y] {
				island = append(island, []int{x, y})
				isChecked[x][y] = true
			}
		}
	}

	res := 0
	for len(island) > 0 {
		size := len(island)
		for idx := 0; idx < size; idx++ {
			l := island[idx]
			i, j := l[0], l[1]
			for l := 0; l < 4; l++ {
				x, y := i+dx[l], j+dy[l]
				if 0 <= x && x < m &&
					0 <= y && y < n &&
					!isChecked[x][y] {
					if A[x][y] == 0 {
						A[x][y] = 1
						island = append(island, []int{x, y})
						isChecked[x][y] = true
					} else {
						return res
					}
				}
			}
		}
		island = island[size:]
		res++
	}
	panic("A with only one island")
}

func searchForIsland(A [][]int) []int {
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A[0]); j++ {
			if A[i][j] == 1 {
				return []int{i, j}
			}
		}
	}
	panic("A without 1")
}

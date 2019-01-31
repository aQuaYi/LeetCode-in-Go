package problem0934

var dx = []int{1, -1, 0, 0}
var dy = []int{0, 0, 1, -1}

func shortestBridge(A [][]int) int {
	m, n := len(A), len(A[0])

	p, q := searchForIsland(A)

	isChecked := [100][100]bool{}
	isChecked[p][q] = true

	island := make([][]int, 0, 1024)
	island = append(island, []int{p, q})

	// bfs 查找岛中全部的点
	for idx := 0; idx < len(island); idx++ {
		i, j := island[idx][0], island[idx][1]
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

	bridge := island
	length := 0

	// 此时， bridge 中所有的点，在 isChecked 中都被标记为 true

	// bfs 查找最短的桥
	for len(bridge) > 0 {
		size := len(bridge)
		for idx := 0; idx < size; idx++ {
			land := bridge[idx]
			i, j := land[0], land[1]
			for l := 0; l < 4; l++ {
				x, y := i+dx[l], j+dy[l]
				if 0 <= x && x < m &&
					0 <= y && y < n &&
					!isChecked[x][y] {
					if A[x][y] == 1 {
						return length
					}
					A[x][y] = 1
					bridge = append(bridge, []int{x, y})
					isChecked[x][y] = true
				}
			}
		}
		bridge = bridge[size:]
		length++
	}

	panic("A has only one island")
}

func searchForIsland(A [][]int) (int, int) {
	for i := 0; i < len(A); i++ {
		for j := 0; j < len(A[0]); j++ {
			if A[i][j] == 1 {
				return i, j
			}
		}
	}
	panic("A has NO island.")
}

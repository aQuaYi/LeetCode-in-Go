package problem0529

var dx = []int{-1, -1, -1, 0, 0, 1, 1, 1}
var dy = []int{-1, 0, 1, -1, 1, -1, 0, 1}

func updateBoard(board [][]byte, click []int) [][]byte {
	x, y := click[0], click[1]

	// 如果点击的是 'M' 点，修改后，可以直接返回
	if board[x][y] == 'M' {
		board[x][y] = 'X'
		return board
	}

	return bfs(board, x, y)
}

// 如果 (x,y) 点是 'E' 点，则修改所有与 (x,y) 相邻的 'E'，并返回修改后的 board
// 使用 bfs 搜索
func bfs(board [][]byte, x, y int) [][]byte {
	m, n := len(board), len(board[0])
	isAdded := make([]bool, m*n)

	xs := make([]int, 1, m*n)
	ys := make([]int, 1, m*n)
	xs[0] = x
	ys[0] = y
	isAdded[x*n+y] = true

	for len(xs) > 0 {
		x = xs[0]
		y = ys[0]
		xs = xs[1:]
		ys = ys[1:]

		bombs := getBumbs(board, x, y)

		if bombs > 0 {
			board[x][y] = byte(bombs) + '0'
			continue
		}

		board[x][y] = 'B'

		// 把 (x,y) 周围的 'E' 点，放入队列
		for k := 0; k < 8; k++ {
			i := x + dx[k]
			j := y + dy[k]
			if 0 <= i && i < m &&
				0 <= j && j < n &&
				board[i][j] == 'E' &&
				!isAdded[i*n+j] {
				xs = append(xs, i)
				ys = append(ys, j)
				isAdded[i*n+j] = true
			}
		}
	}

	return board
}

// 获取 (x,y) 点周围的炸弹个数
func getBumbs(board [][]byte, x, y int) int {
	m, n := len(board), len(board[0])
	bombs := 0
	for k := 0; k < 8; k++ {
		i := x + dx[k]
		j := y + dy[k]
		if 0 <= i && i < m &&
			0 <= j && j < n &&
			board[i][j] == 'M' {
			bombs++
		}
	}
	return bombs
}

package Problem0419

func countBattleships(board [][]byte) int {
	res := 0
	if len(board) == 0 || len(board[0]) == 0 {
		return res
	}
	m, n := len(board), len(board[0])

	r := make([][]bool, m)
	for i := 0; i < m; i++ {
		r[i] = make([]bool, n)
	}

	ds := [][]int{
		{-1, 0},
		{1, 0},
		{0, 1},
		{0, -1},
	}

	bfs := func(queue [][]int) {
		for len(queue) > 0 {
			p := queue[0]
			queue = queue[1:]
			for _, d := range ds {
				i, j := p[0]+d[0], p[1]+d[1]
				if 0 <= i && i < m &&
					0 <= j && j < n &&
					!r[i][j] &&
					board[i][j] != '.' {
					r[i][j] = true
					queue = append(queue, []int{i, j})
				}
			}
		}
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if !r[i][j] && board[i][j] == 'X' {
				res++
				bfs([][]int{{i, j}})
			}
		}
	}

	return res
}

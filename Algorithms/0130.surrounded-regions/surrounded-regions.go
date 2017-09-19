package Problem0130

func solve(board [][]byte) {
	m := len(board)
	if m <= 2 {
		return
	}

	n := len(board[0])
	if n <= 2 {
		return
	}

	isUncaptured := make([][]bool, m)
	for i := 0; i < m; i++ {
		isUncaptured[i] = make([]bool, n)
	}

	idxM := make([]int, 0, m*n)
	idxN := make([]int, 0, m*n)

	bfs := func(i, j int) {
		isUncaptured[i][j] = true
		idxM = append(idxM, i)
		idxN = append(idxN, j)

		for len(idxM) > 0 {
			i := idxM[0]
			idxM = idxM[1:]
			j := idxN[0]
			idxN = idxN[1:]

			if 0 <= i-1 && board[i-1][j] == 'O' && !isUncaptured[i-1][j] {
				idxM = append(idxM, i-1)
				idxN = append(idxN, j)
				isUncaptured[i-1][j] = true
			}

			if 0 <= j-1 && board[i][j-1] == 'O' && !isUncaptured[i][j-1] {
				idxM = append(idxM, i)
				idxN = append(idxN, j-1)
				isUncaptured[i][j-1] = true
			}

			if i+1 < m && board[i+1][j] == 'O' && !isUncaptured[i+1][j] {
				idxM = append(idxM, i+1)
				idxN = append(idxN, j)
				isUncaptured[i+1][j] = true
			}

			if j+1 < n && board[i][j+1] == 'O' && !isUncaptured[i][j+1] {
				isUncaptured[i][j+1] = true
				idxM = append(idxM, i)
				idxN = append(idxN, j+1)
			}
		}
	}

	for j := 0; j < n; j++ {
		if board[0][j] == 'O' && !isUncaptured[0][j] {
			bfs(0, j)
		}
		if board[m-1][j] == 'O' && !isUncaptured[m-1][j] {
			bfs(m-1, j)
		}
	}

	for i := 0; i < m; i++ {
		if board[i][0] == 'O' && !isUncaptured[i][0] {
			bfs(i, 0)
		}
		if board[i][n-1] == 'O' && !isUncaptured[i][n-1] {
			bfs(i, n-1)
		}
	}

	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			if board[i][j] == 'O' && !isUncaptured[i][j] {
				board[i][j] = 'X'
			}
		}
	}
	return
}

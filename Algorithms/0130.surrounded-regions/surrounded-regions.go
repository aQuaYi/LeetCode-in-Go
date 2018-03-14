package problem0130

func solve(board [][]byte) {
	m := len(board)
	if m <= 2 {
		return
	}

	n := len(board[0])
	if n <= 2 {
		return
	}

	// isEscaped[i][j] == true，表示 (i,j) 点是联通的。
	// 也可以用来查看点 (i,j) 是否已经检查过了
	isEscaped := make([][]bool, m)
	for i := 0; i < m; i++ {
		isEscaped[i] = make([]bool, n)
	}

	// idxM，idxN 分别记录点 (i,j) 的坐标值
	idxM := make([]int, 0, m*n)
	idxN := make([]int, 0, m*n)

	bfs := func(i, j int) {
		isEscaped[i][j] = true
		idxM = append(idxM, i)
		idxN = append(idxN, j)

		for len(idxM) > 0 {
			i := idxM[0]
			j := idxN[0]
			idxM = idxM[1:]
			idxN = idxN[1:]

			// 依次对 (i,j) 的上左下右点进行，bfs检查

			if 0 <= i-1 && board[i-1][j] == 'O' && !isEscaped[i-1][j] {
				idxM = append(idxM, i-1)
				idxN = append(idxN, j)
				isEscaped[i-1][j] = true
			}

			if 0 <= j-1 && board[i][j-1] == 'O' && !isEscaped[i][j-1] {
				idxM = append(idxM, i)
				idxN = append(idxN, j-1)
				isEscaped[i][j-1] = true
			}

			if i+1 < m && board[i+1][j] == 'O' && !isEscaped[i+1][j] {
				idxM = append(idxM, i+1)
				idxN = append(idxN, j)
				isEscaped[i+1][j] = true
			}

			if j+1 < n && board[i][j+1] == 'O' && !isEscaped[i][j+1] {
				idxM = append(idxM, i)
				idxN = append(idxN, j+1)
				isEscaped[i][j+1] = true
			}
		}
	}

	// 联通的区域一定会到达四周，所以从4周开始检查所有的联通区域
	for j := 0; j < n; j++ {
		if board[0][j] == 'O' && !isEscaped[0][j] {
			bfs(0, j)
		}
		if board[m-1][j] == 'O' && !isEscaped[m-1][j] {
			bfs(m-1, j)
		}
	}
	for i := 0; i < m; i++ {
		if board[i][0] == 'O' && !isEscaped[i][0] {
			bfs(i, 0)
		}
		if board[i][n-1] == 'O' && !isEscaped[i][n-1] {
			bfs(i, n-1)
		}
	}

	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			// 修改内部被占领的 'O'
			if board[i][j] == 'O' && !isEscaped[i][j] {
				board[i][j] = 'X'
			}
		}
	}
}

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

	idxM := make([]int, 0, m*n)
	idxN := make([]int, 0, m*n)

	bfs := func(i, j int) bool {
		idxM = append(idxM, i)
		idxN = append(idxN, j)

		for len(idxM) > 0 {
			i := idxM[0]
			idxM = idxM[1:]
			j := idxN[0]
			idxN = idxN[1:]

			if i-1 == 0 && board[i-1][j] == 'O' {

				return true
			}

			if 0 == j-1 && board[i][j-1] == 'O' {

				return true
			}

			if (i == 0 || i == m-1 || j == 0 || j == n-1) && board[i][j] == 'O' {

				return true
			}

			if i+1 < m && board[i+1][j] == 'O' {
				idxM = append(idxM, i+1)
				idxN = append(idxN, j)
			}

			if j+1 < n && board[i][j+1] == 'O' {
				idxM = append(idxM, i)
				idxN = append(idxN, j+1)
			}
		}

		return false
	}

	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			if board[i-1][j] == 'X' && board[i][j-1] == 'X' && board[i][j] == 'O' {

				if !bfs(i, j) {
					board[i][j] = 'X'
				}

				// bfs 以后要清空 idxM 和 idxN
				idxM = idxM[len(idxM):]
				idxN = idxN[len(idxN):]
			}
		}
	}
	return
}

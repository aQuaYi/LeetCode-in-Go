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

	var dfs func(int, int) bool

	dfs = func(i, j int) bool {
		if i == 0 || i == m-1 || j == 0 || j == n-1 {
			return board[i][j] == 'O'
		}

		if i < 0 || i >= m || j < 0 || j >= n || board[i][j] == 'X' {
			return false
		}

		board[i][j] = 'X'
		if dfs(i-1, j) || dfs(i+1, j) || dfs(i, j-1) || dfs(i, j+1) {
			board[i][j] = 'O'
			return true
		}

		return false
	}

	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			if board[i][j] == 'O' {
				dfs(i, j)
			}
		}
	}
}
